test-all:
	go test ./...
export-all:
	go run ./cmd/main.go all
observe:
	go run ./cmd/main.go observe
build-linux:
	go build -o dbexport ./cmd
build-windows:
	GOOS=windows go build -o dbexport-win64.exe ./cmd
build: build-linux build-windows
