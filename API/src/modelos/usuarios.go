package modelos

import (
	"errors"
	"strings"
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

	usuario.formatar()
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
	
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O campo Senha é obrigatório e não pode estar vazio!!")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)  
}
