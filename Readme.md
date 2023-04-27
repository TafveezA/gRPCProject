 # To remove error
 go mod init github.com/akhil/grpc-demo-yt
 # To initialize 
 go mod tidy

 # To update go.sum and go.mod
 protoc --go_out=. --go-grpc_out=. proto/greet.proto