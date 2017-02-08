package ptt

import (
	"errors"
	urlpkg "net/url"
	"strings"
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
	if u.Host != "www.ptt.cc" {
		err := errors.New("Hostname is not www.ptt.cc")
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
