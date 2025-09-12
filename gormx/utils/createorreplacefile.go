package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateOrReplaceFile(path, code string) error {
	// Ensure the folder path exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Check if the file exists
	if _, err := os.Stat(path); err == nil {
		// File exists, delete it
		if err := os.Remove(path); err != nil {
			return fmt.Errorf("failed to delete existing file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		// If there's an error other than file not existing, return it
		return fmt.Errorf("failed to check file existence: %w", err)
	}

	// Create the file
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Write the code to the file
	if _, err := file.WriteString(code); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
