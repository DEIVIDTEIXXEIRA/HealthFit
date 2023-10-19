package router

import (
	"API/src/Router/rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
