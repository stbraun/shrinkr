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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stbraun/shrinkr/util"
	"golang.org/x/net/html"
)

// shrinkCmd represents the shrink command
var shrinkCmd = &cobra.Command{
	Use:   "shrink",
	Short: "Looks for a block of text below the article and removes it.",
	Long: `The command checks for the existence of text below the article and removes it. 
In many cases references to other articles and other kind of overhead can be found here. 
These artifacts may consume much more memory and disk space than the article.  
Removing them can therefore shrink the size of the file quite a bit.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("filename missing")
			os.Exit(-1)
		}
		filename := args[0]
		fmt.Println("shrink called for " + filename)
		file := util.OpenFile(filename)
		defer func() { _ = file.Close() }()

		doc := util.ParseHTML(file)
		articleFound := shrinkDocument(doc)
		if !articleFound {
			fmt.Fprintln(os.Stderr, "No <article> tag found.")
		}
		ofile, err := os.Create("./testdata/output.html")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		err = html.Render(ofile, doc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		fmt.Printf("%+v\n", doc)

	},
}

func shrinkDocument(rootNode *html.Node) bool {
	body := util.LookupBody(rootNode)
	var result bool = false
	var nodesToBeRemoved []*html.Node
	var lookupArticle func(*html.Node) bool
	lookupArticle = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "article" {
			nodesToBeRemoved = listSiblings(n)
			return true
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if lookupArticle(c) {
				nodesToBeRemoved = append(nodesToBeRemoved, listSiblings(c)...)
				result = true
			} else {
				r := c
				nodesToBeRemoved = append(nodesToBeRemoved, r)
			}
		}
		for _, r := range nodesToBeRemoved {
			p := r.Parent
			fmt.Printf("parent: %+v\n", p)
			fmt.Printf("child: %+v\n", r)
			if p != nil {
				p.RemoveChild(r)
			}
		}
		return result
	}
	return lookupArticle(body)
}

func listSiblings(n *html.Node) []*html.Node {
	var l []*html.Node
	s := n.NextSibling
	for s != nil {
		l = append(l, s)
		s = s.NextSibling
	}
	return l
}

func init() {
	rootCmd.AddCommand(shrinkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shrinkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shrinkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
