// Package cmd provides commands for the application.
// Copyright 2020 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author Don B. Stringham <donbstringham@icloud.com>
//
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/acarl005/stripansi"
	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
	logr "github.com/spf13/jwalterweatherman"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var (
	wf    *aw.Workflow
	query string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list known servers",
	Long:  "List the known servers in ~/.ssh/config",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			query = ""
		} else {
			query = args[0]
		}
		wf = aw.New()
		wf.Run(run)
	},
}

func run() {
	wf.Args()

	out, err := exec.Command("/usr/local/bin/storm", "list").Output()

	if err != nil {
		logr.FATAL.Println(err)
	}

	buf := strings.Split(string(out), "\n\n")

	for _, server := range buf {
		if strings.Contains(server, "Listing entries:") {
			continue
		}

		if strings.Contains(server, "(*) General options:") {
			continue
		}

		if server == "" {
			continue
		}

		srvLine := strings.Split(server, "\n")
		srvName := strings.Split(srvLine[0], " -> ")
		srvNameClean := strings.TrimSpace(srvName[0])
		srvNameClean = stripansi.Strip(srvNameClean)
		u := "ssh://" + srvNameClean
		wf.NewItem(srvNameClean).Subtitle(u).Arg(u).Valid(true)
	}

	fmt.Printf("%s", query)

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching server found", "Try a different query")
	wf.SendFeedback()
}
