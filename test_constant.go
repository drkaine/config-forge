package main

const (
	WrongChoiceError    = "Wrong choice"
	GoodChoiceError     = "Good choice"
	ApacheInput         = "apache"
	BadInput            = "bad"
	Name                = "default-test"
	ServerName          = "test.net"
	DocumentRoot        = "/var/www/monsite.com/public_html"
	ApacheConfigContent = `
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
