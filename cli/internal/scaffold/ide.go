// ide.go

package scaffold

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
)

type IDEInfo struct {
	IDE  config.IDE
	Name string
	Path string
}

// open ide function

func OpenIDE(projectDir string, ide config.IDE) (bool, error) {
	if ide == config.IDENone {
		return false, nil
	}

	projectDir, err := filepath.Abs(projectDir)

	if err != nil {
		return false, fmt.Errorf("resolve project path: %w", err)
	}

	info, err := FindIDE(ide)

	if err != nil {
		return false, err
	}

	cmd := exec.Command(info.Path, projectDir)

	if err := cmd.Start(); err != nil {
		return false, fmt.Errorf("open %s: %w", info.Name, err)
	}

	// detaching this process from generator

	if cmd.Process != nil {
		_ = cmd.Process.Release()
	}

	return true, nil

}

func DetectIDE() []config.IDE {
	cnd := []config.IDE{
		config.IDEVscode,
		config.IDECursor,
	}

	var found []config.IDE

	for _, ide := range cnd {
		if _, err := FindIDE(ide); err == nil {
			found = append(found, ide)
		}
	}

	return found

}

// find ide function

func FindIDE(ide config.IDE) (*IDEInfo, error) {
	switch ide {
	case config.IDEVscode:
		return findVsCode()

	case config.IDECursor:
		return findCursor()

	default:
		return nil, fmt.Errorf("unknown IDE")
	}
}

// find vs code function

func findVsCode() (*IDEInfo, error) {
	if runtime.GOOS == "windows" {

		userProfile := os.Getenv("USERPROFILE")

		cnd := []string{
			filepath.Join(userProfile,
				"AppData",
				"Local",
				"Programs",
				"Microsoft VS Code",
				"Code.exe",
			),

			`C:\Program Files\Microsoft VS Code\Code.exe`,
			`C:\Program Files (x86)\Microsoft VS Code\Code.exe`,
		}

		for _, path := range cnd {
			if _, err := os.Stat(path); err == nil {
				return &IDEInfo{
					IDE:  config.IDEVscode,
					Name: "VS Code",
					Path: path,
				}, nil
			}
		}

	}

	// fall back to path

	if path, err := exec.LookPath("code"); err == nil {
		return &IDEInfo{
			IDE:  config.IDEVscode,
			Name: "VS Code",
			Path: path,
		}, nil
	}

	return nil, fmt.Errorf("VS Code not found")
}

// cursor function

func findCursor() (*IDEInfo, error) {
	if runtime.GOOS == "windows" {

		userProfile := os.Getenv("USERPROFILE")

		cnd := []string{
			filepath.Join(userProfile,
				"AppData",
				"Local",
				"Programs",
				"Cursor",
				"Cursor.exe",
			),

			`C:\Program Files\Cursor\Cursor.exe`,
		}

		for _, path := range cnd {
			if _, err := os.Stat(path); err == nil {
				return &IDEInfo{
					IDE:  config.IDECursor,
					Name: "Cursor",
					Path: path,
				}, nil
			}
		}
	}

	if path, err := exec.LookPath("cursor"); err == nil {
		return &IDEInfo{
			IDE:  config.IDECursor,
			Name: "Cursor",
			Path: path,
		}, nil
	}

	return nil, fmt.Errorf("Cursor not found")

}
