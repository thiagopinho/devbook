package security

import "golang.org/x/crypto/bcrypt"

// Hash é diferente de criptografia, quando criptogramos um dado, temos um jeito de descriptografar, não há como desfazer um hash, a partir de um momento que colocamos um hash em um dado, é irreversível.

// Hash recebe uma string e coloca um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compara uma senha e um hash e retorna se elas são iguais
func VerifyPassword(hashPassword string, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(passwordString))
}
