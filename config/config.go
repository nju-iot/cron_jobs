package config

import (
	"log"
	"path/filepath"

	"github.com/go-ini/ini"
)

const (
	appINIFilePath = "config/app.ini"
)

var (
	DBConf    *Database
	RedisConf *RedisConfig
	LogConf   *LogConfig
)

type LogConfig struct {
	LogLevel   string
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}
type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

type Database struct {
	DriverName string
	User       string
	Password   string
	DBHostname string
	DBPort     string
	DBName     string
}

func LoadConfig() {
	absPath, err := filepath.Abs(appINIFilePath)
	if err != nil {
		panic(err)
	}
	cfg, err := ini.Load(absPath)
	if err != nil {
		panic(err)
	}
	LogConf = new(LogConfig)
	DBConf = new(Database)
	RedisConf = new(RedisConfig)
	mapTo("Log", LogConf, cfg)
	mapTo("Database", DBConf, cfg)
	mapTo("Redis", RedisConf, cfg)
}

func mapTo(section string, v interface{}, cfg *ini.File) {
	if cfg == nil || section == "" {
		log.Fatalf("section=%v, iniFile=%v", section, cfg)
		return
	}
	if err := cfg.Section(section).MapTo(v); err != nil {
		panic(err)
	}
}
