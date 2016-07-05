// Package internal supports the options and transport packages.
package internal

import (
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
)

// DialSettings holds information needed to establish a connection with a
// Google API service.
type DialSettings struct {
	Endpoint     string
	Scopes       []string
	TokenSource  oauth2.TokenSource
	UserAgent    string
	HTTPClient   *http.Client
	GRPCDialOpts []grpc.DialOption
	GRPCConn     *grpc.ClientConn
}
