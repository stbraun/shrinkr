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
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/stbraun/shrinkr/util"
	"golang.org/x/net/html"
)

var (
	outfileName *string
	outfilePath *string
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
		var outFile string

		if len(args) != 1 {
			fmt.Println("filename missing")
			os.Exit(-1)
		}
		filename := args[0]
		fmt.Println("shrink called for " + filename)
		file := util.OpenFile(filename)
		defer func() { _ = file.Close() }()

		doc := util.ParseHTML(file)
		if !util.HasArticleElement(doc) {
			fmt.Fprintln(os.Stderr, "No <article> element in ", filename)
			os.Exit(-1)
		}
		title := util.LookupTitle(doc)
		fmt.Printf("Title: %+v\n", title)
		shrinkDocument(doc)

		if _, err := os.Stat(*outfilePath); os.IsNotExist(err) {
			err := os.Mkdir(*outfilePath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		if *outfileName != "" {
			outFile = filepath.Join(*outfilePath, *outfileName)
		} else {
			outFile = filepath.Join(*outfilePath, title+".html")
		}
		if Verbose {
			fmt.Println("Creating output file: ", outFile)
		}
		ofile, err := os.Create(outFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		err = html.Render(ofile, doc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	},
}

func shrinkDocument(rootNode *html.Node) bool {
	body := util.LookupBody(rootNode)
	var result bool = false
	var nodesToBeRemoved []*html.Node
	var lookupArticle func(*html.Node) bool
	lookupArticle = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "article" {
			nodesToBeRemoved = util.ListSiblingsOfNode(n)
			return true
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if lookupArticle(c) {
				nodesToBeRemoved = append(nodesToBeRemoved, util.ListSiblingsOfNode(c)...)
				result = true
			} else {
				r := c
				nodesToBeRemoved = append(nodesToBeRemoved, r)
			}
		}
		for _, r := range nodesToBeRemoved {
			p := r.Parent
			if Verbose {
				fmt.Printf("parent: %+v\n", p)
				fmt.Printf("child: %+v\n", r)
			}
			if p != nil {
				p.RemoveChild(r)
			}
		}
		return result
	}
	return lookupArticle(body)
}

func init() {
	rootCmd.AddCommand(shrinkCmd)

	outfileName = shrinkCmd.PersistentFlags().String("outfile", "", "The name of the output file.")
	outfilePath = shrinkCmd.PersistentFlags().String("outpath", "./", "The path where the output file shall be written.")
}
