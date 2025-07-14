# 🤖 WhatsBot

Automatize o envio e o recebimento de mensagens do WhatsApp com uma API leve, eficiente e escrita em Go!

## 📦 Sobre o Projeto

O **WhatsBot** é uma aplicação backend desenvolvida com **Go** que fornece endpoints para integração com o WhatsApp via Webhooks e envio de mensagens. É ideal para automações, bots, notificações e integrações com sistemas externos.

## 🚀 Funcionalidades

- ✅ Recebimento de mensagens via webhook (`/api/webhook`)
- 📤 Envio de mensagens autenticadas (`/api/send`)
- 🧪 Endpoint de teste (`/ping`)
- 🔐 Middleware de autenticação
- 🐳 Suporte completo a Docker e Docker Compose

## 🛠️ Tecnologias Utilizadas

- [Go](https://golang.org/) — Backend
- [Chi](https://github.com/go-chi/chi) — Gerenciador de rotas HTTP
- Docker + Docker Compose — Contêineres e orquestração
- Middleware customizado para autenticação

## 📁 Estrutura do Projeto

```bash
.
├── main.go               # Arquivo principal da aplicação
├── controllers/          # Lógica dos endpoints
├── middleware/           # Middleware de autenticação
├── Dockerfile            # Container Docker da aplicação
├── compose.yaml          # Orquestração com Docker Compose
└── .env.example          # Exemplo de variáveis de ambiente
