package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func process_val(data string) any {
	if strings.Contains(data, "\"") {
		data = strings.ReplaceAll(data, "\"", "")
		return data
	}

	if strings.Contains(data, "true") {
		return true
	}

	if strings.Contains(data, "false") {
		return false
	}

	if strings.ContainsAny(data, "0123456789") {
		data = strings.TrimSpace(data)
		data_int, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println("Error parsing integer")
		}
		return data_int
	}

	return nil
}

func main() {

	filename_ptr := flag.String("f", "", "parse a file")
	flag.Parse()

	byte_data, err := os.ReadFile(*filename_ptr)
	if err != nil {
		fmt.Println("Error reading the file")
	}

	file_content := string(byte_data)
	allLines := strings.Split(file_content, "\n")

	//trim prefix and suffix of '{' and '}'
	allLines = allLines[1 : len(allLines)-2]

	kv_map := make(map[string]any)

	for _, line := range allLines {

		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, ",", "")

		kv := strings.Split(line, ":")

		k := kv[0]
		v := kv[1]

		kv_map[k] = process_val(v)
	}

	for k, v := range kv_map {
		fmt.Printf("%s : %v \n", k, v)
	}

}
