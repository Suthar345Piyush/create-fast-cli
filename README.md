# ⚡ create-fast-cli

> Scaffold production-ready Go CLI projects in seconds — answer few questions, get a fully wired, opinionated Go CLI boilerplate with your chosen framework, features, and it automatically open the IDE.

```bash
go run github.com/Suthar345Piyush/create-fast-cli/cli@latest
```

---

## Table of Contents

- [What it does](#what-it-does)
- [App Types & Supported Tech Stack](#app-types--supported-tech-stack)
- [Overall Tech Stack](#overall-tech-stack)
- [Getting Started](#getting-started)
- [Folder Structure](#folder-structure)

---

## What it does

`create-fast-cli` is a CLI scaffolding tool — think `create-react-app` but for Go CLI applications. Run one command, answer a short wizard, and walk away with a project that already has:

- A working CLI skeleton (Cobra or urfave/cli)
- Structured logging, config file support, shell completions
- A Bubbletea TUI layer with Lipgloss styling
- Testify test suite, GitHub Actions CI, and GoReleaser config
- Your project opened in VS Code or Cursor automatically

---

## App Types & Supported Tech Stack

Each app type generates a different set of commands and starter logic. All types support both **Cobra** and **urfave/cli** frameworks unless noted.

| App Type | Generated Commands | CLI Framework | TUI | Logging | Config | Testing |
|---|---|---|---|---|---|---|
| **Dev Tool** | `run` | Cobra · urfave/cli | Bubbletea + Lipgloss | Uber Zap | Viper | Testify |
| **Git Client** | `status`, `log`, `diff` | Cobra · urfave/cli | Bubbletea + Lipgloss | Uber Zap | Viper | Testify |
| **File Explorer** | `ls`, `find`, `explore` | Cobra · urfave/cli | Bubbletea + Lipgloss | Uber Zap | Viper | Testify |
| **Kubernetes Tool** | `pods`, `describe`, `cluster` | Cobra · urfave/cli | Bubbletea + Lipgloss | Uber Zap | Viper | Testify |
| **AI Assistant** | `chat`, `ask` | Cobra · urfave/cli | Bubbletea + Lipgloss | Uber Zap | Viper | Testify |
| **System Monitor** | `cpu`, `mem`, `disk`, `proc` | Cobra · urfave/cli | Bubbletea + Lipgloss | Uber Zap | Viper | Testify |
| **Package Manager** | `install`, `remove`, `update`, `list`, `search` | Cobra · urfave/cli | Bubbletea + Lipgloss | Uber Zap | Viper | Testify |

> All optional features (TUI, Logging, Config, Completions, Testing) can be toggled on or off individually during the wizard.

---

## Overall Tech Stack

These are the libraries and tools used to **build** `create-fast-cli` itself — not the generated projects.

| Layer | Technology | Purpose |
|---|---|---|
| **Language** | Go 1.26 | Core language |
| **CLI Framework** | [Cobra](https://github.com/spf13/cobra) v1.10.2 | `create` command and root command |
| **Interactive Prompts** | [huh](https://github.com/charmbracelet/huh) v2.0.3 | 13-question wizard UI |
| **Terminal UI** | [Bubbletea](https://github.com/charmbracelet/bubbletea) v2.0.7 | Progress spinner during scaffold |
| **Styling** | [Lipgloss](https://github.com/charmbracelet/lipgloss) v2.0.3 | Summary box, step colours, badges |
| **Config** | [Viper](https://github.com/spf13/viper) v1.21.0 | Persists preferred IDE & framework to `~/.fastcli.yaml` |
| **Logging** | [Uber Zap](https://github.com/uber-go/zap) v1.28.0 | Structured debug/info logs (`--verbose`) |
| **Templates** | `text/template` + `embed.FS` | Renders `.tmpl` files into real Go source at generation time |
| **Filesystem** | `os` · `path/filepath` | Writes rendered files to disk |
| **IDE Detection** | `os/exec` | Detects and launches VS Code / Cursor |
| **Testing** | [Testify](https://github.com/stretchr/testify) v1.9.0 | Unit + golden-file tests for renderer, writer, IDE, config |
| **Packaging** | [GoReleaser](https://goreleaser.com) | Multi-arch binary releases (Linux, macOS, Windows × amd64/arm64) |
| **CI/CD** | GitHub Actions | Lint, test matrix, and release pipeline |
| **Linting** | golangci-lint | `errcheck`, `staticcheck`, `govet`, `revive`, `misspell` and more |

---

## Getting Started

### Prerequisites

- Go **1.25** or later — [install Go](https://go.dev/dl/)
- Git

### Option 1 — Run directly (no clone needed)

```bash
go run github.com/Suthar345Piyush/create-fast-cli/cli@latest
```

This downloads, compiles, and runs the latest version in one step. No installation required.


### Option 2 — Clone and run locally

```bash
# 1. Clone the repository
git clone https://github.com/Suthar345Piyush/create-fast-cli.git
cd cli

# 2. Install dependencies
go mod tidy

# 3. Run the wizard
go run . create

# 4. (Optional) Build a local binary
go build -o create-fast-cli ./cli
./create-fast-cli create
```

### Wizard walkthrough

When you run the command you'll be asked 13 questions:

```
⚡ FastCLI Starter
  A CLI ecosystem starter for Go — answer a few questions to scaffold your project.

  1.  Project name          e.g. my-dev-tool
  2.  Go module path        e.g. github.com/you/my-dev-tool
  3.  What are you building? Dev tool / Git client / File explorer / ...
  4.  CLI framework          Cobra (recommended) · urfave/cli
  5.  Include TUI?           yes / no
  6.  Include logging?       yes / no
  7.  Include config?        yes / no
  8.  Include completions?   yes / no
  9.  Include testing?       yes / no
  10. Output directory       e.g. ~/projects
  11. Open in IDE?           VS Code · Cursor · Don't open
  12. Review your choices    — confirm before writing any files
```

After confirmation, you'll see:

```
  ✅ Validating output directory
  ✅ Rendering templates
  ✅ Writing project files
  ✅ Saving preferences
  ✅ Opening IDE

  ✓ Project "my-dev-tool" created successfully!

  📁 Location:
       /Users/you/projects/my-dev-tool

  🚀 Next steps:
       cd "my-dev-tool"
       go mod tidy
       go run . --help
```

### Running tests

```bash
# Run all tests
go test ./... -v -race

# Run only scaffold tests
go test ./test/... -v -run TestRender
go test ./test/... -v -run TestWrite

# Run only prompt/config tests
go test ./test/... -v -run TestDefaultConfig
go test ./test/... -v -run TestValidate

# With coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## Folder Structure

```
create-fast-cli/
│
├── main.go                          # Binary entrypoint → calls cmd.Execute()
├── go.mod
├── go.sum
├── .goreleaser.yaml                 # Multi-arch release config
├── .golangci.yml                    # Linter config
│
├── cmd/                             # Cobra commands (thin — all logic in internal/)
│   ├── root.go                      # Root command, --verbose flag, Viper init
│   └── create.go                    # `create` subcommand → runs wizard + scaffold
│
├── internal/
│   │
│   ├── config/                      # Central data model
│   │   ├── schema.go                # ProjectConfig struct, AppType/Framework/IDE constants
│   │   ├── defaults.go              # DefaultConfig(), label helpers
│   │   └── viper.go                 # Read/write ~/.fastcli.yaml preferences
│   │
│   ├── prompt/                      # Interactive wizard (huh + Bubbletea)
│   │   ├── questions.go             # All huh form group definitions + validators
│   │   ├── survey.go                # Orchestrates the form flow, prints summary
│   │   ├── tui.go                   # Bubbletea ProgressModel (spinner + step list)
│   │   ├── progress.go              # PrintStep / PrintStepOK / PrintStepErr helpers
│   │   └── styles.go                # Lipgloss colour tokens and style definitions
│   │
│   ├── scaffold/                    # Project generation pipeline
│   │   ├── generator.go             # Orchestrates all steps, prints live progress
│   │   ├── renderer.go              # Walks embed FS, runs text/template → []RenderedFile
│   │   ├── writer.go                # Writes []RenderedFile to disk (os + filepath)
│   │   └── ide.go                   # Detects + launches VS Code / Cursor (os/exec)
│   │
│   ├── templates/                   # Embedded template filesystem
│   │   ├── embed.go                 # //go:embed all:files — exposes FS to renderer
│   │   ├── loader.go                # SubFS() resolves common + typed sub-trees
│   │   └── files/
│   │       ├── common/              # Shared across ALL app types and frameworks
│   │       │   ├── .gitignore.tmpl
│   │       │   ├── README.md.tmpl
│   │       │   ├── Makefile.tmpl
│   │       │   └── .github/
│   │       │       └── workflows/
│   │       │           ├── ci.yml.tmpl
│   │       │           └── release.yml.tmpl
│   │       │
│   │       ├── dev-tool/
│   │       │   ├── cobra/           # main.go, go.mod, cmd/, internal/config/, internal/logger/
│   │       │   └── urfave/          # main.go, go.mod, cmd/, internal/config/, internal/logger/
│   │       │
│   │       ├── git-client/
│   │       │   ├── cobra/           # status.go, log.go, diff.go + internals
│   │       │   └── urfave/          # git.go (status/log/diff) + internals
│   │       │
│   │       ├── ai-assistant/
│   │       │   ├── cobra/           # chat.go, ask.go + internals
│   │       │   └── urfave/          # ai.go (chat/ask) + internals
│   │       │
│   │       ├── file-explorer/
│   │       │   ├── cobra/           # ls.go, find.go + internals
│   │       │   └── urfave/          # fs.go (ls/find/explore) + internals
│   │       │
│   │       ├── k8s-tool/
│   │       │   ├── cobra/           # pods.go + internals
│   │       │   └── urfave/          # k8s.go (pods/cluster/describe) + internals
│   │       │
│   │       ├── system-monitor/
│   │       │   ├── cobra/           # cpu.go, disk.go + internals
│   │       │   └── urfave/          # monitor.go (cpu/mem/disk/proc) + internals
│   │       │
│   │       └── pkg-manager/
│   │           ├── cobra/           # remove.go (remove/list/update) + internals
│   │           └── urfave/          # pkg.go (install/remove/update/list/search) + internals
│   │
│   ├── completions/
│   │   └── completions.go           # bash / zsh / fish completion via cobra.GenX
│   │
│   └── logger/
│       └── logger.go                # Uber Zap wrapper used by create-fast-cli itself
│
├── pkg/
│   └── utils/
│       ├── strings.go               # SlugFunc, PascalCase, CamelCase, TrimExt, IsEmpty
│       └── fs.go                    # MkdirAll, WriteFile, CopyFile, DirIsEmpty
│
├── test/
│   ├── scaffold_test.go             # Renderer, writer, IDE, EnsureDir tests (Testify)
│   └── prompt_test.go               # Config defaults, labels, validators, utils (Testify)
│
└── .github/
    └── workflows/
        ├── ci.yml                   # Lint + test matrix (Go 1.25/1.26 × Linux/macOS/Windows)
        └── release.yml              # GoReleaser fires on v* tags
```

---

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Write tests if applicable
5. Submit a pull request


**Made with ❤️ by Piyush Suthar**

[![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Suthar345Piyush)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/piyush-suthar-641a0826a/)
[![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://x.com/piyushtwtz)