server {
	root /var/www/html;
	server_name api.backend.mama.sh www.api.backend.mama.sh;

	location / {
		proxy_pass http://127.0.0.1:5080;
	}


    listen [::]:443 ssl; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/api.backend.mama.sh/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/api.backend.mama.sh/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot


}
server {
    if ($host = www.api.backend.mama.sh) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


    if ($host = api.backend.mama.sh) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


	listen 80 ;
	listen [::]:80 ;
	server_name api.backend.mama.sh www.api.backend.mama.sh;
    return 404; # managed by Certbot




}
