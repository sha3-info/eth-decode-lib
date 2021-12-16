package transaction_decode

import (
	"encoding/hex"
	"fmt"
	"testing"
)

const txEIP1559 = "0x02f901930181ec8459682f008510de3558118303e13f947a250d5630b4cf539739df2c5dacb4c659f2488d80b9012438ed173900000000000000000000000000000000000000000000000000000001980fc5530000000000000000000000000000000000000000000001407024467bbd548d0f00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000d538caee4b94706204d5e8ea3fe05cb70a77b52c0000000000000000000000000000000000000000000000000000000061b84fbf0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000afbf03181833ab4e8dec24d708a2a24c2baaa4a4c001a04b5e414fdb768af77d159f9b0d554d8c4b8ebe8fb07057cc6042fd80dac01032a00f58e9cba787fee2162810a196a199493ee662fba5ba8f9d2b47b26fae1fc589"
const txEIP155 = "0xf8ad834521a0851279e9144b8305573094dac17f958d2ee523a2206206994597c13d831ec780b844a9059cbb000000000000000000000000a7d9dcecc8740732b6b1d8911d660cdc1677c436000000000000000000000000000000000000000000000000000000001e3c055026a00a1063f2bbc0ad421541bf71d88fca36eddcea49c08825cc0449514978fc540ea07e1708ce8daf372e10fc3630070d902e995e3a7eeda986c3755802447aed72f9"

func TestDecodeRawTransactionBytesEIP1559(t *testing.T) {
	// 0xd4706062b4a0f167113ee3f03798841289d112e6e8ced9f4c4f1152e68c91c46
	txBytes, err := hex.DecodeString(txEIP1559[2:])
	if err != nil {
		fmt.Println(err)
		return
	}
	tx, addr, err := DecodeRawTransactionBytes(txBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("chain id:", tx.ChainId())
	fmt.Println("tx hash:", tx.Hash().String())
	fmt.Println("tx gas fee cap:", tx.GasFeeCap())
	fmt.Println("tx gas tip cap:", tx.GasTipCap())
	fmt.Println("tx gas price:", tx.GasPrice())
	fmt.Println("tx gas limit:", tx.Gas())
	fmt.Println("tx type:", tx.Type())
	fmt.Println("tx to:", tx.To())
	fmt.Println("tx nonce:", tx.Nonce())
	fmt.Println("tx value:", tx.Value())
	fmt.Println("tx data:", hex.EncodeToString(tx.Data()))
	if addr != nil {
		fmt.Println("tx signer:", addr.String())
	}
}

func TestDecodeRawTransactionHexEIP1559(t *testing.T) {
	// 0xd4706062b4a0f167113ee3f03798841289d112e6e8ced9f4c4f1152e68c91c46
	tx, addr, err := DecodeRawTransactionHex(txEIP1559)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("chain id:", tx.ChainId())
	fmt.Println("tx hash:", tx.Hash().String())
	fmt.Println("tx gas fee cap:", tx.GasFeeCap())
	fmt.Println("tx gas tip cap:", tx.GasTipCap())
	fmt.Println("tx gas price:", tx.GasPrice())
	fmt.Println("tx gas limit:", tx.Gas())
	fmt.Println("tx type:", tx.Type())
	fmt.Println("tx to:", tx.To())
	fmt.Println("tx nonce:", tx.Nonce())
	fmt.Println("tx value:", tx.Value())
	fmt.Println("tx data:", hex.EncodeToString(tx.Data()))
	if addr != nil {
		fmt.Println("tx signer:", addr.String())
	}
}

func TestDecodeRawTransactionBytesEIP155(t *testing.T) {
	// 0x4c552aa8b23fdfe7f28e1d3b69dd161f7983e1a3e6707400d263f35fc2794141
	txBytes, err := hex.DecodeString(txEIP155[2:])
	if err != nil {
		fmt.Println(err)
		return
	}
	tx, addr, err := DecodeRawTransactionBytes(txBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("chain id:", tx.ChainId())
	fmt.Println("tx hash:", tx.Hash().String())
	fmt.Println("tx gas fee cap:", tx.GasFeeCap())
	fmt.Println("tx gas tip cap:", tx.GasTipCap())
	fmt.Println("tx gas price:", tx.GasPrice())
	fmt.Println("tx gas limit:", tx.Gas())
	fmt.Println("tx type:", tx.Type())
	fmt.Println("tx to:", tx.To())
	fmt.Println("tx nonce:", tx.Nonce())
	fmt.Println("tx value:", tx.Value())
	fmt.Println("tx data:", hex.EncodeToString(tx.Data()))
	if addr != nil {
		fmt.Println("tx signer:", addr.String())
	}
}

func TestDecodeRawTransactionHexEIP155(t *testing.T) {
	// 0x4c552aa8b23fdfe7f28e1d3b69dd161f7983e1a3e6707400d263f35fc2794141
	tx, addr, err := DecodeRawTransactionHex(txEIP155)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("chain id:", tx.ChainId())
	fmt.Println("tx hash:", tx.Hash().String())
	fmt.Println("tx gas fee cap:", tx.GasFeeCap())
	fmt.Println("tx gas tip cap:", tx.GasTipCap())
	fmt.Println("tx gas price:", tx.GasPrice())
	fmt.Println("tx gas limit:", tx.Gas())
	fmt.Println("tx type:", tx.Type())
	fmt.Println("tx to:", tx.To())
	fmt.Println("tx nonce:", tx.Nonce())
	fmt.Println("tx value:", tx.Value())
	fmt.Println("tx data:", hex.EncodeToString(tx.Data()))
	if addr != nil {
		fmt.Println("tx signer:", addr.String())
	}
}
