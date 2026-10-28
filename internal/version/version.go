// Package version expõe a versão do neuryxctl.
package version

// Version é definida no build com:
//
//	go build -ldflags "-X github.com/neuryx-os/neuryxctl/internal/version.Version=1.0.0"
var Version = "dev"

// String retorna a versão atual.
func String() string {
	return Version
}
