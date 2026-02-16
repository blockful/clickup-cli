# clickup-cli

A command-line interface for ClickUp, optimized for AI agents. All output is JSON by default.

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
# 1. Authenticate with your ClickUp API token
clickup auth login --token pk_YOUR_TOKEN

# 2. List your workspaces
clickup workspace list

# 3. List spaces in a workspace
clickup space list --workspace 1234567

# 4. List tasks in a list
clickup task list --list 900100200300
```

## Command Reference

| Command | Description |
|---------|-------------|
| `clickup auth login` | Authenticate with API token |
| `clickup auth whoami` | Show current user |
| `clickup workspace list` | List workspaces |
| `clickup space list` | List spaces |
| `clickup space get` | Get space details |
| `clickup space create` | Create a space |
| `clickup folder list` | List folders |
| `clickup folder get` | Get folder details |
| `clickup folder create` | Create a folder |
| `clickup list list` | List lists |
| `clickup list get` | Get list details |
| `clickup list create` | Create a list |
| `clickup task list` | List tasks (with filters) |
| `clickup task get` | Get task details |
| `clickup task create` | Create a task |
| `clickup task update` | Update a task |
| `clickup task delete` | Delete a task |
| `clickup comment list` | List task comments |
| `clickup comment create` | Add a comment |

## Global Flags

- `--token` — API token (overrides config file)
- `--workspace` — Default workspace ID
- `--format` — Output format: `json` (default) or `text`
- `--verbose` — Enable verbose output

## Configuration

Config is stored in `~/.clickup-cli.yaml`:

```yaml
token: pk_12345...
workspace: "1234567"
```

## Output Format

All commands output valid JSON. Errors are formatted as:

```json
{
  "error": "description of what went wrong",
  "code": "ERROR_CODE"
}
```

## Documentation

- [Architecture](docs/architecture.md)
- [API Mapping](docs/api.md)
- [Decision Records](docs/decisions/)

## License

MIT — see [LICENSE](LICENSE)
