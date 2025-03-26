package config

type LoggerConfig struct {
	SysLogger string            `mapstructure:"sysLogger"`
	AppLogger string            `mapstructure:"appLogger"`
	Named     map[string]string `mapstructure:"named"`
}
