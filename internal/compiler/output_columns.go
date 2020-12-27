package compiler

import (
	"errors"
	"fmt"

	"github.com/kyleconroy/sqlc/internal/sql/catalog"

	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/astutils"
	"github.com/kyleconroy/sqlc/internal/sql/lang"
	"github.com/kyleconroy/sqlc/internal/sql/sqlerr"
)

// OutputColumns determines which columns a statement will output
func (c *Compiler) OutputColumns(stmt ast.Node) ([]*catalog.Column, error) {
	qc, err := buildQueryCatalog(c.catalog, stmt)
	if err != nil {
		return nil, err
	}
	cols, err := outputColumns(qc, stmt)
	if err != nil {
		return nil, err
	}

	catCols := make([]*catalog.Column, 0, len(cols))
	for _, col := range cols {
		catCols = append(catCols, &catalog.Column{
			Name:      col.Name,
			Type:      ast.TypeName{Name: col.DataType},
			IsNotNull: col.NotNull,
			IsArray:   col.IsArray,
			Comment:   col.Comment,
			Length:    col.Length,
		})
	}
	return catCols, nil
}

func hasStarRef(cf *ast.ColumnRef) bool {
	for _, item := range cf.Fields.Items {
		if _, ok := item.(*ast.A_Star); ok {
			return true
		}
	}
	return false
}

// Compute the output columns for a statement.
//
// Return an error if column references are ambiguous
// Return an error if column references don't exist
func outputColumns(qc *QueryCatalog, node ast.Node) ([]*Column, error) {
	tables, err := sourceTables(qc, node)
	if err != nil {
		return nil, err
	}

	var targets *ast.List
	switch n := node.(type) {
	case *ast.DeleteStmt:
		targets = n.ReturningList
	case *ast.InsertStmt:
		targets = n.ReturningList
	case *ast.SelectStmt:
		targets = n.TargetList

		if n.GroupClause != nil {
			for _, item := range n.GroupClause.Items {
				ref, ok := item.(*ast.ColumnRef)
				if !ok {
					continue
				}

				if err := findColumnForRef(ref, tables, n); err != nil {
					return nil, err
				}
			}
		}

		// For UNION queries, targets is empty and we need to look for the
		// columns in Largs.
		if len(targets.Items) == 0 && n.Larg != nil {
			return outputColumns(qc, n.Larg)
		}
	case *ast.CallStmt:
		targets = &ast.List{}
	case *ast.TruncateStmt:
		targets = &ast.List{}
	case *ast.UpdateStmt:
		targets = n.ReturningList
	default:
		return nil, fmt.Errorf("outputColumns: unsupported node type: %T", n)
	}

	var cols []*Column

	for _, target := range targets.Items {
		res, ok := target.(*ast.ResTarget)
		if !ok {
			continue
		}
		switch n := res.Val.(type) {

		case *ast.A_Expr:
			name := ""
			if res.Name != nil {
				name = *res.Name
			}
			switch {
			case lang.IsComparisonOperator(astutils.Join(n.Name, "")):
				// TODO: Generate a name for these operations
				cols = append(cols, &Column{Name: name, DataType: "bool", NotNull: true})
			case lang.IsMathematicalOperator(astutils.Join(n.Name, "")):
				cols = append(cols, &Column{Name: name, DataType: "int", NotNull: true})
			default:
				cols = append(cols, &Column{Name: name, DataType: "any", NotNull: false})
			}

		case *ast.BoolExpr:
			name := ""
			if res.Name != nil {
				name = *res.Name
			}
			notNull := false
			if n.Boolop == ast.BoolExprTypeNot && len(n.Args.Items) == 1 {
				sublink, ok := n.Args.Items[0].(*ast.SubLink)
				if ok && sublink.SubLinkType == ast.EXISTS_SUBLINK {
					notNull = true
					if name == "" {
						name = "not_exists"
					}
				}
			}
			cols = append(cols, &Column{Name: name, DataType: "bool", NotNull: notNull})

		case *ast.CaseExpr:
			name := ""
			if res.Name != nil {
				name = *res.Name
			}
			// TODO: The TypeCase code has been copied from below. Instead, we
			// need a recurse function to get the type of a node.
			if tc, ok := n.Defresult.(*ast.TypeCast); ok {
				if tc.TypeName == nil {
					return nil, errors.New("no type name type cast")
				}
				name := ""
				if ref, ok := tc.Arg.(*ast.ColumnRef); ok {
					name = astutils.Join(ref.Fields, "_")
				}
				if res.Name != nil {
					name = *res.Name
				}
				// TODO Validate column names
				col := toColumn(tc.TypeName)
				col.Name = name
				cols = append(cols, col)
			} else if aconst, ok := n.Defresult.(*ast.A_Const); ok {
				tn, err := ParseTypeName(aconst.Val)
				if err != nil {
					return nil, err
				}
				cols = append(cols, &Column{Name: name, DataType: dataType(tn), NotNull: true})
			} else {
				cols = append(cols, &Column{Name: name, DataType: "any", NotNull: false})
			}

		case *ast.CoalesceExpr:
			name := "coalesce"
			if res.Name != nil {
				name = *res.Name
			}
			var firstColumn *Column
			var shouldNotBeNull bool
			for _, arg := range n.Args.Items {
				if _, ok := arg.(*ast.A_Const); ok {
					shouldNotBeNull = true
					continue
				}
				if ref, ok := arg.(*ast.ColumnRef); ok {
					columns, err := outputColumnRefs(res, tables, ref)
					if err != nil {
						return nil, err
					}
					for _, c := range columns {
						if firstColumn == nil {
							firstColumn = c
						}
						shouldNotBeNull = shouldNotBeNull || c.NotNull
					}
				}
			}
			if firstColumn != nil {
				firstColumn.NotNull = shouldNotBeNull
				firstColumn.skipTableRequiredCheck = true
				cols = append(cols, firstColumn)
			} else {
				cols = append(cols, &Column{Name: name, DataType: "any", NotNull: false})
			}

		case *ast.ColumnRef:
			if hasStarRef(n) {
				// TODO: This code is copied in func expand()
				for _, t := range tables {
					scope := astutils.Join(n.Fields, ".")
					if scope != "" && scope != t.Rel.Name {
						continue
					}
					for _, c := range t.Columns {
						cname := c.Name
						if res.Name != nil {
							cname = *res.Name
						}
						cols = append(cols, &Column{
							Name:       cname,
							Type:       c.Type,
							Scope:      scope,
							Table:      c.Table,
							TableAlia