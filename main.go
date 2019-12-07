package gocc

import (
	"github.com/tinyzimmer/go-cruise-control/pkg/client"
	"github.com/tinyzimmer/go-cruise-control/pkg/types"
)

// New creates a new CruiseControl client with the supplied options
func New(opts *types.ClientOptions) (client.CruiseControlClient, error) {
	return client.New(opts)
}
