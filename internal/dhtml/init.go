package dhtml

import (
	"github.com/json-iterator/go"
	"math/rand"
	"os"
	"time"
)

var env = os.Getenv
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	rand.Seed(time.Now().Unix())
}
