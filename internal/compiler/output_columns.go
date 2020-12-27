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

				if err := findColumnForRef(ref, tables, 