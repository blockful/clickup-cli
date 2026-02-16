# AGENTS.md — AI Agent Guide for clickup-cli

## What This Is

A CLI wrapping the **complete ClickUp API** (135+ commands). All output is JSON. No interactive prompts. Designed for you.

## Setup

```bash
# Authenticate (one-time)
clickup auth login --token pk_YOUR_TOKEN

# Verify
clickup auth whoami
```

Config saved to `~/.clickup-cli.yaml`. All subsequent commands use the stored token automatically.

## Key Patterns

```bash
# Pattern: clickup <resource> <verb> --flags
clickup task create --list 123 --name "My task" --priority 2
clickup task get --id abc123
clickup task update --id abc123 --status "done"
clickup task delete --id abc123

# Search across workspace
clickup task search --workspace 1234 --assignee 567 --status "in progress"

# Nested resources
clickup task dependency add --task abc --depends-on def
clickup comment reply create --comment-id 123 --text "Reply text"
clickup goal key-result create --goal-id uuid --name "Target" --type number
```

## Output Format

**Success** — raw JSON from ClickUp API (object or array).

**Error** — structured:
```json
{"error": "task not found", "code": "NOT_FOUND", "status": 404}
```

**Exit codes**: 0 = success, non-zero = error. Parse `error` field from stdout on failure.

## Common Workflows

### Discover workspace structure
```bash
clickup workspace list                          # Get workspace IDs
clickup space list --workspace 1234             # Get spaces
clickup folder list --space 5678                # Get folders
clickup list list --folder 9012                 # Get lists
clickup list list --space 5678                  # Folderless lists
```

### Task management
```bash
clickup task list --list 123 --status "open" --subtasks
clickup task create --list 123 --name "Task" --assignee 456 --due-date 1700000000000
clickup task update --id abc --assignees-add 789 --priority 1
clickup task search --workspace 1234 --tag "urgent" --include-closed
```

### Time tracking
```bash
clickup time-entry start --workspace 1234 --task abc --description "Working"
clickup time-entry stop --workspace 1234
clickup time-entry list --workspace 1234 --start-date 1700000000000 --end-date 1700100000000
```

### Comments & docs
```bash
clickup comment create --task abc --text "Status update: done"
clickup comment reply create --comment-id 123 --text "Thanks"
clickup doc create --workspace 1234 --name "Meeting notes"
clickup doc page-create --workspace 1234 --doc docid --name "Page 1" --content "# Hello"
```

## Tips

- **IDs are strings** — always pass them as strings, even if numeric
- **Timestamps are Unix milliseconds** — not seconds
- **`--workspace` defaults** from config — set it once with `auth login`
- **Pagination**: use `--page` for task lists, `--cursor` where supported
- **Custom fields**: use `clickup custom-field list --list 123` to discover fields, then `clickup custom-field set --task abc --field uuid --value '{"option":"value"}'`
- **No confirmation prompts** on delete — safe for automation
- **`--format text`** for human-readable output when debugging

## Full Reference

See [docs/api.md](docs/api.md) for every command, every flag, and every API endpoint mapping.
