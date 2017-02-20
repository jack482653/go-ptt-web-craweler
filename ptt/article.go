package ptt

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Comment struct {
	PushTag        string
	PushUserID     string
	PushContent    string
	PushIpdatetime string
}

type Article struct {
	Title               string
	Author              string
	Date                string
	Content             string
	Ip                  string
	Comments            []Comment
	All, Count, P, B, N int
}

func NewArticle(url string) (*Article, error) {
	if r, err := IsUrlValid(url); r != true {
		return nil, errors.New(fmt.Sprintf("Error: url %s invalid: %v", url, err))
	}
	doc, err := GetDocument(url)
	if err != nil {
		return nil, err
	}
	a := &Article{}
	// check if 404 not found
	not_found := doc.Find("div.bbs-content").Text()
	if not_found == "404 - Not Found." {
		return nil, errors.New(fmt.Sprintf("Error: url %s not found", url))
	}
	// get selector of main content
	main_content := doc.Find("div#main-content")
	// get selector of article metaline
	main_content.Find("div.article-metaline").Each(func(i int, s *goquery.Selection) {
		k := s.Find("span.article-meta-tag").Text()
		v := s.Find("span.article-meta-value").Text()
		switch k {
		case "作者":
			a.Author = v
		case "標題":
			a.Title = v
		case "時間":
			a.Date = v
		}
		// remove article metaline
		s.Remove()
	})
	// remove remain article metaline
	main_content.Find("div.article-metaline-right").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})
	// get selector of pushes
	pushes := main_content.Find("div.push")
	a.Comments = make([]Comment, pushes.Size())
	pushes.Each(func(i int, push *goquery.Selection) {
		push_tag := strings.Trim(push.Find("span.push-tag").Text(), " \t\n\r")
		push_user_id := strings.Trim(push.Find("span.push-userid").Text(), " \t\n\r")
		push_content := strings.Trim(push.Find("span.push-content").Text(), ": \t\n\r")
		push_ipdatetime := strings.Trim(push.Find("span.push-ipdatetime").Text(), " \t\n\r")
		switch push_tag {
		case "推":
			a.P += 1
		case "噓":
			a.B += 1
		default:
			a.N += 1
		}
		a.Comments[i] = Comment{push_tag, push_user_id, push_content, push_ipdatetime}
		push.Remove()
	})
	// count: 推噓文相抵後的數量; all: 推文總數
	a.All = a.P + a.B + a.N
	a.Count = a.P - a.B
	// get ip
	html, err := main_content.Html()
	if err != nil {
		return nil, err
	}
	r, err := regexp.Compile("(※ 發信站: ).*")
	if err != nil {
		return nil, err
	}
	ip := r.FindString(html)
	r, err = regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+")
	if err != nil {
		return nil, err
	}
	a.Ip = r.FindString(ip)
	// remove redundant f2 class and remain text of others class
	main_content.Find("*").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "※ 發信站:") || strings.Contains(text, "※ 文章網址:") || strings.Contains(text, "※ 編輯:") {
			s.Remove()
		} else {
			s.ReplaceWithHtml(text)
		}
	})
	content, err := main_content.Html()
	if err != nil {
		return nil, err
	}
	a.Content = strings.Trim(content, "-\t\n\r")
	return a, nil
}

func (c *Comment) String() string {
	return fmt.Sprintf("%q %q: %q\t%q", c.PushTag, c.PushUserID, c.PushContent, c.PushIpdatetime)
}

func (a *Article) String() string {
	var buffer bytes.Buffer
	meta := fmt.Sprintf("%q\n作者: %q, 日期: %q\n", a.Title, a.Author, a.Date)
	buffer.WriteString(meta)
	content := fmt.Sprintf("%q\n來源: %q\n", a.Content, a.Ip)
	buffer.WriteString(content)
	push_info := fmt.Sprintf("推文數: %v, 噓文數: %v, 其他: %v\n", a.P, a.B, a.N)
	buffer.WriteString(push_info)
	for _, c := range a.Comments {
		buffer.WriteString(fmt.Sprintf("%q\n", c))
	}
	return buffer.String()
}
