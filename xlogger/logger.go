package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gitlab.marathon.edu.vn/pkg/go/xcontext"
)

const (
	contextKey       = "logger"
	RequestIDKey int = 0
)

var (
	fieldMap = logrus.FieldMap{
		logrus.FieldKeyMsg: "message",
	}
)

func CtxPrefix(ctx context.Context, prf string) *logrus.Entry {
	v := ctx.Value(contextKey)
	if log, ok := v.(*logrus.Entry); ok {
		log = log.WithField("request-id", fmt.Sprint(ctx.Value(RequestIDKey)))
		return log
	}
	return initLog(prf)
}

func CToL(ctx context.Context, label string) *logrus.Entry {
	// CToL stands for Context-To-Log
	v := ctx.Value(contextKey)
	traceID := cast.ToString(ctx.Value(xcontext.KeyContextID.String()))
	if log, ok := v.(*logrus.Entry); ok {
		return log.WithField("label", label).WithField("request-id", fmt.Sprint(traceID))
	}

	return initLog(label).WithField("request-id", fmt.Sprint(traceID))
}

func LToC(parent context.Context, logger *logrus.Entry) context.Context {
	// LToC stands for Log-To-Context
	return context.WithValue(parent, contextKey, logger)
}

func initLog(label string) *logrus.Entry {

	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{FieldMap: fieldMap})

	logger.SetLevel(logrus.DebugLevel)

	return logger.WithField("label", label)
}
