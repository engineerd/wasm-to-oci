# WASM to OCI

The goal of this project is to propose an implementation of storing WebAssembly modules in OCI registries.

This project is built with the [ORAS project](https://github.com/deislabs/oras), and currently works with:

- [Distribution (open source, version 2.7+)](https://github.com/docker/distribution)
- [Azure Container Registry](https://docs.microsoft.com/en-us/azure/container-registry/)

# Usage

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

$ wasmer run test.wasm
Hello from WebAssembly!
```

# The OCI manifest

Since this is using ORAS to interact with OCI registries, here is a generated OCI manifest:

```
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
