module metalctl

go 1.13

require (
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/phogolabs/parcello v0.8.2
	github.com/spf13/cobra v0.0.5
	sigs.k8s.io/kustomize/api v0.3.2
	sigs.k8s.io/yaml v1.2.0
)

replace sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v0.2.6-0.20200128020138-506cb9af40aa
