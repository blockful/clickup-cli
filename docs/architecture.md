# Architecture

## Overview

```mermaid
graph TD
    A[main.go] --> B[cmd/root.go]
    B --> C[cmd/auth.go]
    B --> D[cmd/workspace.go]
    B --> E[cmd/space.go]
    B --> F[cmd/folder.go]
    B --> G[cmd/list.go]
    B --> H[cmd/task.go]
    B --> I[cmd/comment.go]
    
    C --> J[internal/api/client.go]
    D --> J
    E --> J
    F --> J
    G --> J
    H --> J
    I --> J
    
    J --> K[ClickUp API v2]
    
    C --> L[internal/config/config.go]
    B --> L
    
    C --> M[internal/output/output.go]
    D --> M
    E --> M
    F --> M
    G --> M
    H --> M
    I --> M
```

## Layers

1. **cmd/** — Cobra command definitions, flag parsing, validation
2. **internal/api/** — HTTP client, API type definitions, request/response handling
3. **internal/config/** — Viper-based config file management
4. **internal/output/** — JSON output formatting, error formatting

## Design Principles

- JSON-first output for AI agent consumption
- Consistent error format across all commands
- Config file for persistent auth, flag overrides for one-off usage
- Thin command layer — business logic lives in api package
