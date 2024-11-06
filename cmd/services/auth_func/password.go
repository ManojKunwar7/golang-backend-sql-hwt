package auth_func

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func CompareHashedPassword(hashed []byte, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, plain)
	return err == nil
}
