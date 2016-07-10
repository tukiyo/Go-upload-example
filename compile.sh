GOOS=windows GOARCH=386 go build -o uploader.exe uploader.go
GOOS=linux GOARCH=386 go build -o uploader uploader.go
GOOS=darwin GOARCH=amd64 go build -o uploader-osx uploader.go
