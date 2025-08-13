// SPDX-License-Identifier: EUPL-1.2

package mail

import (
	"azugo.io/core/config"
	"azugo.io/core/validation"
	"github.com/spf13/viper"
)

type Configuration struct {
	MailHost       string `mapstructure:"mail_host" validate:"required"`
	MailPort       int    `mapstructure:"mail_port" validate:"required"`
	MailPassword   string `mapstructure:"mail_password"`
	MailUser       string `mapstructure:"mail_user"`
	MailSSL        bool   `mapstructure:"mail_ssl"`
	MailSkipVerify bool   `mapstructure:"mail_skip_verify"`
	SenderMail     string `mapstructure:"sender_mail" validate:"email"`
	SenderMailName string `mapstructure:"sender_mail_name"`
}

func (c *Configuration) Bind(prefix string, v *viper.Viper) {
	psw, _ := config.LoadRemoteSecret("MAIL_PASSWORD")

	v.SetDefault(prefix+".mail_password", psw)
	v.SetDefault(prefix+".mail_port", 25)
	v.SetDefault(prefix+".mail_ssl", false)
	v.SetDefault(prefix+".mail_skip_verify", false)

	_ = v.BindEnv(prefix+".mail_password", "MAIL_PASSWORD")
	_ = v.BindEnv(prefix+".mail_port", "MAIL_PORT")
	_ = v.BindEnv(prefix+".mail_host", "MAIL_HOST")
	_ = v.BindEnv(prefix+".mail_user", "MAIL_USER")
	_ = v.BindEnv(prefix+".mail_ssl", "MAIL_SSL")
	_ = v.BindEnv(prefix+".mail_skip_verify", "MAIL_SKIP_VERIFY")
	_ = v.BindEnv(prefix+".sender_mail", "SENDER_MAIL")
	_ = v.BindEnv(prefix+".sender_mail_name", "SENDER_MAIL_NAME")
}

// Validate mail configuration section.
func (c *Configuration) Validate(valid *validation.Validate) error {
	return valid.Struct(c)
}
