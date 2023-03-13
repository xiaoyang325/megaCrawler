package crawlers

import (
	"hash/fnv"
	"sync"
	"time"

	"megaCrawler/crawlers/config"

	"github.com/jpillora/go-tld"
)

var (
	WebMap    = make(map[string]*WebsiteEngine)
	nextTime  = time.Now().Add(3 * time.Second)
	timeMutex = sync.RWMutex{}
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// Register 注册插件控制器
func Register(service string, name string, baseURL string) *WebsiteEngine {
	k, err := tld.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	engine := NewEngine(service, *k)
	if c, ok := config.Configs[service]; !ok {
		engine.Config = &config.Config{
			ID:       service,
			LastIter: time.Time{},
			Disabled: false,
			Name:     name,
		}
		config.Configs[service] = *engine.Config
	} else {
		engine.Config = &c
	}
	WebMap[service] = engine

	return engine
}

func StartAll() {
	for service, engine := range WebMap {
		serviceHash := hash(service)
		engine := engine
		if serviceHash%uint32(Shard.Total) == uint32(Shard.Number) {
			go func() {
				timeMutex.Lock()
				_, err := engine.Scheduler.Every(168).Hour().StartAt(nextTime).Do(StartEngine, engine, false)
				if err != nil {
					Sugar.Error(err)
					return
				}
				nextTime = nextTime.Add(1 * time.Minute)
				engine.Scheduler.StartAsync()
				timeMutex.Unlock()
			}()
		} else {
			delete(WebMap, service)
		}
	}
}
