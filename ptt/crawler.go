package ptt

type Crawler interface {
	Parse(url string) error
}
