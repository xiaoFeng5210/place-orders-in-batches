### 依赖安装
```bash
go mod tidy
```

### 运行
```bash
go run .
```

### 打包
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .
```

### 文件权限
```bash
chmod +x main
```

### 服务重启
```bash
sudo systemctl restart go-mock-order.service
```
