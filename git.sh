#！ /bin/bash
git add . &&
git commit -m "自动提交" &&
echo "已提交修改" &&
git pull &&
git push &&
echo "修改已上传"
