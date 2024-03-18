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
			fmt.Fprintln(os.Stderr, "file does not exist: "+filename)
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
		fmt.Fprintln(os.Stderr, "Error parsing document:", err)
		os.Exit(-1)
	}
	return doc
}

// Search the given HTML tree for <head> and return it.
// Panic if it is not found in the expected place.
func LookupHead(root *html.Node) *html.Node {
	return LookupTopLevel(root, "head")
}

// Search the given HTML tree for <body> and return it.
// Panic if it is not found in the expected place.
func LookupBody(root *html.Node) *html.Node {
	return LookupTopLevel(root, "body")
}

// Search the given HTML tree for <node> and return it.
// Panic if it is not found in the expected place.
func LookupTopLevel(root *html.Node, node string) *html.Node {
	r := root.FirstChild
	if r.Data != "html" {
		panic("<html> tag not found")
	}
	for n := r.FirstChild; n != nil; n = n.NextSibling {
		if n.Data == node {
			return n
		}
	}
	panic("tag <body> not found")
}

// Looks for an <article> element in the HTML tree.
// Returns tree if an element was found.
func HasArticleElement(rootNode *html.Node) bool {
	body := LookupBody(rootNode)
	var lookupArticle func(*html.Node) bool
	lookupArticle = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "article" {
			return true
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if lookupArticle(c) {
				return true
			}
		}
		return false
	}
	return lookupArticle(body)
}

// Looks for the <title> and returns it.
// Panics if <title> not found.
func LookupTitle(rootNode *html.Node) string {
	head := LookupHead(rootNode)
	for c := head.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "title" {
			title := c.FirstChild.Data
			return title
		}
	}
	panic("title tag not found")
}

// Determines the siblings of the given node.
// Currently, only next siblings are considered.
func ListSiblingsOfNode(n *html.Node) []*html.Node {
	var l []*html.Node
	s := n.NextSibling
	for s != nil {
		l = append(l, s)
		s = s.NextSibling
	}
	// TODO add also previous siblings if any.
	return l
}

// Create the given directory if it does not exist.
// Panic if creation fails.
func CreateDirIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
