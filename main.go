package main

import (
	"bufio"
	"fmt"
	"github.com/AAAleg/wordcount/wordCount"
	"io"
	"log"
	"os"
)

func consoleReader() string {
	var reader *bufio.Reader

	reader = bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}

func main() {
	fmt.Println(wordCount.WordCount(consoleReader()))
}