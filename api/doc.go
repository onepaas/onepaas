// Package api OnePaaS API
//
// The purpose of OnePaaS is to provide an application that is using your infrastructure easily.
//
// Terms Of Service:
// There are no TOS at this moment, use at your own risk we take no responsibility.
//
//     Version: unknown-version
//     License: Apache 2.0 https://www.apache.org/licenses/LICENSE-2.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import _ "embed"

//go:embed swagger.json
// SwaggerJson represents swagger.json
var SwaggerJson []byte
