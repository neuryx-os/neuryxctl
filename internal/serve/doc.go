// Package serve gerenciará as unidades systemd do llama-server, uma por
// modelo servido.
//
// Conforme o ADR 0007, não há gateway próprio: cada instância expõe
// diretamente a API compatível com a OpenAI do llama-server.
//
// Comandos previstos: "serve enable" e "serve status". Ainda não implementado.
package serve
