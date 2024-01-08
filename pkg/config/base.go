package config

import (
	"log"

	postgresDB "github.com/duongnln96/blog-realworld/pkg/adapter/postgres"
	scyllaDB "github.com/duongnln96/blog-realworld/pkg/adapter/scylladb"
	"github.com/spf13/viper"
)

type otherConfig map[string]interface{}

func (m otherConfig) Get(key string) interface{} {
	return m[key]
}

type Configs struct {
	Other             otherConfig
	PostgresConfigMap postgresDB.PosgreSQLDBConfigMap
	ScyllaDBConfigMap scyllaDB.ScyllaDBConfigMap
}

func readAllConfig(configPath string) *Configs {
	log.Println("Load config from path ", configPath)

	var onceConfigs = new(Configs)

	onceConfigs.Other = make(map[string]interface{})
	readConfig(configPath, "other", &onceConfigs.Other)

	readConfig(configPath, "postgres", &(onceConfigs.PostgresConfigMap))

	readConfig(configPath, "scylladb", &(onceConfigs.ScyllaDBConfigMap))

	return onceConfigs
}

func readConfig(configPath, configName string, result interface{}) error {

	viperI := viper.New()

	viperI.AutomaticEnv()
	viperI.AddConfigPath(configPath)
	viperI.SetConfigType("yaml")
	viperI.SetConfigName(configName)

	err := viperI.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = nil
		} else {
			// Config file was found but another error was produced
			log.Panicf("[ERROR] reading %s.yaml %s", configName, err.Error())
		}
	}

	err = viperI.Unmarshal(&result)
	if err != nil {
		log.Printf("viperI.Unmarshal %s %s", configName, err.Error())
		return err
	}

	return nil
}

var onceConfigs *Configs

func LoadConfig(configPath string) *Configs {
	if onceConfigs != nil {
		return onceConfigs
	}

	onceConfigs = readAllConfig(configPath)
	return onceConfigs
}

func GetConfig() *Configs {
	if onceConfigs == nil {
		log.Fatal("Config still not load")
	}

	return onceConfigs
}
