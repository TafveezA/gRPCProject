// package main

// import (
// 	"log"

// 	"google.golang.org/grpc"
// )

// const (
// 	port = ":8080"
// )

// func main() {
// 	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials)
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)

// 	}
// 	defer conn.Close()
// 	// client := pb.NewGreetServiceClient(conn)
// 	// names:= &pb.NameList{
// 	// 	Names: [] string {"tafveez","Afzal","Bob"}
// 	// }

// }
