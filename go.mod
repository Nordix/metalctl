module metalctl

go 1.13

require (
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	k8s.io/klog v1.0.0
	sigs.k8s.io/cluster-api v0.0.0-00010101000000-000000000000
	sigs.k8s.io/kustomize v2.0.3+incompatible
	sigs.k8s.io/kustomize/api v0.3.2
	sigs.k8s.io/kustomize/v3 v3.3.1
	sigs.k8s.io/yaml v1.1.0
)

replace sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v0.2.6-0.20200128020138-506cb9af40aa
