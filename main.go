package main

import (
	"flag"
	"fmt"
	"os"
)


func main() {
	flag.Parse()
	fileName := flag.Args()[0]
	contents := readFile(fileName)
	fmt.Println(contents)
}

func readFile(fileName string) string {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	result := ""
	buf := make([]byte, 1)
	for {
		n, err := fp.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		result += string(buf)
	}

	return result
}
