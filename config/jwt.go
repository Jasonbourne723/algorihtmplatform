package config

type Jwt struct {
	Secret           string `yaml:"secret" mapstructure:"secret"`
	Issuer           string `yaml:"issuer" mapstructure:"issuer"`
	Audience         string `yaml:"audience" mapstructure:"audience"`
	ExpireMinutes    int    `yaml:"expire_minutes" mapstructure:"expire_minutes"`
	ClockSkewMinutes int    `yaml:"clock_skew_minutes" mapstructure:"clock_skew_minutes"`
}
