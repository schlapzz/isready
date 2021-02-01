package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Timeout  time.Duration
}

func openMongo(ctx context.Context, c MongoConnection) error {

	connString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?connectTimeoutMS=%d", c.Username, c.Password, c.Host, c.Port, c.Database, c.Timeout.Milliseconds())

	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(ctx)
	return client.Connect(ctx)

}
