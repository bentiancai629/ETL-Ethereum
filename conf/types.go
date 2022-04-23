package conf

type Config struct {
	ChainConfig        *ChainConfig
	DBConfig           DBConfig
	ChainListenConfig  *ChainListenConfig
	TokenAddressConfig *TokenAddressConfig
}

type ChainConfig struct {
	BackwardBlockNumber uint64
}

type DBConfig struct {
	URL      string
	User     string
	Password string
	Scheme   string
	Debug    bool
}

type ChainListenConfig struct {
	ListenSlot uint64
	BatchSize  uint64
	Defer      uint64
	From       uint64
	To         uint64
	Url        string
}

type TokenAddressConfig struct {
	TokenList []TokenInfo
}

type TokenInfo struct {
	Name     string
	Symbol   string
	Address  string
	Decimals uint8
}
