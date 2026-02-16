# Changelog

## [1.0.0] - 2026-02-16

First release of `clickup-cli` — a production-quality CLI covering the **complete ClickUp API** (135+ commands), optimized for AI agents.

### Command Groups (27 resource groups)

- **Auth** — `login`, `whoami`
- **Workspaces** — `list`, `plan`, `seats`
- **Spaces** — full CRUD (list, get, create, update, delete)
- **Folders** — full CRUD
- **Lists** — full CRUD (folder and folderless)
- **Tasks** — full CRUD + `search`, `add-to-list`, `remove-from-list`, `merge`, `time-in-status`
- **Task Dependencies** — `add`, `remove`
- **Task Links** — `add`, `remove`
- **Comments** — full CRUD on tasks, lists, views
- **Comment Replies** — threaded reply `list` and `create`
- **Docs (v3 API)** — doc CRUD + page CRUD (list, get, create, update)
- **Custom Fields** — `list` (workspace/space/folder/list scope), `set`, `remove`
- **Tags** — full CRUD + task `add`/`remove`
- **Checklists** — `create`, `update`, `delete`
- **Checklist Items** — `create`, `update`, `delete`
- **Time Entries** — full CRUD + `start`, `stop`, `current`, `history`
- **Time Entry Legacy** — task-level `list`, `create`, `update`, `delete`
- **Time Entry Tags** — `add`, `remove`, `update`
- **Views** — full CRUD + `tasks`
- **Goals** — full CRUD
- **Goal Key Results** — `create`, `update`, `delete`
- **Webhooks** — full CRUD
- **Members** — `list` (task/list)
- **User Groups** — `list`, `create`, `update`, `delete`
- **Users** — `invite`, `get`, `update`, `remove`
- **Guests** — `invite`, `get`, `edit`, `remove` + `add-to-task/list/folder`, `remove-from-task/list/folder`
- **Roles** — `list` custom roles
- **Custom Task Types** — `list`
- **Shared Hierarchy** — `list`
- **Templates** — `list`, `create-task`, `create-list`, `create-folder`
- **Attachments** — `create` (file upload)

### Agent-Optimized Design

- **JSON-first output** — every command outputs valid JSON. Errors: `{"error":"message","code":"ERROR_CODE"}`
- **Every API parameter exposed** — no capability loss vs. the raw API
- **No interactive prompts** — fully flag-driven for agent/automation use
- **Config persistence** — token and default workspace in `~/.clickup-cli.yaml`
- **Consistent patterns** — `clickup <resource> <verb> [flags]` everywhere

### Developer Experience

- CI with GitHub Actions (Go 1.22 + 1.23, golangci-lint)
- Cross-platform binaries (Linux, macOS, Windows — amd64 + arm64)
- Table-driven tests with httptest mocking
- Conventional commits, issue templates, PR template

[1.0.0]: https://github.com/blockful/clickup-cli/releases/tag/v1.0.0
