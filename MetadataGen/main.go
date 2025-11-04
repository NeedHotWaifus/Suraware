package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var legitimateCompanies = []string{
	"Microsoft Corporation",
	"Adobe Systems Incorporated",
	"Oracle Corporation",
	"Intel Corporation",
	"NVIDIA Corporation",
	"Google LLC",
	"Apple Inc.",
	"Cisco Systems, Inc.",
	"VMware, Inc.",
	"Symantec Corporation",
}

var legitimateProducts = []string{
	"Windows Update Assistant",
	"System Configuration Utility",
	"Microsoft Office Update",
	"Adobe Flash Player Installer",
	"Java Update Scheduler",
	"NVIDIA Graphics Driver",
	"Windows Security Update",
	"System Maintenance Tool",
	"Software Update Service",
	"Configuration Manager",
}

var descriptions = []string{
	"Provides automatic updates for Windows",
	"System maintenance and optimization tool",
	"Keeps your software up to date",
	"Essential system component for Windows",
	"Manages system configuration and updates",
	"Ensures system security and stability",
	"Automatic software update manager",
	"System diagnostics and repair utility",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	company := legitimateCompanies[rand.Intn(len(legitimateCompanies))]
	product := legitimateProducts[rand.Intn(len(legitimateProducts))]
	description := descriptions[rand.Intn(len(descriptions))]

	// Generate random version
	major := rand.Intn(10) + 1
	minor := rand.Intn(20)
	patch := rand.Intn(100)
	build := rand.Intn(9999)
	version := fmt.Sprintf("%d.%d.%d.%d", major, minor, patch, build)

	// Generate year for copyright
	year := time.Now().Year()

	// Create .syso resource file content
	manifest := generateManifest(product, description)

	// Generate rsrc command
	fmt.Println("[*] Generating fake metadata...")
	fmt.Printf("    Company: %s\n", company)
	fmt.Printf("    Product: %s\n", product)
	fmt.Printf("    Version: %s\n", version)
	fmt.Printf("    Description: %s\n", description)
	fmt.Println()

	// Create version info JSON for rsrc tool
	versionInfo := fmt.Sprintf(`{
	"FixedFileInfo": {
		"FileVersion": {
			"Major": %d,
			"Minor": %d,
			"Patch": %d,
			"Build": %d
		},
		"ProductVersion": {
			"Major": %d,
			"Minor": %d,
			"Patch": %d,
			"Build": %d
		},
		"FileFlagsMask": "3f",
		"FileFlags ": "00",
		"FileOS": "040004",
		"FileType": "01",
		"FileSubType": "00"
	},
	"StringFileInfo": {
		"Comments": "%s",
		"CompanyName": "%s",
		"FileDescription": "%s",
		"FileVersion": "%s",
		"InternalName": "%s",
		"LegalCopyright": "Copyright (C) %d %s. All rights reserved.",
		"OriginalFilename": "svchost.exe",
		"ProductName": "%s",
		"ProductVersion": "%s"
	},
	"VarFileInfo": {
		"Translation": {
			"LangID": "0409",
			"CharsetID": "04B0"
		}
	},
	"IconPath": "icon.ico",
	"ManifestPath": "manifest.xml"
}`, major, minor, patch, build, major, minor, patch, build,
		description, company, description, version,
		"svchost.exe", year, company, product, version)

	// Write version info
	err := os.WriteFile("versioninfo.json", []byte(versionInfo), 0644)
	if err != nil {
		fmt.Printf("Error writing versioninfo.json: %v\n", err)
		os.Exit(1)
	}

	// Write manifest
	err = os.WriteFile("manifest.xml", []byte(manifest), 0644)
	if err != nil {
		fmt.Printf("Error writing manifest.xml: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[+] Generated versioninfo.json")
	fmt.Println("[+] Generated manifest.xml")
	fmt.Println()
	fmt.Println("[*] Next steps:")
	fmt.Println("    1. Place icon.ico in this directory (Windows system icon)")
	fmt.Println("    2. Install: go install github.com/tc-hib/go-winres@latest")
	fmt.Println("    3. Run: go-winres make")
	fmt.Println("    4. This creates rsrc_windows_amd64.syso")
	fmt.Println("    5. Build normally - Go will include the resource file")
}

func generateManifest(productName, description string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
  <assemblyIdentity
    version="1.0.0.0"
    processorArchitecture="*"
    name="%s"
    type="win32"
  />
  <description>%s</description>
  <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
      <requestedPrivileges>
        <requestedExecutionLevel level="asInvoker" uiAccess="false"/>
      </requestedPrivileges>
    </security>
  </trustInfo>
  <compatibility xmlns="urn:schemas-microsoft-com:compatibility.v1">
    <application>
      <supportedOS Id="{e2011457-1546-43c5-a5fe-008deee3d3f0}"/>
      <supportedOS Id="{35138b9a-5d96-4fbd-8e2d-a2440225f93a}"/>
      <supportedOS Id="{4a2f28e3-53b9-4441-ba9c-d69d4a4a6e38}"/>
      <supportedOS Id="{1f676c76-80e1-4239-95bb-83d0f6d0da78}"/>
      <supportedOS Id="{8e0f7a12-bfb3-4fe8-b9a5-48fd50a15a9a}"/>
    </application>
  </compatibility>
</assembly>`, productName, description)
}
