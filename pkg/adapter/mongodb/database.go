package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseI interface {
	Close() error
	GetConn() MongoDBConnI
	GetCollection(collName string, opts *CollectionOptions) *MongoDBCollection
}

type mongodbDatabase struct {
	conn     MongoDBConnI
	database *mongo.Database
}

var _ DatabaseI = (*mongodbDatabase)(nil)

func NewDatabase(dbConfig MongoDBConfig, opts *DatabaseOptions) (*mongodbDatabase, error) {
	conn := NewMongoDBConn(dbConfig)

	var dbopts = NewDatabaseOptions()
	if opts != nil {
		dbopts = opts
	}

	mongodb := conn.GetDatabase(dbConfig.DbName, dbopts.rawMongoOptions())

	return &mongodbDatabase{
		conn:     conn,
		database: mongodb,
	}, nil
}

func (m *mongodbDatabase) Close() error {
	return m.conn.Close()
}

func (m *mongodbDatabase) GetConn() MongoDBConnI {
	return m.conn
}

func (m *mongodbDatabase) GetRawMongoDatabase() *mongo.Database {
	return m.database
}

func (m *mongodbDatabase) GetCollection(collName string, opts *CollectionOptions) *MongoDBCollection {

	var collopts = NewCollectionOptions()
	if opts != nil {
		collopts = opts
	}

	collection := m.database.Collection(collName, collopts.rawCollectionOptions())

	return &MongoDBCollection{
		collection,
	}
}
