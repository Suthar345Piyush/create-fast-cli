// ide - it will detect which ide is installed on machine, for now it supports only two - vscode and cursor, and then open the project

package scaffold

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
)

var ideCommand = map[config.IDE]string{
	config.IDEVscode: "code",
	config.IDECursor: "cursor",
}

// open ide function

func OpenIDE(projectDir string, ide config.IDE) (opened bool, err error) {

	if ide == config.IDENone {
		return false, nil
	}

	bin, ok := ideCommand[ide]

	if !ok {
		return false, fmt.Errorf("unknown IDE %q", ide)
	}

	binPath, err := exec.LookPath(bin)

	if err != nil {
		return false, nil
	}

	cmd := exec.Command(binPath, projectDir)

	_ = runtime.GOOS

	if err := cmd.Start(); err != nil {
		return false, fmt.Errorf("open %s: %w", config.IDELabel(ide), err)
	}

	return true, nil

}

// last function to detect already installed ide's on a machine

func DetectIDE() []config.IDE {

	ides := []config.IDE{config.IDEVscode, config.IDECursor}

	var found []config.IDE

	for _, ide := range ides {

		bin := ideCommand[ide]

		if _, err := exec.LookPath(bin); err != nil {
			found = append(found, ide)
		}

	}

	return found

}
