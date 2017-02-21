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
		{"Board not existed", args{"Gossiping2"}, 0, true},
		{"Board normal", args{"Gossiping"}, 20190, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/%s/index.html", PTT_BBS_ROOT, tt.args.board)
			resp_file_path := fmt.Sprintf("testcases/board/%s/index.htm", tt.args.board)
			defer gock.Off()
			gock.New(url).MatchHeader("Cookie", "over18=1").Reply(200).File(resp_file_path)
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

func TestGetLatestPageWithEmptyInput(t *testing.T) {
	got, err := GetLatestPage("")
	st.Refute(t, err, nil)
	st.Assert(t, got, 0)
}

func TestGetLatestPageWithServerError(t *testing.T) {
	defer gock.Off()
	url := fmt.Sprintf("%s/Gossiping/index.html", PTT_BBS_ROOT)
	gock.New(url).MatchHeader("Cookie", "over18=1").Reply(404)
	got, err := GetLatestPage("Gossiping")
	st.Refute(t, err, nil)
	st.Assert(t, got, 0)
}
