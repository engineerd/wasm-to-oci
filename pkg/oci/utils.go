package oci

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/containerd/containerd/remotes"
	"github.com/containerd/containerd/remotes/docker"
	auth "oras.land/oras-go/pkg/auth/docker"
	orascnt "oras.land/oras-go/pkg/content"
	orasctx "oras.land/oras-go/pkg/context"
)

func newORASContext(proxyURL string, insecure, useHTTP bool) (context.Context, remotes.Resolver, *orascnt.Memorystore) {
	var transport *http.Transport
	ctx := orasctx.Background()
	memoryStore := orascnt.NewMemoryStore()
	cli, err := auth.NewClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: Error loading auth file: %v\n", err)
	}

	if insecure {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	if proxyURL != "" {
		if transport == nil {
			transport = &http.Transport{}
		}
		proxyURLParse, _ := url.Parse(proxyURL)
		transport.Proxy = http.ProxyURL(proxyURLParse)
	}

	if transport == nil {
		transport = &http.Transport{}
	}
	client := &http.Client{Transport: transport}

	resolver, err := cli.Resolver(context.Background(), client, useHTTP)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: Error loading resolver: %v\n", err)
		resolver = docker.NewResolver(docker.ResolverOptions{})
	}

	return ctx, resolver, memoryStore
}
