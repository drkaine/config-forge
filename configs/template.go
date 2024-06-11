package configs

const (
	APACHE_CONFIG_TEMPLATE = `
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
	NGINX_CONFIG_TEMPLATE = `
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
)
