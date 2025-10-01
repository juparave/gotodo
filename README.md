# gotodo

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/juparave/gotodo)](https://goreportcard.com/report/github.com/juparave/gotodo)
[![Build Status](https://github.com/juparave/gotodo/workflows/CI/badge.svg)](https://github.com/juparave/gotodo/actions)

A simple, filesystem-aware todo CLI tool written in Go. Manage your todos with colors, Git repository integration, and per-project lists.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Features

- ✅ **Git-aware**: Automatically uses the repository root for todo files when inside a Git repo
- 🎨 **Colored output**: Beautiful CLI interface with [Lipgloss](https://github.com/charmbracelet/lipgloss)
- 📁 **Per-directory/project**: Todo lists are scoped to directories or Git repos
- 🔍 **Discovery**: Commands automatically find the appropriate todo file
- 🛠️ **Simple commands**: Init, add, list, done, remove with intuitive flags
- 📝 **Multi-word tasks**: Support for tasks with spaces
- ⚡ **Fast**: Lightweight and efficient

## Installation

### Prerequisites

- Go 1.20 or later
- Git (for repository detection)

### Install from source

Clone the repository:

```bash
git clone https://github.com/juparave/gotodo.git
cd gotodo
```

Build and install:

```bash
go install ./cmd/gotodo
```

This installs `gotodo` to your `$GOBIN` or `$GOPATH/bin` directory.

### Install with go install

```bash
go install github.com/juparave/gotodo/cmd/gotodo@latest
```

### Manual build

```bash
# From repository root
go build -o gotodo ./cmd/gotodo

# Move to PATH (example for macOS/Linux)
sudo mv gotodo /usr/local/bin/
```

## Usage

gotodo uses a `.gotodo.json` file:
- **Inside Git repos**: At the repository root
- **Outside Git repos**: In the current directory

### Commands

```bash
gotodo <command> [arguments]
```

| Command | Description | Flags |
|---------|-------------|-------|
| `init` | Initialize a new todo file | - |
| `add <task>` | Add a new todo item | - |
| `list` | List all todos | `--done-limit`, `--long`, `--file` |
| `done <id\|n>` | Mark todo as done (id or open index) | `--file` |
| `rm <id\|n>` | Remove a todo (id or open index) | `--force`, `--yes`, `--file` |

### Global Flags

- `--file <path>`: Override the default todo file path

## Examples

### Basic workflow

```bash
# Initialize (only needed once per project/repo)
gotodo init

# Add some todos
gotodo add "Write documentation"
gotodo add "Implement feature X"
gotodo add "Fix bug in parser"

# List todos
gotodo list

# Mark first open todo as done
gotodo done 1

# Remove a todo (with confirmation)
gotodo rm 2

# Force remove without confirmation
gotodo rm 3 --force
```

### Advanced usage

```bash
# Show only last 5 done todos
gotodo list --done-limit 5

# Show timestamps for done todos
gotodo list --long

# Use a specific file
gotodo list --file /path/to/custom/.gotodo.json

# Work from any subdirectory (Git repo aware)
cd subdir/
gotodo add "Subdir task"  # Adds to repo root .gotodo.json
```

### Sample output

```
Todos

Open:
  1. Write documentation
  2. Implement feature X

Done:
  1. Fix bug in parser  (2025-09-30 10:30:00)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development

```bash
# Clone and setup
git clone https://github.com/juparave/gotodo.git
cd gotodo

# Run tests
go test ./...

# Build
go build ./cmd/gotodo

# Format code
gofmt -s -w .
goimports -w .
go vet ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
