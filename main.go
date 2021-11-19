package main

import (
	"github.com/spf13/cobra"
	"gserver/cmd"
)

func init() {

}

func main() {
	var err error
	rootCmd := &cobra.Command{
		Use: "gserver",
	}
	rootCmd.AddCommand(cmd.ServerRun())
	rootCmd.AddCommand(cmd.MgRun())
	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
