# ClickUp CLI Audit Report

Generated: 2026-02-16

## Summary

- **Endpoint coverage**: 134/135 implemented
- **Missing endpoint**: POST `/v2/oauth/token` (Get Access Token) — OAuth flow, not applicable to CLI
- **Total spec parameters** (query + body, excl. path/header): 526
- **Directly implemented** (individual flags): 440
- **Covered via JSON flags** (nested objects passed as `--flag '{}'`): 79
- **Truly missing**: 7

> **Note**: Many ClickUp API parameters are deeply nested objects (e.g., `features.due_dates.enabled`,
> `grouping.dir`, `settings.show_subtasks`). The CLI handles these via JSON string flags like
> `--features '{"due_dates":{"enabled":true}}'`. These are marked as "via JSON flag" below.

## Missing Endpoints

1 endpoint(s) not found in CLI:

| # | Method | Path | Summary | Tag |
|---|--------|------|---------|-----|
| 1 | POST | `/v2/oauth/token` | Get Access Token | Authorization |

## Parameter Coverage Per Endpoint

### Attachments

#### ✅ `POST` `/v2/task/{task_id}/attachment`
**Create Task Attachment** | `CreateTaskAttachment`
API: `internal/api/attachments.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `attachment`

### Authorization

#### ❌ `POST` `/v2/oauth/token`
**Get Access Token** | `GetAccessToken`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 0 |
| Via JSON flag | 0 |
| Missing | 3 |

**Body params**: `client_id`, `client_secret`, `code`
**⚠️ Missing**: `client_id`, `client_secret`, `code`

#### ✅ `GET` `/v2/user`
**Get Authorized User** | `GetAuthorizedUser`
API: `internal/api/auth.go`

No parameters (besides path params).

### Comments

#### ✅ `GET` `/v2/task/{task_id}/comment`
**Get Task Comments** | `GetTaskComments`
API: `internal/api/comments.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 4 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`, `start`, `start_id`

#### ✅ `POST` `/v2/task/{task_id}/comment`
**Create Task Comment** | `CreateTaskComment`
API: `internal/api/comments.go`

| Metric | Count |
|--------|-------|
| Total params | 6 |
| Direct flags | 6 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `comment_text`, `assignee`, `group_assignee`, `notify_all`

#### ✅ `GET` `/v2/view/{view_id}/comment`
**Get Chat View Comments** | `GetChatViewComments`
API: `internal/api/comments.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `start`, `start_id`

#### ✅ `POST` `/v2/view/{view_id}/comment`
**Create Chat View Comment** | `CreateChatViewComment`
API: `internal/api/comments.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `comment_text`, `notify_all`

#### ✅ `GET` `/v2/list/{list_id}/comment`
**Get List Comments** | `GetListComments`
API: `internal/api/comments.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `start`, `start_id`

#### ✅ `POST` `/v2/list/{list_id}/comment`
**Create List Comment** | `CreateListComment`
API: `internal/api/comments.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `comment_text`, `assignee`, `notify_all`

#### ✅ `PUT` `/v2/comment/{comment_id}`
**Update Comment** | `UpdateComment`
API: `internal/api/comments.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 4 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `comment_text`, `assignee`, `group_assignee`, `resolved`

#### ✅ `DELETE` `/v2/comment/{comment_id}`
**Delete Comment** | `DeleteComment`
API: `internal/api/comments.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/comment/{comment_id}/reply`
**Get Threaded Comments** | `GetThreadedComments`
API: `internal/api/comments.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/comment/{comment_id}/reply`
**Create Threaded Comment** | `CreateThreadedComment`
API: `internal/api/comments.go`

No parameters (besides path params).

### Custom Fields

#### ✅ `GET` `/v2/list/{list_id}/field`
**Get List Custom Fields** | `GetAccessibleCustomFields`
API: `internal/api/custom_fields.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/folder/{folder_id}/field`
**Get Folder Custom Fields** | `getFolderAvailableFields`
API: `internal/api/custom_fields.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/space/{space_id}/field`
**Get Space Custom Fields** | `getSpaceAvailableFields`
API: `internal/api/custom_fields.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/team/{team_id}/field`
**Get Workspace Custom Fields** | `getTeamAvailableFields`
API: `internal/api/custom_fields.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/task/{task_id}/field/{field_id}`
**Set Custom Field Value** | `SetCustomFieldValue`
API: `internal/api/custom_fields.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

#### ✅ `DELETE` `/v2/task/{task_id}/field/{field_id}`
**Remove Custom Field Value** | `RemoveCustomFieldValue`
API: `internal/api/custom_fields.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

### Custom Task Types

#### ✅ `GET` `/v2/team/{team_id}/custom_item`
**Get Custom Task Types** | `GetCustomItems`
API: `internal/api/custom_task_types.go`

No parameters (besides path params).

### Folders

#### ✅ `GET` `/v2/space/{space_id}/folder`
**Get Folders** | `GetFolders`
API: `internal/api/templates.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `archived`

#### ✅ `POST` `/v2/space/{space_id}/folder`
**Create Folder** | `CreateFolder`
API: `internal/api/templates.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`

#### ✅ `GET` `/v2/folder/{folder_id}`
**Get Folder** | `GetFolder`
API: `internal/api/lists.go`

No parameters (besides path params).

#### ✅ `PUT` `/v2/folder/{folder_id}`
**Update Folder** | `UpdateFolder`
API: `internal/api/lists.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`

#### ✅ `DELETE` `/v2/folder/{folder_id}`
**Delete Folder** | `DeleteFolder`
API: `internal/api/lists.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/space/{space_id}/folder_template/{template_id}`
**Create Folder from template** | `CreateFolderFromTemplate`
API: `internal/api/templates.go`

| Metric | Count |
|--------|-------|
| Total params | 32 |
| Direct flags | 12 |
| Via JSON flag | 20 |
| Missing | 0 |

**Body params**: `name`, `options`
**Nested body params**: `options.return_immediately`, `options.content`, `options.time_estimate`, `options.automation`, `options.include_views`, `options.old_due_date`, `options.old_start_date`, `options.old_followers`, `options.comment_attachments`, `options.recur_settings`, `options.old_tags`, `options.old_statuses`, `options.subtasks`, `options.custom_type`, `options.old_assignees`, `options.attachments`, `options.comment`, `options.old_status`, `options.external_dependencies`, `options.internal_dependencies`, `options.priority`, `options.custom_fields`, `options.old_checklists`, `options.relationships`, `options.old_subtask_assignees`, `options.start_date`, `options.due_date`, `options.remap_start_date`, `options.skip_weekends`, `options.archived`

### Goals

#### ✅ `GET` `/v2/team/{team_id}/goal`
**Get Goals** | `GetGoals`
API: `internal/api/goals.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_completed`

#### ✅ `POST` `/v2/team/{team_id}/goal`
**Create Goal** | `CreateGoal`
API: `internal/api/goals.go`

| Metric | Count |
|--------|-------|
| Total params | 6 |
| Direct flags | 6 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `due_date`, `description`, `multiple_owners`, `owners`, `color`

#### ✅ `GET` `/v2/goal/{goal_id}`
**Get Goal** | `GetGoal`
API: `internal/api/goals.go`

No parameters (besides path params).

#### ✅ `PUT` `/v2/goal/{goal_id}`
**Update Goal** | `UpdateGoal`
API: `internal/api/goals.go`

| Metric | Count |
|--------|-------|
| Total params | 6 |
| Direct flags | 6 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `due_date`, `description`, `rem_owners`, `add_owners`, `color`

#### ✅ `DELETE` `/v2/goal/{goal_id}`
**Delete Goal** | `DeleteGoal`
API: `internal/api/goals.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/goal/{goal_id}/key_result`
**Create Key Result** | `CreateKeyResult`
API: `internal/api/goals.go`

| Metric | Count |
|--------|-------|
| Total params | 8 |
| Direct flags | 8 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `owners`, `type`, `steps_start`, `steps_end`, `unit`, `task_ids`, `list_ids`

#### ✅ `PUT` `/v2/key_result/{key_result_id}`
**Edit Key Result** | `EditKeyResult`
API: `internal/api/goals.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `steps_current`, `note`

#### ✅ `DELETE` `/v2/key_result/{key_result_id}`
**Delete Key Result** | `DeleteKeyResult`
API: `internal/api/goals.go`

No parameters (besides path params).

### Guests

#### ✅ `POST` `/v2/team/{team_id}/guest`
**Invite Guest To Workspace** | `InviteGuestToWorkspace`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 7 |
| Direct flags | 7 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `email`, `can_edit_tags`, `can_see_time_spent`, `can_see_time_estimated`, `can_create_views`, `can_see_points_estimated`, `custom_role_id`

#### ✅ `GET` `/v2/team/{team_id}/guest/{guest_id}`
**Get Guest** | `GetGuest`
API: `internal/api/members.go`

No parameters (besides path params).

#### ✅ `PUT` `/v2/team/{team_id}/guest/{guest_id}`
**Edit Guest On Workspace** | `EditGuestOnWorkspace`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 6 |
| Direct flags | 6 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `can_see_points_estimated`, `can_edit_tags`, `can_see_time_spent`, `can_see_time_estimated`, `can_create_views`, `custom_role_id`

#### ✅ `DELETE` `/v2/team/{team_id}/guest/{guest_id}`
**Remove Guest From Workspace** | `RemoveGuestFromWorkspace`
API: `internal/api/members.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/task/{task_id}/guest/{guest_id}`
**Add Guest To Task** | `AddGuestToTask`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 4 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_shared`, `custom_task_ids`, `team_id`
**Body params**: `permission_level`

#### ✅ `DELETE` `/v2/task/{task_id}/guest/{guest_id}`
**Remove Guest From Task** | `RemoveGuestFromTask`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_shared`, `custom_task_ids`, `team_id`

#### ✅ `POST` `/v2/list/{list_id}/guest/{guest_id}`
**Add Guest To List** | `AddGuestToList`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_shared`
**Body params**: `permission_level`

#### ✅ `DELETE` `/v2/list/{list_id}/guest/{guest_id}`
**Remove Guest From List** | `RemoveGuestFromList`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_shared`

#### ✅ `POST` `/v2/folder/{folder_id}/guest/{guest_id}`
**Add Guest To Folder** | `AddGuestToFolder`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_shared`
**Body params**: `permission_level`

#### ✅ `DELETE` `/v2/folder/{folder_id}/guest/{guest_id}`
**Remove Guest From Folder** | `RemoveGuestFromFolder`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_shared`

### Lists

#### ✅ `GET` `/v2/folder/{folder_id}/list`
**Get Lists** | `GetLists`
API: `internal/api/lists.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `archived`

#### ✅ `POST` `/v2/folder/{folder_id}/list`
**Create List** | `CreateList`
API: `internal/api/lists.go`

| Metric | Count |
|--------|-------|
| Total params | 8 |
| Direct flags | 8 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `content`, `markdown_content`, `due_date`, `due_date_time`, `priority`, `assignee`, `status`

#### ✅ `GET` `/v2/space/{space_id}/list`
**Get Folderless Lists** | `GetFolderlessLists`
API: `internal/api/lists.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `archived`

#### ✅ `POST` `/v2/space/{space_id}/list`
**Create Folderless List** | `CreateFolderlessList`
API: `internal/api/lists.go`

| Metric | Count |
|--------|-------|
| Total params | 8 |
| Direct flags | 8 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `content`, `markdown_content`, `due_date`, `due_date_time`, `priority`, `assignee`, `status`

#### ✅ `GET` `/v2/list/{list_id}`
**Get List** | `GetList`
API: `internal/api/tasks.go`

No parameters (besides path params).

#### ✅ `PUT` `/v2/list/{list_id}`
**Update List** | `UpdateList`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 9 |
| Direct flags | 9 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `content`, `markdown_content`, `due_date`, `due_date_time`, `priority`, `assignee`, `status`, `unset_status`

#### ✅ `DELETE` `/v2/list/{list_id}`
**Delete List** | `DeleteList`
API: `internal/api/tasks.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/list/{list_id}/task/{task_id}`
**Add Task To List** | `AddTaskToList`
API: `internal/api/tasks.go`

No parameters (besides path params).

#### ✅ `DELETE` `/v2/list/{list_id}/task/{task_id}`
**Remove Task From List** | `RemoveTaskFromList`
API: `internal/api/tasks.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/folder/{folder_id}/list_template/{template_id}`
**Create List From Template in Folder** | `CreateFolderListFromTemplate`
API: `internal/api/templates.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `options`

#### ✅ `POST` `/v2/space/{space_id}/list_template/{template_id}`
**Create List From Template in Space.** | `CreateSpaceListFromTemplate`
API: `internal/api/templates.go`

| Metric | Count |
|--------|-------|
| Total params | 32 |
| Direct flags | 12 |
| Via JSON flag | 20 |
| Missing | 0 |

**Body params**: `name`, `options`
**Nested body params**: `options.return_immediately`, `options.content`, `options.time_estimate`, `options.automation`, `options.include_views`, `options.old_due_date`, `options.old_start_date`, `options.old_followers`, `options.comment_attachments`, `options.recur_settings`, `options.old_tags`, `options.old_statuses`, `options.subtasks`, `options.custom_type`, `options.old_assignees`, `options.attachments`, `options.comment`, `options.old_status`, `options.external_dependencies`, `options.internal_dependencies`, `options.priority`, `options.custom_fields`, `options.old_checklists`, `options.relationships`, `options.old_subtask_assignees`, `options.start_date`, `options.due_date`, `options.remap_start_date`, `options.skip_weekends`, `options.archived`

### Members

#### ✅ `GET` `/v2/task/{task_id}/member`
**Get Task Members** | `GetTaskMembers`
API: `internal/api/members.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/list/{list_id}/member`
**Get List Members** | `GetListMembers`
API: `internal/api/members.go`

No parameters (besides path params).

### Roles

#### ✅ `GET` `/v2/team/{team_id}/customroles`
**Get Custom Roles** | `GetCustomRoles`
API: `internal/api/roles.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_members`

### Shared Hierarchy

#### ✅ `GET` `/v2/team/{team_id}/shared`
**Shared Hierarchy** | `SharedHierarchy`
API: `internal/api/shared.go`

No parameters (besides path params).

### Spaces

#### ✅ `GET` `/v2/team/{team_id}/space`
**Get Spaces** | `GetSpaces`
API: `internal/api/spaces.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `archived`

#### ✅ `POST` `/v2/team/{team_id}/space`
**Create Space** | `CreateSpace`
API: `internal/api/spaces.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `multiple_assignees`, `features`

#### ✅ `GET` `/v2/space/{space_id}`
**Get Space** | `GetSpace`
API: `internal/api/lists.go`

No parameters (besides path params).

#### ✅ `PUT` `/v2/space/{space_id}`
**Update Space** | `UpdateSpace`
API: `internal/api/lists.go`

| Metric | Count |
|--------|-------|
| Total params | 27 |
| Direct flags | 9 |
| Via JSON flag | 18 |
| Missing | 0 |

**Body params**: `name`, `color`, `private`, `admin_can_manage`, `multiple_assignees`, `features`
**Nested body params**: `features.due_dates`, `features.due_dates.enabled`, `features.due_dates.start_date`, `features.due_dates.remap_due_dates`, `features.due_dates.remap_closed_due_date`, `features.time_tracking`, `features.time_tracking.enabled`, `features.tags`, `features.tags.enabled`, `features.time_estimates`, `features.time_estimates.enabled`, `features.checklists`, `features.checklists.enabled`, `features.custom_fields`, `features.custom_fields.enabled`, `features.remap_dependencies`, `features.remap_dependencies.enabled`, `features.dependency_warning`, `features.dependency_warning.enabled`, `features.portfolios`, `features.portfolios.enabled`

#### ✅ `DELETE` `/v2/space/{space_id}`
**Delete Space** | `DeleteSpace`
API: `internal/api/lists.go`

No parameters (besides path params).

### Tags

#### ✅ `GET` `/v2/space/{space_id}/tag`
**Get Space Tags** | `GetSpaceTags`
API: `internal/api/tags.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/space/{space_id}/tag`
**Create Space Tag** | `CreateSpaceTag`
API: `internal/api/tags.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 4 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `tag`
**Nested body params**: `tag.name`, `tag.tag_fg`, `tag.tag_bg`

#### ✅ `PUT` `/v2/space/{space_id}/tag/{tag_name}`
**Edit Space Tag** | `EditSpaceTag`
API: `internal/api/tags.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 2 |
| Via JSON flag | 2 |
| Missing | 0 |

**Body params**: `tag`
**Nested body params**: `tag.name`, `tag.fg_color`, `tag.bg_color`

#### ✅ `DELETE` `/v2/space/{space_id}/tag/{tag_name}`
**Delete Space Tag** | `DeleteSpaceTag`
API: `internal/api/tags.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `tag`

#### ✅ `POST` `/v2/task/{task_id}/tag/{tag_name}`
**Add Tag To Task** | `AddTagToTask`
API: `internal/api/tags.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

#### ✅ `DELETE` `/v2/task/{task_id}/tag/{tag_name}`
**Remove Tag From Task** | `RemoveTagFromTask`
API: `internal/api/tags.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

### Task Checklists

#### ✅ `POST` `/v2/task/{task_id}/checklist`
**Create Checklist** | `CreateChecklist`
API: `internal/api/checklists.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `name`

#### ✅ `PUT` `/v2/checklist/{checklist_id}`
**Edit Checklist** | `EditChecklist`
API: `internal/api/checklists.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `position`

#### ✅ `DELETE` `/v2/checklist/{checklist_id}`
**Delete Checklist** | `DeleteChecklist`
API: `internal/api/checklists.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/checklist/{checklist_id}/checklist_item`
**Create Checklist Item** | `CreateChecklistItem`
API: `internal/api/checklists.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `assignee`

#### ✅ `PUT` `/v2/checklist/{checklist_id}/checklist_item/{checklist_item_id}`
**Edit Checklist Item** | `EditChecklistItem`
API: `internal/api/checklists.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 4 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `assignee`, `resolved`, `parent`

#### ✅ `DELETE` `/v2/checklist/{checklist_id}/checklist_item/{checklist_item_id}`
**Delete Checklist Item** | `DeleteChecklistItem`
API: `internal/api/checklists.go`

No parameters (besides path params).

### Task Relationships

#### ✅ `POST` `/v2/task/{task_id}/dependency`
**Add Dependency** | `AddDependency`
API: `internal/api/relationships.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 1 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `depends_on`, `depedency_of`
**⚠️ Missing**: `depedency_of`

#### ✅ `DELETE` `/v2/task/{task_id}/dependency`
**Delete Dependency** | `DeleteDependency`
API: `internal/api/relationships.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 4 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `depends_on`, `dependency_of`, `custom_task_ids`, `team_id`

#### ✅ `POST` `/v2/task/{task_id}/link/{links_to}`
**Add Task Link** | `AddTaskLink`
API: `internal/api/relationships.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

#### ✅ `DELETE` `/v2/task/{task_id}/link/{links_to}`
**Delete Task Link** | `DeleteTaskLink`
API: `internal/api/relationships.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

### Tasks

#### ✅ `GET` `/v2/list/{list_id}/task`
**Get Tasks** | `GetTasks`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 23 |
| Direct flags | 22 |
| Via JSON flag | 0 |
| Missing | 1 |

**Query params**: `archived`, `include_markdown_description`, `page`, `order_by`, `reverse`, `subtasks`, `statuses`, `include_closed`, `include_timl`, `assignees`, `watchers`, `tags`, `due_date_gt`, `due_date_lt`, `date_created_gt`, `date_created_lt`, `date_updated_gt`, `date_updated_lt`, `date_done_gt`, `date_done_lt`, `custom_fields`, `custom_field`, `custom_items`
**⚠️ Missing**: `statuses`

#### ✅ `POST` `/v2/list/{list_id}/task`
**Create Task** | `CreateTask`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 21 |
| Direct flags | 21 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `description`, `assignees`, `archived`, `group_assignees`, `tags`, `status`, `priority`, `due_date`, `due_date_time`, `time_estimate`, `start_date`, `start_date_time`, `points`, `notify_all`, `parent`, `markdown_content`, `links_to`, `check_required_custom_fields`, `custom_fields`, `custom_item_id`

#### ✅ `GET` `/v2/task/{task_id}`
**Get Task** | `GetTask`
API: `internal/api/attachments.go`

| Metric | Count |
|--------|-------|
| Total params | 5 |
| Direct flags | 5 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`, `include_subtasks`, `include_markdown_description`, `custom_fields`

#### ✅ `PUT` `/v2/task/{task_id}`
**Update Task** | `UpdateTask`
API: `internal/api/attachments.go`

| Metric | Count |
|--------|-------|
| Total params | 25 |
| Direct flags | 25 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `custom_item_id`, `name`, `description`, `markdown_content`, `status`, `priority`, `due_date`, `due_date_time`, `parent`, `time_estimate`, `start_date`, `start_date_time`, `points`, `assignees`, `group_assignees`, `watchers`, `archived`
**Nested body params**: `assignees.add`, `assignees.rem`, `group_assignees.add`, `group_assignees.rem`, `watchers.add`, `watchers.rem`

#### ✅ `DELETE` `/v2/task/{task_id}`
**Delete Task** | `DeleteTask`
API: `internal/api/attachments.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

#### ✅ `GET` `/v2/team/{team_Id}/task`
**Get Filtered Team Tasks** | `GetFilteredTeamTasks`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 23 |
| Direct flags | 22 |
| Via JSON flag | 0 |
| Missing | 1 |

**Query params**: `page`, `order_by`, `reverse`, `subtasks`, `space_ids[]`, `project_ids[]`, `list_ids[]`, `statuses[]`, `include_closed`, `assignees[]`, `tags[]`, `due_date_gt`, `due_date_lt`, `date_created_gt`, `date_created_lt`, `date_updated_gt`, `date_updated_lt`, `date_done_gt`, `date_done_lt`, `custom_fields`, `parent`, `include_markdown_description`, `custom_items[]`
**⚠️ Missing**: `statuses[]`

#### ✅ `POST` `/v2/task/{task_id}/merge`
**Merge Tasks** | `mergeTasks`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 0 |
| Via JSON flag | 0 |
| Missing | 1 |

**Body params**: `source_task_ids`
**⚠️ Missing**: `source_task_ids`

#### ✅ `GET` `/v2/task/{task_id}/time_in_status`
**Get Task's Time in Status** | `GetTask'sTimeinStatus`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

#### ✅ `GET` `/v2/task/bulk_time_in_status/task_ids`
**Get Bulk Tasks' Time in Status** | `GetBulkTasks'TimeinStatus`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `task_ids`, `custom_task_ids`, `team_id`

#### ✅ `POST` `/v2/list/{list_id}/taskTemplate/{template_id}`
**Create Task From Template** | `CreateTaskFromTemplate`
API: `internal/api/templates.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`

### Templates

#### ✅ `GET` `/v2/team/{team_id}/taskTemplate`
**Get Task Templates** | `GetTaskTemplates`
API: `internal/api/templates.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `page`

### Time Tracking

#### ✅ `GET` `/v2/team/{team_Id}/time_entries`
**Get time entries within a date range** | `Gettimeentrieswithinadaterange`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 14 |
| Direct flags | 14 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `start_date`, `end_date`, `assignee`, `include_task_tags`, `include_location_names`, `include_approval_history`, `include_approval_details`, `space_id`, `folder_id`, `list_id`, `task_id`, `custom_task_ids`, `team_id`, `is_billable`

#### ✅ `POST` `/v2/team/{team_Id}/time_entries`
**Create a time entry** | `Createatimeentry`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 14 |
| Direct flags | 14 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `description`, `tags`, `start`, `stop`, `end`, `billable`, `duration`, `assignee`, `tid`
**Nested body params**: `tags[].name`, `tags[].tag_fg`, `tags[].tag_bg`

#### ✅ `GET` `/v2/team/{team_id}/time_entries/{timer_id}`
**Get singular time entry** | `Getsingulartimeentry`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 4 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_task_tags`, `include_location_names`, `include_approval_history`, `include_approval_details`

#### ✅ `DELETE` `/v2/team/{team_id}/time_entries/{timer_id}`
**Delete a time Entry** | `DeleteatimeEntry`
API: `internal/api/time_entries.go`

No parameters (besides path params).

#### ✅ `PUT` `/v2/team/{team_id}/time_entries/{timer_id}`
**Update a time Entry** | `UpdateatimeEntry`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 10 |
| Direct flags | 10 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `description`, `tags`, `tag_action`, `start`, `end`, `tid`, `billable`, `duration`

#### ✅ `GET` `/v2/team/{team_id}/time_entries/{timer_id}/history`
**Get time entry history** | `Gettimeentryhistory`
API: `internal/api/time_entries.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/team/{team_id}/time_entries/current`
**Get running time entry** | `Getrunningtimeentry`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `assignee`

#### ✅ `DELETE` `/v2/team/{team_id}/time_entries/tags`
**Remove tags from time entries** | `Removetagsfromtimeentries`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `time_entry_ids`, `tags`

#### ✅ `GET` `/v2/team/{team_id}/time_entries/tags`
**Get all tags from time entries** | `Getalltagsfromtimeentries`
API: `internal/api/time_entries.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/team/{team_id}/time_entries/tags`
**Add tags from time entries** | `Addtagsfromtimeentries`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `time_entry_ids`, `tags`

#### ✅ `PUT` `/v2/team/{team_id}/time_entries/tags`
**Change tag names from time entries** | `Changetagnamesfromtimeentries`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 4 |
| Direct flags | 4 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `new_name`, `tag_bg`, `tag_fg`

#### ✅ `POST` `/v2/team/{team_Id}/time_entries/start`
**Start a time Entry** | `StartatimeEntry`
API: `internal/api/time_entries.go`

| Metric | Count |
|--------|-------|
| Total params | 7 |
| Direct flags | 7 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `description`, `tags`, `tid`, `billable`
**Nested body params**: `tags[].name`

#### ✅ `POST` `/v2/team/{team_id}/time_entries/stop`
**Stop a time Entry** | `StopatimeEntry`
API: `internal/api/time_entries.go`

No parameters (besides path params).

### Time Tracking (Legacy)

#### ✅ `GET` `/v2/task/{task_id}/time`
**Get tracked time** | `Gettrackedtime`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

#### ✅ `POST` `/v2/task/{task_id}/time`
**Track time** | `Tracktime`
API: `internal/api/tasks.go`

| Metric | Count |
|--------|-------|
| Total params | 5 |
| Direct flags | 5 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `start`, `end`, `time`

#### ✅ `PUT` `/v2/task/{task_id}/time/{interval_id}`
**Edit time tracked** | `Edittimetracked`
API: `internal/api/time_tracking_legacy.go`

| Metric | Count |
|--------|-------|
| Total params | 5 |
| Direct flags | 5 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`
**Body params**: `start`, `end`, `time`

#### ✅ `DELETE` `/v2/task/{task_id}/time/{interval_id}`
**Delete time tracked** | `Deletetimetracked`
API: `internal/api/time_tracking_legacy.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `custom_task_ids`, `team_id`

### User Groups

#### ✅ `POST` `/v2/team/{team_id}/group`
**Create Group** | `CreateUserGroup`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `handle`, `members`

#### ✅ `PUT` `/v2/group/{group_id}`
**Update Group** | `UpdateTeam`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 5 |
| Direct flags | 5 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `handle`, `members`
**Nested body params**: `members.add`, `members.rem`

#### ✅ `DELETE` `/v2/group/{group_id}`
**Delete Group** | `DeleteTeam`
API: `internal/api/members.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/group`
**Get Groups** | `GetTeams1`
API: `internal/api/members.go`

| Metric | Count |
|--------|-------|
| Total params | 2 |
| Direct flags | 2 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `team_id`, `group_ids`

### Users

#### ✅ `POST` `/v2/team/{team_id}/user`
**Invite User To Workspace** | `InviteUserToWorkspace`
API: `internal/api/users.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `email`, `admin`, `custom_role_id`

#### ✅ `GET` `/v2/team/{team_id}/user/{user_id}`
**Get User** | `GetUser`
API: `internal/api/users.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `include_shared`

#### ✅ `PUT` `/v2/team/{team_id}/user/{user_id}`
**Edit User On Workspace** | `EditUserOnWorkspace`
API: `internal/api/users.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `username`, `admin`, `custom_role_id`

#### ✅ `DELETE` `/v2/team/{team_id}/user/{user_id}`
**Remove User From Workspace** | `RemoveUserFromWorkspace`
API: `internal/api/users.go`

No parameters (besides path params).

### Views

#### ✅ `GET` `/v2/team/{team_id}/view`
**Get Workspace (Everything level) Views** | `GetTeamViews`
API: `internal/api/views.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/team/{team_id}/view`
**Create Workspace (Everything level) View** | `CreateTeamView`
API: `internal/api/views.go`

| Metric | Count |
|--------|-------|
| Total params | 9 |
| Direct flags | 9 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `type`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`

#### ✅ `GET` `/v2/space/{space_id}/view`
**Get Space Views** | `GetSpaceViews`
API: `internal/api/views.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/space/{space_id}/view`
**Create Space View** | `CreateSpaceView`
API: `internal/api/views.go`

| Metric | Count |
|--------|-------|
| Total params | 9 |
| Direct flags | 9 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `type`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`

#### ✅ `GET` `/v2/folder/{folder_id}/view`
**Get Folder Views** | `GetFolderViews`
API: `internal/api/views.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/folder/{folder_id}/view`
**Create Folder View** | `CreateFolderView`
API: `internal/api/views.go`

| Metric | Count |
|--------|-------|
| Total params | 9 |
| Direct flags | 9 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `type`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`

#### ✅ `GET` `/v2/list/{list_id}/view`
**Get List Views** | `GetListViews`
API: `internal/api/views.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/list/{list_id}/view`
**Create List View** | `CreateListView`
API: `internal/api/views.go`

| Metric | Count |
|--------|-------|
| Total params | 9 |
| Direct flags | 9 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `name`, `type`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`

#### ✅ `GET` `/v2/view/{view_id}`
**Get View** | `GetView`
API: `internal/api/comments.go`

No parameters (besides path params).

#### ✅ `PUT` `/v2/view/{view_id}`
**Update View** | `UpdateView`
API: `internal/api/comments.go`

| Metric | Count |
|--------|-------|
| Total params | 38 |
| Direct flags | 19 |
| Via JSON flag | 19 |
| Missing | 0 |

**Body params**: `name`, `type`, `parent`, `grouping`, `divide`, `sorting`, `filters`, `columns`, `team_sidebar`, `settings`
**Nested body params**: `parent.id`, `parent.type`, `grouping.field`, `grouping.dir`, `grouping.collapsed`, `grouping.ignore`, `divide.field`, `divide.dir`, `divide.collapsed`, `sorting.fields`, `filters.op`, `filters.fields`, `filters.search`, `filters.show_closed`, `columns.fields`, `team_sidebar.assignees`, `team_sidebar.assigned_comments`, `team_sidebar.unassigned_tasks`, `settings.show_task_locations`, `settings.show_subtasks`, `settings.show_subtask_parent_names`, `settings.show_closed_subtasks`, `settings.show_assignees`, `settings.show_images`, `settings.collapse_empty_columns`, `settings.me_comments`, `settings.me_subtasks`, `settings.me_checklists`

#### ✅ `DELETE` `/v2/view/{view_id}`
**Delete View** | `DeleteView`
API: `internal/api/comments.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/view/{view_id}/task`
**Get View Tasks** | `GetViewTasks`
API: `internal/api/views.go`

| Metric | Count |
|--------|-------|
| Total params | 1 |
| Direct flags | 1 |
| Via JSON flag | 0 |
| Missing | 0 |

**Query params**: `page`

### Webhooks

#### ✅ `GET` `/v2/team/{team_id}/webhook`
**Get Webhooks** | `GetWebhooks`
API: `internal/api/webhooks.go`

No parameters (besides path params).

#### ✅ `POST` `/v2/team/{team_id}/webhook`
**Create Webhook** | `CreateWebhook`
API: `internal/api/webhooks.go`

| Metric | Count |
|--------|-------|
| Total params | 6 |
| Direct flags | 6 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `endpoint`, `events`, `space_id`, `folder_id`, `list_id`, `task_id`

#### ✅ `PUT` `/v2/webhook/{webhook_id}`
**Update Webhook** | `UpdateWebhook`
API: `internal/api/webhooks.go`

| Metric | Count |
|--------|-------|
| Total params | 3 |
| Direct flags | 3 |
| Via JSON flag | 0 |
| Missing | 0 |

**Body params**: `endpoint`, `events`, `status`

#### ✅ `DELETE` `/v2/webhook/{webhook_id}`
**Delete Webhook** | `DeleteWebhook`
API: `internal/api/webhooks.go`

No parameters (besides path params).

### Workspaces

#### ✅ `GET` `/v2/team`
**Get Authorized Workspaces** | `GetAuthorizedTeams`
API: `internal/api/custom_task_types.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/team/{team_id}/seats`
**Get Workspace seats** | `GetWorkspaceseats`
API: `internal/api/workspaces.go`

No parameters (besides path params).

#### ✅ `GET` `/v2/team/{team_id}/plan`
**Get Workspace Plan** | `GetWorkspaceplan`
API: `internal/api/workspaces.go`

No parameters (besides path params).

## All Missing Parameters

7 total missing parameters across 5 endpoints:

### `POST` `/v2/oauth/token` — Get Access Token
- `client_id`
- `client_secret`
- `code`

### `POST` `/v2/task/{task_id}/dependency` — Add Dependency
- `depedency_of`

### `GET` `/v2/list/{list_id}/task` — Get Tasks
- `statuses`

### `GET` `/v2/team/{team_Id}/task` — Get Filtered Team Tasks
- `statuses[]`

### `POST` `/v2/task/{task_id}/merge` — Merge Tasks
- `source_task_ids`
