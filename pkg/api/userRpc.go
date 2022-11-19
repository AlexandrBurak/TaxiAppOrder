package api

import (
	"context"
	"time"

	"github.com/zxcghoulhunter/InnoTaxi-Order/internal/model"
	"github.com/zxcghoulhunter/InnoTaxi-Order/internal/service/user"
	pb "github.com/zxcghoulhunter/InnoTaxi-Order/pkg/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	Service user.UserService
}

//TODO:тут структура с реквестом и грейсфул с каналом,реквест отправляется в отдельный сервис с беснонечным циклом и там реализуется worker pool
func (s *UserServer) OrderTaxi(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResult, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	driver, err := s.Service.GetDriver(ctxWithTimeout)
	if err != nil {
		return nil, err
	}
	order := model.Order{
		UserPhone:   in.User.Phone,
		DriverPhone: driver.Phone,
		From:        in.From,
		To:          in.To,
		TaxiType:    driver.TaxiType,
		Date:        time.Now().String(),
		Status:      "InProgress",
	}
	err = s.Service.SaveOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	orderResult := pb.Order{
		UserPhone:   in.User.Phone,
		DriverPhone: driver.Phone,
		From:        in.From,
		To:          in.To,
		TaxiType:    driver.TaxiType,
		Date:        time.Now().String(),
		Status:      "InProgress",
	}
	return &pb.OrderResult{Order: &orderResult}, nil
}
func (s *UserServer) GetOrdersByUser(ctx context.Context, in *pb.OrderRequest) (*pb.ListOfOrders, error) {
	orders, err := s.Service.GetOrdersByUser(ctx, in.User.Phone)
	if err != nil {
		return nil, err
	}
	return &pb.ListOfOrders{Orders: orders}, nil
}

func (s *UserServer) RateLastTrip(ctx context.Context, in *pb.RateTripByUserRequest) (*pb.RateResponse, error) {
	answer, err := s.Service.RateLastTrip(ctx, in.Phone, in.Rate)
	if err != nil {
		return nil, err
	}
	return &pb.RateResponse{Answer: answer}, nil
}
