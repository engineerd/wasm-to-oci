package oci

import (
	"io/ioutil"

	"github.com/deislabs/oras/pkg/oras"
	log "github.com/sirupsen/logrus"
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

	log.Infof("Pulled: %v", ref)
	log.Infof("Size: %v", desc.Size)
	log.Infof("Digest: %v", manifest.Digest)

	return nil
}
