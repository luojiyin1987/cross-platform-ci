# cross-platform-ci

A small experiment for turning a GitHub repository into a reproducible workspace that can later be built and tested across Linux, Windows, and macOS.

## Current milestone: repository ingestion

The first milestone deliberately stays small:

1. parse a GitHub repository URL;
2. create a temporary workspace;
3. shallow-clone the repository with Git;
4. resolve the checked-out branch and commit SHA;
5. list the repository root.

## Requirements

- Go 1.23 or newer
- Git available on `PATH`

## Run

```bash
go run ./cmd/cross-platform-ci inspect https://github.com/lint-md/parser
```

Example output:

```text
Repository
  owner:  lint-md
  name:   parser
  branch: main
  commit: <commit-sha>

Workspace
  path: <temporary-path>

Files
  package.json
  src/
  ...
```

The workspace is temporary and is removed when `inspect` exits.

## Roadmap

- [x] GitHub URL parsing
- [x] shallow repository checkout
- [x] branch and commit resolution
- [x] root file inspection
- [ ] project detection (`package.json`, `go.mod`, `Cargo.toml`, ...)
- [ ] build-plan generation
- [ ] Linux / Windows / macOS runners
- [ ] artifact collection
