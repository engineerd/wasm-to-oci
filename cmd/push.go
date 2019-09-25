package main

import (
	"github.com/engineerd/wasm-to-oci/pkg/oci"
	"github.com/spf13/cobra"
)

type pushOptions struct {
	module string
	ref    string
}

func newPushCmd() *cobra.Command {
	var opts pushOptions
	cmd := &cobra.Command{
		Use:   "push <module> <reference> [options]",
		Short: "Pushes a WASM module from an OCI registry",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.module = args[0]
			opts.ref = args[1]
			return opts.run()
		},
	}

	return cmd
}

func (p *pushOptions) run() error {
	return oci.Push(p.ref, p.module)
}
