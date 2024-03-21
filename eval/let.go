package eval

import (
	"errors"

	"github.com/killthebuddh4/gecko/types"
)

var Let types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	name, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New(":: Let :: identifier is not a string")
	}

	if len(arguments) < 2 {
		return nil, errors.New(":: Let :: no value to define")
	}

	value := arguments[len(arguments)-1]

	err := types.EditName(scope.Parent, name, value)

	if err != nil {
		return nil, err
	}

	return value, nil
}
