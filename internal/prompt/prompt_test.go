package prompt

import (
	"strings"
	"testing"
)

func TestSystemPromptNotEmpty(t *testing.T) {
	if SystemPrompt("") == "" {
		t.Fatal("SystemPrompt() should not be empty")
	}
}

func TestSystemPromptContainsFormat(t *testing.T) {
	p := SystemPrompt("")
	if !strings.Contains(p, "COMMAND") {
		t.Error("SystemPrompt should mention COMMAND")
	}
	if !strings.Contains(p, "EXPLANATION") {
		t.Error("SystemPrompt should mention EXPLANATION")
	}
}

func TestSystemPromptContainsOSContext(t *testing.T) {
	p := SystemPrompt("")
	if !strings.Contains(p, "user is on") {
		t.Error("SystemPrompt should contain OS-specific context")
	}
}

func TestSystemPromptCustomOverride(t *testing.T) {
	custom := "You are a helpful DevOps assistant. Respond with COMMAND: and EXPLANATION: format."
	p := SystemPrompt(custom)

	if !strings.Contains(p, "DevOps assistant") {
		t.Error("custom prompt should replace the default base prompt")
	}
	// OS context should still be appended
	if !strings.Contains(p, "user is on") {
		t.Error("OS context should still be appended to custom prompt")
	}
	// Default base prompt content should NOT be present
	if strings.Contains(p, "terminal command expert") {
		t.Error("default base prompt should be replaced by custom prompt")
	}
}

func TestSystemPromptEmptyUsesDefault(t *testing.T) {
	p := SystemPrompt("")
	if !strings.Contains(p, "terminal command expert") {
		t.Error("empty custom prompt should use the default base prompt")
	}
}
