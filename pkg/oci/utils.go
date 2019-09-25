package oci

import (
	"context"
	"fmt"
	"os"

	"github.com/containerd/containerd/remotes"
	"github.com/containerd/containerd/remotes/docker"
	auth "github.com/deislabs/oras/pkg/auth/docker"
	orascnt "github.com/deislabs/oras/pkg/content"
	orasctx "github.com/deislabs/oras/pkg/context"
)

func newORASContext() (context.Context, remotes.Resolver, *orascnt.Memorystore) {
	ctx := orasctx.Background()
	memoryStore := orascnt.NewMemoryStore()
	cli, err := auth.NewClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: Error loading auth file: %v\n", err)
	}
	resolver, err := cli.Resolver(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: Error loading resolver: %v\n", err)
		resolver = docker.NewResolver(docker.ResolverOptions{})
	}

	return ctx, resolver, memoryStore
}
