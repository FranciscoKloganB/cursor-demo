//go:build tools

package tools

// Import the packages required for development dependencies.
// These imports ensure that the tools are included in the go.mod file.

import (
	// Development
	_ "github.com/sqlc-dev/sqlc" // SQL code generator

	// Testing
	_ "github.com/gavv/httpexpect/v2" // E2E HTTP and REST API testing
	_ "github.com/go-resty/resty/v2"  // Simple HTTP and REST client library for Go
	_ "github.com/stretchr/testify"   // Assertions API, Mocking API, plus test suite interfaces and functions
)
