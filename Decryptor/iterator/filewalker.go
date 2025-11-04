package filewalker

import (
	"Sura-Decryptor/configuration"
	"Sura-Decryptor/decryption"
	"encoding/base64"
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

func DecryptDirectory(dirPath string) {
	var wg sync.WaitGroup
	noteFile := d("ZGVjcnlwdGlvbiBpbnN0cnVjdGlvbnMudHh0")

	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			basePath := strings.ToLower(filepath.Base(path))
			for _, excluded := range configuration.ExcludedDirectories {
				if strings.EqualFold(basePath, excluded) {
					return filepath.SkipDir
				}
			}
		}

		if !info.IsDir() {
			fileExt := filepath.Ext(path)
			fileName := strings.ToLower(filepath.Base(path))

			if fileName == noteFile {
				time.Sleep(time.Millisecond * time.Duration(50+len(path)%30))
				_ = os.Remove(path)
			} else if strings.HasSuffix(fileExt, configuration.EncryptedExtension) {
				wg.Add(1)
				go func() {
					defer wg.Done()
					decryption.DecryptFile(path)
					_ = os.Rename(path, strings.TrimSuffix(path, configuration.EncryptedExtension))
				}()
			}
		}

		return nil
	})

	wg.Wait()
}
