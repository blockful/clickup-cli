# ADR 002: JSON-First Output

## Status
Accepted

## Context
This CLI is designed primarily for AI agents that need to parse command output programmatically.

## Decision
All commands output valid JSON by default. Errors use a consistent `{"error":"...","code":"..."}` format.

## Consequences
- Every output can be reliably parsed by AI agents
- Errors are structured and machine-readable
- Human readability is secondary (use `--format text` when needed)
