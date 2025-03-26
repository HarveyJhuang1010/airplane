package config

type RestConfig struct {
	ListenAddress string   `mapstructure:"listenAddress"`
	Port          int      `mapstructure:"port" validate:"gte=1,lte=65535"`
	Trace         bool     `mapstructure:"trace"`
	AllowOrigins  []string `mapstructure:"allowOrigins"`
	AllowHeaders  []string `mapstructure:"allowHeaders"`
	ExposeHeaders []string `mapstructure:"exposeHeaders"`
	AllowMethods  []string `mapstructure:"allowMethods"`
}
