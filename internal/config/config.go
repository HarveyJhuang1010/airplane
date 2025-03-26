package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var self Config

// XXX: due to main function not in the root/cmd/ constantly, this might cause issue.
func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return path.Join(filepath.Dir(d), "..")
}

var ConfigFileName string

func newConfig() *Config {
	if ConfigFileName == "" {
		ConfigFileName = "config"
	}

	// For Read config file
	cfgDir := path.Join(rootDir(), "config")
	fmt.Println("config dir", cfgDir)
	viper.AddConfigPath(cfgDir)
	viper.AddConfigPath("/config")
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("failed to read config", err)
	}

	// For Read environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Watch config file
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.BindEnv("database.password", "DATABASE_PASSWORD")
	viper.BindEnv("redis.ca", "REDIS_CA")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("rest.port", "REST_PORT")

	if err := viper.Unmarshal(&self); err != nil {
		panic(errors.Wrap(err, "failed to marshal config"))
	}

	return &self
}

func NewConfigWithoutBindEnv() *Config {
	if err := LoadConfig("config/common.yaml"); err != nil {
		return nil
	}
	return &self
}

func LoadConfig(file string) error {
	path, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.Unmarshal(&self)

	return nil
}

type Config struct {
	Env      string          `mapstructure:"env"`
	Version  string          `mapstructure:"ver"`
	Database *DatabaseConfig `mapstructure:"database"`
	Logger   *LoggerConfig   `mapstructure:"logger"`
	Redis    *RedisConfig    `mapstructure:"redis"`
	Rest     *RestConfig     `mapstructure:"rest"`
	Kafka    *KafkaConfig    `mapstructure:"kafka"`
}

func (c *Config) GetServerName() string {
	return viper.GetString("serverName")
}
