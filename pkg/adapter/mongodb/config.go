package mongodb

import (
	"encoding/json"
	"time"
)

type MongoDBConfig struct {
	// config name
	Name string `mapstructure:"name" yaml:"name"`

	// mongodb config

	Hosts       []string      `mapstructure:"hosts" yaml:"hosts"`
	RSName      *string       `mapstructure:"rs_name" yaml:"rs_name"`
	AuthSource  string        `mapstructure:"auth_source" yaml:"auth_source"`
	UserName    string        `mapstructure:"user_name" yaml:"user_name"`
	Password    string        `mapstructure:"password" yaml:"password"`
	Timeout     time.Duration `mapstructure:"timeout" yaml:"timeout"`
	IsSSLEnable bool          `mapstructure:"is_ssl_enable" yaml:"is_ssl_enable"`
	PoolLimit   *uint64       `mapstructure:"pool_limit" yaml:"pool_limit"`
	ReadPref    string        `mapstructure:"read_pref" yaml:"read_pref"`
	DbName      string        `mapstructure:"db_name" yaml:"db_name"`
}

func (m MongoDBConfig) PrettyPrint() string {
	bytes, _ := json.MarshalIndent(m, "", " ")
	return string(bytes)
}
