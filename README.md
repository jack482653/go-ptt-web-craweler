# go-ptt-web-craweler
## ptt 網路版爬蟲

idea by [ptt-web-crawler](https://github.com/ChinHui-Chen/ptt-web-crawler)

## 特色
* 使用 go 開發
* 多執行緒

## 用法

### import package
```Go
import "github.com/jack482653/pttCrawler/ptt"
```

### 抓取文章
```Go
a := &ptt.Article{}
// parse 文章
err := a.Parse("https://www.ptt.cc/bbs/Gossiping/M.1483256619.A.753.html")
// 印出來！
fmt.Println(a)
}
```

#### Result
```
"[問卦] 有沒有聯俄抗中的八卦"
作者: "F7mini158 (愛撫夭五八)", 日期: "Sun Jan  1 15:43:24 2017"
"當年冷戰時期，美國人季新吉\n\n跟尼克森大統領採行聯中抗蘇的伎\n\n倆，現在聯俄抗中的局勢逐漸成形\n\n請問各位30cmD罩杯，該如何看待\n\n未來的國際局勢，有沒有相關的八卦\n\n呢"
來源: "223.137.155.254"
推文數: 2, 噓文數: 1, 其他: 4
{"→" "kent" "都快八國聯軍了" "01/01 15:44"}
{"噓" "ILoveKMT" "欸欸 俄羅斯使館被驅逐了" "01/01 15:44"}
{"推" "Fongin" "阿扁聯俄抗中 goo.gl/vnLKQ5" "01/01 15:46"}
{"→" "inspire0201" "言之過早。季辛吉在DC建制派還有一大群徒子徒孫呢" "01/01 15:49"}
{"推" "bt9527" "表示甚麼? 表示中國國力已經反超俄羅斯了阿" "01/01 16:10"}
{"→" "bt9527" "接下來20年 老美沒有成功壓制或裂解中國 中國超美會成定局" "01/01 16:11"}
{"→" "shjyug" "妙禪有靈性佛，法輪功李洪志有法身,宋七力有分身,莊圓佛舞" "01/01 17:03"}
```

## 簡單測試

	$ make
	$ cd bin; ./pttCrawler
