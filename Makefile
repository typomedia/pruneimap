build: tidy
	go build -ldflags "-s -w" -o dist/ .

run: tidy
	go run main.go

compile: tidy
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-s -w" -o dist/pruneimap-linux-arm .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/pruneimap-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/pruneimap-windows-amd64.exe .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/pruneimap-macos-amd64 .
	GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w" -o dist/pruneimap-freebsd-amd64 .

tidy:
	go mod tidy

loc:
	go install github.com/boyter/scc/v3@latest
	scc --exclude-dir vendor --exclude-dir bin .