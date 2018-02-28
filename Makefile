build:
	go build

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

build-docker: build-linux
	docker build -t fresh8/goreplay-installer .

push-docker: build-docker
	docker push fresh8/goreplay-installer
