package kts

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/ktask/internal/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type M map[string]interface{}

func (t M) Encode() []byte {
	dt, err := json.Marshal(t)
	utils.MustNotError(err)
	return dt
}
