package main

import (
	"context"
	pojo "goProject/pojo"
	pb "goProject/service"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"net"
)

const (
	dsn = "root:zqayy20011013@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
)

type service struct {
	pb.UnimplementedDBServer
}

func (s *service) Show(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var user pojo.User
	if result := db.First(&user, req.Id); result.Error != nil {
		return nil, result.Error
	}
	var response = pb.UserResponse{Id: &user.Id, Name: &user.Name, Age: &user.Age, Email: &user.Email, Password: &user.PassWord}
	return &response, nil
}
func (s *service) delete(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("id=?", req.Id).Delete(&pojo.User{}); result.Error != nil {
			return result.Error
		}
		return nil
	})
	return nil, nil
}
func (s *service) search(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var user pojo.User
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("name=?", req.Name).Delete(&pojo.User{}); result.Error != nil {
			return result.Error
		}
		return nil
	})
	var response = pb.UserResponse{Id: &user.Id, Name: &user.Name, Age: &user.Age, Email: &user.Email, Password: &user.PassWord}
	return &response, nil
}
func (s *service) translationsearchauto(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var user pojo.User
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("name=?", req.Name).Delete(&pojo.User{}); result.Error != nil {
			return result.Error
		}
		return nil
	})
	var response = pb.UserResponse{Id: &user.Id, Name: &user.Name, Age: &user.Age, Email: &user.Email, Password: &user.PassWord}
	return &response, nil
}
func (s *service) translationsearch(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var user pojo.User
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("name=?", req.Name).Delete(&pojo.User{}); result.Error != nil {
			return result.Error
		}
		return nil
	})
	var response = pb.UserResponse{Id: &user.Id, Name: &user.Name, Age: &user.Age, Email: &user.Email, Password: &user.PassWord}
	return &response, nil
}
func (s *service) Like(ctx context.Context, req *pb.UserRequest) (*pb.Responses, error) {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var user []pojo.User
	db.Transaction(func(tx *gorm.DB) error {
		db.Raw("select * from users where match(name)Against (?)", req.Name).Scan(&user)
		return nil
	})
	var list []*pb.UserResponse
	for _, p := range user {
		t := pb.UserResponse{Id: &p.Id, Name: &p.Name, Age: &p.Age}
		list = append(list, &t)
	}
	response := pb.Responses{Response: list}

	return &response, nil
}
func main() {
	listen, _ := net.Listen("tcp", ":9090")
	grpcServer := grpc.NewServer()
	pb.RegisterDBServer(grpcServer, &service{})
	grpcServer.Serve(listen)

}
