# ADR 001: Go as Implementation Language

## Status
Accepted

## Context
We need a CLI tool that is fast, produces a single binary, and is easy to distribute.

## Decision
Use Go for the CLI implementation.

## Consequences
- Single binary distribution, no runtime dependencies
- Fast startup time (critical for AI agent workflows)
- Strong standard library for HTTP and JSON
- cobra/viper ecosystem for CLI and config
