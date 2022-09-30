package model

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Users representa um usuário utilizando a rede social
type Users struct {
	ID       uint64    `json:"id,omitempty"` //omitempty = se o campo ID estiver em branco, ele não vai passar pro json
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (user *Users) Preparar(etapa string) error {
	if erro := user.validar(etapa); erro != nil {
		return erro
	}
	if erro := user.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (user *Users) validar(etapa string) error {

	if user.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if user.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("o e-mail inserido é inválido")
	}
	if user.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}
	if etapa == "cadastro" && user.Senha == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (user *Users) formatar(etapa string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "cadastro" {
		hashPassword, erro := security.Hash(user.Senha)
		if erro != nil {
			return erro
		}
		user.Senha = string(hashPassword)
	}
	return nil
}
