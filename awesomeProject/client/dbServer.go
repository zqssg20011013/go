package main

import (
	pb "awesomeProject/service"
	"context"
	_ "context"
	"fmt"
	_ "fmt"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/grpc/credentials/insecure"
	"log"
	_ "log"
)

// 初始化
func initconn(url string) pb.DBClient {
	conn, error := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if error != nil {
		log.Fatalf("connect wrong")
		defer conn.Close()
	}
	client := pb.NewDBClient(conn)
	return client
}
func show(id *uint32) *pb.UserResponse {
	client := initconn("127.0.0.1:9090")
	result, error := client.Show(context.Background(), &pb.UserRequest{Id: id})
	if error != nil {
		fmt.Print(error)
	}
	return result
}
func search(name string) *pb.UserResponse {
	client := initconn("127.0.0.1:9090")
	result, error := client.Search(context.Background(), &pb.UserRequest{Name: &name})
	if error != nil {
		fmt.Print(error)
	}
	return result
}
func like(name string) *pb.UserResponse {
	client := initconn("127.0.0.1:9090")
	result, error := client.Like(context.Background(), &pb.UserRequest{Name: &name})
	if error != nil {
		fmt.Print(error)
	}
	print(result.String())
	return result
}
func Translationsearch(name string) *pb.UserResponse {
	client := initconn("127.0.0.1:9090")
	result, error := client.Translationsearch(context.Background(), &pb.UserRequest{Name: &name})
	if error != nil {
		fmt.Print(error)
	}
	return result
}
func Translationsearchauto(name string) *pb.UserResponse {
	client := initconn("127.0.0.1:9090")
	result, error := client.Translationsearchauto(context.Background(), &pb.UserRequest{Name: &name})
	if error != nil {
		fmt.Print(error)
	}
	return result
}
func delete(Id *uint32) *pb.UserResponse {
	client := initconn("127.0.0.1:9090")
	result, error := client.Delete(context.Background(), &pb.UserRequest{Id: Id})
	if error != nil {
		fmt.Print(error)
	}
	return result
}
