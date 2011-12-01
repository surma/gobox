package ifconfig

import (
	"errors"
)

var (
	ErrInvalidIfaceName = errors.New("Invalid interface name")
	ErrNoPortAvailable  = errors.New("Could not find a free port")
)
