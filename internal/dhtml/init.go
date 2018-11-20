package dhtml

import (
	"os"
	"github.com/json-iterator/go"
	"math/rand"
	"time"
)

var env = os.Getenv
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	rand.Seed(time.Now().Unix())
}
