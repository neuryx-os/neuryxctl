# neuryxctl

CLI do Neuryx — setup de GPU, gestão de modelos e serviços de inferência.

**Status:** Em desenvolvimento

## Propósito

O `neuryxctl` é a ferramenta de administração do Neuryx. Ele será responsável
por:

- detectar a GPU e instalar o driver no primeiro boot, a partir do repositório
  oficial da NVIDIA
  ([ADR 0004](https://github.com/neuryx-os/research/blob/main/adr/0004-driver-nvidia-no-primeiro-boot.md));
- baixar modelos GGUF, verificar seus hashes e manter um inventário;
- criar e gerenciar unidades systemd do `llama-server`, uma por modelo
  ([ADR 0007](https://github.com/neuryx-os/research/blob/main/adr/0007-sem-gateway-e-monitoramento-proprios.md)).

É escrito em Go
([ADR 0005](https://github.com/neuryx-os/research/blob/main/adr/0005-neuryxctl-em-go.md)),
usando apenas a biblioteca padrão nesta etapa.

## Comandos

| Comando | Descrição | Estado |
| --- | --- | --- |
| `neuryxctl version` | Mostra a versão | Implementado |
| `neuryxctl help` | Mostra a ajuda | Implementado |
| `neuryxctl gpu detect` | Detecta as GPUs presentes | Não implementado |
| `neuryxctl gpu setup` | Instala o driver da GPU | Não implementado |
| `neuryxctl model pull` | Baixa um modelo GGUF | Não implementado |
| `neuryxctl model list` | Lista os modelos instalados | Não implementado |
| `neuryxctl model verify` | Verifica o hash dos modelos | Não implementado |
| `neuryxctl serve enable` | Ativa o `llama-server` para um modelo | Não implementado |
| `neuryxctl serve status` | Mostra o estado dos serviços de inferência | Não implementado |

Comandos não implementados imprimem `comando ainda não implementado` e saem com
código 2.

## Desenvolvimento

Requer Go (versão em `go.mod`).

```bash
make lint    # gofmt e go vet
make test    # go test ./...
make build   # gera bin/neuryxctl
make build VERSION=0.1.0
./bin/neuryxctl version
```

## Estrutura

```
neuryxctl/
├── cmd/neuryxctl/       # ponto de entrada
├── internal/
│   ├── version/         # versão definida via -ldflags (padrão "dev")
│   ├── gpu/             # detecção de GPU e setup de driver
│   ├── model/           # download, hash e inventário de modelos
│   └── serve/           # unidades systemd do llama-server
└── .github/workflows/   # CI: vet, test, build
```

## Relação com outros repositórios

- [`neuryx-packages`](https://github.com/neuryx-os/neuryx-packages): empacota o `neuryxctl` para o OBS.
- [`neuryx-tuning`](https://github.com/neuryx-os/neuryx-tuning): ajustes que dependem do driver instalado pelo `neuryxctl`.
- [`research`](https://github.com/neuryx-os/research): ADRs 0004, 0005, 0006 e 0007.
- [`lyra-os-linux/vegad`](https://github.com/lyra-os-linux/vegad): daemon do Vega, também em Go; a gestão gráfica será um módulo de IA do Vega (ADR 0006).

## Licença

[GPL-3.0](LICENSE).
