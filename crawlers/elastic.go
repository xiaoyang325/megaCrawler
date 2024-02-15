package crawlers

import (
	"context"
	"net/url"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	rawNewsChannel   chan news
	rawReportChannel chan report
	rawExpertChannel chan expert
)

func getElasticConsumerChannel() (newsChannel chan news, reportChannel chan report, expertChannel chan expert) {
	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://kb.mingyuaigc.com/es/"},
		Username:  "elastic",
		Password:  passwd,
	})
	if err != nil {
		panic(err)
	}

	newsChannel = make(chan news)
	reportChannel = make(chan report)
	expertChannel = make(chan expert)

	for i := 0; i < 4; i++ {
		go func() {
			for {
				document := <-newsChannel
				Sugar.Debug("Sending to elasticsearch news " + document.URL)
				_, err := typedClient.Index("scsm_news").
					Id(url.QueryEscape(document.URL)).
					Request(document).
					Do(context.TODO())
				if err != nil {
					Sugar.Error("Failed to send message to news ", err)
				}
			}
		}()
	}
	for i := 0; i < 4; i++ {
		go func() {
			for {
				document := <-reportChannel
				Sugar.Debug("Sending to elasticsearch report " + document.URL)
				_, err := typedClient.Index("scsm_report").
					Id(url.QueryEscape(document.URL)).
					Request(document).
					Do(context.TODO())
				if err != nil {
					Sugar.Error("Failed to send message to report ", err)
				}
			}
		}()
	}
	for i := 0; i < 4; i++ {
		go func() {
			for {
				document := <-expertChannel
				Sugar.Debug("Sending to elasticsearch expert " + document.URL)
				_, err := typedClient.Index("scsm_expert").
					Id(url.QueryEscape(document.URL)).
					Request(document).
					Do(context.TODO())
				if err != nil {
					Sugar.Error("Failed to send message to expert ", err)
				}
			}
		}()
	}

	return newsChannel, reportChannel, expertChannel
}
