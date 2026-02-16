# ClickUp CLI — Full Audit Report

**Generated**: 2026-02-16  
**Spec**: `clickup-api-v2-reference.json` (135 endpoints)  
**CLI**: `clickup-cli/` (Go, Cobra)

---

## Executive Summary

| Metric | Value |
|--------|-------|
| Endpoints in spec | 135 |
| Endpoints implemented | 134 |
| Endpoints missing | 1 |
| Total parameters (query + body) | 526 |
| Directly implemented (individual flags) | 443 |
| Covered via JSON flags (nested objects) | 79 |
| Field name mismatch (implemented but wrong name) | 1 |
| Truly missing | 3 |
| **Coverage rate** | **99.2%** |

### Key Findings

1. **134/135 endpoints implemented** — only OAuth token endpoint missing (expected, CLI uses API key auth)
2. **98.7% parameter coverage** — nearly all spec params have CLI flags
3. **1 field name mismatch**: `POST /v2/task/{task_id}/merge` — spec says `source_task_ids`, CLI sends `merge_with`
4. **3 truly missing parameters** (see details below)

> **JSON flags note**: Many ClickUp parameters are deeply nested objects (e.g., `features.due_dates.enabled`,
> `grouping.dir`). The CLI handles these via JSON string flags like `--features '{"due_dates":{"enabled":true}}'`.
> These are fully functional but pass the entire object rather than exposing each sub-field as a flag.

---

## Missing Endpoints

| Method | Path | Summary | Tag |
|--------|------|---------|-----|
| POST | `/v2/oauth/token` | Get Access Token | Authorization |

*The OAuth token endpoint is intentionally omitted — the CLI authenticates via API key.*

---

## Issues Found

### 🐛 Field Name Mismatch

**`POST /v2/task/{task_id}/merge`** — Merge Tasks
- Spec requires: `source_task_ids` (array of task IDs to merge)
- CLI sends: `merge_with` (via `--merge-with` flag)
- **Impact**: This may cause the merge API call to fail or be ignored by ClickUp
- **Fix**: Rename `MergeWith` → `SourceTaskIDs` with `json:"source_task_ids"` in `internal/api/tasks.go`

---

## Full Parameter Coverage

### Attachments

**✅ `POST` `/v2/task/{task_id}/attachment`** — Create Task Attachment
  - API: `internal/api/attachments.go`
  - Params: 3/3 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `attachment`

### Authorization

**❌ `POST` `/v2/oauth/token`** — Get Access Token
  - Params: 0/3 (0%)
  - Body: `client_id`, `client_secret`, `code`
  - **⚠️ MISSING**: `client_id`, `client_secret`, `code`

**✅ `GET` `/v2/user`** — Get Authorized User
  - API: `internal/api/auth.go`
  - Params: —

### Comments

**✅ `GET` `/v2/task/{task_id}/comment`** — Get Task Comments
  - API: `internal/api/comments.go`
  - Params: 4/4 (100%)
  - Query: `custom_task_ids`, `team_id`, `start`, `start_id`

**✅ `POST` `/v2/task/{task_id}/comment`** — Create Task Comment
  - API: `internal/api/comments.go`
  - Params: 6/6 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `comment_text`, `assignee`, `group_assignee`, `notify_all`

**✅ `GET` `/v2/view/{view_id}/comment`** — Get Chat View Comments
  - API: `internal/api/comments.go`
  - Params: 2/2 (100%)
  - Query: `start`, `start_id`

**✅ `POST` `/v2/view/{view_id}/comment`** — Create Chat View Comment
  - API: `internal/api/comments.go`
  - Params: 2/2 (100%)
  - Body: `comment_text`, `notify_all`

**✅ `GET` `/v2/list/{list_id}/comment`** — Get List Comments
  - API: `internal/api/comments.go`
  - Params: 2/2 (100%)
  - Query: `start`, `start_id`

**✅ `POST` `/v2/list/{list_id}/comment`** — Create List Comment
  - API: `internal/api/comments.go`
  - Params: 3/3 (100%)
  - Body: `comment_text`, `assignee`, `notify_all`

**✅ `PUT` `/v2/comment/{comment_id}`** — Update Comment
  - API: `internal/api/comments.go`
  - Params: 4/4 (100%)
  - Body: `comment_text`, `assignee`, `group_assignee`, `resolved`

**✅ `DELETE` `/v2/comment/{comment_id}`** — Delete Comment
  - API: `internal/api/comments.go`
  - Params: —

**✅ `GET` `/v2/comment/{comment_id}/reply`** — Get Threaded Comments
  - API: `internal/api/comments.go`
  - Params: —

**✅ `POST` `/v2/comment/{comment_id}/reply`** — Create Threaded Comment
  - API: `internal/api/comments.go`
  - Params: —

### Custom Fields

**✅ `GET` `/v2/list/{list_id}/field`** — Get List Custom Fields
  - API: `internal/api/custom_fields.go`
  - Params: —

**✅ `GET` `/v2/folder/{folder_id}/field`** — Get Folder Custom Fields
  - API: `internal/api/custom_fields.go`
  - Params: —

**✅ `GET` `/v2/space/{space_id}/field`** — Get Space Custom Fields
  - API: `internal/api/custom_fields.go`
  - Params: —

**✅ `GET` `/v2/team/{team_id}/field`** — Get Workspace Custom Fields
  - API: `internal/api/custom_fields.go`
  - Params: —

**✅ `POST` `/v2/task/{task_id}/field/{field_id}`** — Set Custom Field Value
  - API: `internal/api/custom_fields.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

**✅ `DELETE` `/v2/task/{task_id}/field/{field_id}`** — Remove Custom Field Value
  - API: `internal/api/custom_fields.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

### Custom Task Types

**✅ `GET` `/v2/team/{team_id}/custom_item`** — Get Custom Task Types
  - API: `internal/api/custom_task_types.go`
  - Params: —

### Folders

**✅ `GET` `/v2/space/{space_id}/folder`** — Get Folders
  - API: `internal/api/templates.go`
  - Params: 1/1 (100%)
  - Query: `archived`

**✅ `POST` `/v2/space/{space_id}/folder`** — Create Folder
  - API: `internal/api/templates.go`
  - Params: 1/1 (100%)
  - Body: `name`

**✅ `GET` `/v2/folder/{folder_id}`** — Get Folder
  - API: `internal/api/lists.go`
  - Params: —

**✅ `PUT` `/v2/folder/{folder_id}`** — Update Folder
  - API: `internal/api/lists.go`
  - Params: 1/1 (100%)
  - Body: `name`

**✅ `DELETE` `/v2/folder/{folder_id}`** — Delete Folder
  - API: `internal/api/lists.go`
  - Params: —

**✅ `POST` `/v2/space/{space_id}/folder_template/{template_id}`** — Create Folder from template
  - API: `internal/api/templates.go`
  - Params: 32/32 (100%)
  - Body: `name`, `options`
  - Nested: 30 sub-fields (via JSON flag)

### Goals

**✅ `GET` `/v2/team/{team_id}/goal`** — Get Goals
  - API: `internal/api/goals.go`
  - Params: 1/1 (100%)
  - Query: `include_completed`

**✅ `POST` `/v2/team/{team_id}/goal`** — Create Goal
  - API: `internal/api/goals.go`
  - Params: 6/6 (100%)
  - Body: `name`, `due_date`, `description`, `multiple_owners`, `owners`, `color`

**✅ `GET` `/v2/goal/{goal_id}`** — Get Goal
  - API: `internal/api/goals.go`
  - Params: —

**✅ `PUT` `/v2/goal/{goal_id}`** — Update Goal
  - API: `internal/api/goals.go`
  - Params: 6/6 (100%)
  - Body: `name`, `due_date`, `description`, `rem_owners`, `add_owners`, `color`

**✅ `DELETE` `/v2/goal/{goal_id}`** — Delete Goal
  - API: `internal/api/goals.go`
  - Params: —

**✅ `POST` `/v2/goal/{goal_id}/key_result`** — Create Key Result
  - API: `internal/api/goals.go`
  - Params: 8/8 (100%)
  - Body: `name`, `owners`, `type`, `steps_start`, `steps_end`, `unit`, `task_ids`, `list_ids`

**✅ `PUT` `/v2/key_result/{key_result_id}`** — Edit Key Result
  - API: `internal/api/goals.go`
  - Params: 2/2 (100%)
  - Body: `steps_current`, `note`

**✅ `DELETE` `/v2/key_result/{key_result_id}`** — Delete Key Result
  - API: `internal/api/goals.go`
  - Params: —

### Guests

**✅ `POST` `/v2/team/{team_id}/guest`** — Invite Guest To Workspace
  - API: `internal/api/members.go`
  - Params: 7/7 (100%)
  - Body: `email`, `can_edit_tags`, `can_see_time_spent`, `can_see_time_estimated`, `can_create_views`, `can_see_points_estimated`, `custom_role_id`

**✅ `GET` `/v2/team/{team_id}/guest/{guest_id}`** — Get Guest
  - API: `internal/api/members.go`
  - Params: —

**✅ `PUT` `/v2/team/{team_id}/guest/{guest_id}`** — Edit Guest On Workspace
  - API: `internal/api/members.go`
  - Params: 6/6 (100%)
  - Body: `can_see_points_estimated`, `can_edit_tags`, `can_see_time_spent`, `can_see_time_estimated`, `can_create_views`, `custom_role_id`

**✅ `DELETE` `/v2/team/{team_id}/guest/{guest_id}`** — Remove Guest From Workspace
  - API: `internal/api/members.go`
  - Params: —

**✅ `POST` `/v2/task/{task_id}/guest/{guest_id}`** — Add Guest To Task
  - API: `internal/api/members.go`
  - Params: 4/4 (100%)
  - Query: `include_shared`, `custom_task_ids`, `team_id`
  - Body: `permission_level`

**✅ `DELETE` `/v2/task/{task_id}/guest/{guest_id}`** — Remove Guest From Task
  - API: `internal/api/members.go`
  - Params: 3/3 (100%)
  - Query: `include_shared`, `custom_task_ids`, `team_id`

**✅ `POST` `/v2/list/{list_id}/guest/{guest_id}`** — Add Guest To List
  - API: `internal/api/members.go`
  - Params: 2/2 (100%)
  - Query: `include_shared`
  - Body: `permission_level`

**✅ `DELETE` `/v2/list/{list_id}/guest/{guest_id}`** — Remove Guest From List
  - API: `internal/api/members.go`
  - Params: 1/1 (100%)
  - Query: `include_shared`

**✅ `POST` `/v2/folder/{folder_id}/guest/{guest_id}`** — Add Guest To Folder
  - API: `internal/api/members.go`
  - Params: 2/2 (100%)
  - Query: `include_shared`
  - Body: `permission_level`

**✅ `DELETE` `/v2/folder/{folder_id}/guest/{guest_id}`** — Remove Guest From Folder
  - API: `internal/api/members.go`
  - Params: 1/1 (100%)
  - Query: `include_shared`

### Lists

**✅ `GET` `/v2/folder/{folder_id}/list`** — Get Lists
  - API: `internal/api/lists.go`
  - Params: 1/1 (100%)
  - Query: `archived`

**✅ `POST` `/v2/folder/{folder_id}/list`** — Create List
  - API: `internal/api/lists.go`
  - Params: 8/8 (100%)
  - Body: `name`, `content`, `markdown_content`, `due_date`, `due_date_time`, `priority`, `assignee`, `status`

**✅ `GET` `/v2/space/{space_id}/list`** — Get Folderless Lists
  - API: `internal/api/lists.go`
  - Params: 1/1 (100%)
  - Query: `archived`

**✅ `POST` `/v2/space/{space_id}/list`** — Create Folderless List
  - API: `internal/api/lists.go`
  - Params: 8/8 (100%)
  - Body: `name`, `content`, `markdown_content`, `due_date`, `due_date_time`, `priority`, `assignee`, `status`

**✅ `GET` `/v2/list/{list_id}`** — Get List
  - API: `internal/api/tasks.go`
  - Params: —

**✅ `PUT` `/v2/list/{list_id}`** — Update List
  - API: `internal/api/tasks.go`
  - Params: 9/9 (100%)
  - Body: `name`, `content`, `markdown_content`, `due_date`, `due_date_time`, `priority`, `assignee`, `status`, `unset_status`

**✅ `DELETE` `/v2/list/{list_id}`** — Delete List
  - API: `internal/api/tasks.go`
  - Params: —

**✅ `POST` `/v2/list/{list_id}/task/{task_id}`** — Add Task To List
  - API: `internal/api/tasks.go`
  - Params: —

**✅ `DELETE` `/v2/list/{list_id}/task/{task_id}`** — Remove Task From List
  - API: `internal/api/tasks.go`
  - Params: —

**✅ `POST` `/v2/folder/{folder_id}/list_template/{template_id}`** — Create List From Template in Folder
  - API: `internal/api/templates.go`
  - Params: 2/2 (100%)
  - Body: `name`, `options`

**✅ `POST` `/v2/space/{space_id}/list_template/{template_id}`** — Create List From Template in Space.
  - API: `internal/api/templates.go`
  - Params: 32/32 (100%)
  - Body: `name`, `options`
  - Nested: 30 sub-fields (via JSON flag)

### Members

**✅ `GET` `/v2/task/{task_id}/member`** — Get Task Members
  - API: `internal/api/members.go`
  - Params: —

**✅ `GET` `/v2/list/{list_id}/member`** — Get List Members
  - API: `internal/api/members.go`
  - Params: —

### Roles

**✅ `GET` `/v2/team/{team_id}/customroles`** — Get Custom Roles
  - API: `internal/api/roles.go`
  - Params: 1/1 (100%)
  - Query: `include_members`

### Shared Hierarchy

**✅ `GET` `/v2/team/{team_id}/shared`** — Shared Hierarchy
  - API: `internal/api/shared.go`
  - Params: —

### Spaces

**✅ `GET` `/v2/team/{team_id}/space`** — Get Spaces
  - API: `internal/api/spaces.go`
  - Params: 1/1 (100%)
  - Query: `archived`

**✅ `POST` `/v2/team/{team_id}/space`** — Create Space
  - API: `internal/api/spaces.go`
  - Params: 3/3 (100%)
  - Body: `name`, `multiple_assignees`, `features`

**✅ `GET` `/v2/space/{space_id}`** — Get Space
  - API: `internal/api/lists.go`
  - Params: —

**✅ `PUT` `/v2/space/{space_id}`** — Update Space
  - API: `internal/api/lists.go`
  - Params: 27/27 (100%)
  - Body: `name`, `color`, `private`, `admin_can_manage`, `multiple_assignees`, `features`
  - Nested: 21 sub-fields (via JSON flag)

**✅ `DELETE` `/v2/space/{space_id}`** — Delete Space
  - API: `internal/api/lists.go`
  - Params: —

### Tags

**✅ `GET` `/v2/space/{space_id}/tag`** — Get Space Tags
  - API: `internal/api/tags.go`
  - Params: —

**✅ `POST` `/v2/space/{space_id}/tag`** — Create Space Tag
  - API: `internal/api/tags.go`
  - Params: 4/4 (100%)
  - Body: `tag`
  - Nested: 3 sub-fields (via JSON flag)

**✅ `PUT` `/v2/space/{space_id}/tag/{tag_name}`** — Edit Space Tag
  - API: `internal/api/tags.go`
  - Params: 4/4 (100%)
  - Body: `tag`
  - Nested: 3 sub-fields (via JSON flag)

**✅ `DELETE` `/v2/space/{space_id}/tag/{tag_name}`** — Delete Space Tag
  - API: `internal/api/tags.go`
  - Params: 1/1 (100%)
  - Body: `tag`

**✅ `POST` `/v2/task/{task_id}/tag/{tag_name}`** — Add Tag To Task
  - API: `internal/api/tags.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

**✅ `DELETE` `/v2/task/{task_id}/tag/{tag_name}`** — Remove Tag From Task
  - API: `internal/api/tags.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

### Task Checklists

**✅ `POST` `/v2/task/{task_id}/checklist`** — Create Checklist
  - API: `internal/api/checklists.go`
  - Params: 3/3 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `name`

**✅ `PUT` `/v2/checklist/{checklist_id}`** — Edit Checklist
  - API: `internal/api/checklists.go`
  - Params: 2/2 (100%)
  - Body: `name`, `position`

**✅ `DELETE` `/v2/checklist/{checklist_id}`** — Delete Checklist
  - API: `internal/api/checklists.go`
  - Params: —

**✅ `POST` `/v2/checklist/{checklist_id}/checklist_item`** — Create Checklist Item
  - API: `internal/api/checklists.go`
  - Params: 2/2 (100%)
  - Body: `name`, `assignee`

**✅ `PUT` `/v2/checklist/{checklist_id}/checklist_item/{checklist_item_id}`** — Edit Checklist Item
  - API: `internal/api/checklists.go`
  - Params: 4/4 (100%)
  - Body: `name`, `assignee`, `resolved`, `parent`

**✅ `DELETE` `/v2/checklist/{checklist_id}/checklist_item/{checklist_item_id}`** — Delete Checklist Item
  - API: `internal/api/checklists.go`
  - Params: —

### Task Relationships

**✅ `POST` `/v2/task/{task_id}/dependency`** — Add Dependency
  - API: `internal/api/relationships.go`
  - Params: 4/4 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `depends_on`, `depedency_of`

**✅ `DELETE` `/v2/task/{task_id}/dependency`** — Delete Dependency
  - API: `internal/api/relationships.go`
  - Params: 4/4 (100%)
  - Query: `depends_on`, `dependency_of`, `custom_task_ids`, `team_id`

**✅ `POST` `/v2/task/{task_id}/link/{links_to}`** — Add Task Link
  - API: `internal/api/relationships.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

**✅ `DELETE` `/v2/task/{task_id}/link/{links_to}`** — Delete Task Link
  - API: `internal/api/relationships.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

### Tasks

**✅ `GET` `/v2/list/{list_id}/task`** — Get Tasks
  - API: `internal/api/tasks.go`
  - Params: 23/23 (100%)
  - Query: `archived`, `include_markdown_description`, `page`, `order_by`, `reverse`, `subtasks`, `statuses`, `include_closed`, `include_timl`, `assignees`, `watchers`, `tags`, `due_date_gt`, `due_date_lt`, `date_created_gt`, `date_created_lt`, `date_updated_gt`, `date_updated_lt`, `date_done_gt`, `date_done_lt`, `custom_fields`, `custom_field`, `custom_items`

**✅ `POST` `/v2/list/{list_id}/task`** — Create Task
  - API: `internal/api/tasks.go`
  - Params: 21/21 (100%)
  - Body: `name`, `description`, `assignees`, `archived`, `group_assignees`, `tags`, `status`, `priority`, `due_date`, `due_date_time`, `time_estimate`, `start_date`, `start_date_time`, `points`, `notify_all`, `parent`, `markdown_content`, `links_to`, `check_required_custom_fields`, `custom_fields`, `custom_item_id`

**✅ `GET` `/v2/task/{task_id}`** — Get Task
  - API: `internal/api/attachments.go`
  - Params: 5/5 (100%)
  - Query: `custom_task_ids`, `team_id`, `include_subtasks`, `include_markdown_description`, `custom_fields`

**✅ `PUT` `/v2/task/{task_id}`** — Update Task
  - API: `internal/api/attachments.go`
  - Params: 25/25 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `custom_item_id`, `name`, `description`, `markdown_content`, `status`, `priority`, `due_date`, `due_date_time`, `parent`, `time_estimate`, `start_date`, `start_date_time`, `points`, `assignees`, `group_assignees`, `watchers`, `archived`
  - Nested: 6 sub-fields (via JSON flag)

**✅ `DELETE` `/v2/task/{task_id}`** — Delete Task
  - API: `internal/api/attachments.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

**✅ `GET` `/v2/team/{team_Id}/task`** — Get Filtered Team Tasks
  - API: `internal/api/tasks.go`
  - Params: 23/23 (100%)
  - Query: `page`, `order_by`, `reverse`, `subtasks`, `space_ids[]`, `project_ids[]`, `list_ids[]`, `statuses[]`, `include_closed`, `assignees[]`, `tags[]`, `due_date_gt`, `due_date_lt`, `date_created_gt`, `date_created_lt`, `date_updated_gt`, `date_updated_lt`, `date_done_gt`, `date_done_lt`, `custom_fields`, `parent`, `include_markdown_description`, `custom_items[]`

**✅ `POST` `/v2/task/{task_id}/merge`** — Merge Tasks
  - API: `internal/api/tasks.go`
  - Params: 0/1 (0%)
  - Body: `source_task_ids`
  - **🐛 MISMATCH**: `source_task_ids`

**✅ `GET` `/v2/task/{task_id}/time_in_status`** — Get Task's Time in Status
  - API: `internal/api/tasks.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

**✅ `GET` `/v2/task/bulk_time_in_status/task_ids`** — Get Bulk Tasks' Time in Status
  - API: `internal/api/tasks.go`
  - Params: 3/3 (100%)
  - Query: `task_ids`, `custom_task_ids`, `team_id`

**✅ `POST` `/v2/list/{list_id}/taskTemplate/{template_id}`** — Create Task From Template
  - API: `internal/api/templates.go`
  - Params: 1/1 (100%)
  - Body: `name`

### Templates

**✅ `GET` `/v2/team/{team_id}/taskTemplate`** — Get Task Templates
  - API: `internal/api/templates.go`
  - Params: 1/1 (100%)
  - Query: `page`

### Time Tracking

**✅ `GET` `/v2/team/{team_Id}/time_entries`** — Get time entries within a date range
  - API: `internal/api/time_entries.go`
  - Params: 14/14 (100%)
  - Query: `start_date`, `end_date`, `assignee`, `include_task_tags`, `include_location_names`, `include_approval_history`, `include_approval_details`, `space_id`, `folder_id`, `list_id`, `task_id`, `custom_task_ids`, `team_id`, `is_billable`

**✅ `POST` `/v2/team/{team_Id}/time_entries`** — Create a time entry
  - API: `internal/api/time_entries.go`
  - Params: 14/14 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `description`, `tags`, `start`, `stop`, `end`, `billable`, `duration`, `assignee`, `tid`
  - Nested: 3 sub-fields (via JSON flag)

**✅ `GET` `/v2/team/{team_id}/time_entries/{timer_id}`** — Get singular time entry
  - API: `internal/api/time_entries.go`
  - Params: 4/4 (100%)
  - Query: `include_task_tags`, `include_location_names`, `include_approval_history`, `include_approval_details`

**✅ `DELETE` `/v2/team/{team_id}/time_entries/{timer_id}`** — Delete a time Entry
  - API: `internal/api/time_entries.go`
  - Params: —

**✅ `PUT` `/v2/team/{team_id}/time_entries/{timer_id}`** — Update a time Entry
  - API: `internal/api/time_entries.go`
  - Params: 10/10 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `description`, `tags`, `tag_action`, `start`, `end`, `tid`, `billable`, `duration`

**✅ `GET` `/v2/team/{team_id}/time_entries/{timer_id}/history`** — Get time entry history
  - API: `internal/api/time_entries.go`
  - Params: —

**✅ `GET` `/v2/team/{team_id}/time_entries/current`** — Get running time entry
  - API: `internal/api/time_entries.go`
  - Params: 1/1 (100%)
  - Query: `assignee`

**✅ `DELETE` `/v2/team/{team_id}/time_entries/tags`** — Remove tags from time entries
  - API: `internal/api/time_entries.go`
  - Params: 2/2 (100%)
  - Body: `time_entry_ids`, `tags`

**✅ `GET` `/v2/team/{team_id}/time_entries/tags`** — Get all tags from time entries
  - API: `internal/api/time_entries.go`
  - Params: —

**✅ `POST` `/v2/team/{team_id}/time_entries/tags`** — Add tags from time entries
  - API: `internal/api/time_entries.go`
  - Params: 2/2 (100%)
  - Body: `time_entry_ids`, `tags`

**✅ `PUT` `/v2/team/{team_id}/time_entries/tags`** — Change tag names from time entries
  - API: `internal/api/time_entries.go`
  - Params: 4/4 (100%)
  - Body: `name`, `new_name`, `tag_bg`, `tag_fg`

**✅ `POST` `/v2/team/{team_Id}/time_entries/start`** — Start a time Entry
  - API: `internal/api/time_entries.go`
  - Params: 7/7 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `description`, `tags`, `tid`, `billable`
  - Nested: 1 sub-fields (via JSON flag)

**✅ `POST` `/v2/team/{team_id}/time_entries/stop`** — Stop a time Entry
  - API: `internal/api/time_entries.go`
  - Params: —

### Time Tracking (Legacy)

**✅ `GET` `/v2/task/{task_id}/time`** — Get tracked time
  - API: `internal/api/tasks.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

**✅ `POST` `/v2/task/{task_id}/time`** — Track time
  - API: `internal/api/tasks.go`
  - Params: 5/5 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `start`, `end`, `time`

**✅ `PUT` `/v2/task/{task_id}/time/{interval_id}`** — Edit time tracked
  - API: `internal/api/time_tracking_legacy.go`
  - Params: 5/5 (100%)
  - Query: `custom_task_ids`, `team_id`
  - Body: `start`, `end`, `time`

**✅ `DELETE` `/v2/task/{task_id}/time/{interval_id}`** — Delete time tracked
  - API: `internal/api/time_tracking_legacy.go`
  - Params: 2/2 (100%)
  - Query: `custom_task_ids`, `team_id`

### User Groups

**✅ `POST` `/v2/team/{team_id}/group`** — Create Group
  - API: `internal/api/members.go`
  - Params: 3/3 (100%)
  - Body: `name`, `handle`, `members`

**✅ `PUT` `/v2/group/{group_id}`** — Update Group
  - API: `internal/api/members.go`
  - Params: 5/5 (100%)
  - Body: `name`, `handle`, `members`
  - Nested: 2 sub-fields (via JSON flag)

**✅ `DELETE` `/v2/group/{group_id}`** — Delete Group
  - API: `internal/api/members.go`
  - Params: —

**✅ `GET` `/v2/group`** — Get Groups
  - API: `internal/api/members.go`
  - Params: 2/2 (100%)
  - Query: `team_id`, `group_ids`

### Users

**✅ `POST` `/v2/team/{team_id}/user`** — Invite User To Workspace
  - API: `internal/api/users.go`
  - Params: 3/3 (100%)
  - Body: `email`, `admin`, `custom_role_id`

**✅ `GET` `/v2/team/{team_id}/user/{user_id}`** — Get User
  - API: `internal/api/users.go`
  - Params: 1/1 (100%)
  - Query: `include_shared`

**✅ `PUT` `/v2/team/{team_id}/user/{user_id}`** — Edit User On Workspace
  - API: `internal/api/users.go`
  - Params: 3/3 (100%)
  - Body: `username`, `admin`, `custom_role_id`

**✅ `DELETE` `/v2/team/{team_id}/user/{user_id}`** — Remove User From Workspace
  - API: `internal/api/users.go`
  - Params: —

### Views

**✅ `GET` `/v2/team/{team_id}/view`** — Get Workspace (Everything level) Views
  - API: `internal/api/views.go`
  - Params: —

**✅ `POST` `/v2/team/{team_id}/view`** — Create Workspace (Everything level) View
  - API: `internal/api/views.go`
  - Params: 9/9 (100%)
  - Body: `name`, `type`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`

**✅ `GET` `/v2/space/{space_id}/view`** — Get Space Views
  - API: `internal/api/views.go`
  - Params: —

**✅ `POST` `/v2/space/{space_id}/view`** — Create Space View
  - API: `internal/api/views.go`
  - Params: 9/9 (100%)
  - Body: `name`, `type`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`

**✅ `GET` `/v2/folder/{folder_id}/view`** — Get Folder Views
  - API: `internal/api/views.go`
  - Params: —

**✅ `POST` `/v2/folder/{folder_id}/view`** — Create Folder View
  - API: `internal/api/views.go`
  - Params: 9/9 (100%)
  - Body: `name`, `type`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`

**✅ `GET` `/v2/list/{list_id}/view`** — Get List Views
  - API: `internal/api/views.go`
  - Params: —

**✅ `POST` `/v2/list/{list_id}/view`** — Create List View
  - API: `internal/api/views.go`
  - Params: 9/9 (100%)
  - Body: `name`, `type`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`

**✅ `GET` `/v2/view/{view_id}`** — Get View
  - API: `internal/api/comments.go`
  - Params: —

**✅ `PUT` `/v2/view/{view_id}`** — Update View
  - API: `internal/api/comments.go`
  - Params: 38/38 (100%)
  - Body: `name`, `type`, `parent`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`
  - Nested: 28 sub-fields (via JSON flag)

**✅ `DELETE` `/v2/view/{view_id}`** — Delete View
  - API: `internal/api/comments.go`
  - Params: —

**✅ `GET` `/v2/view/{view_id}/task`** — Get View Tasks
  - API: `internal/api/views.go`
  - Params: 1/1 (100%)
  - Query: `page`

### Webhooks

**✅ `GET` `/v2/team/{team_id}/webhook`** — Get Webhooks
  - API: `internal/api/webhooks.go`
  - Params: —

**✅ `POST` `/v2/team/{team_id}/webhook`** — Create Webhook
  - API: `internal/api/webhooks.go`
  - Params: 6/6 (100%)
  - Body: `endpoint`, `events`, `space_id`, `folder_id`, `list_id`, `task_id`

**✅ `PUT` `/v2/webhook/{webhook_id}`** — Update Webhook
  - API: `internal/api/webhooks.go`
  - Params: 3/3 (100%)
  - Body: `endpoint`, `events`, `status`

**✅ `DELETE` `/v2/webhook/{webhook_id}`** — Delete Webhook
  - API: `internal/api/webhooks.go`
  - Params: —

### Workspaces

**✅ `GET` `/v2/team`** — Get Authorized Workspaces
  - API: `internal/api/custom_task_types.go`
  - Params: —

**✅ `GET` `/v2/team/{team_id}/seats`** — Get Workspace seats
  - API: `internal/api/workspaces.go`
  - Params: —

**✅ `GET` `/v2/team/{team_id}/plan`** — Get Workspace Plan
  - API: `internal/api/workspaces.go`
  - Params: —

---

## All Gaps Summary

### Missing Parameters

- `POST` `/v2/oauth/token` (Get Access Token): `client_id`, `client_secret`, `code`

### Field Name Mismatches

- `POST` `/v2/task/{task_id}/merge` (Merge Tasks): `source_task_ids`