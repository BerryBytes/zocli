package middleware

import (
	"github.com/berrybytes/zocli/pkg/utils/factory"
)

// LoggedIn
//
// this function check if the user is logged in prior to commands being executed.
// i.e. this function will be called at PreRun function,
func LoggedIn(f *factory.Factory) {
	if !f.LoggedIn {
		f.Printer.Fatal(1, "not logged in")
		return
	}
}
