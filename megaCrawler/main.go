package megaCrawler

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kardianos/service"
	"log"
	commandImpl2 "megaCrawler/megaCrawler/commandImpl"
	"megaCrawler/megaCrawler/config"
	"net/http"
	"sync"
	"time"
)

var Logger service.Logger
var Debug = false

// CrawlerManager Program structures.
// Define Start and Stop methods.
type CrawlerManager struct {
	exit chan struct{}
}

func (c *CrawlerManager) Start(_ service.Service) error {
	if service.Interactive() {
		err := Logger.Info("Running in terminal.")
		if err != nil {
			return err
		}
	} else {
		err := Logger.Info("Running under service manager.")
		if err != nil {
			return err
		}
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

func (c *CrawlerManager) run() error {
	err := Logger.Infof("I'm running %v.", service.Platform())
	StartWebServer()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-c.exit:
			ticker.Stop()
			return nil
		}
	}
}

func (c *CrawlerManager) Stop(_ service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	err := Logger.Info("CrawlerManager are Stopping!")
	if err != nil {
		return err
	}
	close(c.exit)
	err = config.Configs.Save()
	if err != nil {
		return err
	}
	return nil
}

// Start is a blocking function that starts the crawler
func Start() {
	svcFlag := flag.String("service", "", "Control the system service.")
	listFlag := flag.Bool("list", false, "List all current registered websites.")
	getFlag := flag.String("get", "", "Get the status of the selected website.")
	startFlag := flag.String("start", "", "Launch the selected website now.")
	debugFlag := flag.Bool("debug", false, "Show debug log that will spam console")
	testFlag := flag.Bool("test", false, "Test connection for every website registered")

	flag.Parse()

	Debug = *debugFlag

	if *testFlag {
		gp := sync.WaitGroup{}
		var n int
		errorMap := map[string]string{}

		TestAndPrintWebsite := func(engine *websiteEngine) {
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
		commandImpl2.List()
		return
	}

	if *getFlag != "" {
		commandImpl2.Get(*getFlag)
		return
	}

	if *startFlag != "" {
		commandImpl2.Start(*startFlag)
		return
	}

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "MegaCrawler",
		DisplayName: "A crawler that update resources periodically",
		Description: "This is a crawler that update resources periodically",
		Option:      options,
	}

	prg := &CrawlerManager{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	Logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		_ = Logger.Error(err)
	}
}
