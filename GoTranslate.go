package main

import (
	"fmt"
	"os"

	trans "github.com/bas24/googletranslatefree"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("supply one more more words and a language to transalte to")
	}

	input := os.Args[1]
	destLang := os.Args[2]
	outfile := os.Args[3]

	result, err := trans.Translate(input, "auto", destLang)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	file.WriteString(result)
	fmt.Println("translation saved to file")
}
