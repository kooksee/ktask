package dhtml

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/kooksee/ktask/internal/kts"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
	"github.com/rs/zerolog"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

type config struct {
	ChromeUrls []string
	SleepTime  int
	chromes    []*cdp.Client

	RedisUrl string
	redis    *redis.Client
}

func (t *config) GetChrome() *cdp.Client {
	return t.chromes[rand.Int31n(int32(len(t.chromes)))]
}

func (t *config) initChrome(url string) *cdp.Client {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	devt := devtool.New(url)
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			panic(err.Error())
		}
	}

	// Initiate a new RPC connection to the Chrome DevTools Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		panic(err.Error())
	}

	return cdp.NewClient(conn)
}

func (t *config) Init() {
	if e := env("chrome_urls"); e != "" {
		t.ChromeUrls = strings.Split(e, ",")
	}

	for _, url := range t.ChromeUrls {
		t.chromes = append(t.chromes, t.initChrome(url))
	}

	if e := env("sleep_time"); e != "" {
		b, _ := big.NewInt(0).SetString(e, 10)
		t.SleepTime = int(b.Int64())
	}

	if e := env("redis_url"); e != "" {
		t.RedisUrl = e
	}

	opt, err := redis.ParseURL(t.RedisUrl)
	utils.MustNotError(err)
	opt.DialTimeout = time.Minute
	opt.PoolTimeout = time.Minute
	opt.PoolSize = 10
	t.redis = redis.NewClient(opt)
	utils.MustNotError(t.redis.Ping().Err())
}

func (t *config) GetPattern(name string) *kts.HtmlPattern {
	var cmd *redis.StringCmd
	utils.MustNotError(utils.Retry(5, func(l zerolog.Logger) error {
		cmd = t.redis.HGet("mworker:pattern", name)
		if err := cmd.Err(); err != nil {
			l.Error().Err(err).Str("mth", "GetPattern.redis.HGet").Msg(err.Error())
		}
		return cmd.Err()
	}))

	p := &kts.HtmlPattern{}
	utils.MustNotError(p.Decode([]byte(cmd.Val())))

	return p
}

func NewConfig() *config {
	return &config{SleepTime: 3}
}
