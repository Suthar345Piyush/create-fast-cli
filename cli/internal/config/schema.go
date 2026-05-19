// data models (structs) for app-types

package config

// type of app user want to build
type AppType string

const (
	AppTypeDevTool        AppType = "dev-tool"
	AppTypeGitClient      AppType = "git-client"
	AppTypeFileExplorer   AppType = "file-explorer"
	AppTypeK8sTool        AppType = "k8s-tool"
	AppTypeAiAssistant    AppType = "ai-assitant"
	AppTypeSystemMoniter  AppType = "system-moniter"
	AppTypePackageManager AppType = "package-manager"
)

// cli framework to use (options - cobra , urfave/cli/v3)

type Framework string

const (
	FrameworkCobra     Framework = "cobra"
	FrameworkUrfaveCli Framework = "urfave/cli"
)

// ide options  (vscodem, cursor)

type IDE string

const (
	IDEVscode IDE = "vscode"
	IDECursor IDE = "cursor"
	IDENone   IDE = "none"
)

// projectconfig struct - central data model for the whole application pipeline
