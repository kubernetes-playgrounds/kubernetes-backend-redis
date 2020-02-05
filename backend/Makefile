# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=backend_binary
BINARY_UNIX=$(BINARY_NAME)_unix

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker build -t guanw0826/kubernetes-backend-redis .

docker-push:
	docker push guanw0826/kubernetes-backend-redis

docker-run:
	docker run -d guanw0826/kubernetes-backend-redis