package utils

import (
	"github.com/storyicon/graphquery"
	"strings"
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func UnMashallHtml(data []byte, pattern string, ret interface{}) {
	res := graphquery.ParseFromString(string(data), pattern)
	if len(res.Errors) != 0 {
		log.Error().Str("method", "UnMashallHtml").Str("error", strings.Join(res.Errors, "->")).Msg("graphquery.ParseFromString error")
		return
	}

	if res.Data == nil {
		log.Error().Str("method", "UnMashallHtml").Msg("graphquery.ParseFromString data is nil")
		return
	}

	dt, err := json.Marshal(res.Data)
	if err != nil {
		log.Error().Str("method", "UnMashallHtml").Err(err).Msg("graphquery.ParseFromString json.Marshal error")
		return
	}

	if err := json.Unmarshal(dt, &ret); err != nil {
		log.Error().Str("method", "UnMashallHtml").Err(err).Msg("graphquery.ParseFromString json.Unmarshal error")
		return
	}
}
