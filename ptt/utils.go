package ptt

import (
	"errors"
	"net/http"
	urlpkg "net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func IsUrlValid(url string) (bool, error) {
	u, err := urlpkg.Parse(url)
	// check if url parse is success
	if err != nil {
		return false, err
	}
	// check protocol
	if u.Scheme != "http" && u.Scheme != "https" {
		err := errors.New("Input url is not http or https protocol")
		return false, err
	}
	// check hostname
	if u.Host != PTT_HOST {
		err := errors.New("Hostname is not " + PTT_HOST)
		return false, err
	}
	// check path
	paths := strings.Split(u.Path[1:], "/")
	if len(paths) != 3 {
		err := errors.New("Path structure error")
		return false, err
	}
	if paths[0] != "bbs" {
		err := errors.New("Path is not started with bbs")
		return false, err
	}
	return true, nil
}

func GetDocument(url string) (*goquery.Document, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "over18=1")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	return doc, err
}
