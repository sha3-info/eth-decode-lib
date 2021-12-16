package input_decode

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
	"reflect"
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
	// decode txInput Payload
	decodedData := bs[4:]

	// recover Method from signature and ABI
	method, err := contractAbi.MethodById(decodedSig)
	if err != nil {
		return nil, nil, err
	}
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

func TxInputDataToPretty(data interface{}) interface{} {
	typeStr := reflect.TypeOf(data).String()

	if typeStr == "[]uint8" {
		return hex.EncodeToString(data.([]uint8))
	} else if typeStr[0] == '[' {
		array := reflect.ValueOf(data).Convert(reflect.TypeOf(data))
		arrayNew := make([]interface{}, 0)
		for i := 0; i < array.Len(); i++ {
			arrayNew = append(arrayNew, TxInputDataToPretty(array.Index(i).Interface()))
		}
		return arrayNew
	} else if typeStr == "uint8" {
		return fmt.Sprintf("%0x", data.(uint8))
	} else if typeStr == "*big.Int" {
		return data.(*big.Int).String()
	} else {
		fmt.Println("unsupported type:", typeStr)
		return ""
	}
}

func TxInputDataToJsonStr(data interface{}) string {
	data = TxInputDataToPretty(data)
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}
