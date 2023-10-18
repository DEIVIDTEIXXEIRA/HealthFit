package rotas

import "net/http"

//Rota representa o formato das rotas utilizadas na API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
