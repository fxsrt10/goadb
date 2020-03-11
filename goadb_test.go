package goadb

import (
	"fmt"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	Debug = true
	adbPath := "C:/Users/Faysal Sharif/go/src/gitlab.com/immersivetechstudios/itsd/adbgolang/"
	adb := NewGoAdb(adbPath)
	fmt.Println("adb path is:")
	fmt.Println(adb.GetAdbPath())
	fmt.Println("adb devices:")
	devices := adb.Devices()
	fmt.Println("connected", len(devices), "device")
	for _, dev := range devices {
		if dev != nil {
			fmt.Println("  -- deviceId", dev.GetDeviceId(), "status", dev.GetStatus())
		}
	}
	fmt.Println("adb version", adb.Version())

	pm, _ := adb.ShellCmd("pm list package -f -3")
	fmt.Println(pm)
	pkgStrings := strings.Split(pm, "\n")
	fmt.Println(len(pkgStrings))
	for _, pkg := range pkgStrings {
		fmt.Println(pkg)
	}
}
