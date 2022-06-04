package entities

var config *Config

type Config struct {
	// Server will recive request
	Server *Server `mapstructure:"server"`
	// How agent will request
	Watchers []*Watcher `mapstructure:"Watchers"`
	// Interval how watch sleep to next request in seconds
	Interval int `mapstructure:"interval"`
}

func GetConfig() *Config {
	return config
}

func SetConfig(c *Config) {
	config = c
}
