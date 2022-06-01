package cmd

import (
	"fmt"
	"github.com/jcmaunsell/rhesus/logger"
	"github.com/jcmaunsell/rhesus/repl"
	"github.com/spf13/cobra"
	"os"
	"os/user"
)

var rootCmd = &cobra.Command{
	Use:   "rhesus",
	Short: "Rhesus is a compiler for the Monkey language",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			panic(fmt.Errorf("could not get username: %w", err))
		}
		cmd.Printf("Hi %s! Welcome to Rhesus.\n", usr.Username)
		cmd.Println("You can enter Monkey commands here.")
		r, err := repl.New()
		if err != nil {
			panic(fmt.Errorf("could not initialize REPL: %w", err))
		}
		r.Start(os.Stdin, os.Stdout)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Service().WithError(err).Error("The CLI exited with an error.")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
