.PHONY: gen
all:
	make gen
	make build
gen:
	cd proto/api/v1 && protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative worker.proto
build:
	go build -o lc500 main.go
