package ptt

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
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
		{"Normal Url 1", args{PTT_BBS_ROOT + "/Key_Mou_Pad/M.1486458543.A.4A9.html"}, true, false},
		{"Wrong Url 1", args{""}, false, true},
		{"Wrong Url 2", args{"123456"}, false, true},
		{"Wrong Url 3", args{"!$#&^%&^&^%%"}, false, true},
		{"Wrong Url 4", args{"postgres://user:pass@host.com:5432/path?k=v#f"}, false, true},
		{"Wrong Url 5", args{"https://www.google.com"}, false, true},
		{"Wrong Url 6", args{PTT_BBS_ROOT + "/Key_Mou_Pad/fake/M.1486458543.A.4A9.html"}, false, true},
		{"Wrong Url 7", args{PTT_ROOT + "/other_function/Key_Mou_Pad/M.1486458543.A.4A9.html"}, false, true},
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

func Test_checkDocType(t *testing.T) {
	type args struct {
		doc_path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"404", args{"testcases/board/400/index.htm"}, NotFound},
		{"500", args{"testcases/board/500/index.htm"}, IntenalErr},
		{"Nornael article", args{"testcases/article/Gossiping/M.1483256619.A.753.htm"}, Normal},
		{"Nornael board", args{"testcases/board/Gossiping/index.htm"}, Normal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.args.doc_path)
			st.Assert(t, err, nil)
			defer file.Close()
			doc, err := goquery.NewDocumentFromReader(file)
			st.Assert(t, err, nil)
			got := checkDocType(doc)
			st.Assert(t, got, tt.want)
		})
	}
}
