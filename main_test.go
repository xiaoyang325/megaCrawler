package main

import (
	"github.com/olekukonko/tablewriter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"megaCrawler/Crawler"
	"megaCrawler/Crawler/Tester"
	"os"
	"sync"
	"testing"
	"time"
)

func TestTester(t *testing.T) {
	Crawler.Test = &Tester.Tester{
		WG: &sync.WaitGroup{},
		News: Tester.Status{
			Name: "News",
		},
		Index: Tester.Status{
			Name: "Index",
		},
		Expert: Tester.Status{
			Name: "Expert",
		},
		Report: Tester.Status{
			Name: "Report",
		},
	}

	buf, err := os.Create("table.txt")
	if err != nil {
		t.Error(err)
		return
	}

	Crawler.Test.WG.Add(1)
	target := os.Getenv("TARGET")
	if target == "" {
		_, _ = buf.WriteString("No target specified.\nFailed to run tests.\n")
		return
	}
	c, ok := Crawler.WebMap[target]
	if !ok {
		_, _ = buf.WriteString("No such target.\nFailed to start.\n")
		return
	}

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

	Crawler.Sugar = logger.Sugar()
	go Crawler.StartEngine(c, true)
	Crawler.Test.WG.Wait()
	time.Sleep(time.Second * 5)

	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"Target", "Total", "Filled", "Success Rate"})
	Crawler.Test.News.FillTable(table)
	Crawler.Test.Index.FillTable(table)
	Crawler.Test.Expert.FillTable(table)
	Crawler.Test.Report.FillTable(table)
	table.Render()
}
