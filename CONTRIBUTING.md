# Contributing to zengin-go

Thank you for considering contributing to zengin-go! We welcome contributions from everyone.

## How to Contribute

### Reporting Issues

If you find a bug or have a suggestion for improvement:

1. Check if the issue already exists in the [Issues](https://github.com/hacomono-lib/zengin-go/issues) section
2. If not, create a new issue with a clear title and description
3. Include steps to reproduce the bug (if applicable)
4. Include the version of Go you're using

### Submitting Pull Requests

1. Fork the repository
2. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Make your changes
4. Install development tools (first time only):
   ```bash
   make install-tools
   ```
5. Run all checks before committing:
   ```bash
   make ci
   ```
   This will run:
   - Code formatting (`make fmt`)
   - Go vet (`make vet`)
   - Linting (`make lint`)
   - Tests with race detector (`make test`)
   - Security scanning (`make security`)

6. Commit your changes with a clear commit message:
   ```bash
   git commit -m "feat: add new feature"
   ```
7. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
8. Create a Pull Request

### Commit Message Guidelines

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

- `feat:` - A new feature
- `fix:` - A bug fix
- `docs:` - Documentation changes
- `style:` - Code style changes (formatting, etc.)
- `refactor:` - Code refactoring
- `test:` - Adding or updating tests
- `chore:` - Maintenance tasks

### Code Style

- Follow standard Go conventions
- Run `go fmt` before committing
- Ensure all tests pass
- Add tests for new features

### Development Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/hacomono-lib/zengin-go.git
   cd zengin-go
   ```

2. Initialize submodules:
   ```bash
   make init-submodule
   ```

3. Install development tools:
   ```bash
   make install-tools
   ```

4. Run tests:
   ```bash
   make test
   ```

5. Run all checks:
   ```bash
   make ci
   ```

6. Run example:
   ```bash
   make example
   ```

### Available Make Targets

#### Local Development
- `make test` - Run tests with race detector and coverage
- `make test-cover` - Run tests and open coverage report in browser
- `make lint` - Run golangci-lint
- `make security` - Run security scan with gosec
- `make fmt` - Format code
- `make vet` - Run go vet
- `make ci` - Run all checks (recommended before committing)
- `make install-tools` - Install development tools
- `make help` - Show all available commands

#### Docker Development
- `make docker-build` - Build Docker images
- `make docker-dev` - Start development environment in Docker
- `make docker-test` - Run tests in Docker
- `make docker-lint` - Run linter in Docker
- `make docker-security` - Run security scan in Docker
- `make docker-example` - Run example in Docker
- `make docker-clean` - Clean up Docker containers and volumes

### Container Development

If you prefer to develop in a containerized environment:

1. Using Docker Compose:
   ```bash
   make docker-dev
   ```

2. Using VS Code Dev Containers:
   - Install the "Dev Containers" extension
   - Reopen the project in a container
   - All tools will be automatically installed

## License

By contributing to zengin-go, you agree that your contributions will be licensed under the MIT License.

