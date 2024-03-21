package eval

import (
	"github.com/killthebuddh4/gecko/types"
)

var Value types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	if len(arguments) == 0 {
		return nil, nil
	}

	return arguments[len(arguments)-1], nil
}
