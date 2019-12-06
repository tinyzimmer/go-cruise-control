package gocc

import (
	ccclient "github.com/tinyzimmer/go-cruise-control/pkg/client"
	"github.com/tinyzimmer/go-cruise-control/pkg/types"
)

func New(opts *types.ClientOptions) (ccclient.CruiseControlClient, error) {
	return ccclient.New(opts)
}
