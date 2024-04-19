package main

var configAccepted = []string{
	"apache",
	"ngnix",
}

const (
	Hello                = "Hello, choose your tools at configure : "
	PrepareConfiguration = "Now go to prepare the configuration"
	WrongChoice          = "Uncorrect choice !"
	AskNameFile          = "How would you name the configuration file ? :"
	AskServerName        = "Wath is the name of the server ? :"
	AskDocumentRoot      = "What is the path to the document root of the project ? :"
	ApachePath           = "/etc/apache2/sites-available"
	ApacheConfig         = `
	<VirtualHost *:80>
	    ServerName ServerNameInput
	    ServerAlias ServerAliasInput
	    DocumentRoot DocumentRootInput

	    <Directory DocumentRootInput>
	        Options Indexes FollowSymLinks
	        AllowOverride All
	        Require all granted
	    </Directory>

	    ErrorLog ${APACHE_LOG_DIR}/monsite_error.log
	    CustomLog ${APACHE_LOG_DIR}/monsite_access.log combined
	</VirtualHost>`
)
