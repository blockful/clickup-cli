# Contributing to clickup-cli

Thank you for your interest in contributing to clickup-cli!

## Filing Issues

- Use GitHub Issues to report bugs or request features
- Search existing issues before creating a new one
- Use the provided issue templates

## Development Setup

1. Install Go 1.22+ from https://go.dev/dl/
2. Clone the repository:
   ```bash
   git clone https://github.com/blockful/clickup-cli.git
   cd clickup-cli
   ```
3. Build:
   ```bash
   go build ./...
   ```
4. Run tests:
   ```bash
   go test -race ./...
   ```

## Commit Message Conventions

We use [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` — new feature
- `fix:` — bug fix
- `docs:` — documentation only
- `test:` — adding or updating tests
- `ci:` — CI/CD changes
- `refactor:` — code change that neither fixes a bug nor adds a feature
- `chore:` — maintenance tasks

## Pull Request Process

1. Fork the repository and create a feature branch from `main`
2. Make your changes with clear, focused commits
3. Ensure all tests pass: `go test -race ./...`
4. Ensure code passes linting: `golangci-lint run`
5. Update documentation if needed
6. Open a PR against `main` using the PR template

## Code Review

- All PRs require at least one approving review
- Address all review comments before merging
- Keep PRs focused and reasonably sized

## Testing Requirements

- Write table-driven tests for new functionality
- All existing tests must pass
- Aim for meaningful coverage of business logic
- Use the `api.ClientInterface` for mocking API calls in tests
