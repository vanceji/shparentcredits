package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	host     string
	schoolId int32
	classId  int32
	token    string
	rootCmd  = &cobra.Command{
		Use:     "shparentcredits",
		Short:   "Credits tool for parents",
		Long:    `This tool can help parents living SH to get credits easily`,
		Version: "v1.0.0",
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
