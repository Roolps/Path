package path

import "errors"

type PathError error

var (
	ErrorIPv4Network error = errors.New("invalid address provided")
	ErrorIPv6Network error = errors.New("invalid address provided")
)

var pathErrors = map[string]PathError{
	"value_error.ipv4network": ErrorIPv4Network,
	"value_error.ipv6network": ErrorIPv6Network,
}
