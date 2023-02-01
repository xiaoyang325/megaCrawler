package Crawler

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kardianos/service"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"megaCrawler/Crawler/Tester"
	"megaCrawler/Crawler/commands"
	"megaCrawler/Crawler/config"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"sync"
	"time"
)

var Sugar *zap.SugaredLogger
var Debug bool
var Threads int
var Kafka bool
var Proxy *url.URL
var Test *Tester.Tester
var Shard struct {
	Total  int
	Number int
}

// Manager Program structures.
// Define Start and Stop methods.
type Manager struct {
	exit chan struct{}
}

func (c *Manager) Start(_ service.Service) error {
	if service.Interactive() {
		Sugar.Info("Running in terminal.")
	} else {
		Sugar.Info("Running under service manager.")
	}
	c.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go func() {
		err := c.run()
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

func (c *Manager) run() error {
	Sugar.Infof("I'm running %v.", service.Platform())
	StartWebServer()

	ticker := time.NewTicker(2 * time.Second)
	StartAll()
	for {
		select {
		case <-c.exit:
			ticker.Stop()
			return nil
		}
	}
}

func (c *Manager) Stop(_ service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	Sugar.Info("CrawlerManager are Stopping!")
	close(c.exit)
	err := config.Configs.Save()
	if err != nil {
		return err
	}
	return nil
}

// Start is a blocking function that starts the crawler
func Start() {
	listFlag := flag.Bool("list", false, "List all current registered websites.")
	getFlag := flag.String("get", "", "Get the status of the selected website.")
	startFlag := flag.String("start", "", "Launch the selected website now.")
	debugFlag := flag.Bool("debug", false, "Show debug log that will spam console")
	testFlag := flag.Bool("test", false, "Test connection for every website registered")
	updateFlag := flag.Bool("update", false, "Update the program to the latest release version")
	passwordFlag := flag.String("password", "", "The password for kafka server")
	threadFlag := flag.Int("thread", 16, "Number of networking thread")
	totalShardFlag := flag.Int("totalShard", 1, "Total number of shard")
	shardFlag := flag.Int("shard", 0, "The shard number")

	flag.Parse()

	Shard.Total = *totalShardFlag
	Shard.Number = *shardFlag

	Debug = *debugFlag
	Threads = *threadFlag

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/debug.jsonl",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	var fileCore zapcore.Core
	ProductionEncoder := zap.NewProductionEncoderConfig()
	DevEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	if *debugFlag {
		fileCore = zapcore.NewCore(
			zapcore.NewJSONEncoder(ProductionEncoder),
			w,
			zap.DebugLevel,
		)
	} else {
		fileCore = zapcore.NewCore(
			zapcore.NewJSONEncoder(ProductionEncoder),
			w,
			zap.InfoLevel,
		)
	}

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if Debug {
			return lvl < zapcore.ErrorLevel
		}
		return lvl < zapcore.ErrorLevel && lvl > zapcore.DebugLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	tree := zapcore.NewTee(
		fileCore,
		zapcore.NewCore(DevEncoder, consoleDebugging, lowPriority),
		zapcore.NewCore(DevEncoder, consoleErrors, highPriority),
	)
	logger := zap.New(tree)

	Sugar = logger.Sugar()
	if !Debug && *passwordFlag == "" {
		Sugar.Warn("Debug mode and kafka is both not on, you might not see some output.")
	}

	if *updateFlag {
		var Updater *updater.Updater

		if runtime.GOOS == "linux" {
			Updater = &updater.Updater{
				Provider: &provider.Github{
					RepositoryURL: "github.com/foxwhite25/megaCrawler",
					ArchiveName:   fmt.Sprintf("megaCrawler_%s_%s.tar.gz", runtime.GOOS, runtime.GOARCH),
				},
				ExecutableName: "megaCrawler",
				Version:        "v2.2.0",
			}
		} else if runtime.GOOS == "windows" {
			Updater = &updater.Updater{
				Provider: &provider.Github{
					RepositoryURL: "github.com/foxwhite25/megaCrawler",
					ArchiveName:   fmt.Sprintf("megaCrawler_%s_%s.zip", runtime.GOOS, runtime.GOARCH),
				},
				ExecutableName: "megaCrawler.exe",
				Version:        "v2.2.0",
			}
		}

		update, err := Updater.Update()
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		switch update {
		case updater.Updated:
			version, err := Updater.GetLatestVersion()
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			log.Printf("Sucessfully Update to version %s.", version)
		case updater.UpToDate:
			log.Printf("Program is already up to date.")
		}
		return
	}

	if *testFlag {
		gp := sync.WaitGroup{}
		var n int
		errorMap := map[string]string{}

		TestAndPrintWebsite := func(engine *WebsiteEngine) {
			for _, url := range engine.UrlProcessor.startingUrls {
				n++
				resp, err := http.Get(url)
				var reason string
				if err != nil {
					reason = fmt.Sprintf("Error: %s", err.Error())
				} else if resp.StatusCode != 200 {
					reason = fmt.Sprintf("Status Code: %d", resp.StatusCode)
				}
				if reason != "" {
					println("[Error] Website:", url, ",", reason)
					errorMap[url] = reason
				}
				gp.Done()
			}
		}

		for _, engine := range WebMap {
			engine.Scheduler.Stop()
			gp.Add(len(engine.UrlProcessor.startingUrls))
			go TestAndPrintWebsite(engine)
		}
		gp.Wait()

		println("\nFinished testing,", len(errorMap), "site received error out of", n, "site")
		println()
		j, _ := json.MarshalIndent(errorMap, "", "    ")
		println(string(j))
		return
	}

	if *listFlag {
		commands.List()
		return
	}

	if *getFlag != "" {
		commands.Get(*getFlag)
		return
	}

	if *startFlag != "" {
		commands.Start(*startFlag)
		return
	}

	svcConfig := &service.Config{
		Name:        "MegaCrawler",
		DisplayName: "A crawler that update resources periodically",
		Description: "This is a crawler that update resources periodically",
	}

	passwd = *passwordFlag

	if *passwordFlag != "" {
		newsChannel, reportChannel, expertChannel = getProducer()
		Kafka = true
	}

	if proxy := os.Getenv("HTTP_PROXY"); proxy != "" {
		if parsedU, err := url.Parse(proxy); err == nil {
			Proxy = parsedU
		} else {
			Sugar.Panicf("Cannot parse proxy in HTTP_PROXY: %s", proxy)
		}
	}

	prg := &Manager{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	err = s.Run()
}
