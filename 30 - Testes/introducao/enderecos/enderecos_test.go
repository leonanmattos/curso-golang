package enderecos

import "testing"

func TestTipoEndereco(t *testing.T) {
	// Começa com Test + nome da funcao
	endereco := "Rua Paulista"
	tipoDeEnderecoEsperado := "avenida"

	tipoEnderecoRecebido := TipoEndereco(endereco)

	if tipoEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo recebido é diferente do esperado")
	}
}
