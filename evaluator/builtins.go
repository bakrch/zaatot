package evaluator

import (
	"fmt"
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				message := fmt.Sprintf("Wrong number of arguments. Got=%d, want=%d", len(args), 1)
				return &object.Error{
					Message: message,
				}
			}
			str, ok := args[0].(*object.String)
			if !ok {
				message := fmt.Sprintf("Argument to `len` not supported, got %T", str)
				return &object.Error{
					Message: message,
				}
			}
			return &object.Integer{Value: int64(len(str.Value))}
		},
	},
}
