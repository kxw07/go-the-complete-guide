package hash

import "golang.org/x/crypto/bcrypt"

func Compare(hashedString, unhashedString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(unhashedString))
	return err == nil
}

func Do(unhashedString string) (string, error) {
	hashValue, err := bcrypt.GenerateFromPassword([]byte(unhashedString), bcrypt.MinCost)
	return string(hashValue), err
}
