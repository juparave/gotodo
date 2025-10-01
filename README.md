# gotodo

Minimal todo CLI proof-of-concept.

Quick start:

Initialize a todo file in the current directory:

```bash
gotodo init
```

Add an item:

```bash
gotodo add "Buy milk"
```

List items (uses colors via lipgloss):

```bash
gotodo list
```

Build & install
-----------------

The program entrypoint is `cmd/gotodo/main.go`.

Create a local executable named `gotodo` (recommended for development):

```bash
# from the repository root
go build -o gotodo ./cmd/gotodo
```

You can then run the binary directly from the project root:

```bash
./gotodo list
```

Install so the binary is available on your $PATH:

# If you're inside this module (recommended for development):
```bash
go install ./cmd/gotodo
```

# If you want to install from elsewhere (module-aware install):
```bash
go install github.com/juparave/gotodo/cmd/gotodo@latest
```

Or build and move into `/usr/local/bin` on macOS (requires sudo):

```bash
go build -o gotodo ./cmd/gotodo
sudo mv gotodo /usr/local/bin/gotodo
```

Notes
- If `go install` doesn't put the binary on your PATH, ensure `$GOBIN` or `$GOPATH/bin` is in your `$PATH`.
- This project uses Go modules; a working Go toolchain (>= 1.20) is required.
