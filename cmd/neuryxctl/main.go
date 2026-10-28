// Comando neuryxctl: CLI do Neuryx.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/neuryx-os/neuryxctl/internal/version"
)

const (
	exitOK             = 0
	exitNotImplemented = 2
)

const usage = `neuryxctl — CLI do Neuryx

Uso:
  neuryxctl <comando> [argumentos]

Comandos:
  version         Mostra a versão
  help            Mostra esta ajuda

Planejados (ainda não implementados):
  gpu detect      Detecta as GPUs presentes
  gpu setup       Instala o driver da GPU
  model pull      Baixa um modelo GGUF
  model list      Lista os modelos instalados
  model verify    Verifica o hash dos modelos
  serve enable    Ativa o llama-server para um modelo
  serve status    Mostra o estado dos serviços de inferência
`

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stdout, usage)
		return exitOK
	}

	switch args[0] {
	case "version", "--version", "-v":
		fmt.Fprintln(stdout, version.String())
		return exitOK
	case "help", "--help", "-h":
		fmt.Fprint(stdout, usage)
		return exitOK
	default:
		fmt.Fprintln(stderr, "comando ainda não implementado")
		return exitNotImplemented
	}
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
