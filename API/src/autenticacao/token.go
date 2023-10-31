package autenticacao

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as devidas permissões
func CriarToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["usuarioId"] = usuarioId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("SecretKey"))
}
