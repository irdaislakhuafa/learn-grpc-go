GO_MAIN := src/cmd/main.go
ENT_SCHEMA_PATH := ./src/schema/psqlentity/schema
PROTO_PATH = src/schema/protobuf

gorun:
	@go run ${GO_MAIN}

gowrun:
	@gow -e=go,mod run ${GO_MAIN}

entgen:
	@ent generate ${ENT_SCHEMA_PATH} --target ${ENT_SCHEMA_PATH}/generated

procom:
	@protoc \
		--proto_path="${PROTO_PATH}" \
		--go_out="${PROTO_PATH}" \
		--go-grpc_out="${PROTO_PATH}" \
		"${PROTO_PATH}"/*.proto