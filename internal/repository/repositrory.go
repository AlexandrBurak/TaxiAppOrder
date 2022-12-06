package repository

import (
	"context"

	"github.com/AlexandrBurak/TaxiAppOrder/internal/config"
	"github.com/AlexandrBurak/TaxiAppOrder/internal/model"
)

type Repository struct {
	db string
}

func (r Repository) GetOrders(ctx context.Context, s string) ([]model.Order, error) {
	return nil, nil
}

func (r Repository) RateLastTrip(ctx context.Context, phone string, rate int32) (string, error) {
	return "asdasdzzzz", nil
}

func (r Repository) GetOrdersByUser(ctx context.Context, phone string) ([]model.Order, error) {
	return nil, nil
}

func (r Repository) SaveOrder(ctx context.Context, order model.Order) error {
	return nil
}
func (r Repository) GetAllOrdersOfDriver(ctx context.Context, phone string) ([]model.Order, error) {
	return nil, nil
}
func (r Repository) RateLastTripByDriver(ctx context.Context, phone string, rate int32) (string, error) {
	return "", nil
}
func NewRepository(cfg config.Config) (*Repository, error) {
	//dbCfg := elasticsearch.Config{
	//	Addresses: []string{"172.17.0.1:9200"},
	//}
	//es, err := elasticsearch.NewClient(dbCfg)
	//if err != nil {
	//	log.Fatalf("Error creating the client: %s", err)
	//}
	//res, err := es.Info()
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//if res.IsError() {
	//	log.Fatalf("Error: %s", res.String())
	//}
	return &Repository{db: "Asd"}, nil

}
