package stopwords

import (
	"bufio"
	"os"
)

const (
	baiduStopwords = "./baidu_stopwords.txt"
	cnStopwords    = "./cn_stopwords.txt"
	hitStopwords   = "./hit_stopwords.txt"
	scuStopwords   = "./scu_stopwords.txt"
)

var BaiduStopwords map[string]bool
var CNStopwords map[string]bool
var HITStopwords map[string]bool
var SCUStopwords map[string]bool

func readLines(filename string) map[string]bool {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(file)

	res := make(map[string]bool)
	for sc.Scan() {
		res[sc.Text()] = true
	}
	return res
}

func init() {
	BaiduStopwords = readLines(baiduStopwords)
	CNStopwords = readLines(cnStopwords)
	HITStopwords = readLines(hitStopwords)
	SCUStopwords = readLines(scuStopwords)
}
