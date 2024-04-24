package configurator

import (
	"config-forge/configs"
	"strings"
)

func ConfigBuilder(nameTool string, informations map[string]string) interface{ Config } {
	switch nameTool {
	case "apache":
		return Apache{
			Name:         informations["name"] + ".conf",
			ServerName:   informations["serverName"],
			DocumentRoot: informations["documentRoot"],
			Path:         configs.APACHE_PATH_REPOSITORY,
			FileContent:  configs.APACHE_CONFIG_CONTENT,
		}
	case "nginx":
		return Nginx{
			Name:         informations["name"],
			ServerName:   informations["serverName"],
			DocumentRoot: informations["documentRoot"],
			Path:         configs.NGINX_PATH_REPOSITORY,
			FileContent:  configs.NGINX_CONFIG_CONTENT,
		}
	}
	return nil
}

type Config interface {
	CustomiseConfigContent() string
	GetName() string
}

type Apache struct {
	Name         string
	ServerName   string
	DocumentRoot string
	Path         string
	FileContent  string
}

func (apache Apache) CustomiseConfigContent() string {
	content := strings.Replace(apache.FileContent, "ServerNameInput", apache.ServerName, 1)
	content = strings.Replace(content, "ServerAliasInput", "www."+apache.ServerName, 1)
	content = strings.Replace(content, "DocumentRootInput", apache.DocumentRoot, 2)

	return content
}

func (apache Apache) GetName() string {
	return apache.Name
}

type Nginx struct {
	Name         string
	ServerName   string
	DocumentRoot string
	Path         string
	FileContent  string
}

func (nginx Nginx) CustomiseConfigContent() string {
	content := strings.Replace(nginx.FileContent, "ServerNameInput", nginx.ServerName, 2)
	content = strings.Replace(content, "DocumentRootInput", nginx.DocumentRoot, 1)

	return content
}

func (nginx Nginx) GetName() string {
	return nginx.Name
}
