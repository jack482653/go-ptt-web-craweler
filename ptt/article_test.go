package ptt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestNewArticle(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *Article
		wantErr bool
	}{
	// TODO: Add test cases.
		{"Test Article Normal", args{"Gossiping_M.1483256619.A.753"}, nil, false},
		{"Test Invalid url", args{"404"}, nil, true},
	}
	for i := range tests {
		if tests[i].wantErr == false {
			input := fmt.Sprintf("testcases/%s.json", tests[i].args.name)
			bytes, err := ioutil.ReadFile(input)
			st.Assert(t, err, nil)
			tests[i].want = &Article{}
			json.Unmarshal(bytes, tests[i].want)
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("https://www.ptt.cc/bbs/Gossiping/%s.html", tt.args.name)
			resp_file_path := fmt.Sprintf("testcases/%s.htm", tt.args.name)
			defer gock.Off()
			gock.New(url).MatchHeader("Cookie", "over18=1").Reply(200).File(resp_file_path)
			got, err := NewArticle(url)
			st.Assert(t, err, nil)
			st.Assert(t, got, tt.want)
		})
	}
}
