package mongodb

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBConnI interface {
	GetDatabase(dbName string, opts *options.DatabaseOptions) *mongo.Database
	GetRawConn() *mongo.Client
	Close() error
}

// mongoDBConn is a singleton struct responsible for managing MongoDB connections.
type mongoDBConn struct {
	client *mongo.Client
}

var (
	onceMutex               sync.RWMutex
	onceByConfigName        map[string]*sync.Once
	onceSessionByConfigName map[string]*mongoDBConn
)

func init() {
	onceMutex = sync.RWMutex{}
	onceByConfigName = make(map[string]*sync.Once)
	onceSessionByConfigName = make(map[string]*mongoDBConn)
}

func mapReadPrefFromConfig(readPreference string) (readPrefMode *readpref.ReadPref) {
	mode, err := readpref.ModeFromString(readPreference)
	if err != nil {
		mode = readpref.PrimaryMode
	}

	readPrefMode, _ = readpref.New(mode)
	return
}

var _ MongoDBConnI = (*mongoDBConn)(nil)

func NewMongoDBConn(config MongoDBConfig) *mongoDBConn {

	if _, ok := onceSessionByConfigName[config.Name]; ok {
		return onceSessionByConfigName[config.Name]
	}

	onceMutex.Lock()
	defer onceMutex.Unlock()

	var adapter = new(mongoDBConn)

	if onceByConfigName[config.Name] == nil {
		onceByConfigName[config.Name] = new(sync.Once)
	}

	onceByConfigName[config.Name].Do(func() {

		log.Printf("[%s][%v] MongoDB [Connecting]\n", config.Name, config.Hosts)

		clientOptions := &options.ClientOptions{
			Hosts:                  config.Hosts,
			MaxPoolSize:            config.PoolLimit,
			ReplicaSet:             config.RSName,
			SocketTimeout:          &config.Timeout,
			ServerSelectionTimeout: &config.Timeout,
		}

		clientOptions.SetReadPreference(mapReadPrefFromConfig(config.ReadPref))

		if config.UserName != "" && config.Password != "" &&
			config.AuthSource != "" {
			clientOptions.SetAuth(options.Credential{
				AuthSource: config.AuthSource,
				Username:   config.UserName,
				Password:   config.Password,
			})
		}

		if config.IsSSLEnable {
			tlsConfig := new(tls.Config)
			clientOptions.SetTLSConfig(tlsConfig)
		}

		ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
		defer cancel()

		err := adapter.connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Cannot establish connection to mongodb %s", err.Error())
		}

		onceSessionByConfigName[config.Name] = adapter
		log.Printf("[%s][%v] MongoDB [Connected]\n", config.Name, config.Hosts)
	})

	return onceSessionByConfigName[config.Name]
}

func (m *mongoDBConn) connect(ctx context.Context, clientOptions *options.ClientOptions) error {
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("mongo.Connect %s", err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("mongo.Connect.Ping %s", err.Error())
	}

	m.client = client

	return nil
}

func (m *mongoDBConn) Close() error {
	if m.client != nil {
		err := m.client.Disconnect(context.Background())
		if err != nil {
			return fmt.Errorf("client.Disconnect %s", err)
		}
	}

	return nil
}

func (m *mongoDBConn) GetDatabase(dbName string, opts *options.DatabaseOptions) *mongo.Database {
	return m.client.Database(dbName, opts)
}

func (m *mongoDBConn) GetRawConn() *mongo.Client {
	return m.client
}
