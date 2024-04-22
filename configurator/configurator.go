package configurator

import (
	"config-forge/configs"
	"strings"
)

func InstanciationConfig(nameTool string, informations map[string]string) Apache {
	configs := Apache{
		Name:         informations["name"] + ".conf",
		ServerName:   informations["serverName"],
		DocumentRoot: informations["documentRoot"],
		Path:         configs.APACHE_PATH_REPOSITORY,
		FileContent:  configs.APACHE_CONFIG_CONTENT,
	}

	return configs
}

type Config interface {
	CustomiseConfigContent(config Config)
}

type Apache struct {
	Name         string
	ServerName   string
	DocumentRoot string
	Path         string
	FileContent  string
}

func (apache *Apache) CustomiseConfigContent() {
	apache.FileContent = strings.Replace(apache.FileContent, "ServerNameInput", apache.ServerName, 1)
	apache.FileContent = strings.Replace(apache.FileContent, "ServerAliasInput", "www."+apache.ServerName, 1)
	apache.FileContent = strings.Replace(apache.FileContent, "DocumentRootInput", apache.DocumentRoot, 2)
}

type Ngnix struct {
	Name         string
	ServerName   string
	DocumentRoot string
	Path         string
	FileContent  string
}

func (ngnix *Ngnix) CustomiseConfigContent() {
	ngnix.FileContent = strings.Replace(ngnix.FileContent, "ServerNameInput", ngnix.ServerName, 2)
	ngnix.FileContent = strings.Replace(ngnix.FileContent, "DocumentRootInput", ngnix.DocumentRoot, 1)
}
