package utils

import (
	url2 "net/url"
	"strings"
)

func UrlCheck(u1, url string) string {
	us, _ := url2.Parse(u1)
	u := us.Scheme + "://" + us.Host

	url = strings.TrimSpace(url)
	url = strings.Replace(url, " ", "", -1)

	if url == "" {
		return ""
	}

	if strings.HasPrefix(url, "data") {
		return ""
	}

	if strings.HasPrefix(url, "//") {
		url = us.Scheme + "://" + strings.Trim(url, "//")
	}

	if strings.HasPrefix(url, "..") {
		u2 := strings.Split(u1, "/")
		u1 = strings.Join(u2[:len(u2)-2], "/")
		url = u1 + strings.Trim(url, "..")
	}

	if strings.HasPrefix(url, ".") {
		u2 := strings.Split(u1, "/")
		u1 = strings.Join(u2[:len(u2)-1], "/")
		url = u1 + strings.Trim(url, ".")
	}

	if !strings.HasPrefix(url, "http") {
		if strings.HasPrefix(url, "/") {
			url = u + url
		} else {
			u2 := strings.Split(u1, "/")
			u1 = strings.Join(u2[:len(u2)-1], "/")
			url = u1 + url
		}
	}
	return url

}

func GetUrlType(url string) string {
	dt := strings.Split(url, "?")
	dt1 := strings.Split(dt[0], ".")
	return dt1[len(dt1)-1]
}
