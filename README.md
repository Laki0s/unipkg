# 🚀 unipkg

![GitHub release (latest by date)](https://img.shields.io/github/v/release/your-org/unipkg)
![License: MIT](https://img.shields.io/badge/License-MIT-green)

**Unified Package Manager CLI for Linux**

`unipkg` provides a single, consistent command‑line interface for installing, removing, updating, and searching packages across multiple Linux distributions and packaging systems (APT, DNF, Pacman, Snap, Flatpak, etc.).

---

## 📋 Table of Contents

1. [✨ Features](#-features)
2. [⚙️ Prerequisites](#️-prerequisites)
3. [📥 Installation](#-installation)
4. [🛠️ Usage](#️-usage)
   - [Global Options](#global-options)
   - [Commands](#commands)
5. [🔧 Configuration](#-configuration)
6. [👩‍💻 Development](#-development)
   - [Project Structure](#project-structure)
   - [Building from Source](#building-from-source)
   - [Running Tests](#running-tests)
   - [🔀 Cross‑Compilation](#-cross‑compilation)
7. [🤝 Contributing](#-contributing)
8. [📜 License](#-license)
9. [👥 Authors & Acknowledgments](#-authors--acknowledgments)

---

## ✨ Features

- 🔍 **Automatic backend detection** based on `/etc/os-release` & available binaries
- 🔄 **Fallback** to universal systems (Snap, Flatpak) if native package manager is absent
- 🎛️ **Unified commands**: `install`, `remove`, `update`, `search`
- 🌱 **Modular architecture**: add new backends (Homebrew, Chocolatey, etc.) easily
- 🗄️ **Extensible config**: map distributions to preferred backends via YAML
- 🏋️ **Minimal dependencies**: single static binary (no external runtime)
- 🐞 **Verbose logging** and debug mode for troubleshooting

## ⚙️ Prerequisites

- A Linux distribution with one or more package managers installed: `apt`, `dnf`, `pacman`, `snap`, `flatpak`, etc.
- **Go 1.21+** (for building from source)
- `git` (for development)

## 📥 Installation

### 🏷️ Pre‑built Releases

Download the latest binary from [GitHub Releases](https://github.com/your-org/unipkg/releases):
```bash
chmod +x unipkg-linux-amd64           # Make it executable
sudo mv unipkg-linux-amd64 /usr/local/bin/unipkg   # Move to $PATH
```

### 🍺 Homebrew (macOS / Linuxbrew)
```bash
brew install your-org/tap/unipkg
```

### 🐙 Snap
```bash
sudo snap install unipkg --edge --classic
```

### 📦 From Source
```bash
git clone https://github.com/your-org/unipkg.git
cd unipkg
./scripts/build.sh                     # Build script
# Binary in bin/unipkg
sudo mv bin/unipkg /usr/local/bin/
```

## 🛠️ Usage

```bash
unipkg [command] [options]
```

### Global Options

| Flag           | Description                        |
| -------------- | ---------------------------------- |
| `-h`, `--help` | Show help for `unipkg` or subcommand |
| `--verbose`    | Enable debug logging               |

### Commands

| Command                                   | Description                                |
| ----------------------------------------- | ------------------------------------------ |
| `install <pkg>`                           | 🎁 Install a package                        |
| `remove <pkg>`                            | 🗑️ Remove a package                         |
| `update [pkg]`                            | 🔄 Update all packages or a specific one    |
| `search <term>`                           | 🔎 Search for a package                     |
| `completion [bash|zsh|fish|powershell]`  | ⚙️ Generate shell completion script         |

#### Examples

```bash
# Install curl
unipkg install curl

# Remove curl
unipkg remove curl

# Update everything
unipkg update

# Update git only
unipkg update git

# Search for "docker"
unipkg search docker
```

## 🔧 Configuration

You can override the default distribution-to-backend mapping by setting:
```bash
export UNIPKG_BACKENDS_CONFIG=/path/to/custom_backends.yaml
```
The YAML file format is:
```yaml
ubuntu: apt
# etc.
``` 

## 👩‍💻 Development

### Project Structure

```plaintext
unipkg/
├── cmd/unipkg     # CLI entrypoint and commands (Cobra)
├── drivers/       # Backend implementations
├── pkg/distro     # Distro detection utils
├── internal/      # Private utilities (exec, logger)
├── configs/       # Default YAML mappings
├── scripts/       # Build & test helper scripts
├── .github/       # CI/CD workflows
├── Dockerfile     # Containerized integration test
├── go.mod         # Go module definition
└── README.md      # This file
```

### Building from Source
```bash
go mod tidy
./scripts/build.sh
```

### Running Tests
```bash
go test ./pkg/distro
go test ./drivers
```

### 🔀 Cross‑Compilation

Using GoReleaser:
```bash
go install github.com/goreleaser/goreleaser@latest
goreleaser release --rm-dist
```

## 🤝 Contributing

Contributions welcome!
1. Fork repo
2. Create branch: `git checkout -b feat/your-feature`
3. Commit changes: `git commit -m "feat: description"`
4. Push branch: `git push origin feat/your-feature`
5. Open PR

Please follow code style and include tests.

## 📜 License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.

## 👥 Authors & Acknowledgments

- **Laki0s_** – initial work
- Inspired by various unified package manager wrappers

🙏 Thank you for using `unipkg`! Feel free to open issues or questions.

