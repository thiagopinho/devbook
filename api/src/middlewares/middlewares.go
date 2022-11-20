package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidateToken(r); erro != nil {
			response.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunction(w, r)
	}
}
