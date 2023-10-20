package repositorio

import (
	"API/src/modelos"
	"database/sql"
)

// Usuario representa um repositório de usuário.
type Usuario struct {
	*sql.DB
}

// NovoRepositorioDeUsuario cria um repositório de usuario.
func NovoRepositorioDeUsuario(db *sql.DB) *Usuario {
	return &Usuario{db}
}

// Criar insere um usuário no banco de daddos
func (u Usuario) Criar(modelos.Usuario) (uint64, error) {
	return 0, nil 
}