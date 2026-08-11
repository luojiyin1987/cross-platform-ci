package repository

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// Clone creates a shallow checkout of repo below parentDir and populates its
// local checkout metadata.
func Clone(ctx context.Context, repo *Repository, parentDir string) error {
	workDir := filepath.Join(parentDir, repo.Owner+"-"+repo.Name)
	cmd := exec.CommandContext(ctx, "git", "clone", "--depth=1", "--", repo.CloneURL, workDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("clone repository: %w: %s", err, strings.TrimSpace(string(output)))
	}

	repo.WorkDir = workDir

	branch, err := gitOutput(ctx, workDir, "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return err
	}
	repo.Branch = branch

	commit, err := gitOutput(ctx, workDir, "rev-parse", "HEAD")
	if err != nil {
		return err
	}
	repo.CommitSHA = commit

	return nil
}

// RootEntries returns the names of files and directories at the repository root.
func RootEntries(repo *Repository) ([]string, error) {
	entries, err := os.ReadDir(repo.WorkDir)
	if err != nil {
		return nil, fmt.Errorf("read repository root: %w", err)
	}

	result := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.Name() == ".git" {
			continue
		}
		name := entry.Name()
		if entry.IsDir() {
			name += string(os.PathSeparator)
		}
		result = append(result, name)
	}
	sort.Strings(result)
	return result, nil
}

func gitOutput(ctx context.Context, workDir string, args ...string) (string, error) {
	cmdArgs := append([]string{"-C", workDir}, args...)
	cmd := exec.CommandContext(ctx, "git", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(string(output)))
	}
	return strings.TrimSpace(string(output)), nil
}
