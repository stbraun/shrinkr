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
	"github.com/spf13/viper"
	"github.com/stbraun/shrinkr/util"
)

// existsCmd represents the exists command
var existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "Looks for an element of type article in a given document.",
	Long: `The command checks for the existence of an element of type article in a given HTML document.
It can be run on documents to decide whether shrinking them may work.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "filename missing")
			os.Exit(-1)
		}
		filename := args[0]
		if viper.GetBool("verbose") {
			fmt.Println("exists called for " + filename)
		}
		file := util.OpenFile(filename)
		defer func() { _ = file.Close() }()

		doc := util.ParseHTML(file)
		result := util.HasArticleElement(doc)
		if result {
			fmt.Println("Document contains an <article> element.")
			os.Exit(0)
		} else {
			fmt.Println("Document does not contain an <article> element.")
			os.Exit(1)
		}
	},
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
