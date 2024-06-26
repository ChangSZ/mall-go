package mongodb

import (
	"context"
	"fmt"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/pkg/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var _ Repo = (*dbRepo)(nil)

type Repo interface {
	i()

	Connection() *mongo.Database
	Disconnect() error
}

type dbRepo struct {
	connection *mongo.Database
	client     *mongo.Client
}

var db *dbRepo

func (d *dbRepo) i() {}

func Init() {
	cfg := configs.Get().MongoDB
	uri := fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	connection := client.Database(cfg.Database)
	db = &dbRepo{client: client, connection: connection}
	log.Info("Connected to mongo db")
}

func DB() *dbRepo {
	if db == nil {
		Init()
	}
	return db
}

func (d *dbRepo) Connection() *mongo.Database {
	return d.connection
}

func (d *dbRepo) Disconnect() error {
	return db.client.Disconnect(context.Background())
}
