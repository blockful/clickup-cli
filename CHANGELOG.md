# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-02-16

### Added

#### Core
- JSON-first output format optimized for AI agent workflows
- Config file support (`~/.clickup-cli.yaml`) for token and workspace persistence
- Global flags: `--token`, `--workspace`, `--format`, `--verbose`
- Structured error output with error codes
- `clickup version` command with build info

#### Authentication
- `auth login` — configure API token with validation
- `auth whoami` — display current user info

#### Workspaces
- `workspace list` — list authorized workspaces

#### Spaces
- `space list`, `space get`, `space create`, `space update`, `space delete`

#### Folders
- `folder list`, `folder get`, `folder create`, `folder update`, `folder delete`

#### Lists
- `list list`, `list get`, `list create`, `list update`, `list delete`
- Folderless list support via `--space` flag

#### Tasks
- `task list` — 24 filter flags matching full ClickUp API surface
- `task get`, `task create`, `task update`, `task delete`
- `task search` — workspace-wide task search
- Full custom field, date range, and assignee filtering support

#### Comments
- `comment list` — task and list level comments
- `comment create`, `comment update`, `comment delete`
- Threaded comment support

#### Docs (v3 API)
- `doc list`, `doc get`, `doc create`
- `doc page-list`, `doc page-get`, `doc page-create`, `doc page-update`

#### Custom Fields
- `custom-field list` — at list, folder, space, or workspace level
- `custom-field set`, `custom-field remove`

#### Tags
- `tag list`, `tag create`, `tag update`, `tag delete`
- `tag add` (to task), `tag remove` (from task)

#### Checklists
- `checklist create`, `checklist update`, `checklist delete`
- `checklist-item create`, `checklist-item update`, `checklist-item delete`

#### Time Tracking
- `time-entry list`, `time-entry get`, `time-entry create`, `time-entry update`, `time-entry delete`
- `time-entry start`, `time-entry stop`, `time-entry current`

#### Views
- `view list`, `view get`, `view create`, `view update`, `view delete`, `view tasks`

#### Goals
- `goal list`, `goal get`, `goal create`, `goal update`, `goal delete`

#### Webhooks
- `webhook list`, `webhook create`, `webhook update`, `webhook delete`

#### Members, Groups & Guests
- `member list` (task and list level)
- `group list`, `group create`, `group update`, `group delete`
- `guest get`, `guest invite`, `guest update`, `guest remove`

[1.0.0]: https://github.com/blockful/clickup-cli/releases/tag/v1.0.0
