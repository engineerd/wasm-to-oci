package main

import (
	"github.com/engineerd/wasm-to-oci/pkg/oci"
	"github.com/spf13/cobra"
)

type pullOptions struct {
	ref     string
	outFile string
}

func newPullCmd() *cobra.Command {
	var opts pullOptions
	cmd := &cobra.Command{
		Use:   "pull <reference> [options]",
		Short: "Pulls a WASM module from an OCI registry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.ref = args[0]
			return opts.run()
		},
	}
	cmd.Flags().StringVarP(&opts.outFile, "out", "o", "module.wasm", "Name of the output module")

	return cmd
}

func (p *pullOptions) run() error {
	return oci.Pull(p.ref, p.outFile)
}
