package utils

import (
	"testing"

	"github.com/spf13/viper"
)

func TestAESGCM(t *testing.T) {
	// Setup config
	viper.Set("security.encryption_key", "12345678901234567890123456789012") // 32 chars

	originalText := "Emergency Material 2026"

	// Test Encrypt
	encrypted, err := Encrypt(originalText)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}
	if encrypted == "" {
		t.Fatal("Encrypted text is empty")
	}

	// Test Decrypt
	decrypted, err := Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if decrypted != originalText {
		t.Errorf("Expected %s, got %s", originalText, decrypted)
	}
}
