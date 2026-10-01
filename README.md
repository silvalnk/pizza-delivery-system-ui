# Pizza Tracker

Aplicação de rastreamento de pedidos de pizza em Go, organizada em **Clean Architecture** com **MVVM** na camada de apresentação. Banco de dados: **PostgreSQL**.

![Acompanhamento do pedido: pizza, status ao vivo e os dados da entrega](docs/images/pedido.jpg)

## Estrutura do projeto (Clean Architecture + MVVM)

```
cmd/                    # Ponto de entrada (wire de dependências)
internal/
  domain/               # Regras de negócio (entities + interfaces)
    entity/             # Entidades (Order, OrderItem, User)
    repository/         # Interfaces dos repositórios
  app/                  # Casos de uso
    usecase/            # OrderService, AuthService
    notification/       # Manager (SSE)
  infra/                # Detalhes técnicos
    config/
    persistence/gorm/   # Implementação dos repositórios (PostgreSQL)
    delivery/http/      # Camada de apresentação HTTP
      handler/          # Handlers (ViewModels)
      router/           # Rotas
      middleware/       # Auth
      session/          # Sessão
      viewmodel/        # Dados para as views
      templates/        # Views (HTML) + static/ (favicon, assets)
docker-compose.yml      # PostgreSQL
```

- **Domain**: entidades e contratos (repositórios).
- **App**: serviços que orquestram a lógica (use cases).
- **Infra**: banco (GORM + PostgreSQL), HTTP (Gin), sessão, notificações.
- **MVVM**: Handlers preparam ViewModels; templates são as Views.

## Como rodar

1. Subir o PostgreSQL (Docker):

```bash
docker compose up -d db
```

2. Configurar o `.env` (ou use o padrão que aponta para o Docker):

```bash
cp .env.example .env
# edite .env se o Postgres estiver em outro host/porta
```

3. Executar a aplicação:

```bash
go run ./cmd
```

Acesse: **http://localhost:8080**

## Variáveis de ambiente

Copie `.env.example` para `.env` e ajuste se precisar. O app carrega `.env` automaticamente.

| Variável             | Padrão (Docker)                                                                 | Descrição                    |
|----------------------|---------------------------------------------------------------------------------|------------------------------|
| `PORT`               | `8080`                                                                          | Porta HTTP                   |
| `DATABASE_URL`       | `postgres://pizzauser:pizzapass@localhost:5432/pizza_tracker?sslmode=disable`  | URL do PostgreSQL            |
| `SESSION_SECRET_KEY` | (valor padrão)                                                                  | Chave para sessão (produção)  |

## Usuário admin

Na primeira execução o app cria um usuário admin padrão:

| Campo    | Valor      |
|----------|------------|
| Usuário  | `admin`    |
| Senha    | `admin123` |

Use para acessar `/login` e depois `/admin`. **Em produção, altere a senha ou crie outro usuário e remova o admin padrão.**
