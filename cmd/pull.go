package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/engineerd/wasm-to-oci/pkg/oci"
	"github.com/engineerd/wasm-to-oci/pkg/tuf"
)

type pullOptions struct {
	ref     string
	outFile string

	sign bool
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
	cmd.Flags().BoolVarP(&opts.sign, "sign", "", false, "Verifies the signature of the WebAssembly module from a trust server")

	return cmd
}

func (p *pullOptions) run() error {
	
	err := oci.Pull(p.ref, p.outFile, regopts)
	if err != nil {
		return err
	}

	if p.sign {
		err = tuf.VerifyFileTrust(p.ref, p.outFile, trustServer, tlscacert, trustDir, timeout)
		if err != nil {
			os.Remove(p.outFile)
			return err
		}
	}

	return nil
}
