package citation

import (
	"math/rand"
)

func getRandomQuote(splitedText []string) string {
	citation := splitedText[rand.Int()%len(splitedText)]

	if citation == "" {
		return getRandomQuote(splitedText)
	}

	return citation
}
