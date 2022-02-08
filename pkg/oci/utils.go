package oci

import (
	"context"

	"oras.land/oras-go/pkg/content"
	orascnt "oras.land/oras-go/pkg/content"
	orasctx "oras.land/oras-go/pkg/context"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func newORASContext(opts content.RegistryOptions) (context.Context, *content.Registry, *orascnt.Memory) {
	ctx := orasctx.Background()
	memoryStore := orascnt.NewMemory()

	registry, err := content.NewRegistry(opts)

	check(err)

	return ctx, registry, memoryStore
}
