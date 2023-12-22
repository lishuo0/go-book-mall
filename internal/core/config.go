package core

import "time"

var GlobalConfig MallConfig

type MallConfig struct {
	Server ServerConfig  `mapstructure:"server"`
	Mysql  []MysqlConfig `mapstructure:"mysql"`
	Logger LoggerConfig  `mapstructure:"logger"`
	Redis  RedisConfig   `mapstructure:"redis"`
}
type ServerConfig struct {
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	IdleTimeout  time.Duration `mapstructure:"idleTimeout"`
}
type MysqlConfig struct {
	Instance      string        `mapstructure:"instance"`
	Dsn           string        `mapstructure:"dsn"`
	TraceLog      bool          `mapstructure:"trace_log"`
	SlowThreshold time.Duration `mapstructure:"slow_threshold"`
}

type LoggerConfig struct {
	LogFile  string `mapstructure:"log_file"`
	LogLevel string `mapstructure:"log_level"`
}

type RedisConfig struct {
	Addr         string        `mapstructure:"addr"`
	Password     string        `mapstructure:"password"`
	Db           int           `mapstructure:"db"`
	DialTimeout  time.Duration `mapstructure:"dialTimeout"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}
