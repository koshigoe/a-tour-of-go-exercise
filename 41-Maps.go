package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _, c := range strings.Fields(s) {
		ret[c]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
