# Tools
GOLANGCI_LINT_VERSION := latest
GOSEC_VERSION := latest

# Test
.PHONY: test
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: test-short
test-short:
	go test -v ./...

.PHONY: test-cover
test-cover: test
	go tool cover -html=coverage.txt

# Example
.PHONY: example
example:
	cd example && go run main.go

# Submodule management
.PHONY: update-submodule
update-submodule:
	git submodule update --remote --merge source-data

.PHONY: init-submodule
init-submodule:
	git submodule update --init --recursive

# Code quality
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found. Run 'make install-tools' to install it."; \
		exit 1; \
	fi
	golangci-lint run --timeout=5m

# Security
.PHONY: security
security:
	@if ! command -v gosec &> /dev/null; then \
		echo "gosec not found. Run 'make install-tools' to install it."; \
		exit 1; \
	fi
	gosec -fmt=text -out=results.txt ./...
	@cat results.txt
	@rm results.txt

# Development tools
.PHONY: install-tools
install-tools:
	@echo "Installing development tools..."
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION); \
	else \
		echo "golangci-lint already installed"; \
	fi
	@if ! command -v gosec &> /dev/null; then \
		echo "Installing gosec..."; \
		go install github.com/securego/gosec/v2/cmd/gosec@$(GOSEC_VERSION); \
	else \
		echo "gosec already installed"; \
	fi
	@echo "All tools installed!"

# CI
.PHONY: ci
ci: fmt vet lint test security

# Docker
.PHONY: docker-build
docker-build:
	docker compose build

.PHONY: docker-dev
docker-dev:
	docker compose run --rm dev

.PHONY: docker-test
docker-test:
	docker compose run --rm test

.PHONY: docker-lint
docker-lint:
	docker compose run --rm lint

.PHONY: docker-security
docker-security:
	docker compose run --rm security

.PHONY: docker-example
docker-example:
	docker compose run --rm example

.PHONY: docker-clean
docker-clean:
	docker compose down -v
	docker system prune -f

# Clean
.PHONY: clean
clean:
	go clean
	rm -f coverage.txt
	rm -f example/example

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo ""
	@echo "Testing:"
	@echo "  test             - Run tests with race detector and coverage"
	@echo "  test-short       - Run tests without race detector"
	@echo "  test-cover       - Run tests and open coverage report in browser"
	@echo ""
	@echo "Development:"
	@echo "  example          - Run example code"
	@echo "  fmt              - Format code"
	@echo "  vet              - Run go vet"
	@echo "  lint             - Run golangci-lint"
	@echo "  security         - Run security scan with gosec"
	@echo "  install-tools    - Install development tools"
	@echo ""
	@echo "Docker:"
	@echo "  docker-build     - Build Docker images"
	@echo "  docker-dev       - Start development environment in Docker"
	@echo "  docker-test      - Run tests in Docker"
	@echo "  docker-lint      - Run linter in Docker"
	@echo "  docker-security  - Run security scan in Docker"
	@echo "  docker-example   - Run example in Docker"
	@echo "  docker-clean     - Clean Docker containers and volumes"
	@echo ""
	@echo "Submodules:"
	@echo "  init-submodule   - Initialize source-data submodule"
	@echo "  update-submodule - Update source-data submodule"
	@echo ""
	@echo "CI/CD:"
	@echo "  ci               - Run all checks (fmt, vet, lint, test, security)"
	@echo ""
	@echo "Maintenance:"
	@echo "  clean            - Clean build artifacts"
	@echo "  help             - Show this help message"

