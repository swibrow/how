package ui

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	commandStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("82"))
	explanationStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	labelStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
	errorStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("196"))
)

type Result struct {
	Command     string
	Explanation string
}

// ParseResponse extracts command and explanation from the LLM response.
func ParseResponse(response string) Result {
	var result Result

	for _, line := range strings.Split(response, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "COMMAND:") {
			result.Command = strings.TrimSpace(strings.TrimPrefix(line, "COMMAND:"))
		} else if strings.HasPrefix(line, "EXPLANATION:") {
			result.Explanation = strings.TrimSpace(strings.TrimPrefix(line, "EXPLANATION:"))
		}
	}

	return result
}

// Display shows the formatted result to the user.
func Display(result Result) {
	fmt.Println()
	fmt.Printf("  %s %s\n", labelStyle.Render("$"), commandStyle.Render(result.Command))
	if result.Explanation != "" {
		fmt.Printf("  %s\n", explanationStyle.Render(result.Explanation))
	}
	fmt.Println()
}

// DisplayQuiet shows only the command (for piping).
func DisplayQuiet(result Result) {
	fmt.Println(result.Command)
}

// DisplayError shows a formatted error message.
func DisplayError(msg string) {
	fmt.Fprintf(os.Stderr, "\n  %s %s\n\n", errorStyle.Render("Error:"), msg)
}

// ConfirmAndRun prompts the user to run the command and executes it.
func ConfirmAndRun(command string) error {
	fmt.Printf("  Run this command? [y/N] ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	input = strings.TrimSpace(strings.ToLower(input))
	if input != "y" && input != "yes" {
		return nil
	}

	return RunCommand(command)
}

// RunCommand executes a command via the shell.
// If the command is not found (exit code 127), it suggests how to install it.
func RunCommand(command string) error {
	fmt.Println()
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	var stderrBuf bytes.Buffer
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 127 {
			cmdName := parseNotFoundCommand(stderrBuf.String(), command)
			if cmdName != "" {
				fmt.Fprintln(os.Stderr)
				fmt.Fprintf(os.Stderr, "  %s %s is not installed.\n", hintStyle.Render("Hint:"), cmdName)
				fmt.Fprintf(os.Stderr, "  %s\n", installSuggestion(cmdName))
			}
		}
	}
	return err
}

var (
	hintStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214"))

	// Matches patterns like "sh: ss: command not found" or "bash: ss: command not found"
	notFoundRe = regexp.MustCompile(`(?:sh|bash):\s*(?:line \d+:\s*)?(\S+):\s*(?:command )?not found`)
	// Matches zsh pattern: "zsh: command not found: ss"
	notFoundZshRe = regexp.MustCompile(`zsh:\s*command not found:\s*(\S+)`)
)

// parseNotFoundCommand extracts the missing command name from shell stderr output.
// Falls back to the first token of the original command.
func parseNotFoundCommand(stderr, command string) string {
	if m := notFoundRe.FindStringSubmatch(stderr); len(m) > 1 {
		return m[1]
	}
	if m := notFoundZshRe.FindStringSubmatch(stderr); len(m) > 1 {
		return m[1]
	}
	// Fallback: first token of the command
	if fields := strings.Fields(command); len(fields) > 0 {
		return fields[0]
	}
	return ""
}

// installSuggestion returns a platform-aware install hint.
func installSuggestion(cmdName string) string {
	switch runtime.GOOS {
	case "darwin":
		return fmt.Sprintf("Install with: brew install %s", cmdName)
	case "linux":
		if _, err := exec.LookPath("apt"); err == nil {
			return fmt.Sprintf("Install with: sudo apt install %s", cmdName)
		}
		if _, err := exec.LookPath("dnf"); err == nil {
			return fmt.Sprintf("Install with: sudo dnf install %s", cmdName)
		}
		if _, err := exec.LookPath("pacman"); err == nil {
			return fmt.Sprintf("Install with: sudo pacman -S %s", cmdName)
		}
		return fmt.Sprintf("Install %s using your system package manager", cmdName)
	default:
		return fmt.Sprintf("Install %s using your system package manager", cmdName)
	}
}
