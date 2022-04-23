package conf

import "os"

type env struct {
	MysqlPassword   string
	RedisPassword   string
	ProfilesActive  string
	ConfigFilePath  string
	NacosURL        string
	NacosNamespceID string
	NacosDataID     string
	NacosGroup      string
}

var Env env

const (
	EthName string = "Ethereum"
	LatName string = "PlatON"
)

func init() {
	Env = env{
		MysqlPassword:   os.Getenv("MYSQL_NFT_PASSWORD"),
		RedisPassword:   os.Getenv("REDIS_NFT_PASSWORD"),
		ProfilesActive:  os.Getenv("PROFILES_ACTIVE"),
		ConfigFilePath:  getEnv("CONFIG_FILE_PATH", ""),
		NacosURL:        getEnv("NACOS_URL", ""),
		NacosNamespceID: getEnv("NACOS_NAMESPACE_ID", "public"),
		NacosDataID:     getEnv("NACOS_DATA_ID", "hqn_monitor"),
		NacosGroup:      getEnv("NACOS_GROUP", "DEFAULT_GROUP"),
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
