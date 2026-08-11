package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/luojiyin1987/cross-platform-ci/internal/repository"
)

const usage = `Usage:
  cross-platform-ci inspect <github-url>`

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 2 || args[0] != "inspect" {
		return fmt.Errorf("%s", usage)
	}

	repo, err := repository.ParseGitHubURL(args[1])
	if err != nil {
		return err
	}

	tempDir, err := os.MkdirTemp("", "cross-platform-ci-")
	if err != nil {
		return fmt.Errorf("create workspace: %w", err)
	}
	defer os.RemoveAll(tempDir)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	if err := repository.Clone(ctx, repo, tempDir); err != nil {
		return err
	}

	entries, err := repository.RootEntries(repo)
	if err != nil {
		return err
	}

	fmt.Println("Repository")
	fmt.Printf("  owner:  %s\n", repo.Owner)
	fmt.Printf("  name:   %s\n", repo.Name)
	fmt.Printf("  branch: %s\n", repo.Branch)
	fmt.Printf("  commit: %s\n", repo.CommitSHA)
	fmt.Println()
	fmt.Println("Workspace")
	fmt.Printf("  path: %s (temporary)\n", repo.WorkDir)
	fmt.Println()
	fmt.Println("Files")
	for _, entry := range entries {
		fmt.Printf("  %s\n", entry)
	}

	return nil
}
