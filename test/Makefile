all:
	protoc --go_out=. --go-grpc_out=. test.proto

clean:
	rm test.pb.go example_grpc.pb.go

rebuild:
	clean
	all

