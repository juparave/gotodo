# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.2] - 2025-12-30

### Changed
- Display todo list by default when running `gotodo` without arguments if a `.gotodo.json` file exists.

## [0.0.1] - 2025-09-30

### Added
- **Git Repository Integration**: Automatically uses `.gotodo.json` at the Git repository root when inside a repo, otherwise in the current directory
- **Colored CLI Output**: Beautiful terminal interface using Lipgloss library
- **Multi-word Tasks**: Support for tasks with spaces in todo descriptions
- **Relative Timestamps**: Show "Xh Ym ago" for open todos with `--long` flag
- **Flexible Commands**:
  - `gotodo init` - Initialize a todo file
  - `gotodo add "task description"` - Add a new todo
  - `gotodo list [--long] [--done-limit N]` - List todos with optional timestamps
  - `gotodo done <id|index>` - Mark todo as completed
  - `gotodo rm <id|index> [--force]` - Remove a todo
- **Cross-platform Support**: Works on Linux, macOS, and Windows
- **Comprehensive Documentation**: README with installation, usage, and examples
- **Unit Tests**: Test coverage for core functionality
- **MIT License**: Open source licensing

### Technical Details
- Built with Go 1.24.2
- Uses Lipgloss for terminal styling
- Git-aware file discovery using `git rev-parse --show-toplevel`
- JSON-based storage with atomic writes for data safety
- Command-line argument parsing with flag support
- Relative time formatting for human-readable timestamps

### Known Limitations
- Initial release with core functionality
- No web sync or advanced features yet
- CLI-only interface (no GUI or web UI)

### Installation
```bash
# Via Go install
go install github.com/juparave/gotodo/cmd/gotodo@latest

# From source
git clone https://github.com/juparave/gotodo.git
cd gotodo
go install ./cmd/gotodo
```

### Usage Examples
```bash
gotodo init
gotodo add "Write documentation"
gotodo list --long
gotodo done 1
gotodo rm 2
```