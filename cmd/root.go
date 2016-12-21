// Copyright © 2016 Thomas Fischer <tdf.tomfischer@gmail.com>
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
	"fmt"
	"os"

	"github.com/fschr/tunes/api"
	"github.com/spf13/cobra"
)

var (
	dir    string
	player string
	port   string
	// RootCmd represents the base command when called without any subcommands
	RootCmd = &cobra.Command{
		Use:   "tunes",
		Short: "a temporary sonos replacement with youtube support",
		Long:  ``,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			api.RunServer(dir, player, port)
		},
	}
)

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&dir, "dir", "", "tunes", "download directory")
	RootCmd.PersistentFlags().StringVarP(&player, "player", "", "mpv", "music player command")
	RootCmd.PersistentFlags().StringVarP(&port, "port", "", "8080", "port to run server on")
}
