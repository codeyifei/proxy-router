tidy:
	go mod tidy

build: tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/proxy-router-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/proxy-router-linux-arm64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/proxy-router-windows-amd64.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o bin/proxy-router-windows-arm64.exe main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/proxy-router-darwin-amd64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/proxy-router-darwin-arm64 main.go
