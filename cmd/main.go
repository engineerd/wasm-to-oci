package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/engineerd/wasm-to-oci/pkg/tuf"
	"github.com/spf13/cobra"
)

var (
	trustServer string
	tlscacert   string
	trustDir    string
	logLevel    string
	timeout     string
)

func main() {
	cmd := &cobra.Command{
		Use: "wasm-to-oci <subcommand> [options]",
	}

	cmd.PersistentFlags().StringVarP(&trustServer, "server", "", tuf.DockerNotaryServer, "The trust server used")
	cmd.PersistentFlags().StringVarP(&tlscacert, "tlscacert", "", "", "Trust certs signed only by this CA")
	cmd.PersistentFlags().StringVarP(&trustDir, "dir", "d", defaultTrustDir(), "Directory where the trust data is persisted to")
	cmd.PersistentFlags().StringVar(&logLevel, "log", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)
	cmd.PersistentFlags().StringVarP(&timeout, "timeout", "t", "5s", `Timeout for the trust server`)

	cmd.AddCommand(newPushCmd(), newPullCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func defaultTrustDir() string {
	homeEnvPath := os.Getenv("HOME")
	if homeEnvPath == "" && runtime.GOOS == "windows" {
		homeEnvPath = os.Getenv("USERPROFILE")
	}

	return filepath.Join(homeEnvPath, ".signy")
}
