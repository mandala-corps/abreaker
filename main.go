package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mandala-corps/abreaker/cmd"
	"github.com/mandala-corps/abreaker/internal/dto"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()

	config := dto.GetConfig()
	if config.Mode == "" {
		panic("please set a mode: Server or Agent")
	}

	switch strings.ToLower(config.Mode) {
	case "agent":
		cmd.AgentExecute(ctx, config)
	default:
		panic("mode not implemented")
	}

}

func init() {
	// config file settings
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error when readin config file: %v", err))
	}

	c := &dto.Config{}
	// unmarshal config in entite struct
	err = viper.UnmarshalExact(c)
	if err != nil {
		panic(fmt.Errorf("cannot unmarshal configs: %v", err))
	}
	// save "global" config
	dto.SetConfig(c)
}
