git add .
git commit -m "Ultimo Commit"
git push
GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go buil -tags lambda.norpc -o bootstrap main.go
del go_lambda.zip
zip go_lambda.zip bootstrap