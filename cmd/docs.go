// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var Path string
var OutputFormat string

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for this tool",
	Long: `Generate the documentation for this command line tool
	
Example: cli-docs docs --path ./docs`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch OutputFormat {
		case "rst":
			err = doc.GenReSTTree(rootCmd, Path)
		default:
			err = doc.GenMarkdownTree(rootCmd, Path)
		}
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	docsCmd.Flags().StringVarP(&Path, "path", "", "./", "Where to put the generated documentation")
	docsCmd.Flags().StringVarP(&OutputFormat, "output-format", "", "markdown", "Format to generate. Expected values: markdown, rst. Default: markdown")
}
