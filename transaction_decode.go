package eth_decode_lib

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeRawTransactionBytes(bs []byte) (*types.Transaction, error) {
	tx := new(types.Transaction)
	err := tx.UnmarshalBinary(bs)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func DecodeRawTransactionHex(hexstr string) (*types.Transaction, error) {
	if len(hexstr)%2 != 0 {
		return nil, errors.New("invalid raw transaction hex")
	}
	if len(hexstr) >= 2 && hexstr[0:2] == "0x" {
		hexstr = hexstr[2:]
	}

	bs, err := hex.DecodeString(hexstr)
	if err != nil {
		return nil, err
	}
	return DecodeRawTransactionBytes(bs)
}
