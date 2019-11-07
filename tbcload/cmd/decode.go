// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"encoding/hex"
	"fmt"

	"github.com/mvphilip/tbcload"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode [string to decode]",
	Short: "decode a string into ascii85(re-map), which tbc file used",
	Long: `tbc file use ascii85 encode and map special ascii code.
For example:
	tbcload decode ",CHr@"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var src []byte

		if len(args) == 1 {
			//if there is 'z' ,length will large than src
			dst := make([]byte, len(args[0])*4)

			src = []byte(args[0])
			if ndst := tbcload.Decode(dst, src); ndst > 0 {
				fmt.Printf("source:%s\n", args[0])
				fmt.Printf("decode:%s\n", dst[:ndst])
				fmt.Printf("dump  :\n%s", hex.Dump(dst[:ndst]))
			} else {
				fmt.Printf("source:%s\n", args[0])
				fmt.Printf("decode error, maybe wrong source string")
			}
		} else {
			cmd.Usage()
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
