package api

import (
	"context"

	"github.com/zxcghoulhunter/InnoTaxi-Order/internal/service/driver"
	pb "github.com/zxcghoulhunter/InnoTaxi-Order/pkg/grpc"
)

type DriverServer struct {
	Service driver.DriverService
}

func (s *DriverServer) RateTripByDriver(ctx context.Context, request *pb.RateTripByDriverRequest) (*pb.Error, error) {
	answer, err := s.Service.RateLastTripByDriver(ctx, request.Phone, request.Rate)
	if err != nil {
		return nil, err
	}
	return &pb.Error{Message: answer}, nil
}

func (s *DriverServer) GetAllOrdersOfDriver(ctx context.Context, phone *pb.DriversPhone) (*pb.AllOrders, error) {
	orders, err := s.Service.GetAllOrdersOfDriver(ctx, phone.Phone)
	if err != nil {
		return nil, err
	}
	return &pb.AllOrders{Orders: orders}, nil
}
