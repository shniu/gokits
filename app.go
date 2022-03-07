package gokits

import (
	"context"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
)

// An application lifecycle manager

// AppInfo is application context value.
type AppInfo interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Endpoint() []string
}

type App struct {
	opts   options
	ctx    context.Context
	cancel func()
	lk     sync.Mutex
}

func New(opts ...Option) *App {
	o := options{
		ctx: context.Background(),
		sigs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		stopTimeout: 10 * time.Second,
	}

	if id, err :=uuid.NewUUID(); err == nil {
		o.id = id.String()
	}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		ctx: ctx,
		cancel: cancel,
		opts: o,
	}
}

// ID returns app instance id.
func (a *App) ID() string { return a.opts.id }

// Name returns service name.
func (a *App) Name() string { return a.opts.name }

// Version returns app version.
func (a *App) Version() string { return a.opts.version }

// Metadata returns service metadata.
func (a *App) Metadata() map[string]string { return a.opts.metadata }

// Endpoint returns endpoints.
func (a *App) Endpoint() []string {
	//if a.instance == nil {
	//	return []string{}
	//}
	//return a.instance.Endpoints
	return []string{}
}

