#！ /bin/bash
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
git add .
git commit -m "ccc"
echo "已提交修改"
git pull
git push
echo "修改已上传"
echo off