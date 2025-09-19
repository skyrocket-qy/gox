package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateOrReplaceFile(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Test case 1: Create a new file
	filePath := filepath.Join(tmpDir, "testfile.txt")
	content := "hello world"

	err := CreateOrReplaceFile(filePath, content)
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

	// Test case 3: Error case - read-only directory
	readOnlyDir := filepath.Join(tmpDir, "readonly")

	err = os.Mkdir(readOnlyDir, 0o555) // Read and execute permissions
	if err != nil {
		t.Fatalf("Failed to create read-only dir: %v", err)
	}

	filePathInReadOnlyDir := filepath.Join(readOnlyDir, "testfile.txt")

	err = CreateOrReplaceFile(filePathInReadOnlyDir, content)
	if err == nil {
		t.Errorf("Expected an error when writing to a read-only directory, but got nil")
	}

	// Test case 4: Error case - path is a non-empty directory
	dirPath := filepath.Join(tmpDir, "dir")

	err = os.Mkdir(dirPath, 0o755)
	if err != nil {
		t.Fatalf("Failed to create dir: %v", err)
	}
	// Create a file in the directory to make it non-empty
	_, err = os.Create(filepath.Join(dirPath, "somefile.txt"))
	if err != nil {
		t.Fatalf("Failed to create file in dir: %v", err)
	}

	err = CreateOrReplaceFile(dirPath, content)
	if err == nil {
		t.Errorf("Expected an error when path is a non-empty directory, but got nil")
	}
}
