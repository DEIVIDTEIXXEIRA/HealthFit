package controllers

import (
	"API/src/modelos"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCriarUsuario(t *testing.T) {
	usuario := modelos.Usuario{
		Nome:   "deivid",
		Nick:   "kobe",
		Email:  "deivid@gmail.com",
		Peso:   87,
		Altura: 180,
		Idade:  25,
		Senha:  "000",
	}

	var b bytes.Buffer
	erro := json.NewEncoder(&b).Encode(usuario)
	if erro != nil {
		t.Error(erro)
	}

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "localhost:3030/usuarios", &b)

	CriarUsuario(response, request)

	if http.StatusOK != response.Code {
		t.Errorf("Status esperado %d, obeteve %d", http.StatusOK, response.Code)
	}

	resposta, erro := io.ReadAll(response.Body)
	if erro != nil {
		t.Error(erro)
	}

	respostaEsperada := fmt.Sprintf("Id inserido: %d", usuario.Id)
	if respostaEsperada != string(resposta) {
		t.Errorf("Resposta esperada %s, obeteve %s", respostaEsperada, string(resposta))
	}
}
