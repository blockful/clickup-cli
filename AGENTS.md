# AGENTS.md — AI Agent Guide for clickup-cli

## What This Is

A Go CLI wrapping the **complete ClickUp API** (134/135 endpoints, 99.3% coverage). All output is JSON. No interactive prompts. Designed for AI agents.

## Setup

```bash
# Authenticate (one-time, saves to ~/.clickup-cli.yaml)
clickup auth login --token pk_YOUR_TOKEN

# Verify
clickup auth whoami
```

## Command Pattern

```
clickup <resource> <verb> [flags]
```

All flags are `--kebab-case`. Required flags are noted in `--help`. IDs are always strings.

## Output Format

**Success** — raw JSON from ClickUp API:
```json
{"id": "abc123", "name": "My Task", "status": {"status": "open"}, ...}
```

**Error** — structured:
```json
{"error": "task not found", "code": "NOT_FOUND", "status": 404}
```

**Exit codes**: 0 = success, non-zero = error. Always parse stdout for the JSON response.

Use `--format text` only for human debugging. Default JSON is what you should parse.

---

## Complete Command Reference

### Workspace Discovery

```bash
clickup workspace list                          # → [{"id":"1234","name":"My Workspace",...}]
clickup workspace plan --workspace 1234         # Get plan details
clickup workspace seats --workspace 1234        # Get seat usage
clickup space list --workspace 1234             # List spaces
clickup space get --id 5678                     # Get single space
clickup folder list --space 5678                # List folders in space
clickup folder get --id 9012                    # Get single folder
clickup list list --folder 9012                 # Lists in folder
clickup list list --space 5678                  # Folderless lists in space
clickup list get --id 456                       # Get single list
```

### Task CRUD

```bash
# List tasks in a list
clickup task list --list 123 --include-closed --subtasks --page 0

# Get single task
clickup task get --id abc123
clickup task get --id abc123 --include-markdown --include-subtasks

# Create task
clickup task create --list 123 --name "Ship feature" \
  --assignee 456 --assignee 789 \
  --priority 2 \
  --status "in progress" \
  --due-date 1700000000000 \
  --tag "urgent" --tag "frontend" \
  --markdown-content "## Description\nRich **markdown** here"

# Update task
clickup task update --id abc123 \
  --name "New name" \
  --status "done" \
  --assignees-add 789 \
  --assignees-rem 456 \
  --priority 1

# Delete task
clickup task delete --id abc123

# Search across workspace
clickup task search --workspace 1234 \
  --assignee 567 \
  --status "in progress" \
  --tag "urgent" \
  --include-closed \
  --include-markdown
```

### Task Advanced Operations

```bash
# Add/remove task from additional lists
clickup task add-to-list --id abc123 --list 456
clickup task remove-from-list --id abc123 --list 456

# Merge tasks
clickup task merge --id abc123 --merge-with def456 --merge-with ghi789

# Time in status
clickup task time-in-status --id abc123
clickup task time-in-status --task-ids abc123 --task-ids def456  # bulk

# Dependencies
clickup task dependency add --task abc123 --depends-on def456
clickup task dependency add --task abc123 --dependency-of ghi789
clickup task dependency remove --task abc123 --depends-on def456

# Links
clickup task link add --task abc123 --links-to def456
clickup task link remove --task abc123 --links-to def456
```

### Custom Task IDs

When your workspace uses custom task IDs (e.g., `PROJ-123`), add `--custom-task-ids` and `--team-id`:

```bash
clickup task get --id "PROJ-123" --custom-task-ids --team-id 1234567
clickup task update --id "PROJ-123" --custom-task-ids --team-id 1234567 --status "done"
clickup comment create --task "PROJ-123" --custom-task-ids --team-id 1234567 --text "Done!"
```

This pattern works on: `task get/update/delete`, `task add-to-list/remove-from-list/merge/time-in-status`, `task dependency/link`, `comment create`, `custom-field set/remove`, `attachment create`, `guest add-to-task/remove-from-task`, `time-entry legacy` commands.

### Comments

```bash
# Create comment on task
clickup comment create --task abc123 --text "Status update: done"

# Create comment on list or view
clickup comment create --list 456 --text "List-level note"
clickup comment create --view-id viewid --text "View comment"

# List comments
clickup comment list --task abc123
clickup comment list --list 456

# Update/delete
clickup comment update --id 789 --text "Updated text"
clickup comment delete --id 789

# Threaded replies
clickup comment reply list --comment-id 789
clickup comment reply create --comment-id 789 --text "Reply text"
```

### Custom Fields

```bash
# Discover fields (scoped to list, folder, space, or workspace)
clickup custom-field list --list 123
clickup custom-field list --space 456
clickup custom-field list --workspace 1234

# Set a field value
clickup custom-field set --task abc123 --field field-uuid --value '"some string"'
clickup custom-field set --task abc123 --field field-uuid --value '123'
clickup custom-field set --task abc123 --field field-uuid --value '{"option_id": "uuid"}'

# Remove a field value
clickup custom-field remove --task abc123 --field field-uuid
```

### Tags

```bash
clickup tag list --space 5678                    # List space tags
clickup tag create --space 5678 --name "urgent"  # Create tag
clickup tag add --task abc123 --name "urgent"    # Add to task
clickup tag remove --task abc123 --name "urgent" # Remove from task
clickup tag update --space 5678 --name "urgent" --new-name "critical"
clickup tag delete --space 5678 --name "old-tag"
```

### Time Tracking

```bash
# Start/stop timer
clickup time-entry start --workspace 1234 --task abc123 --description "Working"
clickup time-entry current --workspace 1234  # Check running timer
clickup time-entry stop --workspace 1234

# Manual time entry
clickup time-entry create --workspace 1234 \
  --task abc123 \
  --start 1700000000000 \
  --duration 3600000 \
  --description "Code review"

# List/get time entries
clickup time-entry list --workspace 1234 --start-date 1700000000000 --end-date 1700100000000
clickup time-entry get --workspace 1234 --id 5678

# Update/delete
clickup time-entry update --workspace 1234 --id 5678 --description "Updated"
clickup time-entry delete --workspace 1234 --id 5678

# History
clickup time-entry history --workspace 1234 --id 5678

# Legacy task-level time tracking
clickup time-entry legacy list --task-id abc123
clickup time-entry legacy create --task-id abc123 --time 3600000 --start 1700000000000

# Time entry tags
clickup time-entry tag add --workspace 1234 --tags "dev,review" --time-entry-ids "5678,9012"
clickup time-entry tag remove --workspace 1234 --tags "dev" --time-entry-ids "5678"
clickup time-entry tag update --workspace 1234 --name "dev" --new-name "development"
```

### Docs (v3 API)

```bash
clickup doc list --workspace 1234
clickup doc get --workspace 1234 --id docid
clickup doc create --workspace 1234 --name "Meeting Notes"

# Pages
clickup doc page-list --workspace 1234 --doc docid
clickup doc page-get --workspace 1234 --doc docid --page pageid
clickup doc page-create --workspace 1234 --doc docid --name "Agenda" --content "# Meeting Agenda"
clickup doc page-update --workspace 1234 --doc docid --page pageid --content "# Updated"
```

### Checklists

```bash
clickup checklist create --task abc123 --name "QA Steps"
clickup checklist update --id checklistid --name "Updated Name"
clickup checklist delete --id checklistid

clickup checklist-item create --checklist checklistid --name "Test login"
clickup checklist-item update --checklist checklistid --id itemid --resolved true
clickup checklist-item delete --checklist checklistid --id itemid
```

### Views

```bash
clickup view list --workspace 1234           # Workspace-level views
clickup view list --space 5678               # Space-level views
clickup view list --folder 9012              # Folder-level views
clickup view list --list 456                 # List-level views
clickup view get --id viewid
clickup view create --space 5678 --name "Sprint Board" --type board
clickup view tasks --id viewid               # Get tasks in view
clickup view update --id viewid --name "Updated"
clickup view delete --id viewid
```

### Goals & Key Results

```bash
clickup goal list --workspace 1234
clickup goal get --id goalid
clickup goal create --workspace 1234 --name "Q1 OKRs" --owners 456
clickup goal update --id goalid --name "Updated"
clickup goal delete --id goalid

clickup goal key-result create --goal-id goalid --name "Ship v2" --type number --steps-start 0 --steps-end 100
clickup goal key-result update --id krid --steps-current 50
clickup goal key-result delete --id krid
```

### People & Access

```bash
# Users
clickup user invite --workspace 1234 --email user@co.com --admin
clickup user get --workspace 1234 --id 456
clickup user update --workspace 1234 --id 456 --username "newname" --custom-role-id 789
clickup user remove --workspace 1234 --id 456

# Members
clickup member list --task abc123
clickup member list --list 456

# Groups
clickup group list --workspace 1234
clickup group create --workspace 1234 --name "Backend Team" --members 456 --members 789
clickup group update --workspace 1234 --id groupid --name "New Name"
clickup group delete --workspace 1234 --id groupid

# Guests
clickup guest invite --workspace 1234 --email guest@co.com
clickup guest get --workspace 1234 --guest-id 456
clickup guest edit --workspace 1234 --guest-id 456 --can-edit-tags
clickup guest remove --workspace 1234 --guest-id 456

# Guest resource access
clickup guest add-to-task --workspace 1234 --guest-id 456 --task abc123 --permission-level read
clickup guest add-to-list --workspace 1234 --guest-id 456 --list 789
clickup guest add-to-folder --workspace 1234 --guest-id 456 --folder 012
clickup guest remove-from-task --workspace 1234 --guest-id 456 --task abc123
clickup guest remove-from-list --workspace 1234 --guest-id 456 --list 789
clickup guest remove-from-folder --workspace 1234 --guest-id 456 --folder 012

# Roles
clickup role list --workspace 1234
```

### Infrastructure

```bash
# Webhooks
clickup webhook list --workspace 1234
clickup webhook create --workspace 1234 --endpoint "https://example.com/hook" --events taskCreated --events taskUpdated
clickup webhook update --id hookid --endpoint "https://new.com/hook" --events taskCreated
clickup webhook delete --id hookid

# Templates
clickup template list --workspace 1234 --page 0
clickup template create-task --list 123 --template-id tmplid --name "From Template"
clickup template create-list --space 5678 --template-id tmplid --name "Sprint List"
clickup template create-folder --space 5678 --template-id tmplid --name "Project Folder"

# Shared hierarchy
clickup shared list --workspace 1234

# Custom task types
clickup custom-task-type list --workspace 1234

# Spaces CRUD
clickup space create --workspace 1234 --name "New Space"
clickup space update --id 5678 --name "Renamed"
clickup space delete --id 5678

# Folders/Lists CRUD
clickup folder create --space 5678 --name "New Folder"
clickup list create --folder 9012 --name "New List"
clickup list create --space 5678 --name "Folderless List"  # folderless
```

---

## Important Patterns

### Timestamps

All timestamps are **Unix milliseconds** (not seconds). Multiply seconds × 1000.

```bash
# Example: 2024-01-15 00:00:00 UTC = 1705276800000
clickup task create --list 123 --name "Task" --due-date 1705276800000
```

### Markdown Content

Use `--markdown-content` (alias: `--markdown-description`) on `task create` and `--markdown-description` on `task update` for rich task descriptions. Retrieve with `--include-markdown` on get/list/search.

```bash
clickup task create --list 123 --name "Task" \
  --markdown-content "## Overview\n\n- Point 1\n- Point 2\n\n**Bold** and *italic*"
```

### Custom Fields JSON

The `--value` flag on `custom-field set` accepts raw JSON:

```bash
# Text field
clickup custom-field set --task abc --field uuid --value '"hello"'

# Number field
clickup custom-field set --task abc --field uuid --value '42'

# Dropdown (by option ID)
clickup custom-field set --task abc --field uuid --value '"option-uuid"'

# Labels (array of option IDs)
clickup custom-field set --task abc --field uuid --value '["opt1","opt2"]'
```

### Filtering with Custom Fields

Use `--custom-fields` JSON array on `task list` and `task search`:

```bash
clickup task list --list 123 --custom-fields '[{"field_id":"uuid","operator":"=","value":"hello"}]'
```

### Pagination

- `task list` and `task search`: use `--page 0`, `--page 1`, etc. (zero-indexed)
- Most list commands return all results; check `--help` for specific pagination flags

### Error Handling

1. Check exit code (0 = success)
2. On failure, parse stdout JSON for `error`, `code`, and `status` fields
3. Common codes: `AUTH_REQUIRED`, `NOT_FOUND`, `RATE_LIMITED`, `MISSING_PARAM`, `API_ERROR`
4. Rate limits (HTTP 429): client auto-retries with backoff (up to 3 times)

### Workspace Default

Set once, use everywhere:

```bash
clickup auth login --token pk_... 
# Workspace is saved to ~/.clickup-cli.yaml
# All --workspace flags auto-resolve from config
```

Or override per-command: `--workspace 9999`

---

## Common Agent Workflows

### Create a task with full metadata

```bash
clickup task create --list 123 \
  --name "Implement login page" \
  --markdown-content "## Acceptance Criteria\n- OAuth support\n- Remember me" \
  --assignee 456 \
  --priority 2 \
  --status "to do" \
  --due-date 1705276800000 \
  --tag "frontend" \
  --time-estimate 14400000
```

Then set custom fields:
```bash
clickup custom-field set --task TASKID --field sprint-field-uuid --value '"Sprint 12"'
```

### Move a task through a workflow

```bash
clickup task update --id abc123 --status "in progress"
clickup time-entry start --workspace 1234 --task abc123 --description "Working"
# ... work happens ...
clickup time-entry stop --workspace 1234
clickup comment create --task abc123 --text "Completed implementation"
clickup task update --id abc123 --status "review"
```

### Bulk search and update

```bash
# Find all overdue tasks
clickup task search --workspace 1234 --due-date-lt $(date +%s000) --status "in progress"

# Then update each (pseudo-code — parse JSON, loop through results)
clickup task update --id TASKID --status "overdue"
```

### Set up dependencies between tasks

```bash
clickup task dependency add --task abc123 --depends-on def456  # abc depends on def
clickup task dependency add --task abc123 --dependency-of ghi789  # ghi depends on abc
```

## Full Flag Reference

See **[docs/api.md](docs/api.md)** for every command, every flag, types, defaults, and API endpoint mappings.
