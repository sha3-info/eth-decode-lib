package transaction_decode

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeRawTransactionBytes(bs []byte) (*types.Transaction, *common.Address, error) {
	tx := new(types.Transaction)
	err := tx.UnmarshalBinary(bs)
	if err != nil {
		return nil, nil, err
	}

	v, r, s := tx.RawSignatureValues()
	if v == nil || r == nil || s == nil {
		return tx, nil, nil
	} else {
		var signer types.Signer

		if tx.Type() == 0 {
			// eip155Signer
			signer = types.NewEIP155Signer(tx.ChainId())
		} else if tx.Type() == 1 {
			// eip2930Signer
			signer = types.NewEIP2930Signer(tx.ChainId())
		} else if tx.Type() == 2 {
			// londonSigner
			signer = types.NewLondonSigner(tx.ChainId())
		} else {
			return nil, nil, errors.New("invalid transaction type")
		}

		addr, err := signer.Sender(tx)
		if err != nil {
			return nil, nil, err
		}
		return tx, &addr, nil
	}
}

func DecodeRawTransactionHex(hexStr string) (*types.Transaction, *common.Address, error) {
	if len(hexStr)%2 != 0 {
		return nil, nil, errors.New("invalid raw transaction hex")
	}
	if len(hexStr) >= 2 && hexStr[0:2] == "0x" {
		hexStr = hexStr[2:]
	}

	bs, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, nil, err
	}
	return DecodeRawTransactionBytes(bs)
}
