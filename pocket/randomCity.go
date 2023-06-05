package pocket

import "../wordz"

func City() string {
	wordz.Words = []string{"NY", "LA", "Jersey", "Miami", "London"}
	wordz.Prefix = ""
	//println(wordz.Random())
	return wordz.Random()
}
