package config

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
	"time"
)

type config struct {
	Cache   *cache.Cache
	Debug   bool
	KaskUrl string

	ossCfg *oss.Config
	OssUrl string
	oss    *oss.Client

	id     string
	isInit bool
}

func (t *config) GetOss() *oss.Client {
	if t.oss == nil {
		panic("please init oss")
	}
	return t.oss
}

func (t *config) IsDebug() bool {
	return t.Debug
}

func (t *config) Init() {

	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	if t.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	ip := utils.IpAddress()
	log.Logger = log.Logger.
		With().
		Str("service_name", "ktask").
		Str("service_id", ip).
		Str("service_ip", ip).
		Caller().
		Logger()

	if t.OssUrl != "" {
		ossCfg := utils.ParseOssUrl(t.OssUrl)
		c, err := oss.New(ossCfg.Endpoint, ossCfg.AccessKeyID, ossCfg.AccessKeySecret)
		utils.MustNotError(err)

		rest, err := c.ListBuckets()
		if err != nil {
			utils.P(rest)
			utils.MustNotError(err)
		}
		t.oss = c
	}

}

var cfg *config
var once sync.Once

func DefaultConfig() *config {
	once.Do(func() {
		cfg = &config{
			Debug:   true,
			Cache:   cache.New(time.Minute*10, time.Minute*30),
			KaskUrl: "http://localhost:8080",
		}

		if e := env("Debug"); e != "" {
			cfg.Debug = e == "true"
		}

		if e := env("KaskUrl"); e != "" {
			cfg.KaskUrl = e
		}

		if e := env("OssUrl"); e != "" {
			cfg.OssUrl = e
		}

	})
	return cfg
}
