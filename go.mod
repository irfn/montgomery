module github.com/irfn/montgomery

go 1.15

require (
	github.com/go-git/go-git/v5 v5.2.0
	github.com/helm/helm v2.17.0+incompatible // indirect
	golang.org/x/tools/gopls v0.6.5 // indirect
	helm.sh/helm/v3 v3.5.2
	k8s.io/client-go v0.20.2
)

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

replace github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
