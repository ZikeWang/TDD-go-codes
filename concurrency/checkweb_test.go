package checkweb

import (
	"reflect"
	"testing"
	"time"
)

func MockWebChecker(url string) bool {
	if url == "ss://baidu.com" {
		return false
	}
	return true
}

func TestWebchecker(t *testing.T) {
	urls := []string{
		"https://tencent.com",
		"https://alibaba.com",
		"ss://baidu.com",
	}

	want := map[string]bool{
		"https://tencent.com": true,
		"https://alibaba.com": true,
		"ss://baidu.com":      false,
	}

	got := CheckWebsite(MockWebChecker, urls)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func slowTimerSimulator(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebchecker(b *testing.B) {
	urls := make([]string, 200)

	for i := 0; i < len(urls); i++ {
		urls[i] = "an url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsite(slowTimerSimulator, urls)
	}
}
