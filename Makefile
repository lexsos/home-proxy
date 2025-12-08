format:
	go fmt ./...
	goimports -w -local "github.com/lexsos/home-proxy" .

tests:
	go test -v ./...

clear:
	rm -rf ./build

build_dir:
	mkdir build

build_linux64: build_dir
	GOOS=linux GOARCH=amd64 go build -o ./build/hproxy
