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
}

func (c Config) SetDefaults() {

}

func (c Config) Validate() error {
	if c.NodeURL == "" {
		return fmt.Errorf("nodeUrl is required")
	}

	return nil
}
