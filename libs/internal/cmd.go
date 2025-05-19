package main

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/internal/abigen"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd := cobra.Command{
		Use: "internal",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger.InitLog(logger.Config{
				VerbosityTerminal: 5,
				UseTerminal:       true,
			})
		},
	}

	rootCmd.AddCommand(abigen.GetCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
