# ADR 003: Command Structure

## Status
Accepted

## Context
Need a consistent command hierarchy that maps to ClickUp's resource model.

## Decision
Use `<resource> <action>` pattern: `clickup task list`, `clickup task create`, etc.

## Consequences
- Predictable command discovery
- Maps naturally to CRUD operations
- Easy for AI agents to construct commands programmatically
- "workspace" used instead of ClickUp's internal "team" terminology
