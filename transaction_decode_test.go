package eth_decode_lib

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestDecodeRawTransaction(t *testing.T) {
	txBytes, err := hex.DecodeString("f8ad833532178512dbf3c4ea830f4240943301ee63fb29f863f2333bd4466acb46cd8323e680b844a9059cbb000000000000000000000000ca1caf2222eea0e046e0f653e647566d8d18df950000000000000000000000000000000000000000781f5e0015f15e2f6c7c000025a09c4acb6a0c938e2316cd8391699239ce319d4411e551a8aacf6c04ef1a3debbfa02e1ef233808769afb5d7cf21950d8af87a2fb91e406940529431444f927189e2")
	if err != nil {
		fmt.Println(err)
	}
	tx, err := DecodeRawTransactionBytes(txBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tx:", *tx)
}

func TestDecodeRawTransactionHex(t *testing.T) {
	tx, err := DecodeRawTransactionHex("0xf8ad833532178512dbf3c4ea830f4240943301ee63fb29f863f2333bd4466acb46cd8323e680b844a9059cbb000000000000000000000000ca1caf2222eea0e046e0f653e647566d8d18df950000000000000000000000000000000000000000781f5e0015f15e2f6c7c000025a09c4acb6a0c938e2316cd8391699239ce319d4411e551a8aacf6c04ef1a3debbfa02e1ef233808769afb5d7cf21950d8af87a2fb91e406940529431444f927189e2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tx:", *tx)
}
