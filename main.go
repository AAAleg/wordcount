package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./wordcount TEXT")
		return
	}
	fmt.Println(len(strings.Fields(os.Args[1])))
	// commit for actions
}