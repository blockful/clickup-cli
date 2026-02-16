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

For **complete flag documentation** with types, defaults, and API parameter mappings, see **[docs/api.md](docs/api.md)**.

| Command | Description |
|---------|-------------|
| **Auth** | |
| `clickup auth login` | Authenticate with API token |
| `clickup auth whoami` | Show current user |
| **Workspaces** | |
| `clickup workspace list` | List workspaces |
| **Spaces** | |
| `clickup space list` | List spaces |
| `clickup space get` | Get space details |
| `clickup space create` | Create a space |
| `clickup space update` | Update a space |
| `clickup space delete` | Delete a space |
| **Folders** | |
| `clickup folder list` | List folders |
| `clickup folder get` | Get folder details |
| `clickup folder create` | Create a folder |
| `clickup folder update` | Update a folder |
| `clickup folder delete` | Delete a folder |
| **Lists** | |
| `clickup list list` | List lists |
| `clickup list get` | Get list details |
| `clickup list create` | Create a list |
| `clickup list update` | Update a list |
| `clickup list delete` | Delete a list |
| **Tasks** | |
| `clickup task list` | List tasks (with filters) |
| `clickup task get` | Get task details |
| `clickup task create` | Create a task |
| `clickup task update` | Update a task |
| `clickup task delete` | Delete a task |
| `clickup task search` | Search tasks across workspace |
| **Comments** | |
| `clickup comment list` | List comments |
| `clickup comment create` | Add a comment |
| `clickup comment update` | Update a comment |
| `clickup comment delete` | Delete a comment |
| **Docs (v3)** | |
| `clickup doc list` | List/search docs |
| `clickup doc get` | Get a doc |
| `clickup doc create` | Create a doc |
| `clickup doc page-list` | List pages in a doc |
| `clickup doc page-get` | Get a page |
| `clickup doc page-create` | Create a page |
| `clickup doc page-update` | Update a page |
| **Custom Fields** | |
| `clickup custom-field list` | List custom fields |
| `clickup custom-field set` | Set a custom field value |
| `clickup custom-field remove` | Remove a custom field value |
| **Tags** | |
| `clickup tag list` | List space tags |
| `clickup tag create` | Create a tag |
| `clickup tag update` | Update a tag |
| `clickup tag delete` | Delete a tag |
| `clickup tag add` | Add tag to task |
| `clickup tag remove` | Remove tag from task |
| **Checklists** | |
| `clickup checklist create` | Create a checklist |
| `clickup checklist update` | Update a checklist |
| `clickup checklist delete` | Delete a checklist |
| `clickup checklist-item create` | Create a checklist item |
| `clickup checklist-item update` | Update a checklist item |
| `clickup checklist-item delete` | Delete a checklist item |
| **Time Tracking** | |
| `clickup time-entry list` | List time entries |
| `clickup time-entry get` | Get a time entry |
| `clickup time-entry create` | Create a time entry |
| `clickup time-entry update` | Update a time entry |
| `clickup time-entry delete` | Delete a time entry |
| `clickup time-entry start` | Start a timer |
| `clickup time-entry stop` | Stop running timer |
| `clickup time-entry current` | Get running timer |
| **Webhooks** | |
| `clickup webhook list` | List webhooks |
| `clickup webhook create` | Create a webhook |
| `clickup webhook update` | Update a webhook |
| `clickup webhook delete` | Delete a webhook |
| **Views** | |
| `clickup view list` | List views |
| `clickup view get` | Get a view |
| `clickup view create` | Create a view |
| `clickup view update` | Update a view |
| `clickup view delete` | Delete a view |
| `clickup view tasks` | Get tasks in a view |
| **Goals** | |
| `clickup goal list` | List goals |
| `clickup goal get` | Get a goal |
| `clickup goal create` | Create a goal |
| `clickup goal update` | Update a goal |
| `clickup goal delete` | Delete a goal |
| **Members & Groups** | |
| `clickup member list` | List members (list/task) |
| `clickup group list` | List user groups |
| `clickup group create` | Create a user group |
| `clickup group delete` | Delete a user group |
| **Guests** | |
| `clickup guest invite` | Invite a guest |
| `clickup guest get` | Get guest details |
| `clickup guest remove` | Remove a guest |

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
