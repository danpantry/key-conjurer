package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const versionString string = `
	Version: 		%s
	Client: 		%s
	Default Hostname:	%s
	Upgrade URL:		%s
`

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Shows Key Conjurer version information.",
	Long:    "Shows Key Conjurer version information.",
	Example: "keyconjurer version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(versionString, Version, ClientName, defaultHost, DownloadURL)
	}}