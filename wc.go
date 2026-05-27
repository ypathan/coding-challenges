/*
testing command :  go build . && ./coding-challenges -l -w -c <<file_path>>
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func countLines(file_content string) {
	all_lines := strings.Split(file_content, "\n")
	fmt.Println("Lines: ", len(all_lines))
}

func countChars(file_content string) {
	var allchars = []rune(file_content)
	fmt.Println("Characters: ", len(allchars))
}

func countWords(file_content string) {
	allWords := strings.Fields(file_content)
	fmt.Println("Words: ", len(allWords))
}

func main() {

	line_ptr := flag.Bool("l", false, "Count lines in a file.")
	char_ptr := flag.Bool("c", false, "Count characters in a file.")
	words_ptr := flag.Bool("w", false, "Count words in a file")

	flag.Parse()

	file_path := flag.Args()[0]
	if len(file_path) == 0 {
		fmt.Println("no file path specified")
		return
	}

	file_bytes, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("Error reading the file")
	}
	file_content := string(file_bytes)
	file_content = strings.TrimSuffix(file_content, "\n")

	if *line_ptr {
		countLines(file_content)
	}

	if *char_ptr {
		countChars(file_content)
	}

	if *words_ptr {
		countWords(file_content)
	}

	fmt.Println("Completed Processing ", file_path)
}
