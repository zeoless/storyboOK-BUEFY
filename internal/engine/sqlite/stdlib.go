package sqlite

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/catalog"
)

// sqlite functions from:
// 		 https://www.sqlite.org/lang_aggfunc.html
// 		 https://www.sqlite.org/lang_mathfunc.html
//		 https://www.sqlite.org/lang_corefunc.html

func defaultSchema(name string) *catalog.Schema {
	s := &catalog.Schema{Name: name}
	s.Funcs = []*catalog.Function{
		// Aggregation Functions
		{
			Name: "AVG",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "real"},
			ReturnTypeNullable: true,
		},
		{
			Name:       "COUNT",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "COUNT",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "GROUP_CONCAT",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "GROUP_CONCAT",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "MAX",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "MIN",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "SUM",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "real"},
			ReturnTypeNullable: true,
		},
		{
			Name: "TOTAL",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},

		// Math Functions
		{
			Name: "ACOS",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ACOSH",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ASIN",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ASINH",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ATAN",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ATAN2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ATANH",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "CEIL",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "CEILING",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "COS",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "COSH",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "DEGREES",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "EXP",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "FLOOR",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "LN",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "LOG",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "LOG10",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "LOG",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "LOG2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "MOD",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name:       "PI",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "POW",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "POWER",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "RADIANS",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "SIN",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "SINH",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "SQRT",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "TAN",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "TANH",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "TRUNC",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},

		// Scalar functions
		{
			Name: "ABS",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name:       "CHANGES",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "CHAR",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "int"},
				},
				{
					Type: &ast.TypeName{Name: "int"},
					Mode: ast.FuncParamVariadic,
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "COALESCE",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
					Mode: ast.FuncParamVariadic,
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "FORMAT",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
					Mode: ast.FuncParamVariadic,
				},
			},
			ReturnType:         &ast.TypeName{Name: "text"},
			ReturnTypeNullable: true,
		},
		{
			Name: "GLOB",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "HEX",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "IFNULL",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "IIF",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "INSTR",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "integer"},
			ReturnTypeNullable: true,
		},
		{
			Name:       "LAST_INSERT_ROWID",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "LENGTH",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "integer"},
			ReturnTypeNullable: true,
		},
		{
			Name: "LIKE",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "LIKE",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "LIKELIHOOD",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "LIKELY",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "LOWER",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "LTRIM",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "LTRIM",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "MAX",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
					Mode: ast.FuncParamVariadic,
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "MIN",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
					Mode: ast.FuncParamVariadic,
				},
			},
			ReturnType:         &ast.TypeName{Name: "any"},
			ReturnTypeNullable: true,
		},
		{
			Name: "NULLIF",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType:         &ast.TypeName{