package modelos

import (
	"API/src/seguranca"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuario da aplicação
type Usuario struct {
	Id     uint64 `json:"id,omitempty"`
	Nome   string `json:"nome,omitempty"`
	Nick   string `json:"nick,omitempty"`
	Email  string `json:"email,omitempty"`
	Peso   uint64 `json:"peso,omitempty"`
	Altura uint64 `json:"altura,omitempty"`
	Idade  uint64 `json:"idade,omitempty"`
	Senha  string `json:"senha,omitempty"`
}

// Preparar formata e verifica espaçoes em branco em campos obrigatórios.
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro 
	}
	return nil 
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O campo Nome é obrigatório e não pode estar vazio!!")
	}
	if usuario.Nick == "" {
		return errors.New("O campo Nick é obrigatório e não pode estar vazio!!")
	}
	if usuario.Email == "" {
		return errors.New("O campo Email é obrigatório e não pode estar vazio!!")
	}
	if usuario.Peso == 0 {
		return errors.New("O campo Peso é obrigatório e não pode estar vazio!!")
	}
	if usuario.Altura == 0 {
		return errors.New("O campo Altura é obrigatório e não pode estar vazio!!")
	}
	if usuario.Idade == 0 {
		return errors.New("O campo Idade é obrigatório e não pode estar vazio!!")
	}
	
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O Formato do email é inválido!!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O campo Senha é obrigatório e não pode estar vazio!!")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email) 
	
	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

	usuario.Senha = string(senhaComHash)
	}
	
	return nil 
}
