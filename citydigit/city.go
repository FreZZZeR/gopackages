package citydigit

import (
	"strings"

	"gopackages/wordz"
)

func City() string {
	wordz.Words = []string{"Beijing", "Brasilia", "Belgrad", "Moscow", "Mexico", "Nairobi", "Jakarta", "London", "Paris", "Cairo"}
	return strings.Split(wordz.Random(), ": ")[1]
}
