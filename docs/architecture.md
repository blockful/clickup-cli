# Architecture

## Overview

clickup-cli is a Go CLI built with [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper). It wraps the ClickUp API (v2 + v3 Docs) with 99.3% coverage — 134/135 endpoints across 27 resource groups.

```
clickup-cli/
├── main.go                          # Entry point
├── cmd/                             # Cobra command definitions (one file per resource)
│   ├── root.go                      # Root command, global flags, config init
│   ├── auth.go                      # auth login, auth whoami
│   ├── workspace.go                 # workspace list/plan/seats
│   ├── space.go                     # space CRUD
│   ├── folder.go                    # folder CRUD
│   ├── list.go                      # list CRUD
│   ├── task.go                      # task CRUD, search, merge, add-to-list, dependency, link, time-in-status
│   ├── comment.go                   # comment CRUD + reply subcommands
│   ├── doc.go                       # doc CRUD + page CRUD (v3 API)
│   ├── checklist.go                 # checklist + checklist-item CRUD
│   ├── custom_field.go              # custom-field list/set/remove
│   ├── tag.go                       # tag CRUD + task tagging
│   ├── time_entry.go                # time-entry CRUD, start/stop/current, history
│   ├── time_entry_legacy.go         # time-entry legacy (task-level tracking)
│   ├── time_entry_tags.go           # time-entry tag add/remove/update
│   ├── view.go                      # view CRUD + tasks
│   ├── goal.go                      # goal CRUD + key-result CRUD
│   ├── webhook.go                   # webhook CRUD
│   ├── member.go                    # member list
│   ├── user.go                      # user invite/get/update/remove
│   ├── role.go                      # role list
│   ├── shared.go                    # shared hierarchy
│   ├── custom_task_type.go          # custom-task-type list
│   ├── template.go                  # template list + create-task/list/folder
│   ├── attachment.go                # attachment create (file upload)
│   ├── relationship.go              # task dependency/link (registered via task.go)
│   └── version.go                   # version command
├── internal/
│   ├── api/                         # HTTP client + API type definitions
│   │   ├── client.go                # Base HTTP client, auth, rate limiting, retries
│   │   ├── tasks.go                 # Task endpoints
│   │   ├── lists.go                 # List endpoints
│   │   ├── spaces.go                # Space endpoints
│   │   ├── folders.go               # Folder endpoints
│   │   ├── comments.go              # Comment endpoints
│   │   ├── docs.go                  # Docs v3 endpoints
│   │   ├── checklists.go            # Checklist endpoints
│   │   ├── custom_fields.go         # Custom field endpoints
│   │   ├── tags.go                  # Tag endpoints
│   │   ├── time_entries.go          # Time entry endpoints
│   │   ├── time_tracking_legacy.go  # Legacy time tracking
│   │   ├── views.go                 # View endpoints
│   │   ├── goals.go                 # Goal endpoints
│   │   ├── webhooks.go              # Webhook endpoints
│   │   ├── workspaces.go            # Workspace endpoints
│   │   ├── members.go               # Member endpoints
│   │   ├── users.go                 # User endpoints
│   │   ├── roles.go                 # Role endpoints
│   │   ├── shared.go                # Shared hierarchy endpoints
│   │   ├── custom_task_types.go     # Custom task type endpoints
│   │   ├── templates.go             # Template endpoints
│   │   ├── attachments.go           # Attachment endpoints
│   │   ├── relationships.go         # Relationship (dependency/link) endpoints
│   │   ├── auth.go                  # Auth/user endpoints
│   │   └── *_test.go               # Table-driven tests with httptest
│   ├── config/                      # Viper-based config management
│   └── output/                      # JSON/text output formatting
├── .github/                         # CI, issue templates, PR template
├── docs/                            # Documentation
│   ├── api.md                       # Complete command & flag reference
│   ├── architecture.md              # This file
│   ├── business-rules.md            # Business rules & invariants
│   └── decisions/                   # Architecture Decision Records
└── go.mod / go.sum
```

## Layers

1. **cmd/** — Cobra command definitions, flag parsing, validation. Thin layer — delegates to `internal/api`.
2. **internal/api/** — HTTP client, request/response types, API call logic. Handles auth headers, rate limiting, retries.
3. **internal/config/** — Viper-based config file management (`~/.clickup-cli.yaml`).
4. **internal/output/** — JSON output formatting, structured error formatting.

## Design Principles

- **JSON-first** — All output is valid JSON for machine consumption. Human-readable text via `--format text`.
- **Consistent error format** — `{"error": "...", "code": "...", "status": N}` across all commands.
- **Config file + flag overrides** — Persistent auth via config, one-off overrides via flags.
- **Thin command layer** — Commands parse flags and call API functions. Business logic lives in `internal/api/`.
- **No interactive prompts** — Everything is flag-driven for agent compatibility.
- **Table-driven tests** — All API functions tested with `httptest` mock servers.

## API Versions

- **v2** — Used for all endpoints except Docs
- **v3** — Used for Docs API (`/api/v3/workspaces/...`)
