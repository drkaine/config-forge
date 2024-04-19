package main

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
)
