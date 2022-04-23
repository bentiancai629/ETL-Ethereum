package models

import "time"

type Chain struct {
	Id                  int64  `gorm:"primaryKey;autoIncrement"`
	Height              uint64 `gorm:"type:bigint(20);not null"`
	BackwardBlockNumber uint64 `gorm:"type:bigint(20);not null"`
	ChainName           string `json:"chainName" gorm:"column:chain_name; type:varchar(64); index; comment:'Chain 名称，用于跟其他相关地方关联'"`
}

/*
	Raw.Address
	Raw.Data
	Raw.BlockNumber
	Raw.TxHash
*/
type Erc20TransferEvent struct {
	Id         int64     `gorm:"primaryKey;autoIncrement"`
	ChainName  string    `json:"chainName" gorm:"column:chain_name; type:varchar(64);"`
	Height     uint64    `gorm:"type:bigint(20);index;not null"`
	TokenName  string    `gorm:"type:varchar(6);index;not null"`
	Hash       string    `gorm:"type:varchar(66);index;uniqueIndex:hash_rawIndex;not null;"`
	RawIndex   uint8     `gorm:"type:tinyint(6); uniqueIndex:hash_rawIndex;not null;"`
	EventType  uint8     `gorm:"type:tinyint(4);not null; comment:' 1-Approval  2-Transfer'"`
	From       string    `gorm:"type:varchar(100);not null"`
	To         string    `gorm:"type:varchar(100);not null"`
	Amount     uint64    `gorm:"type:bigint(20); not null; comment:'erc20转移数量'"` // todo  convert to decimal
	EventJson  []byte    `gorm:"type:text;not null"`
	Status     uint8     `gorm:"type:tinyint(4);not null;comment: '0-未处理, 1-已处理'"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time; type:datetime; default:current_timestamp; comment:'创建时间'"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time; type:datetime; default:current_timestamp; comment:'更新时间'"`
}

type ChainInfo struct {
	ID                                uint64    `json:"id" gorm:"primaryKey;autoIncrement; column:id; comment:'主键自增，ID'"`
	ChainName                         string    `json:"chainName" gorm:"column:chain_name; type:varchar(64); index; comment:'Chain 名称，用于跟其他相关地方关联'"`
	DisplayName                       string    `json:"displayName" gorm:"column:display_name; type:varchar(255); comment:'Chain 显示名称，用于展示'"`
	ChainID                           uint      `json:"chainId" gorm:"column:chain_id;index; comment:'Chain ID，记得做16进制跟10进制的转换'"`
	ExplorerURL                       string    `json:"explorerUrl" gorm:"column:explorer_url; type:text; comment:'Explorer URL，用于MetaMask添加chain，以及自己使用'"`
	RpcURL                            string    `json:"rpcUrl" gorm:"column:rpc_url; type:text; comment:'RPC地址'"`
	InternalRpcURL                    string    `json:"-" gorm:"column:internal_rpc_url; type:text; comment:'内部用RPC地址'"`
	CurrencySymbol                    string    `json:"currencySymbol" gorm:"column:currency_symbol; type:varchar(10); comment:'货币符号'"`
	CreatureContractAddress           string    `json:"creatureContractAddress" gorm:"column:creature_contract_address; type:varchar(255); comment:'Creature合约地址'"`
	ExchangeContractAddress           string    `json:"exchangeContractAddress" gorm:"column:exchange_contract_address; type:varchar(255); comment:'Exchange合约地址'"`
	OperatorContractAddress           string    `json:"operatorContractAddress" gorm:"column:operator_contract_address; type:varchar(255); comment:'Operator合约地址'"`
	ProxyContractAddress              string    `json:"proxyContractAddress" gorm:"column:proxy_contract_address; type:varchar(255); comment:'Proxy合约地址'"`
	ERC20ContractAddress              string    `json:"erc20ContractAddress" gorm:"column:erc20_contract_address; type:varchar(255); comment:'ERC20合约地址'"`
	TokenTransferProxyContractAddress string    `json:"tokenTransferProxyContractAddress" gorm:"column:token_transfer_proxy_contract_address; type:varchar(255); comment:'Token Transfer Proxy合约地址'"`
	CreatorContractAddress            string    `json:"creatorContractAddress" gorm:"column:creator_contract_address; type:varchar(255); comment:'创作者合约地址'"`
	RoyaltyRate                       float32   `json:"royaltyRate" gorm:"column:royalty_rate;type:decimal(4,4); default:0.01; comment:'版税抽成率'"`
	CreatedAt                         time.Time `json:"createdAt"  gorm:"column:created_at; comment:'创建时间'"`
	UpdatedAt                         time.Time `json:"updatedAt"  gorm:"column:updated_at; comment:'更新时间'"`
}
