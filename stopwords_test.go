package stopwords

import (
	"testing"
)

func Test_init(t *testing.T) {
	allStopwords := map[string]map[string]bool{
		"baidu": BaiduStopwords,
		"cn":    CNStopwords,
		"hit":   HITStopwords,
		"scu":   SCUStopwords,
	}
	for name, stw := range allStopwords {
		if len(stw) == 0 {
			t.Errorf("stopwords is empty: %s", name)
		}
	}
}
