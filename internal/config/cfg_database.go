package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"strings"
)

type DatabaseConfig struct {
	Dialect      string              `mapstructure:"dialect"`
	User         string              `mapstructure:"user"`
	Password     string              `mapstructure:"password"`
	Host         string              `mapstructure:"host" validate:"ipv4"`
	Port         int                 `mapstructure:"port" validate:"gte=1,lte=65535"`
	DBName       string              `mapstructure:"dbName"`
	SSLMode      string              `mapstructure:"sslMode"`
	MaxOpenConns int                 `mapstructure:"maxOpenConns"`
	MaxIdleConns int                 `mapstructure:"maxIdleConns"`
	MaxLifetime  int                 `mapstructure:"maxLifeTime"`
	LogLevel     gormlogger.LogLevel `mapstructure:"logLevel"`
}

func (cfg *DatabaseConfig) IsMySQL() bool {
	return strings.EqualFold(cfg.Dialect, "mysql")
}

func (cfg *DatabaseConfig) IsPostgreSQL() bool {
	return strings.EqualFold(cfg.Dialect, "postgres") ||
		strings.EqualFold(cfg.Dialect, "postgresql")
}

func (cfg *DatabaseConfig) Open() gorm.Dialector {
	if cfg.IsMySQL() {
		return mysql.Open(cfg.DSN())
	}

	if cfg.IsPostgreSQL() {
		return postgres.New(postgres.Config{
			DSN:                  cfg.DSN(),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		})
	}

	panic("no match dialect")
}

func (cfg *DatabaseConfig) DSN() string {
	if cfg.IsMySQL() {
		//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&time_zone=UTC",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		)
	}

	if cfg.IsPostgreSQL() {
		//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		return fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			cfg.Host,
			cfg.User,
			cfg.Password,
			cfg.DBName,
			cfg.Port,
			cfg.SSLMode,
		)
	}

	return ""
}
