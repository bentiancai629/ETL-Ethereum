package listener

import (
	"encoding/json"
	"errors"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BlockNonce is an 81-byte vrf proof containing random numbers
// Used to verify the block when receiving the block
type BlockNonce [81]byte

// EncodeNonce converts the given byte to a block nonce.
func EncodeNonce(v []byte) BlockNonce {
	var n BlockNonce
	copy(n[:], v)
	return n
}

func (n BlockNonce) Bytes() []byte {
	return n[:]
}

// MarshalText encodes n as a hex string with 0x prefix.
func (n BlockNonce) MarshalText() ([]byte, error) {
	return hexutil.Bytes(n[:]).MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *BlockNonce) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("BlockNonce", input, n[:])
}

// Header represents a block header in the Ethereum core.
type PlatonHeader struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"            gencodec:"required"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	//Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
	Number   *big.Int   `json:"number"           gencodec:"required"`
	GasLimit uint64     `json:"gasLimit"         gencodec:"required"`
	GasUsed  uint64     `json:"gasUsed"          gencodec:"required"`
	Time     uint64     `json:"timestamp"        gencodec:"required"`
	Extra    []byte     `json:"extraData"        gencodec:"required"`
	Nonce    BlockNonce `json:"nonce"            gencodec:"required"`

	// caches
	sealHash  atomic.Value `json:"-" rlp:"-"`
	hash      atomic.Value `json:"-" rlp:"-"`
	publicKey atomic.Value `json:"-" rlp:"-"`
}

// MarshalJSON marshals as JSON.
func (h PlatonHeader) MarshalJSON() ([]byte, error) {
	return nil, errors.New("not support")
}

// UnmarshalJSON unmarshals from JSON.
func (h *PlatonHeader) UnmarshalJSON(input []byte) error {
	type PlatonHeader struct {
		ParentHash  *common.Hash    `json:"parentHash"       gencodec:"required"`
		Coinbase    *common.Address `json:"miner"            gencodec:"required"`
		Root        *common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash      *common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash *common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Number      *hexutil.Big    `json:"number"           gencodec:"required"`
		GasLimit    *hexutil.Uint64 `json:"gasLimit"         gencodec:"required"`
		GasUsed     *hexutil.Uint64 `json:"gasUsed"          gencodec:"required"`
		Time        *hexutil.Uint64 `json:"timestamp"        gencodec:"required"`
		Extra       *hexutil.Bytes  `json:"extraData"        gencodec:"required"`
		Nonce       *BlockNonce     `json:"nonce"            gencodec:"required"`
	}
	var dec PlatonHeader
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for PlatonHeader")
	}
	h.ParentHash = *dec.ParentHash
	if dec.Coinbase == nil {
		return errors.New("missing required field 'miner' for PlatonHeader")
	}
	h.Coinbase = *dec.Coinbase
	if dec.Root == nil {
		return errors.New("missing required field 'stateRoot' for PlatonHeader")
	}
	h.Root = *dec.Root
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for PlatonHeader")
	}
	h.TxHash = *dec.TxHash
	if dec.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for PlatonHeader")
	}
	h.ReceiptHash = *dec.ReceiptHash
	h.Number = (*big.Int)(dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for PlatonHeader")
	}
	h.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for PlatonHeader")
	}
	h.GasUsed = uint64(*dec.GasUsed)
	if dec.Time == nil {
		return errors.New("missing required field 'timestamp' for PlatonHeader")
	}
	h.Time = uint64(*dec.Time)
	if dec.Extra == nil {
		return errors.New("missing required field 'extraData' for PlatonHeader")
	}
	h.Extra = *dec.Extra
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' for PlatonHeader")
	}
	h.Nonce = *dec.Nonce
	return nil
}
