package eval

import (
	"github.com/killthebuddh4/gecko/types"
)

var SchemaIdentity types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	var schema types.Closure = func(context *types.Trajectory, arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		return raw, nil
	}

	return schema, nil
}
