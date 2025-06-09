package main

import (
	"fmt"
)

type Config struct {
	// NetworkName is the network name
	NetworkName string `mapstructure:"network_name"`
	// NodeURL is the url of the blockchain node's rpc server
	NodeURL string `mapstructure:"node_url"`
	// NetworkSymbol is the token symbol for this network
	NetworkSymbol string `mapstructure:"network_name"`
	NodeToken     string `mapstructure:"node_token"`
	S3Url         string `mapstructure:"s3_url"`
	S3Ssl         bool   `mapstructure:"s3_ssl"`
	S3AccessKey   string `mapstructure:"s3_access_key"`
	S3SecretKey   string `mapstructure:"s3_secret_key"`
	S3Bucket      string `mapstructure:"s3_bucket"`
	S3Service     string `mapstructure:"s3_service"`
	S3RawDataPath string `mapstructure:"s3_raw_data_path"`

	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     int    `mapstructure:"db_port"`
	DBSchema   string `mapstructure:"db_schema"`
	DBParams   string `mapstructure:"db_params"`

	RedisAddr     string `mapstructure:"redis_addr"`
	RedisPassword string `mapstructure:"redis_password"`
}

func (c Config) SetDefaults() {

}

func (c Config) Validate() error {
	if c.NodeURL == "" {
		return fmt.Errorf("nodeUrl is required")
	}

	return nil
}
