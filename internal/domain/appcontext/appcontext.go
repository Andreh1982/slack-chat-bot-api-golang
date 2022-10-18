package appcontext

import (
	"context"
	"slack-message-api/internal/infrastructure/simlogger/wrapper"
)

type Context interface {
	SetLogger(logger wrapper.Logger)
	Logger() wrapper.Logger
}

type appContext struct {
	logger wrapper.Logger
	ctx    context.Context
}

func New(ctx context.Context) Context {
	return &appContext{
		ctx: ctx,
	}
}

func NewBackground() Context {
	ctx := context.Background()
	return &appContext{
		ctx: ctx,
	}
}

func (appContext *appContext) SetLogger(logger wrapper.Logger) {
	appContext.logger = logger
}

func (appContext *appContext) Logger() wrapper.Logger {
	return appContext.logger
}
