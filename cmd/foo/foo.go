package main

import (
	"fmt"
	"github.com/acorn-io/baaah/pkg/restconfig"
	"github.com/ebauman/mink-aggregator-test/pkg/scheme/foo"
	"github.com/ebauman/mink-aggregator-test/pkg/server"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	kubeconfig  string
	kubecontext string
)

func init() {
	rootCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "path to kubeconfig file")
	rootCmd.Flags().StringVar(&kubecontext, "context", "default", "kube context")
}

var rootCmd = &cobra.Command{
	Use:   "fooserver",
	Short: "run apiserver for foos",
	RunE:  app,
}

func app(cmd *cobra.Command, args []string) error {
	cfg, err := restconfig.FromFile(kubeconfig, kubecontext)
	if err != nil {
		log.Fatalf(err.Error())
	}

	kclient, err := client.NewWithWatch(cfg, client.Options{
		Scheme: foo.Scheme,
	})

	svr, err := server.NewFooServer(kclient)
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
