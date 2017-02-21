package ptt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Board struct {
	name     string
	id       string
	articles []Article
}

func NewBoard(url string) error {
	return nil
}

func GetLatestPage(board string) (int, error) {
	if board == "" {
		return 0, errors.New("Board name is empty")
	}
	url := fmt.Sprintf("%s/%s/index.html", PTT_BBS_ROOT, board)
	// check url is valid
	if r, err := IsUrlValid(url); r != true {
		return 0, errors.New(fmt.Sprintf("Error: url %s invalid: %v", url, err))
	}
	// get document of latest page of board
	doc, err := GetDocument(url)
	if err != nil {
		return 0, err
	}
	// check if wrong board name causing 404 not found
	not_found := doc.Find("div.bbs-content").Text()
	if not_found == "404 - Not Found." {
		return 0, errors.New(fmt.Sprintf("Error: url %s not found", url))
	}
	href := ""
	prev_find := false
	ok := false
	// get div.btn-group-paging buttons
	doc.Find("div.btn-group-paging").Find("a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "‹ 上頁" {
			prev_find = true
			href, ok = s.Attr("href")
		}
	})
	if !prev_find {
		return 0, errors.New("Prev page button is not found")
	}
	if ok {
		paths := strings.Split(href[1:], "/")
		filename := paths[2]
		len_f := len(filename)
		// extract digits in the middle of (index)[\d]+(.html)
		r, err := strconv.Atoi(filename[5 : len_f-5])
		if err != nil {
			return 0, err
		}
		return r + 1, nil
	} else {
		return 0, errors.New("Cannot get href attr of div.btn-group-paging")
	}
}
