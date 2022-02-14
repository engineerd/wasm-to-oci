package main

import (
	"os"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"oras.land/oras-go/pkg/content"

	"github.com/engineerd/wasm-to-oci/pkg/tuf"
)

var (
	trustServer string
	tlscacert   string
	trustDir    string
	logLevel    string
	timeout     string

	regopts content.RegistryOptions
)

func main() {

	cmd := &cobra.Command{
		Use: "wasm-to-oci <subcommand> [options]",

		Short: "Manage WebAssemblies on OCI repositories",

		Example: "wasm-to-oci <subcommand> [options]",

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
	cmd.PersistentFlags().StringVarP(&regopts.Username, "username", "u", "", `Username (required)`)
	cmd.PersistentFlags().StringVarP(&regopts.Password, "password", "p", "", `Password (required)`)

	cmd.PersistentFlags().BoolVarP(&regopts.Insecure, "insecure", "", false, "Allow connections to SSL registry without certs")
	cmd.PersistentFlags().BoolVarP(&regopts.PlainHTTP, "use-http", "", false, "Use plain http instead of https")

	cmd.MarkPersistentFlagRequired("username")
	cmd.MarkPersistentFlagRequired("password")

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
