
### 交叉编译
```
GOOS=linux GOARCH=amd64 go build
```

### 守护进程
```
pm2 start ./goletter-go-api --name "name" -- run start
```

### nginx 部署
```
server {
    listen       80;
    server_name go.goletter.cn;

    location / {
	proxy_pass http://127.0.0.1:8185;
        proxy_set_header Host $host;
        proxy_set_header X-Forwarded-For $remote_addr;
    }
}
```