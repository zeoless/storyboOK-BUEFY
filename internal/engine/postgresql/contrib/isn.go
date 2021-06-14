// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/catalog"
)

var funcsIsn = []*catalog.Function{
	{
		Name: "btean13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btean13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btean13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btean13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btean13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btean13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btean13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btean13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btisbn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btisbn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btisbn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btisbncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btisbncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btisbncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btismn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btismn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btismn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btismncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btismncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btismncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btissn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btissn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btissn13cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btissncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btissncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btissncmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btupccmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "btupccmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "ean13_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ean13"},
	},
	{
		Name: "ean13_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "ean13_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "ean13_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "ean13_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "hashean13",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hashisbn",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hashisbn13",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hashismn",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hashismn13",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hashissn",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hashissn13",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "hashupc",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "is_valid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "is_valid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "is_valid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "is_valid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "is_valid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "is_valid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "is_valid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "is_valid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isbn",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "isbn"},
	},
	{
		Name: "isbn13",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "isbn13"},
	},
	{
		Name: "isbn13_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "isbn13"},
	},
	{
		Name: "isbn_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "isbn"},
	},
	{
		Name: "ismn",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ismn"},
	},
	{
		Name: "ismn13",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ismn13"},
	},
	{
		Name: "ismn13_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ismn13"},
	},
	{
		Name: "ismn_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ismn"},
	},
	{
		Name: "isn_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "isn_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "isn_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "isn_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name:       "isn_weak",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isn_weak",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "boolean"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isneq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isngt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "upc"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "issn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
			{
				Type: &ast.TypeName{Name: "upc"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
			{
				Type: &ast.TypeName{Name: "isbn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ean13"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ismn13"},
			},
			{
				Type: &ast.TypeName{Name: "ismn"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "isnle",
		Args: []*catalog.Argument{
	