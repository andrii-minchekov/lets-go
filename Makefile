all: clean generateProto

generateProto:
	protoc -I . app/impl/grpc/snippet.proto --go_out=plugins=grpc:.

clean:
	find . -name "*.pb.go" -exec rm -rf {} \;
