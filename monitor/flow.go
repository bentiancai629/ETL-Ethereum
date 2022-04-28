package monitor

import _ "github.com/cskr/pubsub"

// 区块链处理
type BlockEvent interface{}

// tx转账
type Transformer interface {
	Transforms(event BlockEvent)
}

// 合约事件
type ContractEventEvent struct {
}

func (c *ContractEventEvent) GetEvent() (cc *ContractEventEvent) {
	d := new(ContractEventEvent)
	return d
}

// 合约-转账事件
type ContractEventTransformer struct {
}

func (c *ContractEventTransformer) Transform(event ContractEventEvent) (cEvent *ContractEventEvent) {
	tx := new(ContractEventEvent)
	return tx
}

type TxTransformer struct{}
type TxEvent struct{}

func (t *TxTransformer) Transform(event TxEvent) (txEvent *TxEvent) {
	tx := new(TxEvent)
	return tx
}

type Parser interface {
	Parse(event ContractEventEvent)
}

type ERC20Parser struct{}

func (erc20 *ERC20Parser) Parse(event ContractEventEvent) {
	//do save db
}

type ERC20TokenTransformer struct {
}

func (erc20 *ERC20TokenTransformer) Transform(event ContractEventEvent) (p *Parser) {
	//save database object
	return nil
}

type ERC721TokenTransformer struct {
}

func (erc721 *ERC721TokenTransformer) Transform(event ContractEventEvent) (p *Parser) {
	return nil
}

func main()  {

}