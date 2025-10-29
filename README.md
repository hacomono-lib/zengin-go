# zengin-go

[![Go Reference](https://pkg.go.dev/badge/github.com/hacomono-lib/zengin-go.svg)](https://pkg.go.dev/github.com/hacomono-lib/zengin-go)
[![Test](https://github.com/hacomono-lib/zengin-go/workflows/Test/badge.svg)](https://github.com/hacomono-lib/zengin-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/hacomono-lib/zengin-go)](https://goreportcard.com/report/github.com/hacomono-lib/zengin-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The Go library for Zengin Code.

ZenginCode is datasets of bank codes and branch codes for Japanese financial institutions.

## Installation

```bash
go get github.com/hacomono-lib/zengin-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/hacomono-lib/zengin-go"
)

func main() {
    z, err := zengin.New()
    if err != nil {
        panic(err)
    }

    // Get bank by code
    bank, err := z.GetBank("0001")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Found bank: %s\n", bank.Name)

    // Find banks by name pattern (regex)
    banks, err := z.FindBanksByName(".*みずほ.*")
    if err != nil {
        panic(err)
    }
    for _, bank := range banks {
        fmt.Printf("Found bank: %s\n", bank.Name)
    }

    // Get branch by bank code and branch code
    branch, err := z.GetBranch("0001", "001")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Found branch: %s\n", branch.Name)

    // Find branches by name pattern (regex)
    branches, err := z.FindBranchesByName("0001", ".*本店.*")
    if err != nil {
        panic(err)
    }
    for _, branch := range branches {
        fmt.Printf("Found branch: %s\n", branch.Name)
    }
}
```

## Features

- 🚀 Zero external dependencies (data is embedded using go:embed)
- 📦 Full support for all Japanese banks and branches
- 🔍 Powerful search capabilities with regex support
- 🧪 Comprehensive test coverage
- 🔄 Automated data updates via GitHub Actions

## Development

### Prerequisites

- Go 1.23 or later
- Make

### Setup

```bash
# Clone the repository
git clone https://github.com/hacomono-lib/zengin-go.git
cd zengin-go

# Initialize submodules
make init-submodule

# Install development tools
make install-tools
```

### Commands

```bash
# Run tests
make test

# Run tests and open coverage report
make test-cover

# Run linter
make lint

# Run security scan
make security

# Run all checks (recommended before committing)
make ci

# Run example
make example

# Update source-data submodule
make update-submodule

# Format code
make fmt

# See all available commands
make help
```

### Container Development

For developers who prefer containerized development or want to avoid polluting their local environment:

#### Using Docker Compose

```bash
# Build Docker images
make docker-build

# Start development environment
make docker-dev

# Run tests in container
make docker-test

# Run linter in container
make docker-lint

# Run security scan in container
make docker-security

# Run example in container
make docker-example

# Clean up containers and volumes
make docker-clean
```

#### Using VS Code Dev Containers

1. Install the "Dev Containers" extension in VS Code
2. Open the project in VS Code
3. Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on Mac)
4. Select "Dev Containers: Reopen in Container"

The development container includes:
- Go 1.23 with all development tools
- golangci-lint for code quality
- gosec for security scanning
- Git and GitHub CLI
- Optimized VS Code settings for Go development

### CI/CD

The project uses GitHub Actions for continuous integration:

- **Test workflow**: Runs tests on Go 1.23 and 1.24
- **Lint workflow**: Runs golangci-lint and security scanning
- **Update workflow**: Automatically updates source data daily

All checks must pass before merging pull requests.

## API Reference

### Types

```go
type Bank struct {
    Code string `json:"code"` // Bank code
    Name string `json:"name"` // Bank name
    Kana string `json:"kana"` // Katakana
    Hira string `json:"hira"` // Hiragana
    Roma string `json:"roma"` // Romaji
}

type Branch struct {
    Code string `json:"code"` // Branch code
    Name string `json:"name"` // Branch name
    Kana string `json:"kana"` // Katakana
    Hira string `json:"hira"` // Hiragana
    Roma string `json:"roma"` // Romaji
}
```

### Methods

- `New() (*Zengin, error)` - Create a new Zengin instance
- `GetBank(code string) (*Bank, error)` - Get bank by code
- `FindBanksByName(pattern string) ([]*Bank, error)` - Find banks by name pattern (regex)
- `GetBranch(bankCode, branchCode string) (*Branch, error)` - Get branch by bank code and branch code
- `FindBranchesByName(bankCode, pattern string) ([]*Branch, error)` - Find branches by name pattern (regex)
- `GetAllBanks() map[string]*Bank` - Get all banks
- `GetAllBranches(bankCode string) (map[string]*Branch, error)` - Get all branches for a bank

## Data

This project depends heavily on the following project. Big thanks to zengin-code community.

- https://github.com/zengin-code/source-data

The source data is automatically updated daily via GitHub Actions.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License

