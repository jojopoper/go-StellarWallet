package publicdefine

import (
	"crypto/rand"
)

func SafeRandomBytes(length int) ([]byte, error) {
	//rand Read
	k := make([]byte, length)
	_, err := rand.Read(k)
	return k, err
}
