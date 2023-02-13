package main

import (
	"fmt"
	"megaCrawler/crawlers"
	"megaCrawler/crawlers/tester"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/olekukonko/tablewriter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestTester(t *testing.T) {
	buf, err := os.Create("table.txt")
	if err != nil {
		t.Error(err)
		return
	}

	target := os.Getenv("TARGET")
	if target == "" {
		_, _ = buf.WriteString("No target specified.\nFailed to run tests.\n")
		return
	}
	targets := strings.Split(target, ",")
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/debug.jsonl",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	var fileCore zapcore.Core
	ProductionEncoder := zap.NewProductionEncoderConfig()

	fileCore = zapcore.NewCore(
		zapcore.NewJSONEncoder(ProductionEncoder),
		w,
		zap.DebugLevel,
	)

	logger := zap.New(fileCore)

	crawlers.Sugar = logger.Sugar()

	for _, target := range targets {
		_, _ = fmt.Fprintf(buf, "Testing %s:\n\n", target)

		c, ok := crawlers.WebMap[target]
		if !ok {
			_, _ = fmt.Fprintf(buf, "No such target %s.\n\n", target)
			continue
		}
		c.Test = &tester.Tester{
			WG: &sync.WaitGroup{},
			News: tester.Status{
				Name: "News",
			},
			Index: tester.Status{
				Name: "Index",
			},
			Expert: tester.Status{
				Name: "Expert",
			},
			Report: tester.Status{
				Name: "Report",
			},
		}
		c.Test.WG.Add(1)

		go crawlers.StartEngine(c, true)
		go func() {
			time.Sleep(2 * time.Minute)
			if !c.Test.Done {
				c.Test.WG.Done()
				c.Test.Done = true
			}
		}()
		c.Test.WG.Wait()

		table := tablewriter.NewWriter(buf)
		table.SetHeader([]string{"Field", "Total", "Passed", "Coverage"})

		c.Test.News.FillTable(table)
		c.Test.Index.FillTable(table)
		c.Test.Expert.FillTable(table)
		c.Test.Report.FillTable(table)

		table.Render()

		_, _ = buf.WriteString("\n")
	}
}
