package dto

type Agent struct {
	// Server will recive request
	Server *Server `mapstructure:"server"`
	// How agent will request
	Watchers map[string]*Watcher `mapstructure:"watchers"`
	// Interval how watch sleep to next request in seconds
	Interval int `mapstructure:"interval"`
}
