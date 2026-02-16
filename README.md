# clickup-cli

A production-quality command-line interface for the **complete ClickUp API** — 134/135 endpoints (99.3% coverage), every parameter exposed as a CLI flag. Built for AI agents and automation. JSON output by default.

[![Go](https://github.com/blockful/clickup-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/blockful/clickup-cli/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Why This Exists

Most ClickUp integrations are partial. This CLI wraps **every endpoint** so AI agents can manage ClickUp programmatically — no SDK, no HTTP wrangling, no missing features. Every command returns valid JSON, has no interactive prompts, and supports `--dry-run` on destructive operations.

## Features

- **99.3% API coverage** — 134 of 135 ClickUp API endpoints across 27 resource groups
- **JSON-first output** — every command outputs valid JSON; errors are structured `{"error":"...","code":"..."}`
- **AI-agent optimized** — no interactive prompts, deterministic output, machine-parseable
- **Every flag documented** — see [docs/api.md](docs/api.md) for complete flag→API parameter mapping
- **Markdown content** — `--markdown-content` / `--markdown-description` for rich task descriptions
- **Custom task IDs** — `--custom-task-ids` + `--team-id` for human-readable task references
- **v3 Docs API** — full support for ClickUp Docs with page CRUD

## Installation

### Go install (recommended)

```bash
go install github.com/blockful/clickup-cli@latest
```

### Binary releases

Download pre-built binaries from [GitHub Releases](https://github.com/blockful/clickup-cli/releases) for Linux, macOS, and Windows (amd64 + arm64).

### Build from source

```bash
git clone https://github.com/blockful/clickup-cli.git
cd clickup-cli
go build -o clickup ./
```

## Quick Start

```bash
# 1. Authenticate (saves token to ~/.clickup-cli.yaml)
clickup auth login --token pk_YOUR_TOKEN
clickup auth whoami  # verify

# 2. Explore your workspace
clickup workspace list
clickup space list --workspace 1234567
clickup folder list --space 5678
clickup list list --folder 9012

# 3. Create a task with markdown description
clickup task create --list 900100200300 \
  --name "Ship feature" \
  --priority 2 \
  --status "in progress" \
  --assignee 12345 \
  --markdown-content "## Requirements\n- Fast\n- Reliable"

# 4. Search across workspace
clickup task search --workspace 1234567 --assignee 12345 --include-closed

# 5. Track time
clickup time-entry start --workspace 1234567 --task abc123 --description "Working on feature"
clickup time-entry stop --workspace 1234567

# 6. Upload a file
clickup attachment create --task-id abc123 --file ./screenshot.png

# 7. Human-readable output (for debugging)
clickup task list --list 900100200300 --format text
```

## Command Reference

27 top-level command groups, 134+ total commands. For **complete flag documentation** with types, defaults, and API parameter mappings, see **[docs/api.md](docs/api.md)**.

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
| `task` | `add-to-list`, `remove-from-list` | Multi-list task management |
| `task` | `merge`, `time-in-status` | Merge tasks, get status timing |
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
| `custom-field` | `list`, `set`, `remove` | Custom field values (workspace/space/folder/list scope) |
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

**Precedence:** CLI flags > environment variables (`CLICKUP_TOKEN`) > config file.

## Output Format

**Default: JSON.** Every command outputs valid JSON to stdout. Use `--format text` for human-readable output.

**Success:** raw JSON from the ClickUp API (object or array).

**Error:**
```json
{"error": "task not found", "code": "NOT_FOUND", "status": 404}
```

**Exit codes:** 0 = success, non-zero = error.

## Key Features for Agents

### Markdown Content

Create tasks with rich descriptions using `--markdown-content` or `--markdown-description`:

```bash
clickup task create --list 123 --name "Bug fix" \
  --markdown-content "## Steps to reproduce\n1. Open app\n2. Click button\n\n**Expected:** no crash"
```

Retrieve markdown descriptions with `--include-markdown` on `task get`, `task list`, and `task search`.

### Custom Task IDs

Reference tasks by custom IDs (e.g., `PROJ-123`) instead of ClickUp's internal IDs:

```bash
clickup task get --id "PROJ-123" --custom-task-ids --team-id 1234567
clickup task update --id "PROJ-123" --custom-task-ids --team-id 1234567 --status "done"
```

The `--custom-task-ids` + `--team-id` pattern works on: `task get`, `task update`, `task delete`, `task add-to-list`, `task remove-from-list`, `task merge`, `task time-in-status`, `task dependency add/remove`, `task link add/remove`, `comment create`, `custom-field set/remove`, `attachment create`, `guest add-to-task/remove-from-task`, and `time-entry legacy` commands.

## Documentation

- **[API Reference](docs/api.md)** — Every command, every flag, every API mapping
- **[Agents Guide](AGENTS.md)** — For AI agents using this CLI
- **[Architecture](docs/architecture.md)** — Codebase structure and design
- **[Business Rules](docs/business-rules.md)** — Invariants and conventions
- **[Decision Records](docs/decisions/)** — ADRs for key decisions
- **[Contributing](CONTRIBUTING.md)** — How to contribute

## License

MIT — see [LICENSE](LICENSE)
