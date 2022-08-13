package model

import "time"

// Users representa um usuário utilizando a rede social
type Users struct {
	ID       uint64    `json:"id,omitempty"` //omitempty = se o campo ID estiver em branco, ele não vai passar pro json
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}
