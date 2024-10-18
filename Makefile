clean:
	rm -f /c/Code/grpc_chat/pb/*.pb.*
run:
	go run main.go
gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/computer_message.proto