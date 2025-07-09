# 🔑 Go Keylogger - Professional Monitoring Solution

> **Warning**  
> This project is for **educational purposes only**. Always obtain explicit permission before monitoring any system. Misuse of this software may violate privacy laws in your jurisdiction.

## 🚀 Superior Features Compared to Other Projects

| Feature                 | This Project | Others |
|-------------------------|--------------|--------|
| **Military-grade Encryption** | ✅ AES-256 | ❌ Plain text |
| **Context Awareness**   | ✅ Window titles | ❌ Basic logging |
| **Stealth Operation**   | ✅ No UI/process visible | ❌ Visible processes |
| **Intelligent Logging** | ✅ Context grouping | ❌ Raw streams |
| **Automatic Maintenance** | ✅ Log rotation | ❌ Manual cleanup |
| **Persistence**         | ✅ Registry install | ❌ Temporary |
| **Compiled Binary**     | ✅ No dependencies | ❌ Requires interpreters |
| **Resource Efficiency** | ✅ <5MB memory | ❌ Heavy runtimes |

## 🌟 Key Advantages

### 1. Enterprise-grade Security
- **AES-256 encryption** for log files
- Automatic key rotation
- Secure storage in `%APPDATA%` protected directory
- File permissions lockdown (0600 mode)

### 2. Contextual Intelligence
- Tracks active window titles
- Groups keystrokes by application context
- Smart buffering with 5-second timeout
- Timestamps all events with millisecond precision

### 3. Operational Stealth
- Runs with no visible UI
- Uses generic process name (`systemmonitor.exe`)
- Compiles with `-H=windowsgui` flag to hide console
- Avoids suspicious network activity

### 4. Professional Maintenance
- Automatic log rotation
- Compressed encrypted archives
- Configurable cleanup intervals
- File locking for concurrent-safe operations

### 5. Production-ready Engineering
- Proper error handling throughout
- Memory-efficient design (<5MB RAM usage)
- Windows registry integration
- Single-binary deployment
- Clean code architecture with separation of concerns

## 🛠 Installation Guide

```bash
# Build the executable
scripts\build.bat

# Install as persistent service
scripts\install.bat
```

## 📊 Log Structure

```
[2023-07-15 14:30:22] 
[Window: Chrome - Google Search]
helloworld

[2023-07-15 14:32:45] 
[Window: Microsoft Outlook]
testemail@domain.com
```

## ⚙️ Configuration

Edit `pkg/config/config.go` to modify:

```go
return &Config{
	EncryptLogs:    true,    // Enable AES-256 encryption
	UploadInterval: 60,      // Minutes between maintenance
	StealthMode:    true,    // Hide from task manager
	CaptureScreens: false,   // Screenshot capability (future)
}
```

## 📁 File Locations

- **Executable**: `%APPDATA%\SystemMonitor\systemmonitor.exe`
- **Logs**: `%APPDATA%\SystemMonitor\activity.log`
- **Archives**: `%APPDATA%\SystemMonitor\activity.log.20230715.enc`

## 🗑️ Uninstallation

1. Remove registry key:  
   `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run\SystemMonitor`
2. Delete folder:  
   `%APPDATA%\SystemMonitor`
3. Reboot system

## ⚠️ Legal Disclaimer

This software is provided for educational purposes only. The developers assume no liability for any misuse of this tool. Always comply with local laws and regulations regarding system monitoring and privacy. Unauthorized use on systems you don't own is illegal and unethical.

> "With great power comes great responsibility" - Uncle Ben (Spider-Man)

## 🧠 Technical Highlights

- **Concurrency-safe logging** with mutex locks
- **Low-level Windows API hooks** for maximum compatibility
- **Zero-dependency architecture** (pure Go)
- **Efficient memory management** with buffered writes
- **Automatic key generation** using crypto/rand
- **Window title tracking** via GetForegroundWindow API

## 🚫 Limitations

- Windows-only implementation
- Requires admin privileges for installation
- May trigger antivirus warnings
- No remote logging capabilities (by design)
