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
	if doc_type := checkDocType(doc); doc_type != Normal {
		err_str := fmt.Sprintf(
			"Error: cannot fetch board %s: %s", board, doc_type,
		)
		return 0, errors.New(err_str)
	}
	href := ""
	ok := false
	// get div.btn-group-paging buttons
	doc.Find("div.btn-group-paging").Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if s.Text() == "‹ 上頁" {
			href, ok = s.Attr("href")
			return false
		}
		return true
	})
	if !ok {
		return 0, errors.New("Cannot get href attr of div.btn-group-paging")
	}
	filename := strings.Split(href[1:], "/")[2]
	len_f := len(filename)
	// extract digits in the middle of (index)[\d]+(.html)
	r, err := strconv.Atoi(filename[5 : len_f-5])
	if err != nil {
		return 0, err
	}
	return r + 1, nil
}
