package banco

import (
	"fmt"
	"os"
	"testing"
)

// Função para testar a abertura de conexão com o banco de dados
func TestConectar(t *testing.T) {
	TestesDeStringConexaoBanco := []struct {
		descricao     string
		stringConexao string
	}{
		{descricao: "Teste com o usuario errado", stringConexao: "usuario:senha@/nome?charset=utf8&parseTime=True&loc=Local"},
		{descricao: "teste com a senha errada", stringConexao: "root:senha@/healthfit?charset=utf8&parseTime=True&loc=Local"},
		{descricao: "teste com o nome do banco errado", stringConexao: "root:golang@/banco?charset=utf8&parseTime=True&loc=Local"},
		{descricao: "teste com tudo correto", stringConexao: fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USUARIO"),
			os.Getenv("DB_SENHA"),
			os.Getenv("DB_NOME"),
		)},
	}

	for _, test := range TestesDeStringConexaoBanco {
		resultado := 
	}
}
