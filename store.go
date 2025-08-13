// SPDX-License-Identifier: EUPL-1.2

package api

import (
	"context"

	"azugo.io/core/instrumenter"
	"github.com/nobid-lsp-latvia/lx-go-jsondb"
)

// storeInstrumenter is a instrumenter to pass session data to database.
func (a *App) storeInstrumenter() instrumenter.Instrumenter {
	return instrumenter.Instrumenter(func(_ context.Context, op string, _ ...interface{}) func(err error) {
		if op != jsondb.InstrumentationExec {
			return func(_ error) {}
		}

		return func(_ error) {}
	})
}
