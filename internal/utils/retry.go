package utils

import (
	"github.com/kooksee/ktask/internal/errs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func fibonacci() func() int {
	a1, a2 := 0, 1
	return func() int {
		a1, a2 = a2, a1+a2
		return a1
	}
}

func Retry(num int, fn func(l zerolog.Logger) error) (err error) {
	t := fibonacci()
	var sleepTime = 0
	var l zerolog.Logger
	for i := 0; ; i++ {

		l = log.With().
			Err(err).
			Int("retry_num", i).
			Int("sleep_time", sleepTime).
			Str("method", "Retry").Logger()

		if err = fn(l); err == nil || err == errs.NotFound {
			return err
		}

		sleepTime = t()

		if strings.Contains(err.Error(), "timeout") {
			time.Sleep(time.Second * time.Duration(sleepTime))
			continue
		}

		if i > num {
			return err
		}

		time.Sleep(time.Second * time.Duration(sleepTime))
	}
}
