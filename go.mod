module metalctl

go 1.13

require (
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.6.2 // indirect
	golang.org/x/sys v0.0.0-20200124204421-9fbb57f87de9 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	k8s.io/klog v1.0.0
	sigs.k8s.io/cluster-api v0.2.9
	sigs.k8s.io/yaml v1.1.0
)

replace sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v0.2.6-0.20200128020138-506cb9af40aa
