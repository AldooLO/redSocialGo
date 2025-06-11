git add .
git commit -m "Ultimo Commit"
git push
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o bootstrap main.go
del deployment.zip
zip deployment.zip main