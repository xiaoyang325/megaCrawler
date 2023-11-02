package crawlers

import (
	"crypto/sha256"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/xdg/scram"
)

var (
	SHA256 scram.HashGeneratorFcn = sha256.New
)

type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
	x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
	if err != nil {
		return err
	}
	x.ClientConversation = x.Client.NewConversation()
	return nil
}

func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
	response, err = x.ClientConversation.Step(challenge)
	return
}

func (x *XDGSCRAMClient) Done() bool {
	return x.ClientConversation.Done()
}

var (
	userName      = "xzh"
	passwd        = ""
	newsChannel   chan string
	reportChannel chan string
	expertChannel chan string
)

func getProducer() (newsChannel chan string, reportChannel chan string, expertChannel chan string) {
	conf := sarama.NewConfig()
	conf.Producer.Retry.Max = 10
	conf.Producer.Retry.BackoffFunc = func(retries, maxRetries int) time.Duration {
		const maxInt64 = float64(math.MaxInt64 - 512)
		min := 10 * time.Millisecond
		max := 1 * time.Minute

		minf := float64(min)
		durf := minf * math.Pow(2, float64(retries))
		durf = rand.Float64()*(durf-minf) + minf
		if durf > maxInt64 {
			return max
		}
		return time.Duration(durf)
	}
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Metadata.Full = true
	conf.Version = sarama.V3_2_3_0
	conf.ClientID = "sasl_scram_client"
	conf.Metadata.Full = true
	conf.Net.SASL.Enable = true
	conf.Net.SASL.User = userName
	conf.Net.SASL.Password = passwd
	conf.Net.SASL.Handshake = true
	conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
	conf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256

	newsChannel = make(chan string)
	reportChannel = make(chan string)
	expertChannel = make(chan string)

	syncProducer, err := sarama.NewSyncProducer([]string{
		"HW-kafka-01.bdd-o.jnu.edu.cn:19092",
		"HW-kafka-02.bdd-o.jnu.edu.cn:19092",
		"HW-kafka-03.bdd-o.jnu.edu.cn:19092",
		"HW-kafka-04.bdd-o.jnu.edu.cn:19092",
		"HW-kafka-05.bdd-o.jnu.edu.cn:19092",
		"HW-kafka-06.bdd-o.jnu.edu.cn:19092",
	}, conf)

	if err != nil {
		Sugar.Panic("Failed to create producer: ", err)
	}
	wg := sync.WaitGroup{}

	f := func(topic string, channel chan string) {
		wg.Add(1)
		Sugar.Info("Created producer for ", topic)
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

	ProducerLoop:
		for {
			select {
			case message := <-channel:
				if message == "" {
					continue
				}
				partition, offset, err := syncProducer.SendMessage(&sarama.ProducerMessage{
					Topic: topic,
					Value: sarama.StringEncoder(message),
				})
				if err != nil {
					Sugar.Error("Failed to send message to ", topic, err)
				}
				Sugar.Debugf("Wrote message at partition: %d, offset: %d", partition, offset)
			case <-signals:
				break ProducerLoop
			}
		}
		wg.Done()
	}

	go f("mks_news", newsChannel)
	go f("mks_expert", expertChannel)
	go f("mks_report", reportChannel)
	go func() {
		wg.Wait()
		_ = syncProducer.Close()
	}()

	return
}
