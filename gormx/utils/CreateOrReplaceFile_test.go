package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateOrReplaceFile(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Test case 1: Create a new file
	filePath := filepath.Join(tmpDir, "testfile.txt")
	content := "hello world"

	err = CreateOrReplaceFile(filePath, content)
	if err != nil {
		t.Errorf("Failed to create file: %v", err)
	}

	// Check if the file was created with the correct content
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}

	if string(data) != content {
		t.Errorf("File content is incorrect, got: %s, want: %s", string(data), content)
	}

	// Test case 2: Replace an existing file
	newContent := "hello again"

	err = CreateOrReplaceFile(filePath, newContent)
	if err != nil {
		t.Errorf("Failed to replace file: %v", err)
	}

	// Check if the file was replaced with the new content
	data, err = os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}

	if string(data) != newContent {
		t.Errorf("File content is incorrect, got: %s, want: %s", string(data), newContent)
	}
}
