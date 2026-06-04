// ide - it will detect which ide is installed on machine, for now it supports only two - vscode and cursor, and then open the project

package scaffold

import (
	"fmt"
	"os/exec"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
)

var ideCommand = map[config.IDE]string{
	config.IDEVscode: "code",
	config.IDECursor: "cursor",
}

// open ide function

func OpenIDE(projectDir string, ide config.IDE) (opened bool, err error) {

	fmt.Println("opening ide:", ide)
	fmt.Println("project:", projectDir)

	if ide == config.IDENone {
		return false, nil
	}

	bin, ok := ideCommand[ide]

	if !ok {
		return false, fmt.Errorf("unknown IDE %q", ide)
	}

	fmt.Printf("IDE=%s\n", ide)
	fmt.Printf("DIR=%s\n", projectDir)

	binPath, err := exec.LookPath(bin)

	fmt.Printf("BIN=%s\n", binPath)
	fmt.Printf("ERR=%v\n", err)

	if err != nil {
		return false, fmt.Errorf("%s not found in PATH", bin)
	}

	cmd := exec.Command(binPath, projectDir)
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return false, fmt.Errorf("open %s: %w", config.IDELabel(ide), err)
	}

	_ = cmd.Process.Release()

	return true, nil

}

// last function to detect already installed ide's on a machine

func DetectIDE() []config.IDE {

	ides := []config.IDE{config.IDEVscode, config.IDECursor}

	var found []config.IDE

	for _, ide := range ides {

		bin := ideCommand[ide]

		if _, err := exec.LookPath(bin); err == nil {
			found = append(found, ide)
		}

	}

	return found

}
