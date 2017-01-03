package pttCrawler

type Crawler interface {
	Parse(url string) error
}
