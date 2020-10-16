package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/mvenezia/version-info/pkg/version"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version Information",
		Long:  `Provides the version information`,
		Run: func(cmd *cobra.Command, args []string) {
			versionHandler()
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func versionHandler() {
	log.Print("Version Requested")
	info := version.Get()
	fmt.Printf("Version Information:\n")
	fmt.Printf("\tGit Data:\n")
	fmt.Printf("\t\tTagged Version:\t%s\n", info.GitVersion)
	fmt.Printf("\t\tHash:\t\t%s\n", info.GitCommit)
	fmt.Printf("\t\tTree State:\t%s\n", info.GitTreeState)
	fmt.Printf("\tBuild Data:\n")
	fmt.Printf("\t\tBuild Date:\t%s\n", info.BuildDate)
	fmt.Printf("\t\tGo Version:\t%s\n", info.GoVersion)
	fmt.Printf("\t\tCompiler:\t%s\n", info.Compiler)
	fmt.Printf("\t\tPlatform:\t%s\n\n", info.Platform)
	log.Print("Version Served")
}
