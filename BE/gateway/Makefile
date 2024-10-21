.PHONY: proto
## proto: compiles .proto files
proto:
	@ docker run -v $(PWD):/defs namely/protoc-all -f proto/*.proto -l go -o . --go-source-relative

.PHONY: build
## build: builds server's binary
build:
	@ go build -a -installsuffix cgo -o main .

.PHONY: run
## run: runs the server
run: build
	@ ./main

.PHONY: build-docker-image
## build-docker-image: builds the docker image
build-docker-image:
	@ docker buildx build . -t image-retrieval-gateway

.PHONY: run-docker
## run-docker: runs the server as a Docker container
run-docker: 
	@ docker run -p 8000:8000 image-retrieval-gateway