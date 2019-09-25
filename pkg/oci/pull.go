package oci

import (
	"fmt"
	"io/ioutil"

	"github.com/deislabs/oras/pkg/oras"
)

// Pull pulls a WASM module from an OCI registry given a reference
func Pull(ref, outFile string) error {
	ctx, resolver, store := newORASContext()

	pullOpts := []oras.PullOpt{
		oras.WithAllowedMediaType(ContentLayerMediaType),
		oras.WithPullEmptyNameAllowed(),
	}

	_, layers, err := oras.Pull(ctx, resolver, ref, store, pullOpts...)
	if err != nil {
		return err
	}

	desc := layers[0]
	manifest, contents, _ := store.Get(desc)
	ioutil.WriteFile(outFile, contents, 0755)

	fmt.Printf("\nPulled: %v", ref)
	fmt.Printf("\nSize: %v", desc.Size)
	fmt.Printf("\nDigest: %v", manifest.Digest)

	return nil
}
