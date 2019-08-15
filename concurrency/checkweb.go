package checkweb

//WebChecker is a function type
type WebChecker func(string) bool

type result struct {
	url string
	tf  bool
}

/*
type result struct {
	string
	bool
}
*/

//CheckWebsite write judgement (by MockWebchecker) to map
func CheckWebsite(wc WebChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.tf
		//results[result.string] = result.bool
	}

	return results
}
