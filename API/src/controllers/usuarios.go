package controllers

import "net/http"

//CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))
}

//BuscarUsuario busca um usuário no banco de dados
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