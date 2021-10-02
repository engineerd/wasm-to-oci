# WASM to OCI

The goal of this project is to propose an implementation of storing WebAssembly
modules in OCI registries.

This project is built with the
[ORAS project](https://github.com/oras-project/oras-go), and currently works
with:

- [Distribution (open source, version 2.7+)](https://github.com/docker/distribution)
- [Azure Container Registry](https://docs.microsoft.com/en-us/azure/container-registry/)
- [Google Container Registry](https://cloud.google.com/container-registry/)
- [Google Artifact Registry](https://cloud.google.com/artifact-registry)
- [Harbor Container Registry v2.0](https://github.com/goharbor/harbor/releases/tag/v2.0.0)
- [Bundle Bar](https://bundle.bar/docs/supported-clients/wasm-to-oci/)
- [GitHub Package Registry](https://github.com/features/packages)

> Note that trying to push a WebAssembly module to Docker Hub is not supported
> at the time of writing this document, as Docker Hub does _not_ accept unknown
> artifact types.

> As more registries add support for OCI Artifacts, we will update the list of
> supported registries.

# Usage

- login to your container registry using the `docker` CLI (or other tooling that
  your container registry provides. `wasm-to-oci` will use the credentials in
  `~/.docker/config.json`)

- pushing to an OCI registry:

```
$ ls testdata
.rwxr-xr-x 4.1M radu canonicaljson.wasm
.rwxr-xr-x 1.6M radu  hello.wasm

$ wasm-to-oci push testdata/hello.wasm <oci-registry>.azurecr.io/wasm-to-oci:v1

Pushed: <oci-registry>.azurecr.io/wasm-to-oci:v1
Size: 1624962
Digest: sha256:9c82cbe576ee947c00435ac8053a800a1969f4757ae4a81f870f714674afc91a
```

- pulling from an OCI registry:

```
$ wasm-to-oci pull <oci-registry>.azurecr.io/wasm-to-oci:v1 --out test.wasm

Pulled: <oci-registry>.azurecr.io/wasm-to-oci:v1
Size: 1624962
Digest: sha256:4c7915b4c1f9b0c13f962998e4199ceb00db39a4a7fa4554f40ae0bed83d9510

$ wasmtime test.wasm
Hello from WebAssembly!
```

# How does this work?

This leverages
[the OCI Artifacts proposal](https://github.com/opencontainers/artifacts), whose
goal is to enable the distribution of more cloud native artifacts using existing
registry infrastructure, and uses it to store WebAssembly modules as single
layer blobs in the registry.

This project defines a new set of _unofficial_ media types used to identify a
WebAssembly artifact - the artifacts project also describes the process for
projects to
[apply for an official unique media type](https://github.com/opencontainers/artifacts/blob/master/artifact-authors.md#registering-unique-types-with-iana).

```json
{
  "schemaVersion": 2,
  "config": {
    "mediaType": "application/vnd.wasm.config.v1+json",
    "digest": "sha256:44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a",
    "size": 2
  },
  "layers": [
    {
      "mediaType": "application/vnd.wasm.content.layer.v1+wasm",
      "digest": "sha256:4c7915b4c1f9b0c13f962998e4199ceb00db39a4a7fa4554f40ae0bed83d9510",
      "size": 1624962
    }
  ]
}
```

There is also experimental support for artifact signing with Notary v1 - see
[this article](https://radu-matei.com/blog/wasm-oci-tuf/) for more background on
this topic.
