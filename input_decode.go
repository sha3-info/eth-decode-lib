package eth_decode_lib

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

func DecodeTxInputsBytes(bs []byte, abiJsonStr string) ([]byte, []interface{}, error) {
	// load contract ABI
	contractAbi, err := abi.JSON(strings.NewReader(abiJsonStr))
	if err != nil {
		return nil, nil, err
	}

	// decode txInput method signature
	decodedSig := bs[:4]

	// recover Method from signature and ABI
	method, err := contractAbi.MethodById(decodedSig)
	if err != nil {
		return nil, nil, err
	}

	// decode txInput Payload
	decodedData := bs[4:]

	// unpack method inputs
	data, err := method.Inputs.Unpack(decodedData)
	if err != nil {
		return nil, nil, err
	}
	return decodedSig, data, nil
}

func DecodeTxInputsHex(hexStr string, abiJsonStr string) ([]byte, []interface{}, error) {
	if len(hexStr)%2 != 0 {
		return nil, nil, errors.New("invalid transaction input hex")
	}
	if len(hexStr) >= 2 && hexStr[0:2] == "0x" {
		hexStr = hexStr[2:]
	}

	bs, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, nil, err
	}
	return DecodeTxInputsBytes(bs, abiJsonStr)
}
