module github.com/engineerd/wasm-to-oci

go 1.13

require (
	github.com/agl/ed25519 v0.0.0-20150830182803-278e1ec8e8a6 // indirect
	github.com/bitly/go-hostpool v0.1.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cfssl v1.4.1 // indirect
	github.com/containerd/containerd v1.4.3
	github.com/deislabs/oras v0.9.0
	github.com/docker/cli v20.10.2+incompatible
	github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/go v1.5.1-1
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/opencontainers/image-spec v1.0.1
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.1
	github.com/theupdateframework/notary v0.6.1
	gopkg.in/dancannon/gorethink.v3 v3.0.5 // indirect
	gopkg.in/fatih/pool.v2 v2.0.0 // indirect
	gopkg.in/gorethink/gorethink.v3 v3.0.5 // indirect
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
