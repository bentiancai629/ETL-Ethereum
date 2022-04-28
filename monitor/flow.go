package monitor

import "fmt"

type BlockEvent interface {

}

func GetEvent() {

}

type Transformer interface {
	Transform(event BlockEvent)
}

type ContractEventEvent struct {

}

func (c *ContractEventEvent) GetEvent() (cc *ContractEventEvent) {
	d := new(ContractEventEvent);
	return d;
}

type ContractEventTransformer struct {

}

func (c *ContractEventTransformer) Transform(event ContractEventEvent) (cEvent *ContractEventEvent)  {
	tx := new(ContractEventEvent);
	return tx;
}

type TxTransformer struct {

}

type TxEvent struct {

}


func (t *TxTransformer) Transform(event TxEvent) (txEvent *TxEvent)  {
	tx := new(TxEvent);
	fmt.Println(txEvent.TxHash);
	return tx;
}



type Parser interface {
	Parse(event ContractEventEvent);
}

type ERC20Parser struct {

}

func (erc20 *ERC20Parser) Parse(event ContractEventEvent)  {
	//do save db
}

type ERC20TokenTransformer struct {

}

func (erc20 *ERC20TokenTransformer) Transform(event ContractEventEvent) (p *Parser){
	//save database object

}

type ERC721TokenTransformer struct {

}

func (erc721 *ERC721TokenTransformer)  Transform(event ContractEventEvent) (p *Parser) {

}