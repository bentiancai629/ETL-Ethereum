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

	// tokenJson
	tokenAddressFilePath := Env.ConfigFilePath
	if strings.TrimSpace(tokenAddressFilePath) == "" {
		tokenAddressFilePath = "conf/tokenList.json"
	}
	contentTokenList, err := ioutil.ReadFile(tokenAddressFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed on read config file.")
	}
	err = json.Unmarshal(contentTokenList, &cfg.TokenAddressConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed on unmarshal config")
	}
	logs.Info("config :%v", cfg)

	if Env.MysqlUrl == "" || Env.MySqlScheme == "" || Env.MysqlUser == "" || Env.MysqlPassword == "" || Env.EthRpcUrl == "" {
		return nil, errors.Wrap(err, "need to initialize environment")
	}

	cfg.DBConfig.URL = Env.MysqlUrl
	cfg.DBConfig.Scheme = Env.MySqlScheme
	cfg.DBConfig.User = Env.MysqlUser
	cfg.DBConfig.Password = Env.MysqlPassword
	cfg.ChainListenConfig.Url = Env.EthRpcUrl

	GlobalConfig = &cfg
	return &cfg, nil
}
