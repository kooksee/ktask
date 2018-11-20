package dhtml

import (
	"context"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

type config struct {
	ChromeUrls []string
	SleepTime  int
	chromes    []*cdp.Client
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
}

func NewConfig() *config {
	return &config{SleepTime: 3}
}
