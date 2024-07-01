package templates

const (
	APACHE_CONFIG = `
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
	APACHE_CONFIG_TEST = `
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
