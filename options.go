package auth

import (
	"github.com/fulgurant/session"
	"go.uber.org/zap"
)

type Options struct {
	config         *Config
	logger         *zap.Logger
	sessionHandler *session.Handler
}

func DefaultOptions() *Options {
	return &Options{}
}

// WithConfig sets up stuff from the config
func (o *Options) WithConfig(value *Config) *Options {
	o.config = value
	return o
}

// WithLogger sets the logger for the server
func (o *Options) WithLogger(value *zap.Logger) *Options {
	o.logger = value.With(zap.String("module", "auth-handler"))
	return o
}

// WithSessionHandler sets the session handler for the server
func (o *Options) WithSessionHandler(value *session.Handler) *Options {
	o.sessionHandler = value
	return o
}
