file = test.proto

all: $(TARGET)


$(TARGET): $(OBJS)
	protoc --go_out=. --go-grpc_out=. example.proto

