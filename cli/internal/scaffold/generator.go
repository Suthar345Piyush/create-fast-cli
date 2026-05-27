/*
generator will works like this -

  use render ---> write to the disk ----> opens the IDE with TUI messages step by step

*/

package scaffold

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/config"
	"github.com/Suthar345Piyush/create-fast-cli/cli/internal/prompt"
)

// slice of steps messages, shown on terminal after each step completion

var StepMessage = []string{
	"Validating output directory",
	"Rendering template",
	"Writing project files",
	"Saving preferences",
	"Opening IDE",
}

func Generate(cfg *config.ProjectConfig) error {

	model := prompt.NewProgressModel(StepMessage)

	// p := tea.NewProgram(tea.Model, tea.ProgramOption(func (*tea.Program)))

	p := tea.NewProgram(model)

	errCh := make(chan error, 1)

	//go routine

	go func() {
		errCh <- runSteps(p, cfg)
	}()

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("TUI error: %w", err)
	}

	return <-errCh

}

// function to run step by step each process completion

func runSteps(p *tea.Program, cfg *config.ProjectConfig) error {

	// validating the output directory

	exists, err := OutputDirExists(cfg.OutputDir)

	if err != nil {
		p.Send(prompt.StepMsg{Err: err})
		return err
	}

	if exists {
		err = fmt.Errorf("output directory %q already exists and is not empty", cfg.OutputDir)
		p.Send(prompt.StepMsg{Err: err})
		return err
	}

	p.Send(prompt.StepMsg{Label: StepMessage[0]})

	// rendering the templates

	rendered, err := Render(cfg)

	if err != nil {
		p.Send(prompt.StepMsg{Err: fmt.Errorf("render: %w", err)})
		return err
	}

	p.Send(prompt.StepMsg{Label: StepMessage[1]})

	// write file to disk

	if err := EnsureDict(cfg.OutputDir); err != nil {
		p.Send(prompt.StepMsg{Err: err})
		return err
	}

	if err := Write(cfg.OutputDir, rendered); err != nil {

		//removing partially written directory

		_ = os.RemoveAll(cfg.OutputDir)

		p.Send(prompt.StepMsg{Err: fmt.Errorf("write: %w", err)})

		return err

	}

	p.Send(prompt.StepMsg{Label: StepMessage[2]})

	// user preferences

	if err := config.SavePreferences(cfg); err != nil {
		_ = err
	}

	p.Send(prompt.StepMsg{Label: StepMessage[3]})

	// open the IDE

	opened, err := OpenIDE(cfg.OutputDir, cfg.IDE)

	if err != nil {
		_ = err
	}

	_ = opened

	p.Send(prompt.StepMsg{Label: StepMessage[4]})

	// done message

	p.Send(prompt.DoneMsg{})

	printSuccess(cfg)

	return nil

}

// printSuccess function

func printSuccess(cfg *config.ProjectConfig) {
	fmt.Println()
	fmt.Printf("✅ Project %q created at %s\n", cfg.ProjectName, cfg.OutputDir)
	fmt.Println()
	fmt.Println(" Next steps:")
	fmt.Printf("  cd %s\n", cfg.OutputDir)
	fmt.Println("  go mod tidy")
	fmt.Println("   go run . --help")
	fmt.Println()
}
