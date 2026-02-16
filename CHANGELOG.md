# Changelog

## [1.0.0] - 2026-02-16

First release of `clickup-cli` — a full-featured CLI for the ClickUp API, optimized for AI agents.

### Features

**53 commands** covering the complete ClickUp API surface:

- **Tasks** — Create, read, update, delete, and search tasks across your workspace. Full filter support: status, assignee, tags, due dates, custom fields, date ranges, and more. Subtask and dependency management included.

- **Docs (v3 API)** — Create and manage ClickUp Docs (wiki). Full page CRUD: create, read, update pages within docs. Search docs across your workspace.

- **Spaces, Folders & Lists** — Complete hierarchy management. Create, update, delete at every level. Folderless list support for flat structures.

- **Comments** — Task and list-level comments with threading support. Create, update, delete, and list threaded replies.

- **Custom Fields** — List custom fields at any level (workspace, space, folder, list). Set and remove values on tasks.

- **Tags** — Full tag lifecycle: create, update, delete space tags. Add and remove tags from tasks.

- **Checklists** — Create and manage checklists on tasks. Full checklist item CRUD with assignee and resolution tracking.

- **Time Tracking** — Create, update, delete time entries. Start/stop timers. View running timer. Full date range filtering.

- **Views** — CRUD views at workspace, space, folder, and list levels. Retrieve tasks from any view.

- **Goals** — Create and track goals with key results.

- **Webhooks** — Create, update, delete webhooks for event-driven integrations.

- **Members, Groups & Guests** — Manage workspace members, user groups, and guest access.

### Agent-Optimized Design

- **JSON-first output** — Every command outputs valid JSON by default. Errors are structured: `{"error": "message", "code": "ERROR_CODE"}`
- **Comprehensive flags** — Every ClickUp API parameter is exposed as a CLI flag. No capability loss vs. the raw API.
- **Config persistence** — Token and default workspace saved to `~/.clickup-cli.yaml`
- **Predictable structure** — `clickup <resource> <verb> [flags]` pattern across all commands

### Developer Experience

- CI with GitHub Actions (Go 1.22 + 1.23, golangci-lint)
- Cross-platform binaries (Linux, macOS, Windows — amd64 + arm64)
- Table-driven tests with httptest mocking
- Conventional commits, issue templates, PR template

[1.0.0]: https://github.com/blockful/clickup-cli/releases/tag/v1.0.0
