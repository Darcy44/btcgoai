package main

import "testing"

func TestPadPrivateKeyValidationExamples(t *testing.T) {
    tests := []struct {
        name string
        key  []byte
        want int
    }{
        {"chave vazia", []byte{}, 32},
        {"chave curta", []byte{1}, 32},
        {"chave de 32 bytes", make([]byte, 32), 32},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := padPrivateKey(tt.key, 32)
            if len(got) != tt.want {
                t.Fatalf("esperava %d bytes, recebeu %d", tt.want, len(got))
            }
        })
    }
}