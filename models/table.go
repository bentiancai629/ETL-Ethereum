package models

func (Chain) TableName() string {
	return "t_chain"
}

func (Erc20TransferEvent) TableName() string {
	return "t_erc20transfer_event"
}
func (ChainInfo) TableName() string {
	return "t_chain_info"
}
