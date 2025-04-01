package glog_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/jeffotoni/quick/pkg/glog"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkGlogSlog(b *testing.B) {
	var buf bytes.Buffer
	logger := glog.Set(glog.Config{
		Format: "slog",
		Level:  glog.DEBUG,
		Writer: &buf,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug().Time().Level().Str("trace_id", "abc123").Str("func", "handler").Int("code", 200).Msg("done").Send()
	}
}

func BenchmarkGlogText(b *testing.B) {
	var buf bytes.Buffer
	logger := glog.Set(glog.Config{
		Format: "text",
		Level:  glog.DEBUG,
		Writer: &buf,
		// Separator: " | ",
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug().Time().Level().Str("trace_id", "abc123").Str("func", "handler").Int("code", 200).Msg("done").Send()
	}
}

func BenchmarkGlogJSON(b *testing.B) {
	var buf bytes.Buffer
	logger := glog.Set(glog.Config{
		Format: "json",
		Level:  glog.DEBUG,
		Writer: &buf,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug().Time().Level().Str("trace_id", "abc123").Str("func", "handler").Int("code", 200).Msg("done").Send()
	}
}

func BenchmarkZerologText(b *testing.B) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug().Str("trace_id", "abc123").Str("func", "handler").Int("code", 200).Msg("done")
	}
}

func BenchmarkZerologJSON(b *testing.B) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug().Str("trace_id", "abc123").Str("func", "handler").Int("code", 200).Msg("done")
	}
}
func BenchmarkZapText(b *testing.B) {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(io.Discard), // discard output
		zap.DebugLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("hello", zap.String("key", "val"), zap.Int("count", i))
	}
}

func BenchmarkZapJSON(b *testing.B) {
	encoderCfg := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(io.Discard), // discard output
		zap.DebugLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("json", zap.String("key", "val"), zap.Int("count", i))
	}
}

func BenchmarkLogrusText(b *testing.B) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	logger.SetOutput(io.Discard) // discard output

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"key":   "val",
			"count": i,
		}).Info("hello")
	}
}

func BenchmarkLogrusJSON(b *testing.B) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(io.Discard) // discard output

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"key":   "val",
			"count": i,
		}).Info("json")
	}
}
