// bubble tea TUI design - loading spinner and step - by - step progress for completion

package prompt

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
)

// step message struct

type StepMsg struct {
	Label string
	Err   error
}

// done message type after completion

type DoneMsg struct {
	Err error
}

// tick message for waiting

type tickMsg time.Time

// bubble tea model - progress model , when app is written on the terminal

type ProgressModel struct {
	steps      []string
	done       []bool
	current    int
	spinnerIdx int
	err        error
	finished   bool
}

// spinner animation - simple rotation

var spinnerFrame = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func NewProgressModel(steps []string) ProgressModel {

	return ProgressModel{
		steps: steps,
		done:  make([]bool, len(steps)),
	}
}

// bubble tea interface

// bubble - tea model contains these core functions like - Init, Update and view

// Init is the first function that will be called

// update function works when a message is received

// view shows/render the terminal ui and it rendered after every update

func (pg ProgressModel) Init() tea.Cmd {
	return tick()
}

// update function - it will return model and cmd

func (pg ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tickMsg:
		if pg.finished {
			return pg, nil
		}

		pg.spinnerIdx = (pg.spinnerIdx + 1) % len(spinnerFrame)

		return pg, tick()

	case StepMsg:
		if msg.Err != nil {
			pg.err = msg.Err
			pg.finished = true
			return pg, tea.Quit

		}

		if pg.current < len(pg.done) {
			pg.done[pg.current] = true
			pg.current++
		}

		return pg, nil

	case DoneMsg:
		pg.finished = true
		pg.err = msg.Err
		return pg, tea.Quit

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return pg, tea.Quit
		}

	}

	return pg, nil

}

func (pg ProgressModel) View() tea.View {

	var sb strings.Builder

	sb.WriteString("\n")

	for i, step := range pg.steps {
		switch {
		case pg.done[i]:
			sb.WriteString(StepDoneStyle.Render(" ✅") + step + "\n")

		case i == pg.current && !pg.finished:
			frame := SpinnerStyle.Render(spinnerFrame[pg.spinnerIdx])

			sb.WriteString(fmt.Sprintf("  %s %s\n", frame, ProgressStyle.Render(step)))

		default:
			sb.WriteString(MutedStyle.Render(fmt.Sprintf(" . %s\n", step)))
		}

	}

	if pg.finished && pg.err == nil {

		sb.WriteString("\n" + SuccessStyle.Render(" ✅ Done!") + "\n")

	}

	if pg.err != nil {
		sb.WriteString("\n" + SuccessStyle.Render(" ❎ Error: "+pg.err.Error()) + "\n")
	}

	sb.WriteString("\n")

	return tea.NewView(sb.String())

}

// tick  function

func tick() tea.Cmd {
	return tea.Tick(80*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
