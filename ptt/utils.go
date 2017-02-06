package ptt

import (
	urlpkg "net/url"
	"strings"
)

func IsUrlValid(url string) bool {
	u, err := urlpkg.Parse(url)
	// check if url parse is success
	if err != nil {
		return false
	}
	// check protocal
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	// check hostname
	if u.Host != "www.ptt.cc" {
		return false
	}
	// check path
	paths := strings.Split(u.Path[1:], "/")
	if len(paths) != 3 {
		return false
	}
	if paths[0] != "bbs" {
		return false
	}
	return true
}
