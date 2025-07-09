package logger

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger interface {
	Log(string) error
	Close() error
	EncryptAndReset() error
}

type FileLogger struct {
	file      *os.File
	filePath  string
	encrypt   bool
	writeLock sync.Mutex
}

func NewFileLogger(path string) (*FileLogger, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		file:     file,
		filePath: path,
		encrypt:  true,
	}, nil
}

func (l *FileLogger) Log(msg string) error {
	l.writeLock.Lock()
	defer l.writeLock.Unlock()
	
	_, err := l.file.WriteString(fmt.Sprintf("[%s] %s\n", 
		time.Now().Format("2006-01-02 15:04:05"), msg))
	return err
}

func (l *FileLogger) Close() error {
	return l.file.Close()
}

func (l *FileLogger) EncryptAndReset() error {
	l.writeLock.Lock()
	defer l.writeLock.Unlock()

	if err := l.file.Close(); err != nil {
		return err
	}

	data, err := os.ReadFile(l.filePath)
	if err != nil {
		return err
	}

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return err
	}

	encrypted, err := encryptData(data, key)
	if err != nil {
		return err
	}

	encPath := l.filePath + "." + time.Now().Format("20060102") + ".enc"
	if err := os.WriteFile(encPath, encrypted, 0600); err != nil {
		return err
	}

	if err := os.Truncate(l.filePath, 0); err != nil {
		return err
	}

	l.file, err = os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	return err
}

func encryptData(data, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return []byte(base64.StdEncoding.EncodeToString(ciphertext)), nil
}
