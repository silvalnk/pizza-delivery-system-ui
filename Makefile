BINARY  := pizza-tracker
MAIN    := ./cmd
GO      := go
BINDIR  := bin

.PHONY: run build test vet clean deps docker-up docker-down docker-logs help

run:
	$(GO) run $(MAIN)

build:
	@mkdir -p $(BINDIR)
	$(GO) build -o $(BINDIR)/$(BINARY) $(MAIN)

test:
	$(GO) test ./...

vet:
	$(GO) vet ./...

deps:
	$(GO) mod download
	$(GO) mod tidy

docker-up:
	docker compose up -d db

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

clean:
	rm -rf $(BINDIR)

dev: docker-up
	@echo "Aguardando PostgreSQL..."
	@sleep 3
	$(GO) run $(MAIN)

help:
	@echo "Uso: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run         Executa a aplicação (go run ./cmd)"
	@echo "  build       Compila o binário em $(BINDIR)/$(BINARY)"
	@echo "  test        Executa os testes"
	@echo "  vet         Executa go vet"
	@echo "  deps        Baixa e atualiza dependências"
	@echo "  docker-up   Sobe o PostgreSQL (Docker)"
	@echo "  docker-down Para os containers"
	@echo "  docker-logs Exibe os logs do Docker"
	@echo "  clean       Remove o diretório $(BINDIR)"
	@echo "  dev         Sobe o banco e executa a aplicação"
	@echo "  help        Exibe esta mensagem"
