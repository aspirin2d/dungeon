package model

import (
	"context"
	"log"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	mdb *mongo.Database
	rdb *redis.Client

	logger *zap.SugaredLogger
)

func Initialize() {
	ctx := context.Background()

	// connect to mongodb
	mdb = mongoClient(ctx).Database("dungeon")
	// setup username for user collection, and make them unique
	if _, err := mdb.Collection("users").Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"username": 1}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		panic(err)
	}

	// connect to redis
	rdb = redisClient(ctx)

	// init logger
	l, _ := zap.NewProduction()
	defer l.Sync() // nolint: errcheck
	logger = l.Sugar()

	// load config files
	loadClassesConfig()
	loadRacesConfig()
}

func mongoClient(ctx context.Context) *mongo.Client {
	mc, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		panic(err)
	}

	err = mc.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	log.Println("connected to mongo at:", "mongodb://localhost:27017/")
	return mc
}

func redisClient(ctx context.Context) *redis.Client {
	opts, err := redis.ParseURL("redis://localhost:6379/")
	if err != nil {
		panic(err)
	}

	rc := redis.NewClient(opts)
	err = rc.Ping().Err()
	if err != nil {
		panic(err)
	}

	log.Println("connected to redis at:", "redis://localhost:6379/")
	return rc
}

func LogError(err interface{}) {
	logger.Error(err)
}

func Collection(name string) *mongo.Collection {
	return mdb.Collection(name)
}
