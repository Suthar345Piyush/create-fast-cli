/*
writing all the rendered files to the disk

this write function will take a slice of RenderedFile values and write each of them on disk under the root directory

using standard file permission conventions

For Files - 0644 (owner can read and write to it, everyone else can only read)

For Directories - 0755 (owner has full control(read, write, execute) and everyone else can just read and "enter", but not change the content of it)

(0755, 0644) the leading zero is optional and it indicates the octal number system and it is ignored
*/

package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
)

func Write(rootDir string, files []RenderedFile) error {

	for _, f := range files {
		if err := writeFile(rootDir, f); err != nil {
			return err
		}
	}
	return nil

}

// function for write to single rendered file on disk

func writeFile(rootDir string, rf RenderedFile) error {

	dest := filepath.Join(rootDir, rf.RelPath)

	// parent directory should exist
	// directory permission -> 0755
	// here 'o' for octal representation

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf("Create directory for %q: %w", dest, err)
	}

	// writing to the file
	// permission - 0644

	if err := os.WriteFile(dest, rf.Content, 0o644); err != nil {
		return fmt.Errorf("Write file %q: %w", dest, err)
	}

	return nil

}

// function to create the directory if not already exist

func EnsureDict(dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create directory %q: %w", dir, err)
	}
	return nil
}

//function to check if output directory exists or not

func OutputDirExists(dir string) (bool, error) {

	entries, err := os.ReadDir(dir)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return len(entries) > 0, nil

}
