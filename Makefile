GO_MAIN := src/cmd/main.go
gorun:
	@go run ${GO_MAIN}

gowrun:
	@gow -e=go,mod ${GO_MAIN}