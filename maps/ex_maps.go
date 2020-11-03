package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	out := map[string]int{}
	for _, word := range strings.Fields(s) {
		if num, ok := out[word]; ok {
			out[word] = num + 1
		} else {
			out[word] = 1
		}
	}
	return out
}

func main() {
	wc.Test(WordCount)
}
