package eval

import (
	"errors"

	"github.com/killthebuddh4/gecko/types"
)

var Split types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	str, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New("Split :: split only accepts strings")
	}

	result := []types.Value{}

	for _, c := range str {
		result = append(result, string(c))
	}

	return result, nil
}
