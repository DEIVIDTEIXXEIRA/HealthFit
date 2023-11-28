package modelos

// Senha representa o formato da requisição feita para alterar a senha.
type Senhas struct {
	Atual string `json:"atual"`
	Nova  string `json:"nova"`
}
