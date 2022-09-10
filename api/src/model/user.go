package model

import (
	"errors"
	"strings"
	"time"
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
	user.formatar()
	return nil
}

func (user *Users) validar(etapa string) error {

	if user.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if user.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}
	if user.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if etapa == "cadastro" && user.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (user *Users) formatar() {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
