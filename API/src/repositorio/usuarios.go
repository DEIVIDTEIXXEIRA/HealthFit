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
func (repositorio Usuario) BuscarPorId(usuarioId uint64) (modelos.Usuario, error) {
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

// Editar envia as novas informações do usuario para o banco
func (repositorio Usuario) Editar(Id uint64, usuario modelos.Usuario) error {
	statetement, erro := repositorio.DB.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ?, peso = ?, altura = ?, idade = ?  where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statetement.Close()

	if _, erro = statetement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuario.Peso,
		usuario.Altura,
		usuario.Idade,
		Id,
	); erro != nil {
		return erro
	}
	return nil
}

// Deletar exclui um usuario do banco de dados
func (repositorio Usuario) Deletar(Id uint64) error {
	statement, erro := repositorio.DB.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(Id); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuario mediante a um email
func (repositorio Usuario) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.DB.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Id, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Buscar senha busca a senha atual do usuário pelo Id.
func (repositorio Usuario) BuscarSenha(usuarioId uint64) (string, error) {
	linha, erro := repositorio.DB.Query("select senha from usuarios where id = ?", usuarioId)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuario
func (repositorio Usuario) AtualizarSenha(usuarioId uint64, senha string) error {
	statement, erro := repositorio.DB.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioId); erro != nil {
		return erro
	}

	return nil
}
