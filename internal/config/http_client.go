package config

import (
	"bytes"
	"fmt"
	"github.com/kooksee/ktask/internal/utils"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func (t *config) TaskPost(data []byte) ([]byte, error) {
	return t.Post(t.KaskUrl+"/tasks", data)
}

func (t *config) Post(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	utils.MustNotError(err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		dt, err := ioutil.ReadAll(resp.Body)
		utils.MustNotError(err)
		return nil, fmt.Errorf("url: %s post error,input: %s, output: %s", url, data, dt)
	}

	dt, err := ioutil.ReadAll(resp.Body)
	utils.MustNotError(err)

	return dt, nil
}

func (t *config) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	utils.MustNotError(err)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		dt, err := ioutil.ReadAll(resp.Body)
		utils.MustNotError(err)
		return nil, fmt.Errorf("url: %s get error, output: %s", url, dt)
	}

	dt, err := ioutil.ReadAll(resp.Body)
	utils.MustNotError(err)

	return dt, nil
}
