// Package cmd provides commands for the application.
// Copyright 2020 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author donbstringham <donbstringham@icloud.com>
//
package cmd

import (
	"fmt"

	"github.com/donbstringham/gstorm/ver"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", ver.Version)
	},
}
