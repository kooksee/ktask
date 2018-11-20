package config

import (
	"os"
	"github.com/json-iterator/go"
)

var env = os.Getenv
var json = jsoniter.ConfigCompatibleWithStandardLibrary