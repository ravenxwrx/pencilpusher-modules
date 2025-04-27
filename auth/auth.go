package auth

import (
	"github.com/ravenxwrx/pencilpusher-modules/auth/handler"
	"github.com/ravenxwrx/pencilpusher-sdk/model"
)

type AuthModule struct {
	Name     string
	Version  string
	Prefix   string
	Handlers []model.Handler
}

var _ model.Module = (*AuthModule)(nil)

func (a *AuthModule) GetName() string {
	return a.Name
}

func (a *AuthModule) GetVersion() string {
	return a.Version
}

func (a *AuthModule) GetPrefix() string {
	return a.Prefix
}

func (a *AuthModule) GetHandlers() []model.Handler {
	return a.Handlers
}

func New() *AuthModule {
	return &AuthModule{
		Name:    "auth",
		Version: "v1.0.0",
		Handlers: []model.Handler{
			handler.IndexHandler(),
		},
	}
}
