package entities

var config *Config

type Config struct {
	Mode  string `mapstructure:"interval"`
	Agent *Agent `mapstructure:"agent"`
}

func GetConfig() *Config {
	return config
}

func SetConfig(c *Config) {
	config = c
}
