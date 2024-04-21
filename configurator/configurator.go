package configurator

import (
	"config-forge/configs"
	"strings"
)

type Config interface {
	ImplementConfigContent()
}

type Apache struct {
	Name         string
	ServerName   string
	DocumentRoot string
	Path         string
	FileContent  string
}

func (apache *Apache) ImplementConfigContent() {
	apache.FileContent = strings.Replace(apache.FileContent, "ServerNameInput", apache.ServerName, 1)
	apache.FileContent = strings.Replace(apache.FileContent, "ServerAliasInput", "www."+apache.ServerName, 1)
	apache.FileContent = strings.Replace(apache.FileContent, "DocumentRootInput", apache.DocumentRoot, 2)
}

func InstanciationConfig(name string, serverName string, documentRoot string) Apache {
	configs := Apache{
		Name:         name + ".conf",
		ServerName:   serverName,
		DocumentRoot: documentRoot,
		Path:         configs.APACHE_PATH_REPOSITORY,
		FileContent:  configs.APACHE_CONFIG_CONTENT,
	}

	return configs
}
