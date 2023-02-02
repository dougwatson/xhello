package main

import (
	"fmt"
	"os"
	"github.com/dougwatson/xhello/morestrings"
	"github.com/dougwatson/xbar"
)
func main() {
	word:="Hello,World"
	if len(os.Args)>1{
		word=os.Args[1]
	}
	fmt.Println(morestrings.ReverseRunes(word))
	println("import pkg:",xbar.Hello())
}
