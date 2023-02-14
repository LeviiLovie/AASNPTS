# Running
run:
	@go run main.go
safe:
	@go run main.go -safe

# Build
mac_build:
	@mkdir -p bin/$(VER)
	@GOOS=darwin GOARCH=amd64 go build -o bin/$(VER)/pepstab_macos main.go
linux_build:
	@mkdir -p bin/$(VER)
	@GOOS=linux GOARCH=amd64 go build -o bin/$(VER)/pepstab_linux main.go
windows_build:
	@mkdir -p bin/$(VER)
	@GOOS=windows GOARCH=amd64 go build -o bin/$(VER)/pepstab_windows.exe main.go
build: mac_build linux_build windows_build

# Module
mod:
	@go mod tidy
	@go mod vendor