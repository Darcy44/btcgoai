package main

import (
	"encoding/json"
	"testing"
)

func TestModelsJSON(t *testing.T) {
	const sample = `{
		"wallets": ["DADO_FICTICIO"],
		"ranges": [
			{
				"min": "10",
				"max": "20",
				"status": 1
			}
		],
		"hash160s": [
			"0000000000000000000000000000000000000000"
		]
	}`

	var wallets WalletData
	var ranges RangeData
	var hashes Hash160Data

	if err := json.Unmarshal([]byte(sample), &wallets); err != nil {
		t.Fatalf("erro ao interpretar WalletData: %v", err)
	}

	if err := json.Unmarshal([]byte(sample), &ranges); err != nil {
		t.Fatalf("erro ao interpretar RangeData: %v", err)
	}

	if err := json.Unmarshal([]byte(sample), &hashes); err != nil {
		t.Fatalf("erro ao interpretar Hash160Data: %v", err)
	}

	if len(wallets.Wallets) != 1 || wallets.Wallets[0] != "DADO_FICTICIO" {
		t.Fatal("WalletData não foi interpretado corretamente")
	}

	if len(ranges.Ranges) != 1 {
		t.Fatal("RangeData não foi interpretado corretamente")
	}

	if ranges.Ranges[0].Min != "10" ||
		ranges.Ranges[0].Max != "20" ||
		ranges.Ranges[0].Status != 1 {
		t.Fatal("campos de Range não foram interpretados corretamente")
	}

	if len(hashes.Hash160s) != 1 ||
		hashes.Hash160s[0] != "0000000000000000000000000000000000000000" {
		t.Fatal("Hash160Data não foi interpretado corretamente")
	}

	if json.Valid([]byte(`{"wallets": [`)) {
		t.Fatal("JSON incompleto deveria ser inválido")
	}
}
