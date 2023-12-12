package config

type Eos struct {
	Region       string `yaml:"region" mapstructure:"region"`
	Endpoint     string `yaml:"endpoint" mapstructure:"endpoint"`
	BuketName    string `yaml:"buket_name" mapstructure:"buket_name"`
	AccessId     string `yaml:"access_id" mapstructure:"access_id"`
	AccessSecret string `yaml:"access_secret" mapstructure:"access_secret"`
}
