package cmd

import (
	"github.com/jcmaunsell/rhesus/logger"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Print REPL logs from last run",
	Run: func(cmd *cobra.Command, args []string) {
		logger.PrintLogFile(logger.ServiceLogPath)
	},
}
