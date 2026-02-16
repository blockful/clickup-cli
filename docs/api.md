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

Base URL: `https://api.clickup.com/api`

Note: ClickUp API v2 uses "team" in URLs, but the CLI uses "workspace" for clarity.
