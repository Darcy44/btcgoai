package main

import (
	"fmt"
	"math/big"
	"strings"
)

func parseRangeLimits(minValue, maxValue string) (*big.Int, *big.Int, error) {
	normalize := func(value string) string {
		value = strings.TrimSpace(value)
		value = strings.TrimPrefix(value, "0x")
		value = strings.TrimPrefix(value, "0X")
		return value
	}

	minKey, okMin := new(big.Int).SetString(normalize(minValue), 16)
	maxKey, okMax := new(big.Int).SetString(normalize(maxValue), 16)

	if !okMin || !okMax {
		return nil, nil, fmt.Errorf("limite hexadecimal invalido")
	}

	if minKey.Sign() < 0 || maxKey.Sign() < 0 {
		return nil, nil, fmt.Errorf("limites nao podem ser negativos")
	}

	if minKey.Cmp(maxKey) > 0 {
		return nil, nil, fmt.Errorf("limite minimo maior que o maximo")
	}

	return minKey, maxKey, nil
}
