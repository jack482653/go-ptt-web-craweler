package ptt

import (
	"fmt"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestGetLatestPage(t *testing.T) {
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
			got, err := GetLatestPage(tt.args.board)
			if tt.wantErr {
				st.Refute(t, err, nil)
			} else {
				st.Assert(t, err, nil)
			}
			st.Assert(t, got, tt.want)
		})
	}
}

func TestGetLatestPageWithServerError(t *testing.T) {
	defer gock.Off()
	url := fmt.Sprintf("%s/bbs/Gossiping/index.html", PTT_BBS_ROOT)
	gock.New(url).MatchHeader("Cookie", "over18=1").Reply(404)
	got, err := GetLatestPage("Gossiping")
	fmt.Println(err)
	st.Refute(t, err, nil)
	st.Assert(t, got, 0)
}
