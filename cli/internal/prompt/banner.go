package prompt

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

var (
	bannerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#D8B4FE")).
			Bold(true)

	shadowStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8B5CF6"))

	taglineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#94A3B8"))
)

const bannerText = `
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝

███████╗ █████╗ ███████╗████████╗
██╔════╝██╔══██╗██╔════╝╚══██╔══╝
█████╗  ███████║███████╗   ██║
██╔══╝  ██╔══██║╚════██║   ██║
██║     ██║  ██║███████║   ██║
╚═╝     ╚═╝  ╚═╝╚══════╝   ╚═╝

 ██████╗██╗     ██╗
██╔════╝██║     ██║
██║     ██║     ██║
██║     ██║     ██║
╚██████╗███████╗██║
 ╚═════╝╚══════╝╚═╝
`

func printGoodBanner() {
	fmt.Println()

	fmt.Println(
		lipgloss.NewStyle().
			MarginLeft(2).
			Render(shadowStyle.Render(bannerText)),
	)

	fmt.Println(
		taglineStyle.Render(
			"⚡ Scaffold production-ready Go CLI applications in seconds",
		),
	)

	fmt.Println()
}
