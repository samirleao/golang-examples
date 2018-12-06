package mongodb

import (
	"context"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type contextKey string

// Client the mongoDB client used to run queries
var Client *mongo.Client

func init() {
	var err error
	Client, err = connect()
	if err != nil {
		fmt.Printf("Couldn't connect to mongo: %v", err)
	}
}

const (
	host     = contextKey("host")
	username = contextKey("username")
	password = contextKey("password")
	database = contextKey("database")
)

func connect() (*mongo.Client, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx = context.WithValue(ctx, host, os.Getenv("HOST_ENV"))
	ctx = context.WithValue(ctx, username, os.Getenv("USERNAME_ENV"))
	ctx = context.WithValue(ctx, password, os.Getenv("PASSWORD_ENV"))
	ctx = context.WithValue(ctx, database, os.Getenv("DATABASE_ENV"))
	return config(ctx)
}

func config(ctx context.Context) (*mongo.Client, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value(username),
		ctx.Value(password),
		ctx.Value(host),
		ctx.Value(database),
	)

	client, err := mongo.NewClient(uri)
	if err != nil {
		fmt.Printf("todo: couldn't connect to mongo: %v", err)
		return client, err
	}

	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Couldn't connect ot mongo: %v", err)
		return client, err
	}

	return client, nil
}
