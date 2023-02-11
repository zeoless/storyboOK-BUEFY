package sqlite

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/catalog"
)

// sqlite functions from:
// 		 https://www.sqlite.org/lang_aggfunc.html
// 		 https://www.sq