package main

import (
	"fmt"
	"github.com/ebauman/mink-aggregator-test/pkg/server"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "aggregator",
	Short: "run api server aggregator",
	RunE:  app,
}

func app(cmd *cobra.Command, args []string) error {
	svr, err := server.NewAggregatedServer()
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := svr.Run(cmd.Context()); err != nil {
		return err
	}

	<-cmd.Context().Done()

	return cmd.Context().Err()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
