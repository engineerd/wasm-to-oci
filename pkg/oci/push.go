package oci

import (
	"fmt"
	"io/ioutil"

	"github.com/deislabs/oras/pkg/oras"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// Push pushes a WASM module an OCI registry
func Push(ref, module string) error {
	ctx, resolver, store := newORASContext()

	contents, err := ioutil.ReadFile(module)
	if err != nil {
		return err
	}

	desc := store.Add(module, ContentLayerMediaType, contents)
	layers := []ocispec.Descriptor{desc}

	pushOpts := []oras.PushOpt{
		oras.WithConfigMediaType(ConfigMediaType),
		oras.WithNameValidation(nil),
	}

	manifest, err := oras.Push(ctx, resolver, ref, store, layers, pushOpts...)
	if err != nil {
		return err
	}

	fmt.Printf("\nPushed: %v", ref)
	fmt.Printf("\nSize: %v", desc.Size)
	fmt.Printf("\nDigest: %v", manifest.Digest)

	return nil
}
