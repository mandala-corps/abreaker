package dto

var config *Config

type Config struct {
	Mode   string  `mapstructure:"mode"`
	Agent  *Agent  `mapstructure:"agent"`
	Server *Server `mapstructure:"server"`
}

func GetConfig() *Config {
	return config
}

func SetConfig(c *Config) {
	config = c
}
