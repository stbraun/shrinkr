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

	"github.com/spf13/cobra"
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
		fmt.Println("shrink called")
	},
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
