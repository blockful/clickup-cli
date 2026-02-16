# Creating Relationships Between Tasks

This guide explains how to link two ClickUp tasks using relationship custom fields via the CLI. This is the process for connecting contacts to organizations, deals to contacts, or any two tasks linked by a `list_relationship` field.

## Overview

ClickUp relationships work through **custom fields of type `list_relationship`**. Each relationship field connects a task in one list to tasks in another list. For example, a Contact task might have an "Organization" relationship field that links to tasks in the Organizations list.

The CLI command to set a relationship is:

```bash
clickup custom-field set --task TASK_ID --field FIELD_ID --value '{"add":["TARGET_TASK_ID"]}'
```

The challenge is finding the right `TASK_ID`, `FIELD_ID`, and `TARGET_TASK_ID`. Here's the full process.

## Step-by-Step Process

### Step 1: Identify the source task

You need the task ID of the task you want to add a relationship TO. If you have the task name but not the ID, search for it:

```bash
# Search by name across the workspace
clickup task search --workspace 90132341641 --query "Abdullah Umar"
```

Or if you know which list it's in, list tasks and find it:

```bash
# List tasks in the Contacts list
clickup task list --list 901325492602
```

Parse the JSON output to find the task ID. Example output:

```json
{
  "tasks": [
    {
      "id": "86aff5q93",
      "name": "Abdullah Umar",
      "status": { "status": "to do" }
    }
  ]
}
```

**Result:** Source task ID is `86aff5q93`

### Step 2: Find the relationship field ID

Get the source task's details to see its custom fields:

```bash
clickup task get --id 86aff5q93
```

In the response, look for custom fields with `"type": "list_relationship"`. Each one represents a relationship you can set:

```json
{
  "custom_fields": [
    {
      "id": "4c1ffb77-adfa-41be-b64a-ff3e17dcda9d",
      "name": "Organization",
      "type": "list_relationship",
      "type_config": {
        "subcategory_id": "901325492599"
      }
    },
    {
      "id": "99269f28-c68a-4344-a241-5b10228645bf",
      "name": "Deals",
      "type": "list_relationship",
      "type_config": {
        "subcategory_id": "901325492813"
      }
    }
  ]
}
```

Key fields:
- `id` — the field ID you'll use in the set command
- `name` — human-readable name (e.g., "Organization")
- `type_config.subcategory_id` — the list ID where target tasks live

**Result:** Organization field ID is `4c1ffb77-adfa-41be-b64a-ff3e17dcda9d`, and targets live in list `901325492599` (Organizations list).

### Step 3: Find the target task

Now find the task you want to link to. Use the `subcategory_id` from the field's `type_config` to search the right list:

```bash
# List tasks in the Organizations list
clickup task list --list 901325492599
```

Or search by name:

```bash
clickup task search --workspace 90132341641 --query "Uniswap"
```

Find the target task ID from the results:

```json
{
  "tasks": [
    {
      "id": "86aff4n5k",
      "name": "Uniswap Foundation",
      "status": { "status": "to do" }
    }
  ]
}
```

**Result:** Target task ID is `86aff4n5k`

### Step 4: Create the relationship

```bash
clickup custom-field set \
  --task 86aff5q93 \
  --field 4c1ffb77-adfa-41be-b64a-ff3e17dcda9d \
  --value '{"add":["86aff4n5k"]}'
```

Response:

```json
{ "status": "ok" }
```

### Step 5: Verify (optional)

```bash
clickup task get --id 86aff5q93
```

The relationship field will now show the linked task:

```json
{
  "name": "Organization",
  "type": "list_relationship",
  "value": [
    {
      "id": "86aff4n5k",
      "name": "Uniswap Foundation",
      "url": "https://app.clickup.com/t/86aff4n5k"
    }
  ]
}
```

## Common Operations

### Add multiple relationships at once

```bash
clickup custom-field set --task TASK_ID --field FIELD_ID \
  --value '{"add":["task_id_1","task_id_2","task_id_3"]}'
```

### Remove a relationship

```bash
clickup custom-field set --task TASK_ID --field FIELD_ID \
  --value '{"rem":["task_id_to_remove"]}'
```

### Replace relationships (remove old, add new)

```bash
clickup custom-field set --task TASK_ID --field FIELD_ID \
  --value '{"add":["new_task_id"],"rem":["old_task_id"]}'
```

## Blockful CRM Relationship Map

In the blockful Sales space (workspace `90132341641`), the CRM lists and their relationship fields are:

| List | List ID | Relationship Fields |
|------|---------|-------------------|
| **Contacts** | `901325492602` | Organization, Deals, Tasks, Meetings |
| **Organizations** | `901325492599` | Contacts (inverse) |
| **Deals** | `901325492813` | Contacts, Organization, Tasks |
| **Tasks** | `901325510349` | Contacts, Deals |
| **Meetings** | `901325510351` | Contacts |

### Common CRM workflows

**Link a contact to an organization:**
```bash
# 1. Find contact: clickup task search --workspace 90132341641 --query "John"
# 2. Get field ID from contact's custom fields (Organization field)
# 3. Find org: clickup task list --list 901325492599
# 4. Set relationship:
clickup custom-field set --task CONTACT_ID --field ORG_FIELD_ID --value '{"add":["ORG_TASK_ID"]}'
```

**Link a deal to a contact and organization:**
```bash
# Get the deal's custom fields to find Contacts and Organization field IDs
clickup task get --id DEAL_ID
# Set both relationships
clickup custom-field set --task DEAL_ID --field CONTACTS_FIELD_ID --value '{"add":["CONTACT_ID"]}'
clickup custom-field set --task DEAL_ID --field ORG_FIELD_ID --value '{"add":["ORG_ID"]}'
```

## Tips for AI Agents

1. **Always get the task first** to discover available relationship fields and their IDs. Field IDs are UUIDs that don't change.

2. **Use `type_config.subcategory_id`** to know which list the target tasks come from — this tells you where to search.

3. **Cache field IDs** — for a given list, the relationship field IDs are stable. You can use `clickup custom-field list --list LIST_ID` to get all fields for a list without fetching a specific task.

4. **The value is always JSON** with `add` and/or `rem` arrays of task IDs (strings).

5. **Relationships are bidirectional in the UI** but set from one side via the API. Setting "Organization" on a Contact will automatically show that Contact under the Organization's inverse relationship.

6. **Task IDs are alphanumeric strings** like `86aff5q93` — they come from the task's `id` field, or from ClickUp URLs: `https://app.clickup.com/t/86aff5q93` → ID is `86aff5q93`.
