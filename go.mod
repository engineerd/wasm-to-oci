module github.com/engineerd/wasm-to-oci

go 1.16

require (
	github.com/containerd/containerd v1.5.2
	github.com/docker/cli v20.10.7+incompatible
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v17.12.1-ce+incompatible // indirect
	github.com/docker/go v1.5.1-1.0.20160303222718-d30aec9fd63c
	github.com/opencontainers/image-spec v1.0.1
	github.com/oras-project/oras-go v0.3.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/theupdateframework/notary v0.7.0
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
