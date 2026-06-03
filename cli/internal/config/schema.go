// common data model (structs) for app-types

package config

// type of app user want to build
type AppType string

const (
	AppTypeDevTool        AppType = "dev-tool"
	AppTypeGitClient      AppType = "git-client"
	AppTypeFileExplorer   AppType = "file-explorer"
	AppTypeK8sTool        AppType = "k8s-tool"
	AppTypeAiAssistant    AppType = "ai-assitant"
	AppTypeSystemMonitor  AppType = "system-moniter"
	AppTypePackageManager AppType = "package-manager"
)

// cli framework to use (options - cobra , urfave/cli/v3)

type Framework string

const (
	FrameworkCobra     Framework = "cobra"
	FrameworkUrfaveCli Framework = "urfave"
)

// ide options  (vscodem, cursor)

type IDE string

const (
	IDEVscode IDE = "vscode"
	IDECursor IDE = "cursor"
	IDENone   IDE = "none"
)

// projectconfig struct - central data model for the whole application pipeline

type ProjectConfig struct {
	ProjectName string
	ModulePath  string
	AppType     AppType

	Framework Framework
	Language  string // supports only Go

	UseTUI         bool
	UseLogging     bool
	UseConfig      bool
	UseCompletions bool
	UseTesting     bool

	// tools

	IDE       IDE
	OutputDir string

	// version of the language

	GoVersion string // like 1.26

}
