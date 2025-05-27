# API REST em Go com Gin

Este projeto é uma API REST construída utilizando o framework **Gin** em **Go**. A API permite a criação, leitura, atualização e deleção de registros, integrando-se com um banco de dados **PostgreSQL** e utilizando **Docker**.

## 🛠 Tecnologias e Ferramentas

- **Linguagem**: Go (Golang)
- **Framework Web**: [Gin](https://github.com/gin-gonic/gin) — Framework HTTP para Go
- **Banco de Dados**: PostgreSQL
- **Gerenciamento de Dependências**: Go Modules
- **Containerização**: Docker & Docker Compose
- **ORM**: [GORM](https://gorm.io/) — ORM para Go
- **Driver PostgreSQL**: [pgx](https://github.com/jackc/pgx) integrado ao GORM

## 📦 Bibliotecas Principais

- `github.com/gin-gonic/gin` — Framework web
- `gorm.io/gorm` — ORM para Go
- `gorm.io/driver/postgres` — Driver PostgreSQL para GORM

## 🚀 Como Executar o Projeto

### 🔧 Pré-requisitos

- **Go** instalado (versão 1.21 ou superior)
- **Docker** e **Docker Compose** instalados

## ▶️ Executar com Docker Compose

bash
docker-compose up --build
🔗 API disponível em:
arduino
Copiar
Editar
http://localhost:8080
⚙️ Endpoints principais
GET /alunos — Lista todos os alunos

GET /alunos/:id — Retorna um aluno pelo ID

POST /alunos — Cria um novo aluno

PUT /alunos/:id — Atualiza os dados de um aluno

DELETE /alunos/:id — Remove um aluno

## 🗄️ Estrutura do Projeto
python

.
├── main.go                # Arquivo principal que inicia o servidor
├── routes/                # Arquivos que definem as rotas
├── models/                # Definição dos modelos (entidades)
├── database/              # Conexão com o banco de dados
├── docker-compose.yaml    # Arquivo de configuração do Docker
├── go.mod                 # Gerenciamento de dependências
└── go.sum                 # Checksum das dependências
🐳 Docker Compose
O arquivo docker-compose.yaml sobe dois serviços:

db: Banco de dados PostgreSQL na porta 5432

api: A aplicação Go na porta 8080

## 📚 Comandos Úteis
Rodar localmente sem Docker:
bash

go run main.go

## ✍️ Autor
Desenvolvido por Gustavo Aguilar
