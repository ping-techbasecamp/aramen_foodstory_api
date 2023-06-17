package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func (cfg *AppConfig) Init() {
	envconfig.MustProcess("", &cfg.Env)
	envconfig.MustProcess("", &cfg.Database)
	envconfig.MustProcess("", &cfg.JWT)
	envconfig.MustProcess("", &cfg.Mailbit)
	envconfig.MustProcess("", &cfg.Foodstory)
}

func NewAppConfig() *AppConfig {
	godotenv.Load()
	appCfg := AppConfig{}
	appCfg.Init()
	return &appCfg
}
