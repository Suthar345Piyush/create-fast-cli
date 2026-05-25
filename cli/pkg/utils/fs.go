// some filesystem helper functions

package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// MkdirAll creates the directory

func MkdirAll(dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("mkdir %q: %w", dir, err)
	}
	return nil
}

// writing into the file

func WriteFile(path string, content []byte) error {

	if err := MkdirAll(filepath.Dir(path)); err != nil {
		return err
	}

	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write %q: %w", path, err)
	}
	return nil

}

// copy source file to destination file

func CopyFile(src, dest string) error {

	in, err := os.Open(src)

	if err != nil {
		return fmt.Errorf("open src %q: %w", src, err)
	}

	defer in.Close()

	if err := MkdirAll(filepath.Dir(dest)); err != nil {
		return err
	}

	out, err := os.Create(dest)

	if err != nil {
		return fmt.Errorf("create dest %q: %w", dest, err)
	}

	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy %q -> %q: %w", src, dest, err)
	}

	return nil

}

// it returns true, if the directory is empty  or does not exist

func DirIsEmpty(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)

	if os.IsNotExist(err) {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	return len(entries) == 0, nil

}
