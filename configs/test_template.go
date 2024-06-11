package configs

const (
	APACHE_CONFIG_TEMPLATE_TEST = `
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
	NGINX_CONFIG_TEMPLATE_TEST = `
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
