package stopwords

import (
	_ "embed"
	"strings"
)

//go:embed data/baidu_stopwords.txt
var baiduStopwords string

//go:embed data/cn_stopwords.txt
var cnStopwords string

//go:embed data/hit_stopwords.txt
var hitStopwords string

//go:embed data/scu_stopwords.txt
var scuStopwords string

var BaiduStopwords map[string]bool
var CNStopwords map[string]bool
var HITStopwords map[string]bool
var SCUStopwords map[string]bool

func readLines(data string) map[string]bool {
	res := make(map[string]bool)
	content := strings.Split(data, "\n")
	for i := range content {
		res[content[i]] = true
	}
	return res
}

func init() {
	BaiduStopwords = readLines(baiduStopwords)
	CNStopwords = readLines(cnStopwords)
	HITStopwords = readLines(hitStopwords)
	SCUStopwords = readLines(scuStopwords)
}
