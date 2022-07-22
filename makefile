
gen:
	protoc --go_out=. --go-grpc_out=. api.proto

run:
	go run main.go