LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make: generate-note-api

generate-note-api:
	mkdir "pkg/note_v1"
	protoc --proto_path "api/note_v1" \
	--go_out="pkg/note_v1" --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=protoc-gen-go \
	--go-grpc_out="pkg/note_v1" --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=protoc-gen-go-grpc \
	"api/note_v1/note.proto"

generate-note-api2:
	protoc --proto_path "api/note_v1" --go_out="pkg/note_v1" --go_opt=paths=source_relative --go-grpc_out="pkg/note_v1" --go-grpc_opt=paths=source_relative "api/note_v1/note.proto"	