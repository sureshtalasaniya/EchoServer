package configs

import (
	"strings"
	"github.com/spf13/viper"
	"fmt"
)

type dbConfig struct {
	DbName        string     // DB_NAME
	DbHost        string     // DB_HOST 
	DbPort        int        // DB_PORT
	DbUserName    string     // DB_USERNAME
	DbPassword    string     // DB_PASSWORD
}

type cacheConfig struct {
	Host         string // CACHE_HOST
	PoolSize     int    // CACHE_POOLSIZE
}

type logConfig struct {
	LogFile		   string
	LogLevel 		string
}

type httpConfig struct {
	HostPort string
}

type eslConfig struct {
	Host string
	Port uint
	Password string
	Timeout int
}

// Config - configuration object
type Config struct {
	Cache          cacheConfig
	Log		   logConfig
	HttpConfig httpConfig
	EslConfig	eslConfig
	DBConfig dbConfig
}

var conf *Config

// GetConfig - Function to get Config
func GetConfig() *Config {
	if conf != nil {
		return conf
	}
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	fmt.Println(v.Get("cache.host"))
	fmt.Println(v.Get("cache_host"))

	cacheConf := cacheConfig{
		Host:         v.GetString("cache.host"),
		PoolSize:     v.GetInt("cache.poolsize"),
	}

	logConf := logConfig{
		LogFile:         v.GetString("log.file"),
		LogLevel:     v.GetString("log.level"),
	}

	httpConf := httpConfig{
		HostPort:         v.GetString("http.host"),
	}
	eslConf := eslConfig{
		Host: v.GetString("esl.host"),
		Port:         uint(v.GetInt("esl.port")),
		Password:         v.GetString("esl.password"),
		Timeout:         v.GetInt("esl.timeout"),
	}
	
	dbConf := dbConfig{
		DbName : v.GetString("db.name"),
		DbUserName : v.GetString("db.username"),
		DbHost : v.GetString("db.host"),
		DbPort : v.GetInt("db.port"),
		DbPassword : v.GetString("db.password"),
	}
	
	
	conf = &Config{
		Cache:     cacheConf,
		Log: 	logConf,
		HttpConfig: httpConf,
		EslConfig: eslConf,
		DBConfig:dbConf,
	}
	return conf
}

