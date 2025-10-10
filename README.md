### 打包
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .
```

### 服务重启
```bash
sudo systemctl restart go-mock-order.service
```
