package gokits

import (
	"context"
	"github.com/shniu/gokits/transport"
	"net/url"
	"os"
	"time"
)

// Option for application manager.

type Option func(o *options)

// options for an application.
type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []*url.URL

	ctx  context.Context
	sigs []os.Signal

	stopTimeout time.Duration
	servers     []transport.Server
}

// Set service id.
func ID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

// Set service name.
func Name(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

// Set service version
func Version(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

func Metadata(metadata map[string]string) Option {
	return func(o *options) {
		o.metadata = metadata
	}
}

func Endpoint(endpoints ...*url.URL) Option {
	return func(o *options) {
		o.endpoints = endpoints
	}
}

func Context(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func Signal(sigs ...os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

func StopTimeout(t time.Duration) Option {
	return func(o *options) {
		o.stopTimeout = t
	}
}

func Server(srv ...transport.Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}