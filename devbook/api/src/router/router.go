package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

/* Gerar vai gerar um router com as rotas */
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
