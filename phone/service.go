// SPDX-License-Identifier: EUPL-1.2

package phone

import (
	"git.zzdats.lv/edim/api-sender/interfaces"
	"github.com/nobid-lsp-latvia/lx-go-jsondb"

	"azugo.io/core"
)

func NewService(app *core.App, config *Configuration, store jsondb.Store) (interfaces.SenderService, error) {
	return newPhoneService(app, config, store)
}
