package citydigit

import (
	"strings"

	"gopackages/wordz"
)

func Digit() string {
	wordz.Words = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Zero"}
	return strings.Split(wordz.Random(), ": ")[1]
}
