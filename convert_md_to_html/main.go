package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gomarkdown/markdown"
)


func main() {

	file := "markdown.md"

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	html := markdown.ToHTML(content, nil, nil)

	
	fileOut := "htmlFromMarkdown.html"

	writeErr := ioutil.WriteFile(fileOut, html, 0644)

	if writeErr != nil {
		log.Fatalf("Could not write to %s", fileOut)
	}

	fmt.Printf("HTML created to %s", fileOut)


}