package repository

import "testing"

func TestParseGitHubURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		owner   string
		repo    string
		wantErr bool
	}{
		{name: "repository", input: "https://github.com/lint-md/parser", owner: "lint-md", repo: "parser"},
		{name: "git suffix", input: "https://github.com/lint-md/parser.git", owner: "lint-md", repo: "parser"},
		{name: "trailing slash", input: "https://github.com/lint-md/parser/", owner: "lint-md", repo: "parser"},
		{name: "wrong host", input: "https://example.com/lint-md/parser", wantErr: true},
		{name: "nested path", input: "https://github.com/lint-md/parser/tree/main", wantErr: true},
		{name: "missing repo", input: "https://github.com/lint-md", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGitHubURL(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseGitHubURL(%q) expected error", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseGitHubURL(%q): %v", tt.input, err)
			}
			if got.Owner != tt.owner || got.Name != tt.repo {
				t.Fatalf("ParseGitHubURL(%q) = %s/%s, want %s/%s", tt.input, got.Owner, got.Name, tt.owner, tt.repo)
			}
		})
	}
}
