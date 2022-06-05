package dto

type Watcher struct {
	Addr   string `mapstructure:"addr"`
	Method string `mapstructure:"method"`
}
