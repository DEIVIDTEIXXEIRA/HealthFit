package banco

import (
	"API/src/config"
	"database/sql"
)

// Conectar abre a conexão com o banco de dados e a retorna.
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
