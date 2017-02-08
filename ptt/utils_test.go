package ptt

import (
	"testing"

	"github.com/nbio/st"
)

func TestIsUrlValid(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal Url 1", args{"https://www.ptt.cc/bbs/Key_Mou_Pad/M.1486458543.A.4A9.html"}, true, false},
		{"Wrong Url 1", args{""}, false, true},
		{"Wrong Url 2", args{"123456"}, false, true},
		{"Wrong Url 3", args{"postgres://user:pass@host.com:5432/path?k=v#f"}, false, true},
		{"Wrong Url 4", args{"https://www.google.com"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsUrlValid(tt.args.url)
			if tt.wantErr {
				st.Refute(t, err, nil)
			} else {
				st.Assert(t, err, nil)
			}
			st.Assert(t, got, tt.want)
		})
	}
}
