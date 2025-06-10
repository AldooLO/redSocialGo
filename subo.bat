git add .
git commit -m "Ultimo Commit"
git push
GOOS=linux GOARCH=amd64 go build -o main main.go
del deployment.zip
zip deployment.zip main