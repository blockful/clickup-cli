# clickup-cli

A production-quality command-line interface for the **complete ClickUp API** — 135+ endpoints, every parameter exposed as a CLI flag. Designed for AI agents and automation. All output is JSON by default.

[![Go](https://github.com/blockful/clickup-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/blockful/clickup-cli/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Features

- **Full API coverage** — 135+ commands spanning 27 resource groups
- **JSON-first output** — every command outputs valid JSON; errors are structured `{"error":"...","code":"..."}`
- **AI-agent optimized** — no interactive prompts, deterministic output, `--dry-run` on destructive ops
- **Every flag documented** — see [docs/api.md](docs/api.md) for complete flag→API parameter mapping
- **v3 Docs API** — full support for ClickUp Docs with page CRUD

## Installation

```bash
go install github.com/blockful/clickup-cli@latest
```

Or build from source:

```bash
git clone https://github.com/blockful/clickup-cli.git
cd clickup-cli
go build -o clickup ./
```

## Quick Start

```bash
# Authenticate
clickup auth login --token pk_YOUR_TOKEN

# Explore your workspace
clickup workspace list
clickup space list --workspace 1234567
clickup task list --list 900100200300

# Create a task
clickup task create --list 900100200300 --name "Ship feature" --priority 2 --status "in progress"

# Search across workspace
clickup task search --workspace 1234567 --assignee 12345 --include-closed

# Track time
clickup time-entry start --workspace 1234567 --task abc123 --description "Working on feature"
clickup time-entry stop --workspace 1234567

# Upload a file
clickup attachment create --task-id abc123 --file ./screenshot.png
```

## Command Reference

27 top-level command groups, 135+ total commands. For **complete flag documentation** with types, defaults, and API parameter mappings, see **[docs/api.md](docs/api.md)**.

### Core Hierarchy

| Command | Subcommands | Description |
|---------|-------------|-------------|
| `workspace` | `list`, `plan`, `seats` | List workspaces, get plan & seat info |
| `space` | `list`, `get`, `create`, `update`, `delete` | Manage spaces |
| `folder` | `list`, `get`, `create`, `update`, `delete` | Manage folders |
| `list` | `list`, `get`, `create`, `update`, `delete` | Manage lists |

### Tasks

| Command | Subcommands | Description |
|---------|-------------|-------------|
| `task` | `list`, `get`, `create`, `update`, `delete`, `search` | Full task CRUD + workspace search |
| `task` | `add-to-list`, `remove-from-list`, `merge`, `time-in-status` | Multi-list, merge, status timing |
| `task dependency` | `add`, `remove` | Task dependency management |
| `task link` | `add`, `remove` | Task link management |

### Content & Collaboration

| Command | Subcommands | Description |
|---------|-------------|-------------|
| `comment` | `list`, `create`, `update`, `delete` | Task/list/view comments |
| `comment reply` | `list`, `create` | Threaded comment replies |
| `doc` | `list`, `get`, `create` | ClickUp Docs (v3 API) |
| `doc` | `page-list`, `page-get`, `page-create`, `page-update` | Doc page CRUD |
| `checklist` | `create`, `update`, `delete` | Task checklists |
| `checklist-item` | `create`, `update`, `delete` | Checklist items |
| `attachment` | `create` | File uploads to tasks |

### Custom Fields & Tags

| Command | Subcommands | Description |
|---------|-------------|-------------|
| `custom-field` | `list`, `set`, `remove` | Custom field values |
| `tag` | `list`, `create`, `update`, `delete`, `add`, `remove` | Space tags & task tagging |
| `custom-task-type` | `list` | List custom task types |

### Time Tracking

| Command | Subcommands | Description |
|---------|-------------|-------------|
| `time-entry` | `list`, `get`, `create`, `update`, `delete` | Time entry CRUD |
| `time-entry` | `start`, `stop`, `current` | Timer controls |
| `time-entry` | `history` | Time entry change history |
| `time-entry legacy` | `list`, `create`, `update`, `delete` | Task-level time tracking (legacy) |
| `time-entry tag` | `add`, `remove`, `update` | Time entry tag management |

### Views & Goals

| Command | Subcommands | Description |
|---------|-------------|-------------|
| `view` | `list`, `get`, `create`, `update`, `delete`, `tasks` | View CRUD + task retrieval |
| `goal` | `list`, `get`, `create`, `update`, `delete` | Goal management |
| `goal key-result` | `create`, `update`, `delete` | Key result CRUD |

### People & Access

| Command | Subcommands | Description |
|---------|-------------|-------------|
| `user` | `invite`, `get`, `update`, `remove` | Workspace user management |
| `member` | `list` | List members of list/task |
| `group` | `list`, `create`, `update`, `delete` | User group management |
| `guest` | `invite`, `get`, `edit`, `remove` | Guest workspace access |
| `guest` | `add-to-task`, `add-to-list`, `add-to-folder` | Grant guest access to resources |
| `guest` | `remove-from-task`, `remove-from-list`, `remove-from-folder` | Revoke guest resource access |
| `role` | `list` | List custom roles |

### Infrastructure

| Command | Subcommands | Description |
|---------|-------------|-------------|
| `webhook` | `list`, `create`, `update`, `delete` | Webhook management |
| `template` | `list`, `create-task`, `create-list`, `create-folder` | Template management |
| `shared` | `list` | Shared hierarchy |
| `auth` | `login`, `whoami` | Authentication |

## Global Flags

| Flag | Description |
|------|-------------|
| `--token` | API token (overrides config file and `CLICKUP_TOKEN` env) |
| `--workspace` | Default workspace ID (overrides config) |
| `--format` | Output format: `json` (default) or `text` |
| `--verbose` | Enable verbose output (to stderr) |

## Configuration

Config stored in `~/.clickup-cli.yaml`:

```yaml
token: pk_12345...
workspace: "1234567"
```

Precedence: CLI flags > environment variables (`CLICKUP_TOKEN`) > config file.

## Output Format

All commands output valid JSON. Errors:

```json
{"error": "description of what went wrong", "code": "ERROR_CODE", "status": 400}
```

## Documentation

- **[API Reference](docs/api.md)** — Every command, every flag, every API mapping
- **[Architecture](docs/architecture.md)** — Codebase structure and design
- **[Business Rules](docs/business-rules.md)** — Invariants and conventions
- **[Decision Records](docs/decisions/)** — ADRs for key decisions
- **[Contributing](CONTRIBUTING.md)** — How to contribute
- **[Agents Guide](AGENTS.md)** — For AI agents using this CLI

## License

MIT — see [LICENSE](LICENSE)
