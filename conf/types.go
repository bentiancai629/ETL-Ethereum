package conf

type Config struct {
	ChainConfig       *ChainConfig
	DBConfig          DBConfig
	ChainListenConfig *ChainListenConfig
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
	ListenSlot      uint64
	BatchSize       uint64
	Defer           uint64
	Url             string
	WETHAddr 		string
}
