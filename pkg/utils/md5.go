package utils

import (
	"crypto/md5"
	"fmt"
)


func MD5(bytes []byte) (string, error) {
	hash := md5.New()
	_, err := hash.Write(bytes)
	if err != nil {
		return "", err
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}
