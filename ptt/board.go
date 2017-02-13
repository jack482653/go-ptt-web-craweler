package ptt

import (
	"errors"
)

type Board struct {
	name     string
	id       string
	articles []Article
}

func NewBoard(url string) error {
	return nil
}

func GetLastPage(board string) (int, error) {
	if board == "" {
		return 0, errors.New("Board name is empty")
	}
	return 0, nil
}
