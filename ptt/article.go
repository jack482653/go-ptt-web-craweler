package ptt

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Comment struct {
	pushTag        string
	pushUserID     string
	pushContent    string
	pushIpdatetime string
}

type Article struct {
	id                  string
	title               string
	auther              string
	date                string
	content             string
	ip                  string
	comments            []Comment
	all, count, p, b, n int
}

func (a *Article) Parse(url string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", "over18=1")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return err
	}
	// get selector of main content
	main_content := doc.Find("div#main-content")
	// get selector of article metaline
	main_content.Find("div.article-metaline").Each(func(i int, s *goquery.Selection) {
		k := s.Find("span.article-meta-tag").Text()
		v := s.Find("span.article-meta-value").Text()
		switch k {
		case "作者":
			a.auther = v
		case "標題":
			a.title = v
		case "時間":
			a.date = v
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
	a.comments = make([]Comment, pushes.Size())
	pushes.Each(func(i int, push *goquery.Selection) {
		push_tag := strings.Trim(push.Find("span.push-tag").Text(), " \t\n\r")
		push_user_id := strings.Trim(push.Find("span.push-userid").Text(), " \t\n\r")
		push_content := strings.Trim(push.Find("span.push-content").Text(), ": \t\n\r")
		push_ipdatetime := strings.Trim(push.Find("span.push-ipdatetime").Text(), " \t\n\r")
		switch push_tag {
		case "推":
			a.p += 1
		case "噓":
			a.b += 1
		default:
			a.n += 1
		}
		a.comments[i] = Comment{push_tag, push_user_id, push_content, push_ipdatetime}
		push.Remove()
	})
	// count: 推噓文相抵後的數量; all: 推文總數
	a.all = a.p + a.b + a.n
	a.count = a.p - a.b
	// get ip
	html, err := main_content.Html()
	if err != nil {
		return err
	}
	r, err := regexp.Compile("(※ 發信站: ).*")
	if err != nil {
		return err
	}
	ip := r.FindString(html)
	r, err = regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+")
	if err != nil {
		return err
	}
	a.ip = r.FindString(ip)
	// remove class f2
	main_content.Find("span.f2").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})
	content, err := main_content.Html()
	if err != nil {
		return nil
	}
	a.content = strings.Trim(content, "- \t\n\r")
	return nil
}

func (c *Comment) String() string {
	return fmt.Sprintf("%q %q: %q\t%q", c.pushTag, c.pushUserID, c.pushContent, c.pushIpdatetime)
}

func (a *Article) String() string {
	var buffer bytes.Buffer
	meta := fmt.Sprintf("%q\n作者: %q, 日期: %q\n", a.title, a.auther, a.date)
	buffer.WriteString(meta)
	content := fmt.Sprintf("%q\n來源: %q\n", a.content, a.ip)
	buffer.WriteString(content)
	push_info := fmt.Sprintf("推文數: %v, 噓文數: %v, 其他: %v\n", a.p, a.b, a.n)
	buffer.WriteString(push_info)
	for _, c := range a.comments {
		buffer.WriteString(fmt.Sprintf("%q\n", c))
	}
	return buffer.String()
}
