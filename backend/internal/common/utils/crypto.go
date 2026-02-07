package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/spf13/viper"
)

// Encrypt 使用 AES-GCM 加密字符串
func Encrypt(text string) (string, error) {
	key := viper.GetString("security.encryption_key")
	if len(key) != 32 {
		return "", fmt.Errorf("encryption key must be 32 characters for AES-256")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Seal 将随机 nonce 放在密文前面
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 使用 AES-GCM 解密字符串
func Decrypt(cryptoText string) (string, error) {
	key := viper.GetString("security.encryption_key")
	if len(key) != 32 {
		return "", fmt.Errorf("encryption key must be 32 characters for AES-256")
	}

	data, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
