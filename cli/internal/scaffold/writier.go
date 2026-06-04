/*
writing all the rendered files to the disk

this write function will take a slice of RenderedFile values and write each of them on disk under the root directory

using standard file permission conventions

For Files - 0644 (owner can read and write to it, everyone else can only read)

For Directories - 0755 (owner has full control(read, write, execute) and everyone else can just read and "enter", but not change the content of it)

(0755, 0644) the leading zero is optional and it indicates the octal number system and it can be ignored
*/

package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
)

func Write(rootDir string, files []RenderedFile) error {

	// fmt.Printf("passing %d files to written\n", len(ren))

	fmt.Printf("writing %d files\n", len(files))

	for _, f := range files {

		fmt.Printf("File: %s\n", f.RelPath)

		dest := filepath.Join(rootDir, f.RelPath)

		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return fmt.Errorf("create directory for %q: %w", dest, err)
		}

		if err := os.WriteFile(dest, f.Content, 0o644); err != nil {
			return fmt.Errorf("write file %q: %w", dest, err)
		}

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
