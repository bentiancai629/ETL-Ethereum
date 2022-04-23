package conf

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

var GlobalConfig *Config

func LoadConfig() (*Config, error) {
	//configs
	var cfg Config
	//env := Env.ProfilesActive
	// Use config/config.json
	configFilePath := Env.ConfigFilePath
	if strings.TrimSpace(configFilePath) == "" {
		configFilePath = "conf/config.json"
	}
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed on read config file.")
	}
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed on unmarshal config")
	}

	logs.Info("config :%v", cfg)
	mysqlPassword := Env.MysqlPassword
	if mysqlPassword != "" {
		cfg.DBConfig.Password = mysqlPassword
	}

	//redisPassword := Env.RedisPassword

	GlobalConfig = &cfg
	return &cfg, nil
}
