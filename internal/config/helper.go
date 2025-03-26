package config

import (
	"github.com/spf13/viper"
)

// System environment enum.
const (
	EnvProd = iota
	EnvUat
	EnvDev
	EnvLocal
	EnvCi
)

// System environment string tag.
const (
	EnvProdTag  = "prod"
	EnvUatTag   = "stage"
	EnvDevTag   = "dev"
	EnvLocalTag = "local"
	EnvCiTag    = "ci"
)

var envMap = map[string]int{
	EnvProdTag:  EnvProd,
	EnvUatTag:   EnvUat,
	EnvDevTag:   EnvDev,
	EnvLocalTag: EnvLocal,
	EnvCiTag:    EnvCi,
}

func IsValidEnv(env string) bool {
	_, exist := envMap[env]
	return exist
}

func GetEnv() string {
	return viper.GetString("env")
}

// Environment returns current system environment.
func Environment() int {
	environment := GetEnv()
	switch environment {
	case EnvProdTag:
		return EnvProd
	case EnvUatTag:
		return EnvUat
	case EnvDevTag:
		return EnvDev
	case EnvLocalTag:
		return EnvLocal
	case EnvCiTag:
		return EnvCi
	default:
		panic("Unknown environment.")
	}
}

// IsProduction returns true if we are in production mode.
func IsProduction() bool {
	return GetEnv() == EnvProdTag
}

func IsUat() bool {
	return GetEnv() == EnvUatTag
}

func IsDev() bool {
	return GetEnv() == EnvDevTag
}

func IsLocal() bool {
	return GetEnv() == EnvLocalTag
}

func IsCI() bool {
	return GetEnv() == EnvCiTag
}
