// SPDX-License-Identifier: EUPL-1.2

package phone

import (
	"azugo.io/core/config"
	"azugo.io/core/validation"
	"github.com/spf13/viper"
)

type Configuration struct {
	PhoneAPIKey     string `mapstructure:"phone_api_key"`
	PhoneURL        string `mapstructure:"phone_client_url"`
	SenderPhoneName string `mapstructure:"sender_phone_name"`
	Debug           bool   `mapstructure:"phone_debug"`
}

func (c *Configuration) Bind(prefix string, v *viper.Viper) {
	apiKey, _ := config.LoadRemoteSecret("PHONE_API_KEY")

	v.SetDefault(prefix+".phone_api_key", apiKey)

	_ = v.BindEnv(prefix+".phone_api_key", "PHONE_API_KEY")
	_ = v.BindEnv(prefix+".phone_client_url", "PHONE_CLIENT_URL")
	_ = v.BindEnv(prefix+".sender_phone_name", "SENDER_PHONE_NAME")
	_ = v.BindEnv(prefix+".phone_debug", "PHONE_DEBUG")
}

// Validate phone configuration section.
func (c *Configuration) Validate(valid *validation.Validate) error {
	return valid.Struct(c)
}
