# Business Rules

## BR-001: Output Format

- **BR-001a**: Every command MUST output valid JSON to stdout by default.
- **BR-001b**: When `--format=text` or `--human` is set, output human-readable text to stdout instead.
- **BR-001c**: Error responses MUST be JSON: `{"error": "<message>", "code": "<ERROR_CODE>", "status": <http_status>}`.
- **BR-001d**: Debug/verbose output MUST go to stderr, never stdout.
- **BR-001e**: Exit code 0 for success, non-zero for errors.

## BR-002: Authentication

- **BR-002a**: API token is the primary auth method. Passed via `--token`, `CLICKUP_TOKEN`, or config file.
- **BR-002b**: OAuth2 support is planned but not P0.
- **BR-002c**: If no token is available, the CLI MUST return error code `AUTH_REQUIRED` and exit 1.
- **BR-002d**: Tokens MUST NOT be logged, even in verbose mode. Mask to `pk_****` in any output.

## BR-003: Workspace / Team Terminology

- **BR-003a**: The CLI MUST use "workspace" in all user-facing commands, flags, and output.
- **BR-003b**: Internally, the API client translates "workspace" to "team" for v2 endpoints.
- **BR-003c**: JSON output keys use `workspace_id`, never `team_id`.

## BR-004: ID Handling

- **BR-004a**: All ClickUp IDs MUST be treated as strings throughout the codebase.
- **BR-004b**: IDs MUST be passed as positional arguments or named flags, never inferred.

## BR-005: Pagination

- **BR-005a**: List commands MUST support `--limit` and `--cursor` flags.
- **BR-005b**: Responses MUST include pagination metadata: `{"data": [...], "cursor": "next_cursor", "has_more": true}`.
- **BR-005c**: If the upstream API uses page-based pagination, the client MUST translate to cursor semantics.

## BR-006: Config Precedence

- **BR-006a**: CLI flags override environment variables override config file.
- **BR-006b**: Config file location: `~/.clickup-cli.yaml` (XDG override not in P0).
- **BR-006c**: `clickup config set <key> <value>` writes to config file.
- **BR-006d**: `clickup config get <key>` reads resolved value (after precedence).

## BR-007: Task Operations

- **BR-007a**: `task create` requires `--list-id` and `--name` at minimum.
- **BR-007b**: `task update` accepts any combination of mutable fields; only specified fields are sent to API.
- **BR-007c**: `task search` uses the filtered team tasks endpoint (GET /v2/team/{team_id}/task) with query parameters.
- **BR-007d**: Custom field values on tasks are set via `custom-field set`, not via `task update`.
- **BR-007e**: `task delete` is permanent. The CLI MUST NOT add confirmation prompts (agents can't interact). Use `--dry-run` for safety.

## BR-008: Hierarchical Context

- **BR-008a**: ClickUp hierarchy: Workspace → Space → Folder → List → Task.
- **BR-008b**: Lists can exist directly under Spaces (folderless). The CLI supports both paths.
- **BR-008c**: When `--workspace-id` is omitted, use the default workspace from config.
- **BR-008d**: Commands MUST fail with `MISSING_PARAM` if a required parent ID is not provided and cannot be inferred from config.

## BR-009: Comments

- **BR-009a**: Comments can be attached to tasks, lists, or views. The CLI uses `--task-id`, `--list-id`, or `--view-id` to specify the parent.
- **BR-009b**: Exactly one parent flag is required for `comment create` and `comment list`.

## BR-010: Docs (v3 API)

- **BR-010a**: Doc commands use the v3 base URL.
- **BR-010b**: Doc page content is Markdown. The CLI accepts `--content` as a string or `--content-file` to read from a file.
- **BR-010c**: `doc page update` replaces the entire page content (PUT semantics).

## BR-011: Rate Limiting

- **BR-011a**: On HTTP 429, the client MUST retry with exponential backoff, respecting the `Retry-After` header.
- **BR-011b**: Maximum 3 retries before returning a `RATE_LIMITED` error.
- **BR-011c**: Rate limit info SHOULD be included in verbose stderr output.

## BR-012: HTTP Errors

- **BR-012a**: 4xx errors return the ClickUp error message in the `error` field.
- **BR-012b**: 5xx errors return `"error": "ClickUp API error", "code": "API_ERROR"`.
- **BR-012c**: Network errors return `"error": "<detail>", "code": "NETWORK_ERROR"`.

## BR-013: Destructive Operations

- **BR-013a**: Delete commands MUST NOT prompt for confirmation (breaks agent workflows).
- **BR-013b**: All delete commands MUST support `--dry-run` which shows what would be deleted without executing.

## BR-014: Custom Fields

- **BR-014a**: `custom-field list` can scope to list, folder, space, or workspace level.
- **BR-014b**: `custom-field set` requires `--task-id`, `--field-id`, and `--value`.
- **BR-014c**: The `--value` flag accepts JSON for complex field types (e.g., dropdowns, labels).

## BR-015: Time Entries

- **BR-015a**: `time-entry start` starts a running timer on a task.
- **BR-015b**: `time-entry stop` stops the current running timer.
- **BR-015c**: Duration values are in milliseconds.

## BR-016: Webhooks

- **BR-016a**: `webhook create` requires `--endpoint` (URL) and `--events` (comma-separated event types).
- **BR-016b**: Webhook secrets are returned on create and MUST be displayed in output.

## BR-017: Attachments

- **BR-017a**: `attachment create` requires `--task-id` and `--file` (local path).
- **BR-017b**: File upload uses multipart/form-data encoding.

## BR-018: Guests

- **BR-018a**: Guests can be scoped to task, list, or folder level via `add-to-*` / `remove-from-*` commands.
- **BR-018b**: `--permission-level` defaults to `read`.
- **BR-018c**: Guest invite requires `--email`.

## BR-019: Users

- **BR-019a**: `user invite` requires `--email`.
- **BR-019b**: `user update` can change `--admin`, `--username`, and `--custom-role-id`.

## BR-020: Templates

- **BR-020a**: Templates can create tasks, lists, or folders from saved templates.
- **BR-020b**: `template create-task` requires `--list`, `--template-id`, and `--name`.

## BR-021: Time Entry Legacy

- **BR-021a**: Legacy time tracking operates at the task level (different from workspace-level time entries).
- **BR-021b**: Uses interval IDs for update/delete operations.

## BR-022: Task Dependencies & Links

- **BR-022a**: Dependencies use `--depends-on` or `--dependency-of` to specify direction.
- **BR-022b**: Links are bidirectional between two tasks via `--links-to`.

## BR-023: Goal Key Results

- **BR-023a**: Key results are children of goals. Types: `number`, `percentage`, `automatic`, `boolean`.
- **BR-023b**: `automatic` type requires `--task-ids` or `--list-ids`.
