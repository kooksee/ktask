package dhtml

import (
	"fmt"
	config2 "github.com/kooksee/ktask/internal/config"
	"github.com/kooksee/ktask/internal/utils"
)

func (t *config) TaskPost(data []byte) error {
	c := config2.DefaultConfig()
	d, err := utils.HttpPost(c.KaskUrl+"/tasks", data, nil)
	if err != nil {
		fmt.Println(string(d))
	}
	return err
}
