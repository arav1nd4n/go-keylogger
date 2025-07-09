package config

type Config struct {
	EncryptLogs     bool `json:"encrypt_logs"`
	UploadInterval  int  `json:"upload_interval"`
	StealthMode     bool `json:"stealth_mode"`
	CaptureScreens  bool `json:"capture_screens"`
}

func Load() *Config {
	return &Config{
		EncryptLogs:    true,
		UploadInterval: 60,
		StealthMode:    true,
		CaptureScreens: false,
	}
}
