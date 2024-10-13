gen:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/*.proto
	@python -m grpc_tools.protoc -I api/ --python_out=api/ --pyi_out=api/ --grpc_python_out=api/ api/*.proto