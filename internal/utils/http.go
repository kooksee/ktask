package utils

import (
	"errors"
	"fmt"
	"gitee.com/johng/gf/g/os/gproc"
	"github.com/kooksee/ktask/internal/errs"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HttpGet(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	MustNotError(err)

	// 处理header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return []byte(fmt.Sprintf("url: %s not found", url)), errs.NotFound
	}

	if resp.StatusCode != http.StatusOK {
		dt, err := ioutil.ReadAll(resp.Body)
		MustNotError(err)
		return nil, fmt.Errorf("url: %s get error, output: %s", url, dt)
	}

	dt, err := ioutil.ReadAll(resp.Body)
	MustNotError(err)

	return dt, nil
}

// probe makes am HTTP request to the site and return site infomation.
// If site is not reachable, return non-nil error.
// If site supports for range request, return the file length (should be greater than 0).
func GetHttpRange(url string) (length int64, err error) {
	// Check whether site is reachable
	req, err := http.NewRequest(http.MethodHead, url, nil)
	MustNotError(err)

	// Do HTTP HEAD request with range header to this site
	client := &http.Client{Timeout: time.Second * 5}
	req.Header.Set("Range", "bytes=0-")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return 0, err
	}

	// Collect site infomation
	switch resp.StatusCode {
	case http.StatusPartialContent:
		P("Break-point is supported in this downloading task.")

		attr := resp.Header.Get("Content-Length")
		length, err = strconv.ParseInt(attr, 10, 0)
		return 0, err
	case http.StatusOK, http.StatusRequestedRangeNotSatisfiable:
		P(url, "does not support for range request.")
		// set length to N/A or unknown
		length = 0
	default:
		P("Got unexpected status code", resp.StatusCode)
		err = errors.New("Unexpected error response when access site: " + url)
	}

	return
}

func WGet(fileName, url string) error {
	resp, _ := http.Head(url)
	if resp.StatusCode == http.StatusNotFound {
		return nil
	}

	t := resp.Header.Get("Content-Type")

	t1 := strings.Split(t, "/")
	return gproc.ShellRun(fmt.Sprintf("wget -o /dev/null -O %s -c %s", fileName+"."+t1[len(t1)-1], url))
}
