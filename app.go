// SPDX-License-Identifier: EUPL-1.2

package api

import (
	"git.zzdats.lv/edim/api-sender/interfaces"
	"git.zzdats.lv/edim/api-sender/mail"
	"git.zzdats.lv/edim/api-sender/phone"
	"github.com/nobid-lsp-latvia/lx-go-jsondb"

	"azugo.io/azugo"
	"azugo.io/azugo/server"
	"azugo.io/core/instrumenter"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
)

// App is the application instance.
type App struct {
	*azugo.App

	config      *Configuration
	store       jsondb.Store
	db          *pgxpool.Pool
	phoneClient interfaces.SenderService
	mailClient  interfaces.SenderService
}

// New returns a new application instance.
func New(cmd *cobra.Command, version string) (*App, error) {
	config := NewConfiguration()

	a, err := server.New(cmd, server.Options{
		AppName:       "Sender API service for e-mail and SMS",
		AppVer:        version,
		Configuration: config,
	})
	if err != nil {
		return nil, err
	}

	store, db, err := jsondb.New(a.App, config.Postgres)
	if err != nil {
		return nil, err
	}

	phoneClient, err := phone.NewService(a.App, config.Phone, store)
	if err != nil {
		return nil, err
	}

	mailClient, err := mail.NewService(a.App, config.Mail, store)
	if err != nil {
		return nil, err
	}

	app := &App{
		App:         a,
		config:      config,
		store:       store,
		db:          db,
		phoneClient: phoneClient,
		mailClient:  mailClient,
	}

	app.Instrumentation(instrumenter.CombinedInstrumenter(app.Instrumenter(), app.storeInstrumenter()))

	return app, nil
}

// Start starts the application.
func (a *App) Start() error {
	if err := a.Store().Start(a.BackgroundContext()); err != nil {
		return err
	}

	return a.App.Start()
}

// Config returns application configuration.
//
// Panics if configuration is not loaded.
func (a *App) Config() *Configuration {
	if a.config == nil || !a.config.Ready() {
		panic("configuration is not loaded")
	}

	return a.config
}

// Client returns phone client.
func (a *App) PhoneClient() interfaces.SenderService {
	return a.phoneClient
}

// Client returns mail client.
func (a *App) MailClient() interfaces.SenderService {
	return a.mailClient
}

func (a *App) Store() jsondb.Store {
	return a.store
}
