package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// loadWalletAddresses loads wallet addresses from data/wallets.json
func loadWalletAddresses() ([]string, error) {
	file, err := os.Open("data/wallets.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var walletData WalletData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&walletData); err != nil {
		return nil, err
	}

	return walletData.Wallets, nil
}

func loadWalletHash160s() ([][]byte, error) {
	file, err := os.Open("data/hash160s.json")
	if err != nil {
		return convertAddressesToHash160()
	}
	defer file.Close()

	var hash160Data Hash160Data
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&hash160Data); err != nil {
		return nil, err
	}

	result := make([][]byte, len(hash160Data.Hash160s))

	for i, hexStr := range hash160Data.Hash160s {
		hash160Bytes, err := hex.DecodeString(strings.TrimSpace(hexStr))

		if err != nil {
			return nil, fmt.Errorf(
				"HASH160 inválido na posição %d: %w",
				i+1, err,
			)
		}

		if len(hash160Bytes) != 20 {
			return nil, fmt.Errorf(
				"HASH160 na posição %d tem %d bytes; esperado: 20",
				i+1, len(hash160Bytes),
			)
		}

		result[i] = hash160Bytes
	}

	return result, nil
}

// convertAddressesToHash160 is a fallback function that loads wallet addresses.
func convertAddressesToHash160() ([][]byte, error) {
	_, err := loadWalletAddresses()
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf(
		"conversão de endereços para HASH160 ainda não implementada; crie data/hash160s.json",
	)
}

// loadRanges loads ranges from data/ranges.json
func loadRanges() ([]Range, error) {
	file, err := os.Open("data/ranges.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rangeData RangeData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&rangeData); err != nil {
		return nil, err
	}

	return rangeData.Ranges, nil
}
