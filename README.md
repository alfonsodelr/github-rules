# github-rules

A minimal Go HTTP server used to practice setting up GitHub branch protection rules and CI/CD workflows.

## Project Overview

This is a simple Go HTTP server built with the standard library and [swaggo](https://github.com/swaggo/swag) for API documentation. It serves as a sandbox for configuring GitHub repository rules, automated workflows, and pull request enforcement.

## Getting Started

### Prerequisites

- Go 1.22+
- [golangci-lint](https://golangci-lint.run/usage/install/) (for linting)
- [swag CLI](https://github.com/swaggo/swag#install) (for regenerating Swagger docs)

### Run

```bash
go run server.go
```

Server starts at `http://localhost:8080`. Swagger UI is available at `http://localhost:8080/swagger/index.html`.

### Build

```bash
make build
# or
go build -o bin/server .
```

### Test

```bash
go test ./...
```

### Lint

```bash
make lint
```

## CI/CD

A GitHub Actions workflow (`.github/workflows/ci.yml`) runs on every push and pull request to `main`. It runs three jobs in parallel:

| Job | What it does |
|-----|-------------|
| **Build** | Compiles the server with `go build` |
| **Test** | Runs all tests with `go test ./...` |
| **Lint** | Runs `golangci-lint` |

## Branch Protection Rules

The `main` branch is protected with the following rules:

- **Restrict updates** — only authorized users can push
- **Restrict deletions** — the branch cannot be deleted
- **Require a pull request before merging** — direct pushes to `main` are blocked
- **Required approvals: 1** — at least one reviewer must approve before merging
- **Dismiss stale PR approvals** — approval is dismissed when new commits are pushed
- **Require status checks to pass** — all three CI jobs (Build, Test, Lint) must pass
- **Require branches to be up to date** — the PR branch must be current with `main` before merging
- **Block force pushes** — `git push --force` is not allowed on `main`
