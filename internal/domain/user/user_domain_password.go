package user

import (
	"crypto/md5"
	"encoding/hex"
)

// Implementação do método EncryptPassword para a estrutura UserDomain, que criptografa a senha usando MD5
func (ud *userDomain) EncryptPassword() {
	// Criptografa a senha usando MD5 e armazena o resultado de volta na estrutura UserDomain
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
