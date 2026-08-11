# cross-platform-ci

Build Go projects natively across Linux, Windows, and macOS on both x64 and ARM64 with one reusable GitHub Actions workflow.

## Goal

A caller should configure cross-platform builds once and let this repository expand the build into native GitHub-hosted runners:

| Target | Runner |
| --- | --- |
| Linux x64 | `ubuntu-latest` |
| Linux ARM64 | `ubuntu-24.04-arm` |
| Windows x64 | `windows-latest` |
| Windows ARM64 | `windows-11-arm` |
| macOS x64 | `macos-15-intel` |
| macOS ARM64 | `macos-latest` |

The workflow does not set `GOOS` or `GOARCH`. Each binary is built on a native runner for its target platform.

> Windows ARM64 is currently a GitHub-hosted runner public preview.

## Usage

For a Go command whose `main` package is at the repository root:

```yaml
name: Build

on:
  push:
  pull_request:

jobs:
  build:
    uses: luojiyin1987/cross-platform-ci/.github/workflows/build.yml@main
```

For a command under `cmd/myapp`:

```yaml
jobs:
  build:
    uses: luojiyin1987/cross-platform-ci/.github/workflows/build.yml@main
    with:
      package: ./cmd/myapp
      binary-name: myapp
```

During development, examples use `@main`. Stable consumers should use a version tag such as `@v1` after the first release is cut.

## What the workflow does

For every target it:

1. checks out the caller repository;
2. installs Go from the caller's `go.mod` by default;
3. prints the runner-native `GOOS` and `GOARCH`;
4. runs `go test ./...`;
5. builds the requested main package;
6. uploads the binary as a target-specific artifact.

Artifacts are named:

```text
linux-x64
linux-arm64
windows-x64
windows-arm64
macos-x64
macos-arm64
```

## Inputs

| Input | Default | Description |
| --- | --- | --- |
| `go-version-file` | `go.mod` | File used by `actions/setup-go` to select Go |
| `package` | `.` | Main package passed to `go build` |
| `binary-name` | caller repository name | Output executable name |
| `run-tests` | `true` | Run `go test ./...` before build |

## Repository self-test

This repository uses the reusable workflow to build its own `./cmd/cross-platform-ci` command on all six targets. That gives the workflow an end-to-end native build test before it is published as a stable version.

## Experimental repository inspector

The existing CLI can inspect a GitHub repository by shallow-cloning it into a temporary workspace:

```bash
go run ./cmd/cross-platform-ci inspect https://github.com/lint-md/parser
```

This is experimental and is not the main public interface of the project.

## Roadmap

- [x] Six-target native Go build matrix
- [x] Artifact upload per target
- [x] Self-test the reusable workflow
- [ ] Stable `v1` release/tag
- [ ] Rust support
- [ ] Node native-addon support
- [ ] CMake support
- [ ] Project detection and build-plan generation
