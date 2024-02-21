/*
Copyright © 2024 Stefan Braun sb@action.ms

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package util

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Open a file by given filename.
// Exit or panic in case of error.
func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// File does not exist
			fmt.Println("file does not exist: " + filename)
			os.Exit(-1)
		} else {
			panic(err)
		}
	}
	return file
}

// Parse a given file into a HTML tree.
// Exit program in case of a parse error.
func ParseHTML(file *os.File) *html.Node {
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println("Error parsing document:", err)
		os.Exit(-1)
	}
	return doc
}

// Search the given HTML tree for <body> and return it.
// Panic if it is not found in the expected place.
func LookupBody(root *html.Node) *html.Node {
	r := root.FirstChild
	if r.Data != "html" {
		panic("<html> tag not found")
	}
	for n := r.FirstChild; n != nil; n = n.NextSibling {
		if n.Data == "body" {
			return n
		}
	}
	panic("tag <body> not found")
}
