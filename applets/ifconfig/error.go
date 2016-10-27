package ifconfig

import (
	"flag"

	"errors"
)

var (
	flagSet                 = flag.NewFlagSet("error", flag.PanicOnError)
	ErrInvalidIfaceName     = errors.New("Invalid interface name")
	ErrNoPortAvailable      = errors.New("Could not find a free port")
	ErrInvalidAddressFormat = errors.New("Invalid address format")
	ErrInvalidState         = errors.New("Invalid state")
)
