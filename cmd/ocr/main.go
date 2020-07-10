package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"os"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	fmt.Println("start to ocr from image file", os.Args[1])
	client.SetImage(os.Args[1])
	text, err := client.Text()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ocr successfully, the text is as follows:[[[")
		fmt.Println(text)
		fmt.Println("]]]")

	}
	// Hello, World!
}
