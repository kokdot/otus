package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

var s = "Hello, OTUS!"

func main() {
	s = stringutil.Reverse(s)
	fmt.Println(s)
}
