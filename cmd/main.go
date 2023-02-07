package main

import (
	"os"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/engineerd/wasm-to-oci/pkg/tuf"
)

var (
	trustServer string
	tlscacert   string
	trustDir    string
	logLevel    string
	timeout     string
	proxyURL    string

	insecure bool
	useHTTP  bool
)

func main() {
	cmd := &cobra.Command{
		Use: "wasm-to-oci <subcommand> [options]",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			l, err := log.ParseLevel(logLevel)
			if err != nil {
				return err
			}
			log.SetLevel(l)
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&trustServer, "server", "", tuf.DockerNotaryServer, "The trust server used")
	cmd.PersistentFlags().StringVarP(&tlscacert, "tlscacert", "", "", "Trust certs signed only by this CA")
	cmd.PersistentFlags().StringVarP(&trustDir, "dir", "d", defaultTrustDir(), "Directory where the trust data is persisted to")
	cmd.PersistentFlags().StringVar(&logLevel, "log", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)
	cmd.PersistentFlags().StringVarP(&timeout, "timeout", "t", "5s", `Timeout for the trust server`)
	cmd.PersistentFlags().BoolVarP(&insecure, "insecure", "", false, "Allow connections to SSL registry without certs")
	cmd.PersistentFlags().BoolVarP(&useHTTP, "use-http", "", false, "Use plain http instead of https")
	cmd.PersistentFlags().StringVarP(&proxyURL, "proxy", "", "", `Allow connections to use proxy server`)

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

	return filepath.Join(homeEnvPath, ".wasm-to-oci")
}
