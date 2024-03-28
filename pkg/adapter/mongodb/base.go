package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBCollection struct {
	*mongo.Collection
}
