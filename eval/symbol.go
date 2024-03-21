package eval

import (
	"github.com/killthebuddh4/gecko/types"
)

var Symbol types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return scope.Expression.Operator.Value, nil
}
