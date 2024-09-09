CONTRACT=./proto/jd.proto

all: jd-proto

jd-proto:
	protoc --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative $(CONTRACT)

clean:
	rm -rf ./gen/go/*