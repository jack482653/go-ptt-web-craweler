package ptt

import (
	"errors"
	"fmt"
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
	_, err := GetDocument(url)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
