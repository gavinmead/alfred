package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// LoggerKey contains the context key for the logger
type LoggerKey string

// Key is the name of the key in the context where the logger should be present
const Key LoggerKey = "_logger_"

// GetLogger returns the zap Logger associated with this
// context
func GetLogger(ctx context.Context) *zap.Logger {
	logger := ctx.Value(Key)
	if logger == nil {
		panic(fmt.Sprintf("no logger found at key %s", Key))
	}

	switch v := logger.(type) {
	case *zap.Logger:
		return v
	default:
		panic(fmt.Sprintf("type at at key %s is not a zap logger", Key))
	}
}
