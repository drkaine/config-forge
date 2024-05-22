package configs

var ConfigAccepted = []string{
	"apache",
	"nginx",
}

const (
	HELLO_RETURN                 = "Hello, choose your tools at configure : "
	PREPARE_CONFIGURATION_RETURN = "Now go to prepare the configuration"
	WRONG_CHOICE_RETURN          = "Uncorrect choice !"
	GOODBY_RETURN                = "Goodbye !"
	ASK_NAME_FILE                = "How would you name the configuration file ? :"
	ASK_NAME_SERVER              = "Wath is the name of the server ? :"
	ASK_DOCUMENT_ROOT_PATH       = "What is the path to the document root of the project ? :"
	ASK_CONTINU                  = "Want an other configuration ? (y/n):"
	APACHE_PATH_REPOSITORY       = "/etc/apache2/sites-available"
	APACHE_CONFIG_CONTENT        = `
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
	NGINX_PATH_REPOSITORY = "/etc/nginx/sites-available/"
	NGINX_CONFIG_CONTENT  = `
	server {

		server_name ServerNameInput www.ServerNameInput;
		root DocumentRootInput;
	
		index index.html index.htm index.php;
		charset utf-8;
	
		location / {
			try_files $uri $uri/ /index.php?$query_string;
		}
	
		location = /favicon.ico { access_log off; log_not_found off; }
		location = /robots.txt { access_log off; log_not_found off; }
		error_page 404 /index.php;
	
		location ~ /\.(?!well-known).* {
			deny all;
		}
	}`
	MINIMUM_ARGUMENTS_RETURN     = "Need 5 arguments : go run main.go toolName fileName serverName documentRoot and received less"
	MAXIMUM_ARGUMENTS_RETURN     = "Need 5 arguments : go run main.go toolName fileName serverName documentRoot and received more"
	MINIMUM_FILE_NAME_RETURN     = "fileName need minimum 2 caracter"
	MINIMUM_SERVER_NAME_RETURN   = "serverName need minimum 4 caracter"
	MINIMUM_DOCUMENT_ROOT_RETURN = "documentRoot minimum 2 caracter"
	SANITY_CHECK_PASSED          = "ok"
)
