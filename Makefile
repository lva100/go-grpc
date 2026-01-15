# winget install GnuWin32.Make in PowerShell
include .env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD)"

show-curdir: 
	@echo "CURDIR is:" ${CURDIR};

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

install-deps-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

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
	protoc --proto_path "api/note_v1" --go_out="pkg/note_v1" --go_opt=paths=source_relative --go-grpc_out="pkg/note_v1" \ --go-grpc_opt=paths=source_relative "api/note_v1/note.proto"	

local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v