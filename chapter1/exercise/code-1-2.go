package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + "Index : " + strconv.Itoa(index) + " Arg: " + arg
		sep = "\n"
	}
	fmt.Println(s)
}
