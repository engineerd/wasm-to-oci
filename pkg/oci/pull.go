package oci

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"oras.land/oras-go/pkg/oras"
)

// Pull pulls a Wasm module from an OCI registry given a reference
func Pull(ref, outFile string, proxyURL string, insecure, useHTTP bool) error {
	ctx, resolver, store := newORASContext(proxyURL, insecure, useHTTP)

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
