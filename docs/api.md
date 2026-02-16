# ClickUp CLI — Complete Command & Flag Reference (135+ commands)

> **Base URL:** `https://api.clickup.com/api`
> Docs API (v3) uses `https://api.clickup.com/api/v3/workspaces/`.
> All output is JSON. All timestamps are Unix milliseconds unless noted.

## Global Flags

These flags are available on **every** command:

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--token` | string | `~/.clickup-cli.yaml` | ClickUp API token (overrides config) |
| `--workspace` | string | `~/.clickup-cli.yaml` | Default workspace ID (overrides config) |
| `--format` | string | `json` | Output format: `json` or `text` |
| `--verbose` | bool | `false` | Enable verbose output |

---

## Auth

### `clickup auth login`

Authenticate with a ClickUp API token. Validates by calling the user endpoint, then saves to config.

**API:** `GET /v2/user`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--token` | string | *(prompt)* | — | API token. If omitted, prompts interactively |

### `clickup auth whoami`

Show the currently authenticated user.

**API:** `GET /v2/user`

*No command-specific flags.*

---

## Workspaces

### `clickup workspace list`

List all workspaces (teams) the authenticated user belongs to.

**API:** `GET /v2/team`

*No command-specific flags.*

### `clickup workspace plan`

Get workspace plan details.

**API:** `GET /v2/team/{team_id}/plan`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |

### `clickup workspace seats`

Get workspace seat usage.

**API:** `GET /v2/team/{team_id}/seats`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |

---

## Spaces

### `clickup space list`

List spaces in a workspace.

**API:** `GET /v2/team/{team_id}/space`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |

### `clickup space get`

Get a space by ID.

**API:** `GET /v2/space/{space_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `space_id` (path) | Space ID |

### `clickup space create`

Create a new space in a workspace.

**API:** `POST /v2/team/{team_id}/space`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--name` | string | *(required)* | `name` (body) | Space name |
| `--multiple-assignees` | bool | `false` | `multiple_assignees` (body) | Enable multiple assignees |
| `--features` | string | — | `features` (body) | Space features as JSON object |

### `clickup space update`

Update a space.

**API:** `PUT /v2/space/{space_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `space_id` (path) | Space ID |
| `--name` | string | — | `name` (body) | New name |
| `--multiple-assignees` | bool | — | `multiple_assignees` (body) | Enable/disable multiple assignees |
| `--features` | string | — | `features` (body) | Space features as JSON object |

### `clickup space delete`

Delete a space.

**API:** `DELETE /v2/space/{space_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `space_id` (path) | Space ID |

---

## Folders

### `clickup folder list`

List folders in a space.

**API:** `GET /v2/space/{space_id}/folder`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--space` | string | *(required)* | `space_id` (path) | Space ID |

### `clickup folder get`

Get a folder by ID.

**API:** `GET /v2/folder/{folder_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `folder_id` (path) | Folder ID |

### `clickup folder create`

Create a new folder in a space.

**API:** `POST /v2/space/{space_id}/folder`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--space` | string | *(required)* | `space_id` (path) | Space ID |
| `--name` | string | *(required)* | `name` (body) | Folder name |

### `clickup folder update`

Update a folder.

**API:** `PUT /v2/folder/{folder_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `folder_id` (path) | Folder ID |
| `--name` | string | *(required)* | `name` (body) | New name |

### `clickup folder delete`

Delete a folder.

**API:** `DELETE /v2/folder/{folder_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `folder_id` (path) | Folder ID |

---

## Lists

### `clickup list list`

List lists in a folder, or folderless lists in a space.

**API:** `GET /v2/folder/{folder_id}/list` or `GET /v2/space/{space_id}/list`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--folder` | string | — | `folder_id` (path) | Folder ID (use one of `--folder` or `--space`) |
| `--space` | string | — | `space_id` (path) | Space ID (for folderless lists) |

### `clickup list get`

Get a list by ID.

**API:** `GET /v2/list/{list_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `list_id` (path) | List ID |

### `clickup list create`

Create a new list in a folder or as a folderless list in a space.

**API:** `POST /v2/folder/{folder_id}/list` or `POST /v2/space/{space_id}/list`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--folder` | string | — | `folder_id` (path) | Folder ID (use one of `--folder` or `--space`) |
| `--space` | string | — | `space_id` (path) | Space ID (for folderless list) |
| `--name` | string | *(required)* | `name` (body) | List name |
| `--content` | string | — | `content` (body) | List description/content |
| `--due-date` | int64 | — | `due_date` (body) | Due date (Unix ms) |
| `--priority` | int | — | `priority` (body) | Priority: 1=urgent, 2=high, 3=normal, 4=low |
| `--assignee` | int | — | `assignee` (body) | Assignee user ID |
| `--status` | string | — | `status` (body) | List status |

### `clickup list update`

Update a list.

**API:** `PUT /v2/list/{list_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `list_id` (path) | List ID |
| `--name` | string | — | `name` (body) | New name |
| `--content` | string | — | `content` (body) | New description/content |
| `--due-date` | int64 | — | `due_date` (body) | Due date (Unix ms) |
| `--priority` | int | — | `priority` (body) | Priority |
| `--assignee` | int | — | `assignee` (body) | Assignee user ID |
| `--status` | string | — | `status` (body) | List status |
| `--unset-status` | bool | `false` | `unset_status` (body) | Remove list status |

### `clickup list delete`

Delete a list.

**API:** `DELETE /v2/list/{list_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `list_id` (path) | List ID |

---

## Tasks

### `clickup task list`

List tasks in a list.

**API:** `GET /v2/list/{list_id}/task`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--list` | string | *(required)* | `list_id` (path) | List ID |
| `--status` | string[] | — | `statuses[]` (query) | Filter by status(es) |
| `--assignee` | string[] | — | `assignees[]` (query) | Filter by assignee(s) |
| `--tag` | string[] | — | `tags[]` (query) | Filter by tag(s) |
| `--watchers` | string[] | — | `watchers[]` (query) | Filter by watcher(s) |
| `--page` | int | `0` | `page` (query) | Page number (0-indexed) |
| `--order-by` | string | — | `order_by` (query) | Order by field (e.g. `created`, `updated`, `due_date`) |
| `--reverse` | bool | `false` | `reverse` (query) | Reverse sort order |
| `--subtasks` | bool | `false` | `subtasks` (query) | Include subtasks |
| `--include-closed` | bool | `false` | `include_closed` (query) | Include closed tasks |
| `--archived` | bool | `false` | `archived` (query) | Include archived tasks |
| `--include-markdown` | bool | `false` | `include_markdown_description` (query) | Include markdown description |
| `--include-timl` | bool | `false` | `include_tasks_in_multiple_lists` (query) | Include tasks in multiple lists |
| `--due-date-gt` | int64 | `0` | `due_date_gt` (query) | Due date greater than (Unix ms) |
| `--due-date-lt` | int64 | `0` | `due_date_lt` (query) | Due date less than (Unix ms) |
| `--date-created-gt` | int64 | `0` | `date_created_gt` (query) | Created after (Unix ms) |
| `--date-created-lt` | int64 | `0` | `date_created_lt` (query) | Created before (Unix ms) |
| `--date-updated-gt` | int64 | `0` | `date_updated_gt` (query) | Updated after (Unix ms) |
| `--date-updated-lt` | int64 | `0` | `date_updated_lt` (query) | Updated before (Unix ms) |
| `--date-done-gt` | int64 | `0` | `date_done_gt` (query) | Done after (Unix ms) |
| `--date-done-lt` | int64 | `0` | `date_done_lt` (query) | Done before (Unix ms) |
| `--custom-fields` | string | — | `custom_fields` (query) | Custom fields filter (JSON array) |
| `--custom-items` | int[] | — | `custom_items[]` (query) | Filter by custom task type IDs |

### `clickup task get`

Get a task by ID.

**API:** `GET /v2/task/{task_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `task_id` (path) | Task ID |
| `--custom-task-ids` | bool | `false` | `custom_task_ids` (query) | Interpret `--id` as a custom task ID |
| `--team-id` | string | — | `team_id` (query) | Team/workspace ID (required when `custom-task-ids=true`) |
| `--include-subtasks` | bool | `false` | `include_subtasks` (query) | Include subtasks |
| `--include-markdown` | bool | `false` | `include_markdown_description` (query) | Include markdown description |

### `clickup task create`

Create a new task in a list.

**API:** `POST /v2/list/{list_id}/task`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--list` | string | *(required)* | `list_id` (path) | List ID |
| `--name` | string | *(required)* | `name` (body) | Task name |
| `--description` | string | — | `description` (body) | Plain text description |
| `--markdown-description` | string | — | `markdown_description` (body) | Markdown description |
| `--assignee` | int[] | — | `assignees` (body) | Assignee user IDs |
| `--status` | string | — | `status` (body) | Task status |
| `--priority` | int | — | `priority` (body) | Priority: 1=urgent, 2=high, 3=normal, 4=low |
| `--tag` | string[] | — | `tags` (body) | Tag names |
| `--due-date` | int64 | — | `due_date` (body) | Due date (Unix ms) |
| `--due-date-time` | bool | — | `due_date_time` (body) | Due date includes time |
| `--start-date` | int64 | — | `start_date` (body) | Start date (Unix ms) |
| `--start-date-time` | bool | — | `start_date_time` (body) | Start date includes time |
| `--time-estimate` | int64 | — | `time_estimate` (body) | Time estimate in milliseconds |
| `--notify-all` | bool | `false` | `notify_all` (body) | Notify all assignees |
| `--parent` | string | — | `parent` (body) | Parent task ID (creates subtask) |
| `--links-to` | string | — | `links_to` (body) | Task ID to link to |
| `--custom-fields` | string | — | `custom_fields` (body) | Custom fields as JSON array `[{"id":"...","value":"..."}]` |

### `clickup task update`

Update a task.

**API:** `PUT /v2/task/{task_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `task_id` (path) | Task ID |
| `--name` | string | — | `name` (body) | New name |
| `--description` | string | — | `description` (body) | New description |
| `--status` | string | — | `status` (body) | New status |
| `--priority` | int | — | `priority` (body) | Priority: 1=urgent, 2=high, 3=normal, 4=low |
| `--assignees-add` | int[] | — | `assignees.add` (body) | User IDs to add as assignees |
| `--assignees-rem` | int[] | — | `assignees.rem` (body) | User IDs to remove as assignees |
| `--due-date` | int64 | — | `due_date` (body) | Due date (Unix ms) |
| `--due-date-time` | bool | — | `due_date_time` (body) | Due date includes time |
| `--start-date` | int64 | — | `start_date` (body) | Start date (Unix ms) |
| `--start-date-time` | bool | — | `start_date_time` (body) | Start date includes time |
| `--time-estimate` | int64 | — | `time_estimate` (body) | Time estimate (ms) |
| `--archived` | bool | — | `archived` (body) | Archive/unarchive task |
| `--parent` | string | — | `parent` (body) | Parent task ID |
| `--custom-task-ids` | bool | `false` | `custom_task_ids` (query) | Interpret `--id` as custom task ID |
| `--team-id` | string | — | `team_id` (query) | Team ID (required when `custom-task-ids=true`) |

### `clickup task delete`

Delete a task.

**API:** `DELETE /v2/task/{task_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `task_id` (path) | Task ID |

### `clickup task search`

Search tasks across a workspace.

**API:** `GET /v2/team/{team_id}/task`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--status` | string[] | — | `statuses[]` (query) | Filter by status(es) |
| `--assignee` | string[] | — | `assignees[]` (query) | Filter by assignee(s) |
| `--tag` | string[] | — | `tags[]` (query) | Filter by tag(s) |
| `--page` | int | `0` | `page` (query) | Page number |
| `--order-by` | string | — | `order_by` (query) | Order by field |
| `--reverse` | bool | `false` | `reverse` (query) | Reverse sort order |
| `--subtasks` | bool | `false` | `subtasks` (query) | Include subtasks |
| `--include-closed` | bool | `false` | `include_closed` (query) | Include closed tasks |
| `--include-markdown` | bool | `false` | `include_markdown_description` (query) | Include markdown description |
| `--due-date-gt` | int64 | `0` | `due_date_gt` (query) | Due date greater than (Unix ms) |
| `--due-date-lt` | int64 | `0` | `due_date_lt` (query) | Due date less than (Unix ms) |
| `--date-created-gt` | int64 | `0` | `date_created_gt` (query) | Created after (Unix ms) |
| `--date-created-lt` | int64 | `0` | `date_created_lt` (query) | Created before (Unix ms) |
| `--date-updated-gt` | int64 | `0` | `date_updated_gt` (query) | Updated after (Unix ms) |
| `--date-updated-lt` | int64 | `0` | `date_updated_lt` (query) | Updated before (Unix ms) |
| `--date-done-gt` | int64 | `0` | `date_done_gt` (query) | Done after (Unix ms) |
| `--date-done-lt` | int64 | `0` | `date_done_lt` (query) | Done before (Unix ms) |
| `--custom-fields` | string | — | `custom_fields` (query) | Custom fields filter (JSON array) |
| `--custom-items` | int[] | — | `custom_items[]` (query) | Filter by custom task type IDs |
| `--list-ids` | string[] | — | `list_ids[]` (query) | Filter by list IDs |
| `--project-ids` | string[] | — | `project_ids[]` (query) | Filter by project/folder IDs |
| `--space-ids` | string[] | — | `space_ids[]` (query) | Filter by space IDs |
| `--folder-ids` | string[] | — | `folder_ids[]` (query) | Filter by folder IDs |

---

## Comments

### `clickup comment list`

List comments on a task or list.

**API:** `GET /v2/task/{task_id}/comment` or `GET /v2/list/{list_id}/comment`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--task` | string | — | `task_id` (path) | Task ID (use one of `--task` or `--list`) |
| `--list` | string | — | `list_id` (path) | List ID |

### `clickup comment create`

Add a comment to a task or list.

**API:** `POST /v2/task/{task_id}/comment` or `POST /v2/list/{list_id}/comment`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--task` | string | — | `task_id` (path) | Task ID (use one of `--task` or `--list`) |
| `--list` | string | — | `list_id` (path) | List ID |
| `--text` | string | *(required)* | `comment_text` (body) | Comment text |
| `--assignee` | int | — | `assignee` (body) | Assignee user ID |
| `--notify-all` | bool | `false` | `notify_all` (body) | Notify all |

### `clickup comment update`

Update a comment.

**API:** `PUT /v2/comment/{comment_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `comment_id` (path) | Comment ID |
| `--text` | string | *(required)* | `comment_text` (body) | New comment text |
| `--assignee` | int | — | `assignee` (body) | Reassign comment |
| `--resolved` | bool | — | `resolved` (body) | Mark as resolved/unresolved |

### `clickup comment delete`

Delete a comment.

**API:** `DELETE /v2/comment/{comment_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `comment_id` (path) | Comment ID |

---

## Docs (v3 API)

### `clickup doc list`

List/search docs in a workspace.

**API:** `GET /v3/workspaces/{workspace_id}/docs`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `workspace_id` (path) | Workspace ID |

### `clickup doc get`

Get a doc by ID.

**API:** `GET /v3/workspaces/{workspace_id}/docs/{doc_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `workspace_id` (path) | Workspace ID |
| `--id` | string | *(required)* | `doc_id` (path) | Doc ID |

### `clickup doc create`

Create a doc.

**API:** `POST /v3/workspaces/{workspace_id}/docs`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `workspace_id` (path) | Workspace ID |
| `--name` | string | *(required)* | `name` (body) | Doc name |
| `--visibility` | string | — | `visibility` (body) | Visibility setting |
| `--parent-id` | string | — | `parent.id` (body) | Parent resource ID |
| `--parent-type` | int | `0` | `parent.type` (body) | Parent resource type |

### `clickup doc page-list`

List pages in a doc.

**API:** `GET /v3/workspaces/{workspace_id}/docs/{doc_id}/page_listing`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `workspace_id` (path) | Workspace ID |
| `--doc` | string | *(required)* | `doc_id` (path) | Doc ID |

### `clickup doc page-get`

Get a page from a doc.

**API:** `GET /v3/workspaces/{workspace_id}/docs/{doc_id}/pages/{page_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `workspace_id` (path) | Workspace ID |
| `--doc` | string | *(required)* | `doc_id` (path) | Doc ID |
| `--page` | string | *(required)* | `page_id` (path) | Page ID |

### `clickup doc page-create`

Create a page in a doc.

**API:** `POST /v3/workspaces/{workspace_id}/docs/{doc_id}/pages`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `workspace_id` (path) | Workspace ID |
| `--doc` | string | *(required)* | `doc_id` (path) | Doc ID |
| `--name` | string | *(required)* | `name` (body) | Page name |
| `--content` | string | — | `content` (body) | Page content (markdown) |
| `--content-html` | string | — | `content_html` (body) | Page content (HTML) |
| `--parent-page` | string | — | `parent_page_id` (body) | Parent page ID (for nesting) |

### `clickup doc page-update`

Update a page in a doc.

**API:** `PUT /v3/workspaces/{workspace_id}/docs/{doc_id}/pages/{page_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `workspace_id` (path) | Workspace ID |
| `--doc` | string | *(required)* | `doc_id` (path) | Doc ID |
| `--page` | string | *(required)* | `page_id` (path) | Page ID |
| `--name` | string | — | `name` (body) | New name |
| `--content` | string | — | `content` (body) | New content (markdown) |
| `--content-html` | string | — | `content_html` (body) | New content (HTML) |

---

## Custom Fields

### `clickup custom-field list`

List custom fields accessible at a given scope.

**API:** `GET /v2/list/{list_id}/field` · `GET /v2/folder/{folder_id}/field` · `GET /v2/space/{space_id}/field` · `GET /v2/team/{team_id}/field`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--list` | string | — | `list_id` (path) | List ID (use exactly one scope flag) |
| `--folder` | string | — | `folder_id` (path) | Folder ID |
| `--space` | string | — | `space_id` (path) | Space ID |
| `--workspace` | string | — | `team_id` (path) | Workspace/Team ID |

### `clickup custom-field set`

Set a custom field value on a task.

**API:** `POST /v2/task/{task_id}/field/{field_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--task` | string | *(required)* | `task_id` (path) | Task ID |
| `--field` | string | *(required)* | `field_id` (path) | Custom field UUID |
| `--value` | string | *(required)* | `value` (body) | Value as JSON or plain string. Parsed as JSON first, falls back to string |

### `clickup custom-field remove`

Remove a custom field value from a task.

**API:** `DELETE /v2/task/{task_id}/field/{field_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--task` | string | *(required)* | `task_id` (path) | Task ID |
| `--field` | string | *(required)* | `field_id` (path) | Custom field UUID |

---

## Tags

### `clickup tag list`

List tags available in a space.

**API:** `GET /v2/space/{space_id}/tag`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--space` | string | *(required)* | `space_id` (path) | Space ID |

### `clickup tag create`

Create a tag in a space.

**API:** `POST /v2/space/{space_id}/tag`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--space` | string | *(required)* | `space_id` (path) | Space ID |
| `--name` | string | *(required)* | `tag.name` (body) | Tag name |
| `--fg` | string | `#000000` | `tag.tag_fg` (body) | Foreground color (hex) |
| `--bg` | string | `#000000` | `tag.tag_bg` (body) | Background color (hex) |

### `clickup tag update`

Update a tag in a space.

**API:** `PUT /v2/space/{space_id}/tag/{tag_name}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--space` | string | *(required)* | `space_id` (path) | Space ID |
| `--name` | string | *(required)* | `tag_name` (path) | Current tag name |
| `--new-name` | string | — | `tag.name` (body) | New tag name |
| `--fg` | string | — | `tag.tag_fg` (body) | New foreground color (hex) |
| `--bg` | string | — | `tag.tag_bg` (body) | New background color (hex) |

### `clickup tag delete`

Delete a tag from a space.

**API:** `DELETE /v2/space/{space_id}/tag/{tag_name}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--space` | string | *(required)* | `space_id` (path) | Space ID |
| `--name` | string | *(required)* | `tag_name` (path) | Tag name to delete |

### `clickup tag add`

Add a tag to a task.

**API:** `POST /v2/task/{task_id}/tag/{tag_name}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--task` | string | *(required)* | `task_id` (path) | Task ID |
| `--name` | string | *(required)* | `tag_name` (path) | Tag name |

### `clickup tag remove`

Remove a tag from a task.

**API:** `DELETE /v2/task/{task_id}/tag/{tag_name}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--task` | string | *(required)* | `task_id` (path) | Task ID |
| `--name` | string | *(required)* | `tag_name` (path) | Tag name |

---

## Checklists

### `clickup checklist create`

Create a checklist on a task.

**API:** `POST /v2/task/{task_id}/checklist`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--task` | string | *(required)* | `task_id` (path) | Task ID |
| `--name` | string | *(required)* | `name` (body) | Checklist name |

### `clickup checklist update`

Rename or reorder a checklist.

**API:** `PUT /v2/checklist/{checklist_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `checklist_id` (path) | Checklist ID (UUID) |
| `--name` | string | — | `name` (body) | New name |
| `--position` | int | — | `position` (body) | Position index (0 = top) |

### `clickup checklist delete`

Delete a checklist.

**API:** `DELETE /v2/checklist/{checklist_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `checklist_id` (path) | Checklist ID (UUID) |

---

## Checklist Items

### `clickup checklist-item create`

Create a checklist item.

**API:** `POST /v2/checklist/{checklist_id}/checklist_item`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--checklist` | string | *(required)* | `checklist_id` (path) | Checklist ID |
| `--name` | string | — | `name` (body) | Item name |
| `--assignee` | int | — | `assignee` (body) | Assignee user ID |

### `clickup checklist-item update`

Update a checklist item.

**API:** `PUT /v2/checklist/{checklist_id}/checklist_item/{checklist_item_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--checklist` | string | *(required)* | `checklist_id` (path) | Checklist ID |
| `--id` | string | *(required)* | `checklist_item_id` (path) | Checklist item ID |
| `--name` | string | — | `name` (body) | New name |
| `--resolved` | bool | — | `resolved` (body) | Mark as resolved/unresolved |
| `--assignee` | string | — | `assignee` (body) | Assignee user ID (or `null` to unassign) |
| `--parent` | string | — | `parent` (body) | Parent checklist item ID (for nesting, or `null`) |

### `clickup checklist-item delete`

Delete a checklist item.

**API:** `DELETE /v2/checklist/{checklist_id}/checklist_item/{checklist_item_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--checklist` | string | *(required)* | `checklist_id` (path) | Checklist ID |
| `--id` | string | *(required)* | `checklist_item_id` (path) | Checklist item ID |

---

## Time Tracking

### `clickup time-entry list`

List time entries within a date range. Defaults to last 30 days for the authenticated user.

**API:** `GET /v2/team/{team_id}/time_entries`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--start-date` | string | — | `start_date` (query) | Start date (Unix ms) |
| `--end-date` | string | — | `end_date` (query) | End date (Unix ms) |
| `--assignee` | string | — | `assignee` (query) | User ID to filter by |
| `--space` | string | — | `space_id` (query) | Space ID filter |
| `--folder` | string | — | `folder_id` (query) | Folder ID filter |
| `--list` | string | — | `list_id` (query) | List ID filter |
| `--task` | string | — | `task_id` (query) | Task ID filter |
| `--include-task-tags` | bool | `false` | `include_task_tags` (query) | Include task tags in response |
| `--include-location-names` | bool | `false` | `include_location_names` (query) | Include location names |

> **Note:** Only one location filter (`--space`, `--folder`, `--list`, `--task`) can be used at a time.

### `clickup time-entry get`

Get a single time entry.

**API:** `GET /v2/team/{team_id}/time_entries/{timer_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--id` | string | *(required)* | `timer_id` (path) | Time entry ID |

### `clickup time-entry create`

Create a time entry.

**API:** `POST /v2/team/{team_id}/time_entries`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--start` | int64 | *(required)* | `start` (body) | Start time (Unix ms) |
| `--duration` | int64 | *(required)* | `duration` (body) | Duration in milliseconds |
| `--description` | string | — | `description` (body) | Description |
| `--task` | string | — | `tid` (body) | Task ID to associate with |
| `--billable` | bool | `false` | `billable` (body) | Mark as billable |

### `clickup time-entry update`

Update a time entry.

**API:** `PUT /v2/team/{team_id}/time_entries/{timer_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--id` | string | *(required)* | `timer_id` (path) | Time entry ID |
| `--description` | string | — | `description` (body) | New description |
| `--task` | string | — | `tid` (body) | New task ID |
| `--tag-action` | string | — | `tag_action` (body) | Tag action: `add` or `replace` |

### `clickup time-entry delete`

Delete a time entry.

**API:** `DELETE /v2/team/{team_id}/time_entries/{timer_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--id` | string | *(required)* | `timer_id` (path) | Time entry ID |

### `clickup time-entry start`

Start a timer for the authenticated user.

**API:** `POST /v2/team/{team_id}/time_entries/start`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--task` | string | — | `tid` (body) | Task ID to track time against |
| `--description` | string | — | `description` (body) | Description |
| `--billable` | bool | `false` | `billable` (body) | Mark as billable |

### `clickup time-entry stop`

Stop the currently running timer.

**API:** `POST /v2/team/{team_id}/time_entries/stop`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |

### `clickup time-entry current`

Get the currently running timer.

**API:** `GET /v2/team/{team_id}/time_entries/current`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--assignee` | string | — | `assignee` (query) | User ID (defaults to authenticated user) |

---

## Webhooks

### `clickup webhook list`

List webhooks for a workspace. Returns only webhooks created by the authenticated user.

**API:** `GET /v2/team/{team_id}/webhook`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |

### `clickup webhook create`

Create a webhook.

**API:** `POST /v2/team/{team_id}/webhook`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--endpoint` | string | *(required)* | `endpoint` (body) | Webhook URL |
| `--events` | string[] | *(required)* | `events` (body) | Events to subscribe to (use `*` for all) |

### `clickup webhook update`

Update a webhook.

**API:** `PUT /v2/webhook/{webhook_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `webhook_id` (path) | Webhook ID (UUID) |
| `--endpoint` | string | — | `endpoint` (body) | New webhook URL |
| `--events` | string | — | `events` (body) | Events (use `*` for all) |
| `--status` | string | — | `status` (body) | Status: `active` or `inactive` |

### `clickup webhook delete`

Delete a webhook.

**API:** `DELETE /v2/webhook/{webhook_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `webhook_id` (path) | Webhook ID (UUID) |

---

## Views

### `clickup view list`

List views at workspace, space, folder, or list level. If no scope flag is provided, lists workspace-level views.

**API:** `GET /v2/team/{team_id}/view` · `GET /v2/space/{space_id}/view` · `GET /v2/folder/{folder_id}/view` · `GET /v2/list/{list_id}/view`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID (default scope) |
| `--space` | string | — | `space_id` (path) | Space ID |
| `--folder` | string | — | `folder_id` (path) | Folder ID |
| `--list` | string | — | `list_id` (path) | List ID |

### `clickup view get`

Get a view by ID.

**API:** `GET /v2/view/{view_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `view_id` (path) | View ID |

### `clickup view create`

Create a view. Scope is determined by which flag is provided (defaults to workspace level).

**API:** `POST /v2/team/{team_id}/view` · `POST /v2/space/{space_id}/view` · `POST /v2/folder/{folder_id}/view` · `POST /v2/list/{list_id}/view`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--name` | string | *(required)* | `name` (body) | View name |
| `--type` | string | *(required)* | `type` (body) | View type: `list`, `board`, `calendar`, `table`, `timeline`, `workload`, `activity`, `map`, `conversation`, `gantt` |
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID (default scope) |
| `--space` | string | — | `space_id` (path) | Space ID |
| `--folder` | string | — | `folder_id` (path) | Folder ID |
| `--list` | string | — | `list_id` (path) | List ID |

### `clickup view update`

Update a view.

**API:** `PUT /v2/view/{view_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `view_id` (path) | View ID |
| `--name` | string | — | `name` (body) | New name |

### `clickup view delete`

Delete a view.

**API:** `DELETE /v2/view/{view_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `view_id` (path) | View ID |

### `clickup view tasks`

Get tasks in a view.

**API:** `GET /v2/view/{view_id}/task`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `view_id` (path) | View ID |
| `--page` | int | `0` | `page` (query) | Page number |

---

## Goals

### `clickup goal list`

List goals in a workspace.

**API:** `GET /v2/team/{team_id}/goal`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--include-completed` | bool | `false` | `include_completed` (query) | Include completed goals |

### `clickup goal get`

Get a goal by ID.

**API:** `GET /v2/goal/{goal_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `goal_id` (path) | Goal ID (UUID) |

### `clickup goal create`

Create a goal.

**API:** `POST /v2/team/{team_id}/goal`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--name` | string | *(required)* | `name` (body) | Goal name |
| `--due-date` | int64 | `0` | `due_date` (body) | Due date (Unix ms) |
| `--description` | string | — | `description` (body) | Description |
| `--color` | string | — | `color` (body) | Color hex (e.g. `#FF0000`) |
| `--multiple-owners` | bool | `false` | `multiple_owners` (body) | Allow multiple owners |

### `clickup goal update`

Update a goal.

**API:** `PUT /v2/goal/{goal_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `goal_id` (path) | Goal ID (UUID) |
| `--name` | string | — | `name` (body) | New name |
| `--description` | string | — | `description` (body) | New description |
| `--color` | string | — | `color` (body) | New color hex |

### `clickup goal delete`

Delete a goal.

**API:** `DELETE /v2/goal/{goal_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `goal_id` (path) | Goal ID (UUID) |

---

## Members

### `clickup member list`

List members of a list or task.

**API:** `GET /v2/list/{list_id}/member` or `GET /v2/task/{task_id}/member`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--list` | string | — | `list_id` (path) | List ID (use one of `--list` or `--task`) |
| `--task` | string | — | `task_id` (path) | Task ID |

---

## Groups (User Groups / Teams)

### `clickup group list`

List user groups in a workspace.

**API:** `GET /v2/group?team_id={team_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (query) | Workspace ID |

### `clickup group create`

Create a user group.

**API:** `POST /v2/team/{team_id}/group`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--name` | string | *(required)* | `name` (body) | Group name |
| `--handle` | string | — | `handle` (body) | Group handle/slug |

### `clickup group delete`

Delete a user group.

**API:** `DELETE /v2/group/{group_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--id` | string | *(required)* | `group_id` (path) | Group ID |

---

## Guests

### `clickup guest invite`

Invite a guest to a workspace.

**API:** `POST /v2/team/{team_id}/guest`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--email` | string | *(required)* | `email` (body) | Guest email address |

### `clickup guest get`

Get guest details.

**API:** `GET /v2/team/{team_id}/guest/{guest_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--id` | string | *(required)* | `guest_id` (path) | Guest ID |

### `clickup guest remove`

Remove a guest from a workspace.

**API:** `DELETE /v2/team/{team_id}/guest/{guest_id}`

| Flag | Type | Default | API Param | Description |
|------|------|---------|-----------|-------------|
| `--workspace` | string | *(global)* | `team_id` (path) | Workspace ID |
| `--id` | string | *(required)* | `guest_id` (path) | Guest ID |
