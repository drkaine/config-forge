package configs

const (
	WRONG_CHOICE_ERROR         = "Wrong choice"
	GOOD_CHOICE_ERROR          = "Good choice"
	APACHE_INPUT               = "apache"
	BAD_INPUT                  = "bad"
	NAME_TEST                  = "default-test"
	NAME_FILE_TEST             = "default-test.conf"
	SERVER_NAME_TEST           = "test.net"
	DOCUMENT_ROOT_TEST         = "/var/www/monsite.com/public_html"
	APACHE_CONFIG_CONTENT_TEST = `
	<VirtualHost *:80>
	    ServerName test.net
	    ServerAlias www.test.net
	    DocumentRoot /var/www/monsite.com/public_html

	    <Directory /var/www/monsite.com/public_html>
	        Options Indexes FollowSymLinks
	        AllowOverride All
	        Require all granted
	    </Directory>

	    ErrorLog ${APACHE_LOG_DIR}/monsite_error.log
	    CustomLog ${APACHE_LOG_DIR}/monsite_access.log combined
	</VirtualHost>`
	NGNIX_CONFIG_CONTENT_TEST = `
	server {

		server_name test.net www.test.net;
		root /var/www/monsite.com/public_html;
	
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
