package driver

import (
	"context"

	"github.com/zxcghoulhunter/InnoTaxi-Order/internal/model"
	pb "github.com/zxcghoulhunter/InnoTaxi-Order/pkg/grpc"
)

type RepositoryDriver interface {
	GetOrders(context.Context, string) ([]model.Order, error)
	GetAllOrdersOfDriver(ctx context.Context, phone string) ([]model.Order, error)
	RateLastTripByDriver(ctx context.Context, phone string, rate int32) (string, error)
}

type DriverService struct {
	Repos RepositoryDriver
}

func NewDriverService(repos RepositoryDriver) *DriverService {
	return &DriverService{Repos: repos}
}

func (s *DriverService) GetAllOrdersOfDriver(ctx context.Context, phone string) ([]*pb.Order, error) {
	orders, err := s.Repos.GetAllOrdersOfDriver(ctx, phone)
	if err != nil {
		return nil, nil
	}
	var grpcOrders = make([]*pb.Order, 0)
	for _, order := range orders {
		grpcOrder := pb.Order{
			UserPhone:   order.UserPhone,
			DriverPhone: order.DriverPhone,
			From:        order.From,
			To:          order.To,
			TaxiType:    order.TaxiType,
			Date:        order.Date,
			Status:      order.Status,
			Rating:      order.Rating,
		}
		grpcOrders = append(grpcOrders, &grpcOrder)
	}
	return grpcOrders, nil
}

func (s *DriverService) RateLastTripByDriver(ctx context.Context, phone string, rate int32) (string, error) {
	answer, err := s.Repos.RateLastTripByDriver(ctx, phone, rate)
	if err != nil {
		return "", nil
	}
	return answer, nil
}
