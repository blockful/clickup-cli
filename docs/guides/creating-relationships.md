# Creating Relationships Between Tasks

This guide explains how to link two ClickUp tasks using relationship custom fields via the CLI.

## Overview

ClickUp relationships work through **custom fields of type `list_relationship`**. Each relationship field connects a task in one list to tasks in another list. For example, a Contact task might have an "Organization" relationship field that links to tasks in the Organizations list.

The CLI command to set a relationship is:

```bash
clickup custom-field set --task TASK_ID --field FIELD_ID --value '{"add":["TARGET_TASK_ID"]}'
```

The challenge is finding the right `TASK_ID`, `FIELD_ID`, and `TARGET_TASK_ID`. Here's the full process.

## Step-by-Step Process

### Step 1: Identify the source task

Find the task you want to add a relationship to. If you have the name but not the ID, search for it:

```bash
# Search by name across the workspace
clickup task search --workspace YOUR_WORKSPACE_ID --query "Jane Smith"

# Or list tasks in a specific list
clickup task list --list YOUR_LIST_ID
```

Example output:

```json
{
  "tasks": [
    {
      "id": "abc123xyz",
      "name": "Jane Smith",
      "status": { "status": "active" }
    }
  ]
}
```

### Step 2: Find the relationship field ID

Get the source task's details to see its custom fields:

```bash
clickup task get --id abc123xyz
```

In the response, look for custom fields with `"type": "list_relationship"`:

```json
{
  "custom_fields": [
    {
      "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "name": "Organization",
      "type": "list_relationship",
      "type_config": {
        "subcategory_id": "111222333"
      }
    },
    {
      "id": "f9e8d7c6-b5a4-3210-fedc-ba9876543210",
      "name": "Deals",
      "type": "list_relationship",
      "type_config": {
        "subcategory_id": "444555666"
      }
    }
  ]
}
```

Key fields:
- `id` — the field ID you'll use in the set command
- `name` — human-readable name (e.g., "Organization")
- `type_config.subcategory_id` — the list ID where target tasks live

### Step 3: Find the target task

Use the `subcategory_id` from the field's `type_config` to search the correct list:

```bash
# List tasks in the target list
clickup task list --list 111222333

# Or search by name
clickup task search --workspace YOUR_WORKSPACE_ID --query "Acme Corp"
```

Find the target task ID from the results:

```json
{
  "tasks": [
    {
      "id": "def456uvw",
      "name": "Acme Corp",
      "status": { "status": "active" }
    }
  ]
}
```

### Step 4: Create the relationship

```bash
clickup custom-field set \
  --task abc123xyz \
  --field a1b2c3d4-e5f6-7890-abcd-ef1234567890 \
  --value '{"add":["def456uvw"]}'
```

Response:

```json
{ "status": "ok" }
```

### Step 5: Verify (optional)

```bash
clickup task get --id abc123xyz
```

The relationship field will now show the linked task:

```json
{
  "name": "Organization",
  "type": "list_relationship",
  "value": [
    {
      "id": "def456uvw",
      "name": "Acme Corp",
      "url": "https://app.clickup.com/t/def456uvw"
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

## Discovering Your Relationship Fields

To find all relationship fields available on a list without fetching a specific task:

```bash
clickup custom-field list --list YOUR_LIST_ID
```

This returns all custom fields for that list, including their IDs and types. Filter for `"type": "list_relationship"` to find linkable fields.

## Tips for AI Agents

1. **Always get the task first** to discover available relationship fields and their IDs. Field IDs are UUIDs that don't change.

2. **Use `type_config.subcategory_id`** to know which list the target tasks come from — this tells you where to search.

3. **Cache field IDs** — for a given list, the relationship field IDs are stable. Use `clickup custom-field list --list LIST_ID` to get all fields without fetching a specific task.

4. **The value is always JSON** with `add` and/or `rem` arrays of task IDs (strings).

5. **Relationships are bidirectional in the UI** but set from one side via the API. Setting "Organization" on a Contact will automatically show that Contact under the Organization's inverse relationship.

6. **Task IDs are alphanumeric strings** like `abc123xyz` — extract them from ClickUp URLs: `https://app.clickup.com/t/abc123xyz` → ID is `abc123xyz`.
