package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "wasm-to-oci <subcommand> [options]",
	}

	cmd.AddCommand(newPushCmd(), newPullCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
