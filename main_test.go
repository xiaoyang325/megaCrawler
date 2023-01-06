package main

import (
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
		WG:     &sync.WaitGroup{},
		News:   Tester.Status{},
		Index:  Tester.Status{},
		Expert: Tester.Status{},
		Report: Tester.Status{},
	}
	Crawler.Test.WG.Add(1)
	target := os.Getenv("TARGET")
	if target == "" {
		t.Error("please set TARGET")
		return
	}
	c, ok := Crawler.WebMap[target]
	if !ok {
		t.Error("no such target")
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
	Crawler.Sugar.Info("News: ", Crawler.Test.News.Count, Crawler.Test.News.FilledCount)
	Crawler.Sugar.Info("Index: ", Crawler.Test.Index.Count, Crawler.Test.Index.FilledCount)
	Crawler.Sugar.Info("Expert: ", Crawler.Test.Expert.Count, Crawler.Test.Expert.FilledCount)
	Crawler.Sugar.Info("Report: ", Crawler.Test.Report.Count, Crawler.Test.Report.FilledCount)
}
