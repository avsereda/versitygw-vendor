package s3log

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	logFileMode = 0644
	timeFormat  = time.RFC1123
)

type LogMeta struct {
	BucketOwner string
	ObjectSize  int64
	Action      string
	HttpStatus  int
}

// Common logging interfaces. This interface used by s3api package
// FIXME: The implementation is written by during code vendoring
type AuditLogger interface {
	Log(ctx *fiber.Ctx, err error, body []byte, meta LogMeta)
	HangUp() error
	Shutdown() error
}

type LogConfig struct {
	LogFile      string
	AdminLogFile string
}

// FIXME: Loggers set using in versitygw project to write audit and access logs
// This is simplified version of original one
type Loggers struct {
	S3Logger    AuditLogger
	AdminLogger AuditLogger
}

// Setup all loggers for provided config
func InitLogger(cfg *LogConfig) (*Loggers, error) {
	var loggers Loggers
	if cfg.LogFile != "" {
		fmt.Printf("initializing S3 access logs with '%v' file\n", cfg.LogFile)
		l, err := newSlogLogger(cfg.LogFile)
		if err != nil {
			return nil, err
		}

		loggers.S3Logger = l
	}

	if cfg.AdminLogFile != "" {
		fmt.Printf("initializing admin access logs with '%v' file\n", cfg.AdminLogFile)
		l, err := newSlogLogger(cfg.AdminLogFile)
		if err != nil {
			return nil, err
		}

		loggers.AdminLogger = l
	}

	return &loggers, nil
}

// Simple implementation of file logger with slog package
type slogLogger struct {
	logname string
	slog    *slog.Logger
	file    *os.File
	mu      sync.Mutex
}

func newSlogLogger(logname string) (AuditLogger, error) {
	var logger slogLogger
	if err := logger.openFile(logname); err != nil {
		return nil, err
	}

	return &logger, nil
}

func (f *slogLogger) openFile(logname string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.file != nil { // Close currently open file if any
		if err := f.file.Close(); err != nil {
			return fmt.Errorf("close log file erorr: %w", err)
		}
	}

	var err error
	f.file, err = os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, logFileMode)
	if err != nil {
		return fmt.Errorf("open log file error: %w", err)
	}

	f.slog = slog.New(
		slog.NewJSONHandler(f.file, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		}))

	return nil
}

// HangUp signals to the implementation to flush existing log messages and reopen log file
func (f *slogLogger) HangUp() error {
	return f.openFile(f.logname)
}

func (f *slogLogger) Shutdown() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.file.Close()
}

// Log write log message to the file
func (f *slogLogger) Log(ctx *fiber.Ctx, err error, body []byte, meta LogMeta) {
	f.mu.Lock()
	defer f.mu.Unlock()
	// TODO: This method should implement event logging to the existing sergbase.t_ib_log or file
}
