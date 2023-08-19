/*
Copyright © 2023 rebelice <yangrebelice@gmail.com>

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
	"github.com/rebelice/agt/toolkit/differ"
	"github.com/spf13/cobra"
)

var (
	flags struct {
		grammarFile string
		target      string
	}
)

// differCmd represents the differ command
var differCmd = &cobra.Command{
	Use:   "differ",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if flags.grammarFile == "" {
			panic("grammar file is required")
		}

		if flags.target == "" {
			panic("target grammar rule is required")
		}

		differ := differ.NewDiffer(
			flags.grammarFile,
			flags.target,
		)

		if err := differ.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(differCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// differCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// differCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	differCmd.Flags().StringVarP(&flags.grammarFile, "grammar", "g", "", "Grammar file to parse")

	differCmd.Flags().StringVarP(&flags.target, "target", "t", "", "Target grammar rule for differ")
}
