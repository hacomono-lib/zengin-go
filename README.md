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

The library automatically preloads all bank and branch data at startup. Just import and use:

```go
package main

import (
    "fmt"
    "github.com/hacomono-lib/zengin-go"
)

func main() {
    // Get all banks
    banks := zengin.AllBanks()
    fmt.Printf("Total banks: %d\n", len(banks))

    // Get bank by code
    bank, err := zengin.FindBank("0001")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Bank: %s\n", bank.Name)
    fmt.Printf("Bank has %d branches\n", len(bank.Branches))

    // Get branch by bank code and branch code
    branch, err := zengin.FindBranch("0001", "001")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Branch: %s\n", branch.Name)
    
    // Branch has reference to Bank (bidirectional relationship)
    fmt.Printf("Branch's bank: %s\n", branch.Bank.Name)
    
    // Get all branches for a bank
    allBranches, err := zengin.AllBranches("0001")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Total branches: %d\n", len(allBranches))
}
```

## Features

- 🚀 Zero external dependencies (data is embedded using go:embed)
- 📦 Full support for all Japanese banks and branches
- 🔄 Bidirectional relationship between Bank and Branch
- 🎯 Simple API - just import and use (similar to zengin-rb)
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
    Code     string                 `json:"code"`               // Bank code
    Name     string                 `json:"name"`               // Bank name
    Kana     string                 `json:"kana"`               // Katakana
    Hira     string                 `json:"hira"`               // Hiragana
    Roma     string                 `json:"roma"`               // Romaji
    Branches map[string]*Branch `json:"branches,omitempty"` // Branches (key: branch code)
}

type Branch struct {
    Bank *Bank `json:"-"`    // Reference to parent bank (bidirectional)
    Code string    `json:"code"` // Branch code
    Name string    `json:"name"` // Branch name
    Kana string    `json:"kana"` // Katakana
    Hira string    `json:"hira"` // Hiragana
    Roma string    `json:"roma"` // Romaji
}
```

### Functions

These functions work with a preloaded global instance:

- `AllBanks() map[string]*Bank` - Get all banks
- `FindBank(code string) (*Bank, error)` - Find bank by code
- `FindBranch(bankCode, branchCode string) (*Branch, error)` - Find branch by bank code and branch code
- `AllBranches(bankCode string) (map[string]*Branch, error)` - Get all branches for a bank

### Advanced: Custom Instance

For advanced use cases (e.g., testing, custom data sources), you can create your own instance:

```go
z, err := zengin.New()
if err != nil {
    panic(err)
}
bank, _ := z.FindBank("0001")
```

The instance methods mirror the package-level functions. See [GoDoc](https://pkg.go.dev/github.com/hacomono-lib/zengin-go) for details.

## Data

This project depends heavily on the following project. Big thanks to zengin-code community.

- https://github.com/zengin-code/source-data

The source data is automatically updated daily via GitHub Actions.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License

