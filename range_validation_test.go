package main

import "testing"

func TestParseRangeLimits(t *testing.T) {
	tests := []struct {
		name    string
		min     string
		max     string
		wantErr bool
	}{
		{"hexadecimal valido", "0x10", "0x20", false},
		{"sem prefixo", "10", "20", false},
		{"prefixo maiusculo", "0X10", "0X20", false},
		{"com espacos", " 0x10 ", " 0x20 ", false},
		{"limites iguais", "0x10", "0x10", false},
		{"minimo maior que maximo", "0x20", "0x10", true},
		{"hexadecimal invalido", "xyz", "20", true},
		{"limite vazio", "", "20", true},
		{"minimo negativo", "-10", "20", true},
		{"maximo negativo", "10", "-20", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minKey, maxKey, err := parseRangeLimits(tt.min, tt.max)

			if tt.wantErr {
				if err == nil {
					t.Fatal("esperava um erro, mas nenhum ocorreu")
				}
				return
			}

			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}

			if minKey.Cmp(maxKey) > 0 {
				t.Fatal("o limite minimo ficou maior que o maximo")
			}
		})
	}
}
