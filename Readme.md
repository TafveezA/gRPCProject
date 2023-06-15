 # To initialize
 go mod init github.com/TafveezA/gRPCProject


 # To update go.sum and go.mod
 protoc --go_out=. --go-grpc_out=. proto/greet.proto
