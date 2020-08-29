#！ /bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ark main.go &&
chmod -R 777 ./ark &&
echo "打包完成"