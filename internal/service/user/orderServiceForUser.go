package user

import (
	"context"

	"github.com/AlexandrBurak/TaxiAppOrder/internal/model"
	"github.com/AlexandrBurak/TaxiAppOrder/internal/repository"
	pb "github.com/AlexandrBurak/TaxiAppOrder/pkg/grpc"
)

type (
	RepositoryUser interface {
		SaveOrder(context.Context, model.Order) error
		GetOrdersByUser(ctx context.Context, phone string) ([]model.Order, error)
		RateLastTrip(ctx context.Context, phone string, rate int32) (string, error)
	}
)

type UserService struct {
	Repos  RepositoryUser
	Client pb.DriverServiceClient
}

func NewService(repository repository.Repository, client pb.DriverServiceClient) UserService {
	return UserService{Repos: &repository, Client: client}
}

func (s *UserService) GetDriver(ctx context.Context) (model.Driver, error) {
	drivers, err := s.Client.GetAllFreeDrivers(ctx, &pb.Request{})
	if err != nil {
		return model.Driver{}, nil
	}

	driversChan := make(chan *pb.Driver, 100)
	for _, r := range drivers.Drivers {
		driversChan <- r
	}

	return model.Driver{
		Phone:    drivers.Drivers[0].Phone,
		Email:    drivers.Drivers[0].Email,
		TaxiType: drivers.Drivers[0].TaxiType,
		Status:   drivers.Drivers[0].Status,
		Rating:   drivers.Drivers[0].Rating,
	}, nil
}

func (s *UserService) SaveOrder(ctx context.Context, order model.Order) error {
	err := s.Repos.SaveOrder(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetOrdersByUser(ctx context.Context, phone string) ([]*pb.Order, error) {
	orders, err := s.Repos.GetOrdersByUser(ctx, phone)
	if err != nil {
		return nil, err
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

func (s *UserService) RateLastTrip(ctx context.Context, phone string, rate int32) (string, error) {
	answer, err := s.Repos.RateLastTrip(ctx, phone, rate)
	if err != nil {
		return "", err
	}
	return answer, nil
}
