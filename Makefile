all: clean generateProto generateOrm generateMocks

generateProto:
	protoc -I . app/impl/grpc/snippet.proto --go_out=plugins=grpc:.

generateOrm:
	sqlboiler --config sqlboiler.toml --output app/impl/db/models --wipe --debug psql

generateMocks:
	mockery --all

clean:
	find . -name "*.pb.go" -exec rm -rf {} \;
