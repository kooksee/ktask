package config

import (
	"fmt"
	"github.com/kooksee/ktask/internal/utils"
)

func (t *config) TaskPost(data []byte) error {
	d, err := utils.HttpPost(t.KaskUrl+"/tasks", data, nil)
	if err != nil {
		fmt.Println(string(d))
	}
	return err
}
