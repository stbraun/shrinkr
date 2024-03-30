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
	"strings"

	"github.com/spf13/cobra"
	"github.com/stbraun/shrinkr/util"
	"golang.org/x/net/html"
)

var (
	outfileName      string
	outfilePath      string
	stats            *util.Stats
	doNotReportStats bool
)

// shrinkCmd represents the shrink command
var shrinkCmd = &cobra.Command{
	Use:   "shrink <filename or glob pattern>",
	Short: "Looks for a block of text below the article and removes it.",
	Long: `The command checks for the existence of text below the article and removes it. 
In many cases references to other articles and other kind of overhead can be found here. 
These artifacts may consume much more memory and disk space than the article.  
Removing them can therefore shrink the size of the file quite a bit.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stats = util.NewStats()
		stats.Start()
		files, err := filepath.Glob(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if Verbose {
			listFilesToProcess(files)
		}
		for _, filename := range files {
			processFile(filename)
		}
		stats.Stop()
		if !doNotReportStats {
			reportStatistics(*stats)
		}
	},
}

func listFilesToProcess(files []string) {
	fmt.Printf("\n----------------\n%d files to process\n----------------\n", len(files))
	for _, fn := range files {
		fmt.Printf("\t%s\n", fn)
	}
	fmt.Println("----------------")
}

func reportStatistics(stats util.Stats) {
	fmt.Printf("\n----------\nStatistics\n----------\n")
	fmt.Printf("%d articles were processed in %dms\nreducing the cumulated size by %s from %s to %s\n",
		stats.Count(),
		stats.ElapsedTime(),
		util.FormatFileSize(stats.SizeReducedBy()),
		util.FormatFileSize(stats.CumulatedSizesOfOriginalFiles()),
		util.FormatFileSize(stats.CumulatedSizesOfShrinkedFiles()))
	fmt.Println("----------")
}

func processFile(filename string) error {
	fmt.Printf("shrinking %s...\n", filename)
	file := util.OpenFile(filename)
	defer func() { _ = file.Close() }()

	doc, err := html.Parse(file)
	if err != nil {
		return err
	}
	if !util.HasArticleElement(doc) {
		return fmt.Errorf("no <article> element in %s", filename)
	}
	title := util.LookupTitle(doc)

	ofile, ofileName, err := createOutputFile(outfilePath, outfileName, title)
	if err != nil {
		return err
	}

	shrinkDocument(doc)
	err = html.Render(ofile, doc)
	if err != nil {
		return err
	}
	stats.AddSizes(util.GetFileSize(filename), util.GetFileSize(ofileName))
	return nil
}

func createOutputFile(outPath, outName, title string) (*os.File, string, error) {
	var ofileName string
	util.CreateDirIfNotExist(outPath)
	if len(outName) > 0 {
		ofileName = filepath.Join(outfilePath, outName)
	} else {
		shortenedTitle := shortenTitle(title)
		ofileName = filepath.Join(outfilePath, shortenedTitle+".html")
	}
	fmt.Printf("writing %s...\n", ofileName)
	ofile, err := os.Create(ofileName)
	if err != nil {
		return nil, "", err
	}
	return ofile, ofileName, nil
}

// Cut off trailing meta data and spurious sentences from given title.
func shortenTitle(title string) string {
	shortenedTitle, _, _ := strings.Cut(title, " |")        // cut off meta data
	shortenedTitle, _, _ = strings.Cut(shortenedTitle, ".") // cut off following sentence(s)
	return shortenedTitle
}

func shrinkDocument(rootNode *html.Node) {
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
			if p != nil {
				p.RemoveChild(r)
			}
		}
		return result
	}
	lookupArticle(body)
}

func init() {
	rootCmd.AddCommand(shrinkCmd)

	shrinkCmd.PersistentFlags().StringVar(&outfileName, "outfile", "", "The name of the output file.")
	shrinkCmd.PersistentFlags().StringVar(&outfilePath, "outpath", "./", "The path where the output file shall be written.")
	shrinkCmd.PersistentFlags().BoolVar(&doNotReportStats, "nostats", false, "Suppress reporting of statistics.")
}
