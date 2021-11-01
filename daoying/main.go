package main

import (
	"api/user/config"
	"api/user/global"
	"api/user/handler"
	"api/user/middleware"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	//"api/user/middleware"
	"api/user/pb"
	"api/user/repo"
	//grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":8001"
)

func init() {
	config.InitConfig("user")
	global.InitDB()
	global.GetMyDB().AutoMigrate(repo.LightChangeLog{}, repo.User{}, repo.UserAuth{})
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//grpc.UnaryInterceptor(),
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(middleware.AuthFunc),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(middleware.Recoverf)),
		)),
	)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	pb.RegisterUserServer(s, handler.UserHandler{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
