package main

import (
	"fmt"

	"github.com/jack482653/pttCrawler/ptt"
)

func main() {
	a, err := ptt.NewArticle("https://www.ptt.cc/bbs/Gossiping/M.1483256619.A.753.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
