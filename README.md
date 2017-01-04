# go-ptt-web-craweler
## ptt 網路版爬蟲

idea by [ptt-web-crawler](https://github.com/ChinHui-Chen/ptt-web-crawler)

## 特色
* 使用 go 開發
* 多執行緒

## 用法
### 抓取文章
```go=
import (
	"fmt"
	"github.com/jack482653/pttCrawler/ptt"
)

func main() {
	a := &ptt.Article{}
	// parse article
	err := a.Parse("https://www.ptt.cc/bbs/Gossiping/M.1483256619.A.753.html")
	// print parse result
	fmt.Println(a)
}
```
