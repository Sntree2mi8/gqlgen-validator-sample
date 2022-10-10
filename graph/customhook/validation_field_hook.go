package customhook

import (
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/vektah/gqlparser/v2/ast"
)

func ValidationFieldHook(td *ast.Definition, fd *ast.FieldDefinition, f *modelgen.Field) (*modelgen.Field, error) {
	c := fd.Directives.ForName("validation")
	if c != nil {
		formatConstraint := c.Arguments.ForName("format")
		if formatConstraint != nil {
			f.Tag += " validate:" + formatConstraint.Value.String()
		}
	}
	return f, nil
}
