package lilog

import "context"

type Discard struct {
}

var _ Logger = &Discard{}

func (d *Discard) Trace(msg string, fields ...interface{})                                  {}
func (d *Discard) TraceContext(ctx context.Context, msg string, fields ...interface{})      {}
func (d *Discard) TraceArgs(args map[string]interface{}, msg string, fields ...interface{}) {}
func (d *Discard) TraceContextArgs(ctx context.Context, args map[string]interface{}, msg string, fields ...interface{}) {
}

func (d *Discard) Debug(msg string, fields ...interface{})                                  {}
func (d *Discard) DebugContext(ctx context.Context, msg string, fields ...interface{})      {}
func (d *Discard) DebugArgs(args map[string]interface{}, msg string, fields ...interface{}) {}
func (d *Discard) DebugContextArgs(ctx context.Context, args map[string]interface{}, msg string, fields ...interface{}) {
}

func (d *Discard) Info(msg string, fields ...interface{})                                  {}
func (d *Discard) InfoContext(ctx context.Context, msg string, fields ...interface{})      {}
func (d *Discard) InfoArgs(args map[string]interface{}, msg string, fields ...interface{}) {}
func (d *Discard) InfoContextArgs(ctx context.Context, args map[string]interface{}, msg string, fields ...interface{}) {
}

func (d *Discard) Warn(msg string, fields ...interface{})                                  {}
func (d *Discard) WarnContext(ctx context.Context, msg string, fields ...interface{})      {}
func (d *Discard) WarnArgs(args map[string]interface{}, msg string, fields ...interface{}) {}
func (d *Discard) WarnContextArgs(ctx context.Context, args map[string]interface{}, msg string, fields ...interface{}) {
}

func (d *Discard) Error(err error, msg string, fields ...interface{})                             {}
func (d *Discard) ErrorContext(ctx context.Context, err error, msg string, fields ...interface{}) {}
func (d *Discard) ErrorArgs(args map[string]interface{}, err error, msg string, fields ...interface{}) {
}
func (d *Discard) ErrorContextArgs(ctx context.Context, err error, args map[string]interface{}, msg string, fields ...interface{}) {
}
