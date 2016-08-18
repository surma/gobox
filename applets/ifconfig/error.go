package ifconfig

import (
	"flag"

	"errors"
)

var (
	ErrInvalidIfaceName     = errors.New("Invalid interface name")
	ErrNoPortAvailable      = errors.New("Could not find a free port")
	ErrInvalidAddressFormat = errors.New("Invalid address format")
	ErrInvalidState         = errors.New("Invalid state")
)
