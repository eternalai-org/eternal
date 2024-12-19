package utils

import (
	"strings"

	"eternal/pkg/encrypt"
	"eternal/pkg/eth"
)

// return encryptedPrivateKeyKey, address,  error
func GenerateAddress(secretKey string) (string, string, error) {
	var err error
	privateKey := ""
	encryptedPrivateKeyKey := ""
	address := ""
	privateKey, _, address, err = eth.GenerateAddress()
	if err != nil {
		return "", "", err
	}

	// encryptedWorkerKey := privWorkerKey
	encryptedPrivateKeyKey, err = encrypt.EncryptToString(privateKey, secretKey)
	if err != nil {
		return "", "", err
	}

	return encryptedPrivateKeyKey, strings.ToLower(address), nil
}
