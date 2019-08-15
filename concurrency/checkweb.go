package checkweb

type WebChecker func(string) bool

func CheckWebsite(wc WebChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
