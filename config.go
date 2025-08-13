// SPDX-License-Identifier: EUPL-1.2

package api

import (
	"git.zzdats.lv/edim/api-sender/mail"
	"git.zzdats.lv/edim/api-sender/phone"
	"github.com/nobid-lsp-latvia/lx-go-jsondb"

	"azugo.io/azugo/config"
	"azugo.io/core/validation"
	"github.com/nobid-lsp-latvia/go-idauth"
	"github.com/spf13/viper"
)

// Configuration represents the configuration for the application.
type Configuration struct {
	*config.Configuration `mapstructure:",squash"`

	IDAuth   *idauth.Configuration `mapstruct:"idauth"`
	Postgres *jsondb.Configuration `mapstructure:"postgres"`
	Phone    *phone.Configuration  `mapstructure:"phone"`
	Mail     *mail.Configuration   `mapstructure:"mail"`
}

// NewConfiguration returns a new configuration.
func NewConfiguration() *Configuration {
	return &Configuration{
		Configuration: config.New(),
	}
}

// Core returns the core configuration.
func (c *Configuration) ServerCore() *config.Configuration {
	return c.Configuration
}

// Bind configuration to viper.
func (c *Configuration) Bind(_ string, v *viper.Viper) {
	c.Configuration.Bind("", v)
	c.IDAuth = config.Bind(c.IDAuth, "idauth", v)
	c.Postgres = config.Bind(c.Postgres, "postgres", v)
	c.Phone = config.Bind(c.Phone, "phone", v)
	c.Mail = config.Bind(c.Mail, "mail", v)
}

// Validate application configuration.
func (c *Configuration) Validate(validate *validation.Validate) error {
	if err := c.IDAuth.Validate(validate); err != nil {
		return err
	}

	if err := c.Postgres.Validate(validate); err != nil {
		return err
	}

	if err := c.Phone.Validate(validate); err != nil {
		return err
	}

	if err := c.Mail.Validate(validate); err != nil {
		return err
	}

	return nil
}
