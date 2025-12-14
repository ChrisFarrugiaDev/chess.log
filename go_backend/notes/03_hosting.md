# Apache Hosting Notes — Go Backend + Vue SPA

This virtual host serves:

- **Backend API** from Go server on port **4567**
- **Frontend SPA** (Vue build served via the Go SPA handler)
- **Domain**: `chesslog.chrisfarrugia.dev`
- Optional alternative domain: `iotrack.live`

---

## Apache VirtualHost (HTTP – port 80)

```apache
<VirtualHost *:80>
    ServerName chesslog.chrisfarrugia.dev    

    ProxyPreserveHost On
    ProxyRequests Off

    # ----------------------------------
    # 1. Backend API (Go server)
    # ----------------------------------
    ProxyPass        "/api/"  "http://127.0.0.1:4567/api/"
    ProxyPassReverse "/api/"  "http://127.0.0.1:4567/api/"

    # ----------------------------------
    # 2. Frontend SPA (Vue)
    # Handled entirely by the Go SPA handler
    # ----------------------------------
    ProxyPass        "/"  "http://127.0.0.1:4567/"
    ProxyPassReverse "/"  "http://127.0.0.1:4567/"

    ErrorLog  ${APACHE_LOG_DIR}/chesslog-error.log
    CustomLog ${APACHE_LOG_DIR}/chesslog-access.log combined
</VirtualHost>
```

Save as:

```
/etc/apache2/sites-available/chesslog.chrisfarrugia.dev.conf
```

---

## Required Apache Modules

Enable necessary modules:

```bash
sudo a2enmod proxy
sudo a2enmod proxy_http
sudo a2enmod rewrite
sudo a2enmod headers
```

For HTTPS support:

```bash
sudo a2enmod ssl
```

Restart Apache:

```bash
sudo systemctl restart apache2
```

---

## DNS Settings (Namecheap)

Add:

| Type     | Host      | Value            | TTL   |
|----------|-----------|------------------|--------|
| A Record | chesslog  | 57.129.22.122    | 5 min |

---

## Enable the site

```bash
sudo a2ensite chesslog.chrisfarrugia.dev.conf
sudo systemctl reload apache2
```

Test:

```bash
sudo apache2ctl configtest
```

---

## Add HTTPS (Certbot)

```bash
sudo certbot --apache -d chesslog.chrisfarrugia.dev
```

---

## Notes

- After certbot → **https://chesslog.chrisfarrugia.dev**
- Logs:

```
/var/log/apache2/chesslog-error.log
/var/log/apache2/chesslog-access.log
```

## Run Go Production Server From Your Project Directory

Your current working directory:

```bash
/home/ubuntu/projects/chess.log/go_backend
```

### Step 1 — Create a systemd service pointing to your project directory


```bash
sudo vi /etc/systemd/system/chesslog.service
```

```bash
[Unit]
Description=ChessLog Go Backend
After=network.target

[Service]
Type=simple

# Run from your project directory
WorkingDirectory=/home/ubuntu/projects/chess.log/go_backend

# Load your .env file before starting the Go app
# EnvironmentFile=/home/ubuntu/projects/chess.log/go_backend/.env

# Run your compiled binary
ExecStart=/home/ubuntu/projects/chess.log/go_backend/chesslog

Restart=on-failure
RestartSec=3

# Run as ubuntu user
User=ubuntu

[Install]
WantedBy=multi-user.target
```


### Step 2 — Reload + Start

```bash
sudo systemctl daemon-reload
sudo systemctl start chesslog
sudo systemctl enable chesslog
sudo systemctl restart chesslog.service
```

Check logs:

```bash
journalctl -u chesslog -f
```