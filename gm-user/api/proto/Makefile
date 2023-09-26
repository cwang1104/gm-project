.PHONY: gen-user

gen-user:
	protoc --go_out=./gm --go_opt=paths=source_relative --go-grpc_out=./gm --go-grpc_opt=paths=source_relative login_service.proto
