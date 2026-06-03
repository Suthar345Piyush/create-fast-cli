// lipgloss styles for prompts and components

package prompt

import "charm.land/lipgloss/v2"

var (
	colorPrimary   = lipgloss.Color("#E07B54") // - main
	colorSecondary = lipgloss.Color("#6C91C2") // - secondary
	colorMuted     = lipgloss.Color("#6C6C6C") // - hints
	colorSuccess   = lipgloss.Color("#4CAF50") // - completion message
	colorError     = lipgloss.Color("#E53935") // - errors

	AppTitleStyle = lipgloss.NewStyle().Bold(true).Foreground(colorPrimary).MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().Foreground(colorMuted).Italic(true)

	SectionStyle = lipgloss.NewStyle().Bold(true).Foreground(colorSecondary).MarginTop(1)

	SuccessStyle = lipgloss.NewStyle().Bold(true).Foreground(colorSuccess)

	ErrorStyle = lipgloss.NewStyle().Foreground(colorError)

	MutedStyle = lipgloss.NewStyle().Foreground(colorMuted)

	SummaryBoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colorPrimary).Padding(1, 2).MarginTop(1)

	SummaryKeyStyle = lipgloss.NewStyle().Foreground(colorSecondary).Width(18)

	SummaryValueStyle = lipgloss.NewStyle().Bold(true)

	SpinnerStyle = lipgloss.NewStyle().Foreground(colorPrimary)

	ProgressStyle = lipgloss.NewStyle().Foreground(colorSecondary)

	StepDoneStyle = lipgloss.NewStyle().Foreground(colorSuccess).Bold(true)
)
