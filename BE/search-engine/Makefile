.PHONY: proto
## proto: compiles .proto files
proto:
	@ docker run -v $(PWD):/defs namely/protoc-all -l go -o proto --go-source-relative -d proto

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
	@ docker buildx build . -t image-retrieval-search-engine

.PHONY: run-docker
## run-docker: runs the server as a Docker container
run-docker: 
	@ docker run -p 8002:8002 image-retrieval-search-engine