package ptt

import (
	"fmt"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestGetLastPage(t *testing.T) {
	type args struct {
		board string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Nil input", args{""}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLastPage(tt.args.board)
			if tt.wantErr {
				st.Refute(t, err, nil)
			} else {
				st.Assert(t, err, nil)
			}
			st.Assert(t, got, tt.want)
		})
	}
}

func TestGetLastPageWithServerError(t *testing.T) {
	defer gock.Off()
	gock.New("https://www.ptt.cc/bbs/Gossiping/index.html").MatchHeader("Cookie", "over18=1").Reply(404)
	got, err := GetLastPage("Gossiping")
	fmt.Println(err)
	st.Refute(t, err, nil)
	st.Assert(t, got, 0)
}
