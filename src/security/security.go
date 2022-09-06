package security

import "golang.org/x/crypto/bcrypt"

// recebe uma string e adiciona um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
// compara uma senha e um hash
func VerifiPassword(passwordWithHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordString))
}
