package eval

import (
	"errors"

	"github.com/killthebuddh4/gecko/types"
)

type GeckoError struct {
	message string
}

var Throw types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	signal, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New(":: Throw :: signal not a string")
	}

	return GeckoError{message: signal}, nil
}
