package models

import "time"

type Chain struct {
	Id                  int64  `gorm:"primaryKey;autoIncrement"`
	Height              uint64 `gorm:"type:bigint(20);not null"`
	BackwardBlockNumber uint64 `gorm:"type:bigint(20);not null"`
	ChainName           string `json:"chainName" gorm:"column:chain_name; type:varchar(64); index; comment:'Chain 名称，用于跟其他相关地方关联'"`
}

type Erc20TransferEvent struct {
	Id         int64     `gorm:"primaryKey;autoIncrement"`
	ChainName  string    `json:"chainName" gorm:"column:chain_name; type:varchar(64);"`
	Height     uint64    `gorm:"type:bigint(20);index;not null"`
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
	ChainID                           uint      `json:"chainId" gorm:"column:chain_id;index; comment:'Chain ID，记得做16进制跟10进制的转换'"`
	CreatedAt                         time.Time `json:"createdAt"  gorm:"column:created_at; comment:'创建时间'"`
	UpdatedAt                         time.Time `json:"updatedAt"  gorm:"column:updated_at; comment:'更新时间'"`
}
