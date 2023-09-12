build:
	@make build-linux && make build-windows && make build-mac-m1

build-linux:
	@cd cmd && GOOS=linux GOARCH=amd64 go build -o ../gateway-linux

build-windows:
	@cd cmd && GOOS=windows GOARCH=amd64 go build -o ../gateway-windows.exe

build-mac-m1:
	@cd cmd && GOOS=darwin GOARCH=arm64 go build -o ../gateway-mac-m1

run:
	@go run cmd/main.go