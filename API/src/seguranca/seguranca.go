package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash coloca um hash na senha recebida
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificaSenha verifica se a senha é a mesma que o hash
func VerificaSenha(senhaString, senhaComHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}