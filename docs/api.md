# CLI Command to API Endpoint Mapping

| CLI Command | HTTP Method | API Endpoint |
|-------------|-------------|--------------|
| `auth login` | GET | `/v2/user` |
| `auth whoami` | GET | `/v2/user` |
| `workspace list` | GET | `/v2/team` |
| `space list` | GET | `/v2/team/{team_id}/space` |
| `space get` | GET | `/v2/space/{space_id}` |
| `space create` | POST | `/v2/team/{team_id}/space` |
| `folder list` | GET | `/v2/space/{space_id}/folder` |
| `folder get` | GET | `/v2/folder/{folder_id}` |
| `folder create` | POST | `/v2/space/{space_id}/folder` |
| `list list` | GET | `/v2/folder/{folder_id}/list` |
| `list get` | GET | `/v2/list/{list_id}` |
| `list create` | POST | `/v2/folder/{folder_id}/list` |
| `task list` | GET | `/v2/list/{list_id}/task` |
| `task get` | GET | `/v2/task/{task_id}` |
| `task create` | POST | `/v2/list/{list_id}/task` |
| `task update` | PUT | `/v2/task/{task_id}` |
| `task delete` | DELETE | `/v2/task/{task_id}` |
| `comment list` | GET | `/v2/task/{task_id}/comment` |
| `comment create` | POST | `/v2/task/{task_id}/comment` |
| `doc list` | GET | `/v3/workspaces/{workspace_id}/docs` |
| `doc get` | GET | `/v3/workspaces/{workspace_id}/docs/{doc_id}` |
| `doc create` | POST | `/v3/workspaces/{workspace_id}/docs` |
| `doc page-list` | GET | `/v3/workspaces/{workspace_id}/docs/{doc_id}/page_listing` |
| `doc page-get` | GET | `/v3/workspaces/{workspace_id}/docs/{doc_id}/pages/{page_id}` |
| `doc page-create` | POST | `/v3/workspaces/{workspace_id}/docs/{doc_id}/pages` |
| `doc page-update` | PUT | `/v3/workspaces/{workspace_id}/docs/{doc_id}/pages/{page_id}` |
| `custom-field list` | GET | `/v2/list/{list_id}/field` (or folder/space/team) |
| `custom-field set` | POST | `/v2/task/{task_id}/field/{field_id}` |
| `custom-field remove` | DELETE | `/v2/task/{task_id}/field/{field_id}` |
| `tag list` | GET | `/v2/space/{space_id}/tag` |
| `tag create` | POST | `/v2/space/{space_id}/tag` |
| `tag update` | PUT | `/v2/space/{space_id}/tag/{tag_name}` |
| `tag delete` | DELETE | `/v2/space/{space_id}/tag/{tag_name}` |
| `tag add` | POST | `/v2/task/{task_id}/tag/{tag_name}` |
| `tag remove` | DELETE | `/v2/task/{task_id}/tag/{tag_name}` |
| `checklist create` | POST | `/v2/task/{task_id}/checklist` |
| `checklist update` | PUT | `/v2/checklist/{checklist_id}` |
| `checklist delete` | DELETE | `/v2/checklist/{checklist_id}` |
| `checklist-item create` | POST | `/v2/checklist/{checklist_id}/checklist_item` |
| `checklist-item update` | PUT | `/v2/checklist/{checklist_id}/checklist_item/{item_id}` |
| `checklist-item delete` | DELETE | `/v2/checklist/{checklist_id}/checklist_item/{item_id}` |
| `time-entry list` | GET | `/v2/team/{team_id}/time_entries` |
| `time-entry get` | GET | `/v2/team/{team_id}/time_entries/{timer_id}` |
| `time-entry create` | POST | `/v2/team/{team_id}/time_entries` |
| `time-entry update` | PUT | `/v2/team/{team_id}/time_entries/{timer_id}` |
| `time-entry delete` | DELETE | `/v2/team/{team_id}/time_entries/{timer_id}` |
| `time-entry start` | POST | `/v2/team/{team_id}/time_entries/start` |
| `time-entry stop` | POST | `/v2/team/{team_id}/time_entries/stop` |
| `time-entry current` | GET | `/v2/team/{team_id}/time_entries/current` |
| `webhook list` | GET | `/v2/team/{team_id}/webhook` |
| `webhook create` | POST | `/v2/team/{team_id}/webhook` |
| `webhook update` | PUT | `/v2/webhook/{webhook_id}` |
| `webhook delete` | DELETE | `/v2/webhook/{webhook_id}` |
| `view list` | GET | `/v2/team/{team_id}/view` (or space/folder/list) |
| `view get` | GET | `/v2/view/{view_id}` |
| `view create` | POST | `/v2/team/{team_id}/view` (or space/folder/list) |
| `view update` | PUT | `/v2/view/{view_id}` |
| `view delete` | DELETE | `/v2/view/{view_id}` |
| `view tasks` | GET | `/v2/view/{view_id}/task` |
| `goal list` | GET | `/v2/team/{team_id}/goal` |
| `goal get` | GET | `/v2/goal/{goal_id}` |
| `goal create` | POST | `/v2/team/{team_id}/goal` |
| `goal update` | PUT | `/v2/goal/{goal_id}` |
| `goal delete` | DELETE | `/v2/goal/{goal_id}` |
| `member list` | GET | `/v2/list/{list_id}/member` or `/v2/task/{task_id}/member` |
| `group list` | GET | `/v2/group` |
| `group create` | POST | `/v2/team/{team_id}/group` |
| `group delete` | DELETE | `/v2/group/{group_id}` |
| `guest invite` | POST | `/v2/team/{team_id}/guest` |
| `guest get` | GET | `/v2/team/{team_id}/guest/{guest_id}` |
| `guest remove` | DELETE | `/v2/team/{team_id}/guest/{guest_id}` |

Base URL: `https://api.clickup.com/api`

Note: ClickUp API v2 uses "team" in URLs, but the CLI uses "workspace" for clarity.
Docs API uses v3 endpoints under `/v3/workspaces/`.
