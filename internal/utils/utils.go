package utils

import (
	"fmt"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"regexp"
	"reflect"
)

func MustNotError(err error) {
	if err != nil {
		log.Error().Err(err).Str("method", "MustNotError").Msg("error")
		panic(err.Error())
	}
}

func P(d ... interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(reflect.ValueOf(i).String(),"->",string(dt))
	}
}

func ParseOssUrl(url string) *oss.Config {
	c := &oss.Config{}
	dt := regexp.MustCompile(`oss://(?P<username>.*):(?P<password>.*)@(?P<host>.*)`).FindStringSubmatch(url)
	if len(dt) == 0 {
		panic(fmt.Sprintf("url %s parse error", url))
	}

	c.AccessKeyID = dt[1]
	c.AccessKeySecret = dt[2]
	c.Endpoint = dt[3]
	return c
}
