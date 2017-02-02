all: get-deps build

build:
	go build -o bin/pttCrawler

get-deps:
	go get github.com/PuerkitoBio/goquery
	go get github.com/nbio/st
	go get gopkg.in/h2non/gock.v1
