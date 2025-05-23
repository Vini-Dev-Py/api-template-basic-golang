# 🧱 Go API Template

Template básico para criar APIs em Golang com Docker e PostgreSQL. Ideal para iniciar novos projetos rapidamente com uma estrutura organizada e pronta para desenvolvimento local.

---

## 🚀 Tecnologias

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Air](https://github.com/cosmtrek/air) (hot reload para Go)

---

## ⚙️ Como usar

### Pré-requisitos

- Docker e Docker Compose instalados
- Go instalado (opcional, se for rodar fora do container)

### Passo a passo

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/go-api-template.git
cd go-api-template
docker-compose up --build
```

2. Acesse a API em: http://localhost:8080

🧪 Endpoints de exemplo
GET /health → Verifica se o servidor está online.

🗃️ Variáveis de ambiente
As variáveis estão definidas no .env:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_api
APP_PORT=8080
