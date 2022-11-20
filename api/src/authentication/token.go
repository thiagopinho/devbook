package authentication

import (
	"api/src/config"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken retorna um token assinado com as permissões do usuário
func CreateToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken verifica se o token passado na requisição é válido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
}

// extractToken
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]

	}
	return ""
}
