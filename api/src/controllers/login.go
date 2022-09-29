package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {

	requestBody, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Users
	if erro = json.Unmarshal(requestBody, &usuario); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	db, erro := database.Conectar()
	if erro = json.Unmarshal(requestBody, &usuario); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repository.BuscarPorEmail(usuario.Email)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	fmt.Println(usuarioSalvoNoBanco)

	if erro = security.VerifyPassword(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	w.Write([]byte("Você está logado!"))
}
