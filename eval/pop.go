package eval

import (
	"errors"

	"github.com/killthebuddh4/gecko/types"
)

var Pop types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Pop :: not an array")
	}

	return arr[:len(arr)-1], nil
}
