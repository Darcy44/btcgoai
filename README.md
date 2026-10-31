# BTCGoAI — Laboratório Educacional de Criptografia Bitcoin

## 1. Apresentação

BTCGoAI é um projeto desenvolvido em Go para estudos de estruturas de dados, operações criptográficas e validação de informações relacionadas ao formato de endereços Bitcoin.

O projeto contém rotinas de conversão e validação HASH160, estruturas para leitura de arquivos JSON e testes automatizados com dados de exemplo.

Este repositório deve ser utilizado como ambiente educacional, com dados fictícios e testes controlados.

## 2. Estrutura do projeto

```text
btcgoai/
├── data/
│   ├── hash160s.json
│   ├── ranges.json
│   └── wallets.json
├── bitcoin.go
├── colors.go
├── data.go
├── main.go
├── models.go
├── range_validation.go
├── range_validation_test.go
├── search.go
├── search_test.go
├── data_validation_test.go
├── bench_test.go
├── go.mod
├── go.sum
└── README.md
```

## 3. Requisitos

- Go 1.21 ou superior.
- Windows, Linux ou macOS.
- Arquivos JSON de exemplo na pasta `data`.

## 4. Organização dos dados

Os arquivos JSON utilizados pelo projeto são:

- `data/wallets.json`: contém os endereços armazenados.
- `data/hash160s.json`: contém valores HASH160 em formato hexadecimal.
- `data/ranges.json`: contém os limites e o campo de status das faixas.

O carregamento dos arquivos é realizado pelas funções existentes em `data.go`.

Os valores HASH160 são decodificados e validados para garantir o comprimento esperado de 20 bytes.

## 5. Validação de limites

O arquivo `range_validation.go` contém a função `parseRangeLimits`, responsável por validar limites hexadecimais.

As verificações incluem:

- Remoção de espaços desnecessários.
- Aceitação dos prefixos `0x` e `0X`.
- Rejeição de valores hexadecimais inválidos.
- Rejeição de valores negativos.
- Verificação da ordem dos limites.

## 6. Testes automatizados

Para executar os testes do projeto, utilize:

```bash
go test -count=1 ./...
```

Os testes verificam diferentes aspectos das rotinas implementadas, incluindo validações e cálculos criptográficos com valores de teste.

## 7. Limitações conhecidas

- A conversão de endereços para HASH160 na função `convertAddressesToHash160` ainda não está implementada.
- O campo `Status` existe na estrutura `Range`, mas seu uso como filtro não foi confirmado na implementação examinada.
- A documentação e os testes devem ser atualizados sempre que novas funcionalidades forem implementadas.

## 8. Segurança e uso responsável

Este projeto é destinado ao aprendizado de programação e criptografia.

Utilize exclusivamente dados fictícios, chaves de demonstração e ambientes controlados.

Não utilize o projeto para tentar obter acesso a carteiras de terceiros ou recuperar chaves privadas sem autorização.

Nunca compartilhe chaves privadas ou frases de recuperação.

## 9. Desenvolvimento

Contribuições são bem-vindas, especialmente melhorias em:

- Organização do código.
- Validação de entradas.
- Tratamento de erros.
- Testes automatizados.
- Documentação técnica.

---

**Projeto BTCGoAI — Estudos de Go e criptografia Bitcoin.**