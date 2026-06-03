// defaults configs, common for all app-types
// each template has it's own default values

package config

// these are the default values, prompt layer will overwrite any field the user customises

func DefaultConfig() ProjectConfig {
	return ProjectConfig{

		Language:       "go",
		Framework:      FrameworkCobra,
		UseTUI:         true,
		UseTesting:     true,
		UseLogging:     true,
		UseConfig:      true,
		UseCompletions: true,
		IDE:            IDEVscode,
		GoVersion:      "1.26",
	}
}

// label so that human can know about the app type

func AppTypeLabel(t AppType) string {
	labels := map[AppType]string{
		AppTypeDevTool:        "Dev Tool",
		AppTypeGitClient:      "Git Client",
		AppTypeAiAssistant:    "AI Assistant",
		AppTypeFileExplorer:   "File Explorer",
		AppTypeK8sTool:        "K8s Tool",
		AppTypePackageManager: "Package Manager",
		AppTypeSystemMonitor:  "System Moniter",
	}

	if l, ok := labels[t]; ok {
		return l
	}

	return string(t)

}

// frame work label

func FrameworkLabel(f Framework) string {
	switch f {
	case FrameworkCobra:
		return "Cobra (recommended)"

	case FrameworkUrfaveCli:
		return "urfave/cli"

	default:
		return string(f)

	}
}

// ide label

func IDELabel(i IDE) string {
	switch i {
	case IDEVscode:
		return "VS Code"

	case IDECursor:
		return "Cursor"

	case IDENone:
		return "Don't open"

	default:
		return string(i)

	}
}
