package ui

import (
	"bytes"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestParseResponse(t *testing.T) {
	response := "COMMAND: ls -la\nEXPLANATION: List all files in long format"
	result := ParseResponse(response)

	if result.Command != "ls -la" {
		t.Errorf("command: got %q, want %q", result.Command, "ls -la")
	}
	if result.Explanation != "List all files in long format" {
		t.Errorf("explanation: got %q, want %q", result.Explanation, "List all files in long format")
	}
}

func TestParseResponseCommandOnly(t *testing.T) {
	response := "COMMAND: git status"
	result := ParseResponse(response)

	if result.Command != "git status" {
		t.Errorf("command: got %q, want %q", result.Command, "git status")
	}
	if result.Explanation != "" {
		t.Errorf("explanation: got %q, want empty", result.Explanation)
	}
}

func TestParseResponseEmpty(t *testing.T) {
	result := ParseResponse("")

	if result.Command != "" {
		t.Errorf("command: got %q, want empty", result.Command)
	}
	if result.Explanation != "" {
		t.Errorf("explanation: got %q, want empty", result.Explanation)
	}
}

func TestParseResponseExtraWhitespace(t *testing.T) {
	response := "  COMMAND:   find . -name '*.go'   \n  EXPLANATION:   Find all Go files   "
	result := ParseResponse(response)

	if result.Command != "find . -name '*.go'" {
		t.Errorf("command: got %q, want %q", result.Command, "find . -name '*.go'")
	}
	if result.Explanation != "Find all Go files" {
		t.Errorf("explanation: got %q, want %q", result.Explanation, "Find all Go files")
	}
}

func TestDisplayQuiet(t *testing.T) {
	result := Result{
		Command:     "echo hello",
		Explanation: "Print hello",
	}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	DisplayQuiet(result)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "echo hello") {
		t.Errorf("expected 'echo hello' in output, got: %q", output)
	}
	// Quiet mode should not include the explanation
	if strings.Contains(output, "Print hello") {
		t.Error("quiet mode should not include explanation")
	}
}

func TestParseNotFoundCommandBash(t *testing.T) {
	cases := []struct {
		name    string
		stderr  string
		command string
		want    string
	}{
		{
			name:    "bash style",
			stderr:  "sh: ss: command not found\n",
			command: "ss -tuln",
			want:    "ss",
		},
		{
			name:    "bash with line number",
			stderr:  "bash: line 1: htop: command not found\n",
			command: "htop",
			want:    "htop",
		},
		{
			name:    "zsh style",
			stderr:  "zsh: command not found: rg\n",
			command: "rg foo",
			want:    "rg",
		},
		{
			name:    "fallback to first token",
			stderr:  "some unexpected error\n",
			command: "nonexistent --flag",
			want:    "nonexistent",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := parseNotFoundCommand(tc.stderr, tc.command)
			if got != tc.want {
				t.Errorf("parseNotFoundCommand(%q, %q) = %q, want %q", tc.stderr, tc.command, got, tc.want)
			}
		})
	}
}

func TestInstallSuggestion(t *testing.T) {
	suggestion := installSuggestion("ripgrep")

	switch runtime.GOOS {
	case "darwin":
		if !strings.Contains(suggestion, "brew install ripgrep") {
			t.Errorf("expected brew suggestion on macOS, got: %s", suggestion)
		}
	case "linux":
		if !strings.Contains(suggestion, "ripgrep") {
			t.Errorf("expected ripgrep in suggestion, got: %s", suggestion)
		}
	default:
		if !strings.Contains(suggestion, "ripgrep") {
			t.Errorf("expected ripgrep in suggestion, got: %s", suggestion)
		}
	}
}

func TestRunCommandNotFound(t *testing.T) {
	// Capture stderr to verify the hint is printed
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	err := RunCommand("this_command_does_not_exist_xyz123")

	w.Close()
	os.Stderr = oldStderr

	if err == nil {
		t.Fatal("expected error for nonexistent command")
	}

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "not installed") {
		t.Errorf("expected 'not installed' hint in stderr, got: %q", output)
	}
}
