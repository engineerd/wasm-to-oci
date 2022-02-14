package oci

import (
	log "github.com/sirupsen/logrus"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

// Pull pulls a Wasm module from an OCI registry given a reference
func Pull(ref, outFile string, opts content.RegistryOptions) error {
	ctx, registry, _ := newORASContext(opts)

	// Pull file(s) from registry and save to disk
	log.Infof("Pulling from %s and saving to %s...\n", ref, outFile)
	fileStore := content.NewFile(outFile)
	defer fileStore.Close()
	allowedMediaTypes := []string{ContentLayerMediaType}
	desc, err := oras.Copy(ctx, registry, ref, fileStore, "", oras.WithAllowedMediaTypes(allowedMediaTypes))
	if err != nil {
		return err
	}
	log.Infof("Pulled: %v", ref)
	log.Infof("Size: %v", desc.Size)
	log.Infof("Digest: %v", desc.Digest)

	return nil
}
