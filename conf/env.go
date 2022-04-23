package conf

import "os"

type env struct {
	MysqlUrl       string
	MySqlScheme    string
	MysqlUser      string
	MysqlPassword  string
	EthRpcUrl      string
	RedisPassword  string
	ProfilesActive string
	ConfigFilePath string
}

var Env env

const (
	EthName string = "Ethereum"
)

func init() {
	Env = env{
		MysqlUrl:       os.Getenv("MYSQL_URL"),
		MySqlScheme:    os.Getenv("MYSQL_SCHEME"),
		MysqlUser:      os.Getenv("MYSQL_USER"),
		MysqlPassword:  os.Getenv("MYSQL_PASSWORD"),
		EthRpcUrl:      os.Getenv("ETH_RPC_URL"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		ProfilesActive: os.Getenv("PROFILES_ACTIVE"),
		ConfigFilePath: getEnv("CONFIG_FILE_PATH", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
