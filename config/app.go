package config

type App struct {
	Port        int    `json:"port" yaml:"port" mapstructure:"port"`
	SignalrPort int    `json:"signalr_port" yaml:"signalr_port" mapstructure:"signalr_port"`
	Env         string `json:"env" yaml:"env" mapstructure:"env"`
}
