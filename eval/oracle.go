package eval

import (
	"github.com/killthebuddh4/gecko/types"
)

var Oracle types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return "oracle", nil
}
