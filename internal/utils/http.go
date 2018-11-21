package utils

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"gitee.com/johng/gf/g/os/gproc"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kooksee/ktask/internal/errs"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func HttpPost(url string, data []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	MustNotError(err)

	// 处理header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		dt, err := ioutil.ReadAll(resp.Body)
		MustNotError(err)
		return nil, fmt.Errorf("url: %s post error,input: %s, output: %s", url, data, dt)
	}

	dt, err := ioutil.ReadAll(resp.Body)
	MustNotError(err)

	return dt, nil
}

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

	dt, err := gproc.ShellExec(fmt.Sprintf(`wget --header="User-Agent: Mozilla/5.0 (Windows NT 6.0) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11" --timeout=10 -O %s -c %s`, fileName, url))
	fmt.Println(err, "ok", string(dt))
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "The file is already fully retrieved; nothing to do.") {
		return nil
	}

	if strings.Contains(err.Error(), "ERROR 404: Not Found") || strings.Contains(err.Error(), "ERROR 403: Forbidden") {
		return errs.NotFound
	}

	return err
}

func Http2File(url string, headers map[string]string, fileDir string) error {
	req, err := http.NewRequest("GET", url, nil)
	MustNotError(err)

	// 处理header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		log.Error().Str("url", url).Int("status_code", resp.StatusCode).Msg("nou found")
		return errs.NotFound
	}

	if resp.StatusCode != http.StatusOK {
		dt, err := ioutil.ReadAll(resp.Body)
		MustNotError(err)
		return fmt.Errorf("url: %s get error, output: %s", url, dt)
	}

	fileType := ""
	ct := strings.Split(resp.Header.Get("Content-Type"), "/")
	if len(ct) == 0 || ct[len(ct)-1] == "" {
		fileType = GetUrlType(url)
	} else {
		fileType = ct[len(ct)-1]
	}

	fileName := filepath.Join(fileDir, hex.EncodeToString(Sha256([]byte(url)))) + "." + fileType
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	r := bufio.NewReader(resp.Body)
	_, err = r.WriteTo(f)
	if err != nil {
		return err
	}

	return f.Close()
}

func Http2OSS(url string, headers map[string]string, bkt *oss.Bucket) error {
	req, err := http.NewRequest("GET", url, nil)
	MustNotError(err)

	// 处理header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		log.Error().Str("url", url).Int("status_code", resp.StatusCode).Msg("nou found")
		return errs.NotFound
	}

	if resp.StatusCode != http.StatusOK {
		dt, err := ioutil.ReadAll(resp.Body)
		MustNotError(err)
		return fmt.Errorf("url: %s get error, output: %s", url, dt)
	}

	fileType := ""
	ct := strings.Split(resp.Header.Get("Content-Type"), "/")
	if len(ct) == 0 || ct[len(ct)-1] == "" {
		fileType = GetUrlType(url)
	} else {
		fileType = ct[len(ct)-1]
	}

	fileName := filepath.Join(resp.Request.URL.Host, hex.EncodeToString(Sha256([]byte(url)))) + "." + fileType
	fmt.Println(fileName)
	return bkt.PutObject(fileName, resp.Body)
}
