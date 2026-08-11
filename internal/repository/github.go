package repository

import (
	"fmt"
	"net/url"
	"strings"
)

// Repository identifies a GitHub repository and its local checkout state.
type Repository struct {
	URL       string
	Owner     string
	Name      string
	CloneURL  string
	Branch    string
	CommitSHA string
	WorkDir   string
}

// ParseGitHubURL parses a canonical GitHub HTTPS repository URL.
func ParseGitHubURL(rawURL string) (*Repository, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("parse GitHub URL: %w", err)
	}

	if u.Scheme != "https" || !strings.EqualFold(u.Hostname(), "github.com") {
		return nil, fmt.Errorf("unsupported GitHub URL %q: expected https://github.com/<owner>/<repo>", rawURL)
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf("invalid GitHub repository URL %q", rawURL)
	}

	name := strings.TrimSuffix(parts[1], ".git")
	if name == "" {
		return nil, fmt.Errorf("invalid GitHub repository URL %q", rawURL)
	}

	canonicalURL := fmt.Sprintf("https://github.com/%s/%s", parts[0], name)
	return &Repository{
		URL:      canonicalURL,
		Owner:    parts[0],
		Name:     name,
		CloneURL: canonicalURL + ".git",
	}, nil
}
