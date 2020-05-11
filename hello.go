package main

import (
	"fmt"

	"github.com/yokall/go-hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello world")

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
