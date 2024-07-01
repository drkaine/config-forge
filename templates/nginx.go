package templates

const (
	NGINX_CONFIG = `
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
	NGINX_CONFIG_TEST = `
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
