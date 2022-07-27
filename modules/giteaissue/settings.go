package giteaissue

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Gitea Issus"
)

type Settings struct {
	*cfg.Common

	numberOfTodos int    `help:"Defines number of stories to be displayed. Default is 10" optional:"true"`
	apiKey        string `help:"A Gitea personal access token. Requires at least api access."`
	domain        string `help:"Your Gitea URL."`
	showProject   bool   `help:"Determines whether or not to show the project a given todo is for."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		numberOfTodos: ymlConfig.UInt("numberOfTodos", 10),
		apiKey:        ymlConfig.UString("apiKey", os.Getenv("WTF_GITEA_TOKEN")),
		domain:        ymlConfig.UString("domain", "https://example.com"),
		showProject:   ymlConfig.UBool("showProject", true),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).
		Service(settings.domain).Load()

	return &settings
}
