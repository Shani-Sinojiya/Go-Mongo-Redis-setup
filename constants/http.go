// Package constants provides global constants and shared instances for the application.
package constants

import "accounts.workflecks.com/http"

var (
	HTTPSuccess = http.NewSuccess()
	HTTPErrors  = http.NewHTTPErrors()
)
