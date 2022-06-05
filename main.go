package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mandala-corps/abreaker/cmd"
	"github.com/mandala-corps/abreaker/internal/dependency"
	"github.com/mandala-corps/abreaker/internal/dto"
	"github.com/mandala-corps/abreaker/internal/service"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()
	// inject dependencies
	ctx = registerDependecies(ctx)
	// get config
	config := dto.GetConfig()
	if config.Mode == "" {
		panic("please set a mode: Server or Agent")
	}

	switch strings.ToLower(config.Mode) {
	case "agent":
		cmd.AgentExecute(ctx, config)
	case "server":
		cmd.ExecuteServer(ctx, config)
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

func registerDependecies(ctx context.Context) context.Context {
	d := make(map[dependency.Key]interface{})

	d[dependency.ConfigKey] = dto.GetConfig()
	d[dependency.ServerServiceKey] = service.NewServerService()

	for k, v := range d {
		ctx = context.WithValue(ctx, k, v)
	}

	return ctx
}
