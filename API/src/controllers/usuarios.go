package controllers

import (
	"API/src/banco"
	"API/src/modelos"
	"API/src/repositorio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorio.NovoRepositorioDeUsuario(db)
	repositorio.Criar(usuario)

	
}

// BuscarUsuario busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario"))
}

// EditarUsuario edita as informações de um usuario no banco de dados
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editando usuario"))
}

// DeletarUsuario exclui um usuario do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuario"))
}
