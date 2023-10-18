package modelos

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
