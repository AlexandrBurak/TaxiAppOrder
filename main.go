package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/zxcghoulhunter/InnoTaxi-Order/internal/config"
	"github.com/zxcghoulhunter/InnoTaxi-Order/internal/repository"
	"github.com/zxcghoulhunter/InnoTaxi-Order/internal/service/driver"
	"github.com/zxcghoulhunter/InnoTaxi-Order/internal/service/user"
	"github.com/zxcghoulhunter/InnoTaxi-Order/pkg/api"
	pb "github.com/zxcghoulhunter/InnoTaxi-Order/pkg/grpc"
)

func main() {
	cfg, err := config.GetCfg()
	if err != nil {
		log.Fatalf("Error while loading config: %v", err)
	}
	repos, err := repository.NewRepository(cfg)
	if err != nil {
		log.Fatalf("Error while loading repository: %v", err)
	}

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)

	conn, err := grpc.Dial("localhost:4343", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(err.Error())
	}
	client := pb.NewDriverServiceClient(conn)
	userService := user.NewService(*repos, client)
	driverService := driver.NewDriverService(repos)
	go startGrpcServerForUser(userService, cfg)
	go startGrpcServerForDriver(driverService, cfg)
	<-wait
}

func startGrpcServerForUser(service user.UserService, cfg config.Config) {
	lis, err := net.Listen("tcp", cfg.RPC_PORT_USER)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &api.UserServer{Service: service})
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

}

func startGrpcServerForDriver(service *driver.DriverService, cfg config.Config) {
	lis, err := net.Listen("tcp", cfg.RPC_PORT_DRIVER)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterDriverServiceServer(server, &api.DriverServer{Service: *service})
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

}
