package compiler

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/kyleconroy/sqlc/internal/config"
	"github.com/kyleconroy/sqlc/internal/debug"
	"github.com/kyleconroy/sqlc/internal/metadata"
	"github.com/kyleconroy/sqlc/internal/opts"
	"github.com/kyleconroy/sqlc/internal/source"
	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/astutils"
	"github.com/kyleconroy/sqlc/internal/sql/rewrite"
	"github.com/kyleconroy/sqlc/internal/sql/validate"
)

var ErrUnsupportedStatementType = errors.New("parseQuery: unsupported statement type")

func (c *Compiler) parseQuery(stmt ast.Node, src string, o opts.Parser) (*Query, error) {
	if o.Debug.DumpAST {
		debug.Dump(stmt)
	}
	if err := validate.ParamStyle(stmt); err != nil {
		return nil, err
	}
	numbers, dollar, err := validate.ParamRef(stmt)
	if err != nil {
		return nil, err
	}
	raw, ok := stmt.(*ast.RawStmt)
	if !ok {
		return nil, errors.New("node is not a statement")
	}
	var table *ast.TableName
	switch n := raw.Stmt.(type) {
	case *ast.CallStmt:
	case *ast.SelectStmt:
	case *ast.DeleteStmt:
	case *ast.InsertStmt:
		if err := validate.InsertStmt(n); err != nil {
			return nil, err
		}
		var err error
		table, err = ParseTableName(n.Relation)
		if err != nil {
			return nil, err
		}
	case *ast.TruncateStmt:
	case *ast.UpdateStmt:
	default:
		return nil, ErrUnsupportedStatementType
	}

	rawSQL, err := source.Pluck(src, raw.StmtLocation, raw.StmtLen)
	if err != nil {
		return nil, err
	}
	if rawSQL == "" {
		return nil, errors.New("missing semicolon at end of file")
	}
	if err := validate.FuncCall(c.catalog, c.combo, raw); err != nil {
		return nil, err
	}
	name, cmd, err := metadata.Parse(strings.TrimSpace(rawSQL), c.parser.CommentSyntax())
	if err != nil {
		return nil, err
	}
	raw, namedParams, edits := rewrite.NamedParameters(c.conf.Engine, raw, numbers, dollar)
	if err := validate.Cmd(raw.Stmt, nam