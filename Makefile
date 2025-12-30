DOCKER_IMAGE = lexsos/hproxy

format:
	go fmt ./...
	goimports -w -local "github.com/lexsos/home-proxy" .

tests:
	go test -v -count=1 ./...

clear:
	rm -rf ./build

build_dir:
	mkdir build

build_linux64: clear build_dir
	GOOS=linux GOARCH=amd64 go build -o ./build/hproxy ./cmd/hproxy

image_amd64: build_linux64
	docker build  --no-cache --platform=linux/amd64 --output=type=docker -t $(DOCKER_IMAGE):amd64 .
	docker save $(DOCKER_IMAGE):amd64 > ./build/hproxy_amd64.tar

image_upload: image_amd64
	docker push $(DOCKER_IMAGE):amd64

local:
	docker-compose -f docker-compose.yml up --force-recreate --renew-anon-volumes --build
