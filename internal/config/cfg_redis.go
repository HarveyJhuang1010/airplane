package config

type RedisConfig struct {
	Host     string `mapstructure:"host" validate:"ipv4"`
	Port     int    `mapstructure:"port" validate:"gte=1,lte=65535"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	CA       string `mapstructure:"ca"`
}
