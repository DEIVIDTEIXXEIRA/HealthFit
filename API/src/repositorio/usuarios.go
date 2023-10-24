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
func (repositorio Usuario) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.DB.Prepare(
		"insert into usuarios (nome, nick, email, peso, altura, idade, senha) values(?, ?, ?, ?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Peso, usuario.Altura, usuario.Idade, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdinserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdinserido), nil
}

// BuscarUsuario busca um usuario no banco mediante a seu Id.
func (repositorio Usuario) BuscarUsuario(usuarioId uint64) (modelos.Usuario , error){
	linha, erro := repositorio.DB.Query(
		"select id, nome, nick, email, peso, altura, idade from usuarios where id = ?", usuarioId,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Peso,
			&usuario.Altura,
			&usuario.Idade,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}
