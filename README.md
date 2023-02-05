## MUXWEBAPPV2
Skeleton for a new Go web app based on the Gorilla Mux router and MySQL DB backend.

## Deployment Notes

1. Build for linux (Ubuntu) architecture amd64:

```shell
cd /path/to/<project>/
GOOS=linux GOARCH=amd64 go build -o <executable> -v
```

1. Copy your project excluding the Go source `*.go` to your server in `/path/to/remote/<project>/`.

```shell
rsync -azv -e 'ssh -i yourkey.pem' --exclude '*.go' --exclude 'debug' --exclude '__debug_bin' --exclude '.gitignore' --exclude '.git' --exclude 'config.yaml' --exclude 'go.mod' --exclude 'go.sum' ../<project> user@your_server:/path/to/remote/<project>/
```
3. Create application configuration.

```shell
vim /path/to/remote/<project>/config.yaml

# paste content from app/config/config.yaml
# update values according to the environment
sudo chown <app_user>: /path/to/remote/<project>/config.yaml
sudo chmod 400 /path/to/remote/<project>/config.yaml
```

4. Configure systemd to run the app.

```shell
sudo vim /etc/systemd/system/<project>.service

# add the following
[Unit]
Description=<project> Web App

[Service]
WorkingDirectory=/path/to/remote/<project>/
StandardOutput=append:/var/log/<project>.log
StandardError=append:/var/log/<project>-error.log
ExecStart=/path/to/remote/<project>/<executable>
Environment=CONFIG_FILE=/path/to/remote/<project>/config.yaml
Restart=always
User=<app_user>

[Install]
WantedBy=multi-user.target


# reload systemd configuration
sudo systemctl daemon-reload
sudo systemctl enable <project>.service
sudo systemctl start <project>.service
```

5. Create logrotate configuration

```shell
sudo vim /etc/logrotate.d/<project>.conf

# add the following content
/var/log/<project>.log {
    size 50M
    rotate 10
    compress
    missingok
    notifempty
    copytruncate
}

/var/log/<project>-error.log {
    size 50M
    rotate 10
    compress
    missingok
    notifempty
    copytruncate
}


# test configuration
logrotate -f /etc/logrotate.d/<project>.conf
```

6. Configure nginx as the frontend.

```shell
sudo vim /etc/nginx/sites-available/<project>.conf

# add the following
server {
    listen 80;
    server_name <project_domain>;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl;
    ssl_certificate <path/to/ssl/fullchain.pem>;
    ssl_certificate_key <path/to/ssl//privkey.pem>;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_prefer_server_ciphers on;
    ssl_ciphers 'EECDH+AESGCM:EDH+AESGCM:AES256+EECDH:AES256+EDH';

    server_name <project_domain>;

    access_log /var/log/nginx/<project>_access.log;
    error_log /var/log/nginx/<project>_error.log error;

    location /static/ { alias /path/to/remote/<project>/app/assets/; }
    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}


# enable the website configuration
sudo ln -s /etc/nginx/sites-available/<project>.conf /etc/nginx/sites-enabled
sudo service nginx reload
```
