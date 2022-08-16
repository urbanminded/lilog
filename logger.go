package main

import "context"

type Logger interface {
	Trace(msg string, fields ...interface{})
	TraceContext(ctx context.Context, msg string, fields ...interface{})
	TraceArgs(args map[string]interface{}, msg string, fields ...interface{})
	TraceContextArgs(ctx context.Context, args map[string]interface{}, msg string, fields ...interface{})

	Debug(msg string, fields ...interface{})
	DebugContext(ctx context.Context, msg string, fields ...interface{})
	DebugArgs(args map[string]interface{}, msg string, fields ...interface{})
	DebugContextArgs(ctx context.Context, args map[string]interface{}, msg string, fields ...interface{})

	Info(msg string, fields ...interface{})
	InfoContext(ctx context.Context, msg string, fields ...interface{})
	InfoArgs(args map[string]interface{}, msg string, fields ...interface{})
	InfoContextArgs(ctx context.Context, args map[string]interface{}, msg string, fields ...interface{})

	Warn(msg string, fields ...interface{})
	WarnContext(ctx context.Context, msg string, fields ...interface{})
	WarnArgs(args map[string]interface{}, msg string, fields ...interface{})
	WarnContextArgs(ctx context.Context, args map[string]interface{}, msg string, fields ...interface{})

	Error(err error, msg string, fields ...interface{})
	ErrorContext(ctx context.Context, err error, msg string, fields ...interface{})
	ErrorArgs(args map[string]interface{}, err error, msg string, fields ...interface{})
	ErrorContextArgs(ctx context.Context, err error, args map[string]interface{}, msg string, fields ...interface{})
}
