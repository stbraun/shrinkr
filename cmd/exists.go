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
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

// existsCmd represents the exists command
var existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "Looks for an element of type article in a given document.",
	Long: `The command checks for the existence of an element of type article in a given HTML document.
It can be run on documents to decide whether shrinking them may work.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("filename missing")
			os.Exit(1)
		}
		filename := args[0]
		fmt.Println("exists called for " + filename)
		file, err := os.Open(filename)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				// File does not exist
				fmt.Println("file does not exist: " + filename)
				os.Exit(2)
			} else {
				panic(err)
			}
		}
		defer func() { _ = file.Close() }()

		// TODO
		doc, err := html.Parse(file)
		if err != nil {
			fmt.Println("Error parsing document:", err)
			os.Exit(3)
		}
		result := hasArticleElement(doc)
		fmt.Printf("Document has Article element: %v\n", result)
	},
}

func hasArticleElement(rootNode *html.Node) bool {
	body := lookupBody(rootNode)
	var lookupArticle func(*html.Node) bool
	lookupArticle = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "article" {
			// fmt.Printf("%+v\n", n)
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

func lookupBody(root *html.Node) *html.Node {
	html_ := root.FirstChild
	body := html_.FirstChild.NextSibling
	if body.Data != "body" {
		panic("node body not found")
	}
	return body
}

func init() {
	rootCmd.AddCommand(existsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// existsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// existsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
