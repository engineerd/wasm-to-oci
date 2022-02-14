module github.com/engineerd/wasm-to-oci

go 1.16

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/hcsshim v0.9.2 // indirect
	github.com/containerd/containerd v1.5.9
	github.com/containerd/continuity v0.2.2 // indirect
	github.com/docker/cli v20.10.12+incompatible
	github.com/docker/distribution v2.8.0+incompatible
	github.com/docker/docker v20.10.12+incompatible // indirect
	github.com/docker/go v1.5.1-1.0.20160303222718-d30aec9fd63c
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/miekg/pkcs11 v1.1.1 // indirect
	github.com/opencontainers/image-spec v1.0.2
	github.com/opencontainers/runc v1.1.0 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0
	github.com/theupdateframework/notary v0.7.0
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220204135822-1c1b9b1eba6a // indirect
	google.golang.org/genproto v0.0.0-20220207185906-7721543eae58 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	oras.land/oras-go v1.1.0
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
