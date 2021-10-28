package main

import (
	"encoding/hex"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/engineerd/wasm-to-oci/pkg/oci"
	"github.com/engineerd/wasm-to-oci/pkg/tuf"
)

type pushOptions struct {
	module string
	ref    string

	sign bool
}

func newPushCmd() *cobra.Command {
	var opts pushOptions
	cmd := &cobra.Command{
		Use:   "push <module> <reference> [options]",
		Short: "Pushes a WASM module to an OCI registry",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.module = args[0]
			opts.ref = args[1]
			return opts.run()
		},
	}

	cmd.Flags().BoolVarP(&opts.sign, "sign", "", false, "Signs the WebAssembly module and pushes the metadata to a trust server")
	return cmd
}

func (p *pushOptions) run() error {
	if p.sign {
		target, err := tuf.SignAndPublish(trustDir, trustServer, p.ref, p.module, tlscacert, "", timeout, nil)
		if err != nil {
			return fmt.Errorf("cannot sign and publish trust data: %v", err)
		}
		log.Infof("Pushed trust data for %v: %v\n", p.ref, hex.EncodeToString(target.Hashes["sha256"]))

	}

	return oci.Push(p.ref, p.module, insecure, useHTTP)
}
