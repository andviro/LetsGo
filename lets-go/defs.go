package main

import (
	"fmt"
)

var s1 string // = ""

func main() {
	var s2 string = "Но зачем?.."
	s3, d := "Так лучше.", 42
	fmt.Printf("%q\n%q\n%q\n%d\n", s1, s2, s3, d)
}
