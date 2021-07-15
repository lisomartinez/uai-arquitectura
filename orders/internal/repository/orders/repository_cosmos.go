package orders

import (
	"context"
	"fmt"
	"orders-service/internal/domain/model"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CONNECTIONSTRING = os.Getenv("AZ_COSMOS_ORDER_CONNECTION_STRING")
)

const (
	DB     = "orders"
	ORDERS = "orders"
)

type Repository interface {
	Save(ctx context.Context, order *model.Order) (*model.Order, error)
	GetOrder(ctx context.Context, orderId uint64) (*model.Order, error)
}

type cosmosDbRepository struct {
}

func NewCosmosDbRepository() Repository {
	return &cosmosDbRepository{}
}

func (c cosmosDbRepository) Save(ctx context.Context, order *model.Order) (*model.Order, error) {
	client, err := getMongoClient(ctx)
	if err != nil {
		return nil, err
	}

	collection := client.Database(DB).Collection(ORDERS)

	_, err = collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c cosmosDbRepository) GetOrder(ctx context.Context, orderId uint64) (*model.Order, error) {
	result := &model.Order{}

	filter := bson.D{primitive.E{Key: "order_id", Value: orderId}}

	client, err := getMongoClient(ctx)
	if err != nil {
		return nil, err
	}

	collection := client.Database(DB).Collection(ORDERS)

	err = collection.FindOne(context.TODO(), filter).Decode(result)

	if err != nil {

		return nil, err
	}

	return result, nil
}

func getMongoClient(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(CONNECTIONSTRING).SetDirect(true)
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize connection %w", err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to connect %w", err)
	}

	return client, err

}
