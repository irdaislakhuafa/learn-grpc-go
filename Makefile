GO_MAIN := src/cmd/main.go
ENT_SCHEMA_PATH := ./src/schema/psqlentity/schema

gorun:
	@go run ${GO_MAIN}

gowrun:
	@gow -e=go,mod ${GO_MAIN}

entgen:
	@ent generate ${ENT_SCHEMA_PATH} --target ${ENT_SCHEMA_PATH}/generated