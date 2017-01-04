package ptt

type Board struct {
	name     string
	id       string
	articles []Article
}

func (b *Board) Parse(url string) error {
	return nil
}
