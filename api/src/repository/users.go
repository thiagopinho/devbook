package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario model.Users) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {

		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos os usuarios que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]model.Users, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []model.Users

	for linhas.Next() {

		var usuario model.Users

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

// BuscarPorId traz um usuário no banco de dados
func (repositorio Usuarios) BuscarPorId(ID uint64) (model.Users, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?", ID)

	if erro != nil {
		return model.Users{}, erro
	}
	defer linhas.Close()

	var usuario model.Users

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return model.Users{}, erro
		}
	}
	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario model.Users) error {

	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (model.Users, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return model.Users{}, erro
	}
	defer linha.Close()
	var usuario model.Users
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return model.Users{}, erro
		}
	}
	return usuario, nil
}
