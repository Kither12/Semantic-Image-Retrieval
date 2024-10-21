.PHONY: proto
## proto: compiles .proto files
proto:
	@ docker run -v $(PWD):/defs namely/protoc-all -l python -o proto -d proto
