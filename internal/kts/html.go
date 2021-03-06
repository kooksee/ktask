package kts

import (
	"github.com/kooksee/ktask/internal/utils"
)

type HTMLItem struct {
	URL         string                 `json:"url"`
	PUrl        string                 `json:"purl"`
	Static      bool                   `json:"static,omitempty"`
	Header      map[string]string      `json:"header,omitempty"`
	PatternName string                 `json:"pattern_name"`
	IsList      bool                   `json:"is_list"`
	Extra       map[string]interface{} `json:"data"`
}

func (t *HTMLItem) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *HTMLItem) Encode() []byte {
	dt, err := json.Marshal(t)
	utils.MustNotError(err)
	return dt
}

func (t *HTMLItem) Mock() *HTMLItem {
	return &HTMLItem{
		URL:         "http://roll.news.qq.com",
		PUrl:        "http://roll.news.qq.com",
		Static:      false,
		PatternName: "qq_roll",
		IsList:      true,
	}
}
