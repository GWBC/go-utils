package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AesGCM struct {
	gcm cipher.AEAD
}

func (a *AesGCM) Init(key string) error {
	block, err := aes.NewCipher([]byte(FillStr(key, aes.BlockSize, '0')))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	a.gcm = gcm

	return nil
}

func (a *AesGCM) Encrypt(data string) (string, error) {
	//一次性随机值
	nonce := make([]byte, a.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherData := a.gcm.Seal(nonce, nonce, []byte(data), nonce)
	return base64.StdEncoding.EncodeToString(cipherData), nil
}

func (a *AesGCM) Decrypt(data string) (string, error) {
	cipherData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	size := a.gcm.NonceSize()
	if len(cipherData) < size {
		return "", errors.New("cipherdata too short")
	}

	nonce, cipherAndTag := cipherData[:size], cipherData[size:]

	plaintext, err := a.gcm.Open(nil, nonce, cipherAndTag, nonce)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
