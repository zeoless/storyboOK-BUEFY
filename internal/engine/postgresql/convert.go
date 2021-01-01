
//go:build !windows && cgo
// +build !windows,cgo

package postgresql

import (
	"fmt"

	pg "github.com/pganalyze/pg_query_go/v4"

	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

func convertFuncParamMode(m pg.FunctionParameterMode) (ast.FuncParamMode, error) {
	switch m {
	case pg.FunctionParameterMode_FUNC_PARAM_IN:
		return ast.FuncParamIn, nil
	case pg.FunctionParameterMode_FUNC_PARAM_OUT:
		return ast.FuncParamOut, nil
	case pg.FunctionParameterMode_FUNC_PARAM_INOUT:
		return ast.FuncParamInOut, nil
	case pg.FunctionParameterMode_FUNC_PARAM_VARIADIC:
		return ast.FuncParamVariadic, nil
	case pg.FunctionParameterMode_FUNC_PARAM_TABLE:
		return ast.FuncParamTable, nil
	case pg.FunctionParameterMode_FUNC_PARAM_DEFAULT:
		return ast.FuncParamDefault, nil
	default:
		return -1, fmt.Errorf("parse func param: invalid mode %v", m)
	}
}

func convertSubLinkType(t pg.SubLinkType) (ast.SubLinkType, error) {
	switch t {
	case pg.SubLinkType_EXISTS_SUBLINK:
		return ast.EXISTS_SUBLINK, nil
	case pg.SubLinkType_ALL_SUBLINK:
		return ast.ALL_SUBLINK, nil
	case pg.SubLinkType_ANY_SUBLINK:
		return ast.ANY_SUBLINK, nil
	case pg.SubLinkType_ROWCOMPARE_SUBLINK:
		return ast.ROWCOMPARE_SUBLINK, nil
	case pg.SubLinkType_EXPR_SUBLINK:
		return ast.EXPR_SUBLINK, nil
	case pg.SubLinkType_MULTIEXPR_SUBLINK:
		return ast.MULTIEXPR_SUBLINK, nil
	case pg.SubLinkType_ARRAY_SUBLINK:
		return ast.ARRAY_SUBLINK, nil
	case pg.SubLinkType_CTE_SUBLINK:
		return ast.CTE_SUBLINK, nil
	default:
		return 0, fmt.Errorf("parse sublink type: unknown type %s", t)
	}
}

func convertSetOperation(t pg.SetOperation) (ast.SetOperation, error) {
	switch t {
	case pg.SetOperation_SETOP_NONE:
		return ast.None, nil
	case pg.SetOperation_SETOP_UNION:
		return ast.Union, nil
	case pg.SetOperation_SETOP_INTERSECT:
		return ast.Intersect, nil
	case pg.SetOperation_SETOP_EXCEPT:
		return ast.Except, nil
	default:
		return 0, fmt.Errorf("parse set operation: unknown type %s", t)
	}
}

func convertList(l *pg.List) *ast.List {
	out := &ast.List{}
	for _, item := range l.Items {
		out.Items = append(out.Items, convertNode(item))
	}
	return out
}

func convertSlice(nodes []*pg.Node) *ast.List {
	out := &ast.List{}
	for _, n := range nodes {
		out.Items = append(out.Items, convertNode(n))
	}
	return out
}

func convertValuesList(l [][]*pg.Node) *ast.List {
	out := &ast.List{}
	for _, outer := range l {
		o := &ast.List{}
		for _, inner := range outer {
			o.Items = append(o.Items, convertNode(inner))
		}
		out.Items = append(out.Items, o)
	}
	return out
}

func convert(node *pg.Node) (ast.Node, error) {
	return convertNode(node), nil
}

func convertA_ArrayExpr(n *pg.A_ArrayExpr) *ast.A_ArrayExpr {
	if n == nil {
		return nil
	}
	return &ast.A_ArrayExpr{
		Elements: convertSlice(n.Elements),
		Location: int(n.Location),
	}
}

func convertA_Const(n *pg.A_Const) *ast.A_Const {
	if n == nil {
		return nil
	}
	var val ast.Node
	if n.Isnull {
		val = &ast.Null{}
	} else {
		switch v := n.Val.(type) {
		case *pg.A_Const_Boolval:
			val = convertBoolean(v.Boolval)
		case *pg.A_Const_Bsval:
			val = convertBitString(v.Bsval)
		case *pg.A_Const_Fval:
			val = convertFloat(v.Fval)
		case *pg.A_Const_Ival:
			val = convertInteger(v.Ival)
		case *pg.A_Const_Sval:
			val = convertString(v.Sval)
		}
	}
	return &ast.A_Const{
		Val:      val,
		Location: int(n.Location),
	}
}

func convertA_Expr(n *pg.A_Expr) *ast.A_Expr {
	if n == nil {
		return nil
	}
	return &ast.A_Expr{
		Kind:     ast.A_Expr_Kind(n.Kind),
		Name:     convertSlice(n.Name),
		Lexpr:    convertNode(n.Lexpr),
		Rexpr:    convertNode(n.Rexpr),
		Location: int(n.Location),
	}
}

func convertA_Indices(n *pg.A_Indices) *ast.A_Indices {
	if n == nil {
		return nil
	}
	return &ast.A_Indices{
		IsSlice: n.IsSlice,
		Lidx:    convertNode(n.Lidx),
		Uidx:    convertNode(n.Uidx),
	}
}

func convertA_Indirection(n *pg.A_Indirection) *ast.A_Indirection {
	if n == nil {
		return nil
	}
	return &ast.A_Indirection{
		Arg:         convertNode(n.Arg),
		Indirection: convertSlice(n.Indirection),
	}
}

func convertA_Star(n *pg.A_Star) *ast.A_Star {
	if n == nil {
		return nil
	}
	return &ast.A_Star{}
}

func convertAccessPriv(n *pg.AccessPriv) *ast.AccessPriv {
	if n == nil {
		return nil
	}
	return &ast.AccessPriv{
		PrivName: makeString(n.PrivName),
		Cols:     convertSlice(n.Cols),
	}
}

func convertAggref(n *pg.Aggref) *ast.Aggref {
	if n == nil {
		return nil
	}
	return &ast.Aggref{
		Xpr:           convertNode(n.Xpr),
		Aggfnoid:      ast.Oid(n.Aggfnoid),
		Aggtype:       ast.Oid(n.Aggtype),
		Aggcollid:     ast.Oid(n.Aggcollid),
		Inputcollid:   ast.Oid(n.Inputcollid),
		Aggtranstype:  ast.Oid(n.Aggtranstype),
		Aggargtypes:   convertSlice(n.Aggargtypes),
		Aggdirectargs: convertSlice(n.Aggdirectargs),
		Args:          convertSlice(n.Args),
		Aggorder:      convertSlice(n.Aggorder),
		Aggdistinct:   convertSlice(n.Aggdistinct),
		Aggfilter:     convertNode(n.Aggfilter),
		Aggstar:       n.Aggstar,
		Aggvariadic:   n.Aggvariadic,
		Aggkind:       makeByte(n.Aggkind),
		Agglevelsup:   ast.Index(n.Agglevelsup),
		Aggsplit:      ast.AggSplit(n.Aggsplit),
		Location:      int(n.Location),
	}
}

func convertAlias(n *pg.Alias) *ast.Alias {
	if n == nil {
		return nil
	}
	return &ast.Alias{
		Aliasname: makeString(n.Aliasname),
		Colnames:  convertSlice(n.Colnames),
	}
}

func convertAlterCollationStmt(n *pg.AlterCollationStmt) *ast.AlterCollationStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterCollationStmt{
		Collname: convertSlice(n.Collname),
	}
}

func convertAlterDatabaseSetStmt(n *pg.AlterDatabaseSetStmt) *ast.AlterDatabaseSetStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterDatabaseSetStmt{
		Dbname:  makeString(n.Dbname),
		Setstmt: convertVariableSetStmt(n.Setstmt),
	}
}

func convertAlterDatabaseStmt(n *pg.AlterDatabaseStmt) *ast.AlterDatabaseStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterDatabaseStmt{
		Dbname:  makeString(n.Dbname),
		Options: convertSlice(n.Options),
	}
}

func convertAlterDefaultPrivilegesStmt(n *pg.AlterDefaultPrivilegesStmt) *ast.AlterDefaultPrivilegesStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterDefaultPrivilegesStmt{
		Options: convertSlice(n.Options),
		Action:  convertGrantStmt(n.Action),
	}
}

func convertAlterDomainStmt(n *pg.AlterDomainStmt) *ast.AlterDomainStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterDomainStmt{
		Subtype:   makeByte(n.Subtype),
		TypeName:  convertSlice(n.TypeName),
		Name:      makeString(n.Name),
		Def:       convertNode(n.Def),
		Behavior:  ast.DropBehavior(n.Behavior),
		MissingOk: n.MissingOk,
	}
}

func convertAlterEnumStmt(n *pg.AlterEnumStmt) *ast.AlterEnumStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterEnumStmt{
		TypeName:           convertSlice(n.TypeName),
		OldVal:             makeString(n.OldVal),
		NewVal:             makeString(n.NewVal),
		NewValNeighbor:     makeString(n.NewValNeighbor),
		NewValIsAfter:      n.NewValIsAfter,
		SkipIfNewValExists: n.SkipIfNewValExists,
	}
}

func convertAlterEventTrigStmt(n *pg.AlterEventTrigStmt) *ast.AlterEventTrigStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterEventTrigStmt{
		Trigname:  makeString(n.Trigname),
		Tgenabled: makeByte(n.Tgenabled),
	}
}

func convertAlterExtensionContentsStmt(n *pg.AlterExtensionContentsStmt) *ast.AlterExtensionContentsStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterExtensionContentsStmt{
		Extname: makeString(n.Extname),
		Action:  int(n.Action),
		Objtype: ast.ObjectType(n.Objtype),
		Object:  convertNode(n.Object),
	}
}

func convertAlterExtensionStmt(n *pg.AlterExtensionStmt) *ast.AlterExtensionStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterExtensionStmt{
		Extname: makeString(n.Extname),
		Options: convertSlice(n.Options),
	}
}

func convertAlterFdwStmt(n *pg.AlterFdwStmt) *ast.AlterFdwStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterFdwStmt{
		Fdwname:     makeString(n.Fdwname),
		FuncOptions: convertSlice(n.FuncOptions),
		Options:     convertSlice(n.Options),
	}
}

func convertAlterForeignServerStmt(n *pg.AlterForeignServerStmt) *ast.AlterForeignServerStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterForeignServerStmt{
		Servername: makeString(n.Servername),
		Version:    makeString(n.Version),
		Options:    convertSlice(n.Options),
		HasVersion: n.HasVersion,
	}
}

func convertAlterFunctionStmt(n *pg.AlterFunctionStmt) *ast.AlterFunctionStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterFunctionStmt{
		Func:    convertObjectWithArgs(n.Func),
		Actions: convertSlice(n.Actions),
	}
}

func convertAlterObjectDependsStmt(n *pg.AlterObjectDependsStmt) *ast.AlterObjectDependsStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterObjectDependsStmt{
		ObjectType: ast.ObjectType(n.ObjectType),
		Relation:   convertRangeVar(n.Relation),
		Object:     convertNode(n.Object),
		Extname:    convertString(n.Extname),
	}
}

func convertAlterObjectSchemaStmt(n *pg.AlterObjectSchemaStmt) *ast.AlterObjectSchemaStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterObjectSchemaStmt{
		ObjectType: ast.ObjectType(n.ObjectType),
		Relation:   convertRangeVar(n.Relation),
		Object:     convertNode(n.Object),
		Newschema:  makeString(n.Newschema),
		MissingOk:  n.MissingOk,
	}
}

func convertAlterOpFamilyStmt(n *pg.AlterOpFamilyStmt) *ast.AlterOpFamilyStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterOpFamilyStmt{
		Opfamilyname: convertSlice(n.Opfamilyname),
		Amname:       makeString(n.Amname),
		IsDrop:       n.IsDrop,
		Items:        convertSlice(n.Items),
	}
}

func convertAlterOperatorStmt(n *pg.AlterOperatorStmt) *ast.AlterOperatorStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterOperatorStmt{
		Opername: convertObjectWithArgs(n.Opername),
		Options:  convertSlice(n.Options),
	}
}

func convertAlterOwnerStmt(n *pg.AlterOwnerStmt) *ast.AlterOwnerStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterOwnerStmt{
		ObjectType: ast.ObjectType(n.ObjectType),
		Relation:   convertRangeVar(n.Relation),
		Object:     convertNode(n.Object),
		Newowner:   convertRoleSpec(n.Newowner),
	}
}

func convertAlterPolicyStmt(n *pg.AlterPolicyStmt) *ast.AlterPolicyStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterPolicyStmt{
		PolicyName: makeString(n.PolicyName),
		Table:      convertRangeVar(n.Table),
		Roles:      convertSlice(n.Roles),
		Qual:       convertNode(n.Qual),
		WithCheck:  convertNode(n.WithCheck),
	}
}

func convertAlterPublicationStmt(n *pg.AlterPublicationStmt) *ast.AlterPublicationStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterPublicationStmt{
		Pubname:      makeString(n.Pubname),
		Options:      convertSlice(n.Options),
		Tables:       convertSlice(n.Pubobjects),
		ForAllTables: n.ForAllTables,
		TableAction:  ast.DefElemAction(n.Action),
	}
}

func convertAlterRoleSetStmt(n *pg.AlterRoleSetStmt) *ast.AlterRoleSetStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterRoleSetStmt{
		Role:     convertRoleSpec(n.Role),
		Database: makeString(n.Database),
		Setstmt:  convertVariableSetStmt(n.Setstmt),
	}
}

func convertAlterRoleStmt(n *pg.AlterRoleStmt) *ast.AlterRoleStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterRoleStmt{
		Role:    convertRoleSpec(n.Role),
		Options: convertSlice(n.Options),
		Action:  int(n.Action),
	}
}

func convertAlterSeqStmt(n *pg.AlterSeqStmt) *ast.AlterSeqStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterSeqStmt{
		Sequence:    convertRangeVar(n.Sequence),
		Options:     convertSlice(n.Options),
		ForIdentity: n.ForIdentity,
		MissingOk:   n.MissingOk,
	}
}

func convertAlterSubscriptionStmt(n *pg.AlterSubscriptionStmt) *ast.AlterSubscriptionStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterSubscriptionStmt{
		Kind:        ast.AlterSubscriptionType(n.Kind),
		Subname:     makeString(n.Subname),
		Conninfo:    makeString(n.Conninfo),
		Publication: convertSlice(n.Publication),
		Options:     convertSlice(n.Options),
	}
}

func convertAlterSystemStmt(n *pg.AlterSystemStmt) *ast.AlterSystemStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterSystemStmt{
		Setstmt: convertVariableSetStmt(n.Setstmt),
	}
}

func convertAlterTSConfigurationStmt(n *pg.AlterTSConfigurationStmt) *ast.AlterTSConfigurationStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterTSConfigurationStmt{
		Kind:      ast.AlterTSConfigType(n.Kind),
		Cfgname:   convertSlice(n.Cfgname),
		Tokentype: convertSlice(n.Tokentype),
		Dicts:     convertSlice(n.Dicts),
		Override:  n.Override,
		Replace:   n.Replace,
		MissingOk: n.MissingOk,
	}
}

func convertAlterTSDictionaryStmt(n *pg.AlterTSDictionaryStmt) *ast.AlterTSDictionaryStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterTSDictionaryStmt{
		Dictname: convertSlice(n.Dictname),
		Options:  convertSlice(n.Options),
	}
}

func convertAlterTableCmd(n *pg.AlterTableCmd) *ast.AlterTableCmd {
	if n == nil {
		return nil
	}
	def := convertNode(n.Def)
	columnDef := def.(*ast.ColumnDef)
	return &ast.AlterTableCmd{
		Subtype:   ast.AlterTableType(n.Subtype),
		Name:      makeString(n.Name),
		Newowner:  convertRoleSpec(n.Newowner),
		Def:       columnDef,
		Behavior:  ast.DropBehavior(n.Behavior),
		MissingOk: n.MissingOk,
	}
}

func convertAlterTableMoveAllStmt(n *pg.AlterTableMoveAllStmt) *ast.AlterTableMoveAllStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterTableMoveAllStmt{
		OrigTablespacename: makeString(n.OrigTablespacename),
		Objtype:            ast.ObjectType(n.Objtype),
		Roles:              convertSlice(n.Roles),
		NewTablespacename:  makeString(n.NewTablespacename),
		Nowait:             n.Nowait,
	}
}

func convertAlterTableSpaceOptionsStmt(n *pg.AlterTableSpaceOptionsStmt) *ast.AlterTableSpaceOptionsStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterTableSpaceOptionsStmt{
		Tablespacename: makeString(n.Tablespacename),
		Options:        convertSlice(n.Options),
		IsReset:        n.IsReset,
	}
}

func convertAlterTableStmt(n *pg.AlterTableStmt) *ast.AlterTableStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterTableStmt{
		Relation:  convertRangeVar(n.Relation),
		Cmds:      convertSlice(n.Cmds),
		Relkind:   ast.ObjectType(n.Objtype),
		MissingOk: n.MissingOk,
	}
}

func convertAlterUserMappingStmt(n *pg.AlterUserMappingStmt) *ast.AlterUserMappingStmt {
	if n == nil {
		return nil
	}
	return &ast.AlterUserMappingStmt{
		User:       convertRoleSpec(n.User),
		Servername: makeString(n.Servername),
		Options:    convertSlice(n.Options),
	}
}

func convertAlternativeSubPlan(n *pg.AlternativeSubPlan) *ast.AlternativeSubPlan {
	if n == nil {
		return nil
	}
	return &ast.AlternativeSubPlan{
		Xpr:      convertNode(n.Xpr),
		Subplans: convertSlice(n.Subplans),
	}
}

func convertArrayCoerceExpr(n *pg.ArrayCoerceExpr) *ast.ArrayCoerceExpr {
	if n == nil {
		return nil
	}
	return &ast.ArrayCoerceExpr{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Resulttype:   ast.Oid(n.Resulttype),
		Resulttypmod: n.Resulttypmod,
		Resultcollid: ast.Oid(n.Resultcollid),
		Coerceformat: ast.CoercionForm(n.Coerceformat),
		Location:     int(n.Location),
	}
}

func convertArrayExpr(n *pg.ArrayExpr) *ast.ArrayExpr {
	if n == nil {
		return nil
	}
	return &ast.ArrayExpr{
		Xpr:           convertNode(n.Xpr),
		ArrayTypeid:   ast.Oid(n.ArrayTypeid),
		ArrayCollid:   ast.Oid(n.ArrayCollid),
		ElementTypeid: ast.Oid(n.ElementTypeid),
		Elements:      convertSlice(n.Elements),
		Multidims:     n.Multidims,
		Location:      int(n.Location),
	}
}

func convertBitString(n *pg.BitString) *ast.BitString {
	if n == nil {
		return nil
	}
	return &ast.BitString{
		Str: n.Bsval,
	}
}

func convertBoolExpr(n *pg.BoolExpr) *ast.BoolExpr {
	if n == nil {
		return nil
	}
	return &ast.BoolExpr{
		Xpr:      convertNode(n.Xpr),
		Boolop:   ast.BoolExprType(n.Boolop),
		Args:     convertSlice(n.Args),
		Location: int(n.Location),
	}
}

func convertBoolean(n *pg.Boolean) *ast.Boolean {
	if n == nil {
		return nil
	}
	return &ast.Boolean{
		Boolval: n.Boolval,
	}
}

func convertBooleanTest(n *pg.BooleanTest) *ast.BooleanTest {
	if n == nil {
		return nil
	}
	return &ast.BooleanTest{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Booltesttype: ast.BoolTestType(n.Booltesttype),
		Location:     int(n.Location),
	}
}

func convertCallStmt(n *pg.CallStmt) *ast.CallStmt {
	if n == nil {
		return nil
	}
	rel, err := parseRelationFromNodes(n.Funccall.Funcname)
	if err != nil {
		// TODO: How should we handle errors?
		panic(err)
	}

	return &ast.CallStmt{
		FuncCall: &ast.FuncCall{
			Func:           rel.FuncName(),
			Funcname:       convertSlice(n.Funccall.Funcname),
			Args:           convertSlice(n.Funccall.Args),
			AggOrder:       convertSlice(n.Funccall.AggOrder),
			AggFilter:      convertNode(n.Funccall.AggFilter),
			AggWithinGroup: n.Funccall.AggWithinGroup,
			AggStar:        n.Funccall.AggStar,
			AggDistinct:    n.Funccall.AggDistinct,
			FuncVariadic:   n.Funccall.FuncVariadic,
			Over:           convertWindowDef(n.Funccall.Over),
			Location:       int(n.Funccall.Location),
		},
	}
}

func convertCaseExpr(n *pg.CaseExpr) *ast.CaseExpr {
	if n == nil {
		return nil
	}
	return &ast.CaseExpr{
		Xpr:        convertNode(n.Xpr),
		Casetype:   ast.Oid(n.Casetype),
		Casecollid: ast.Oid(n.Casecollid),
		Arg:        convertNode(n.Arg),
		Args:       convertSlice(n.Args),
		Defresult:  convertNode(n.Defresult),
		Location:   int(n.Location),
	}
}

func convertCaseTestExpr(n *pg.CaseTestExpr) *ast.CaseTestExpr {
	if n == nil {
		return nil
	}
	return &ast.CaseTestExpr{
		Xpr:       convertNode(n.Xpr),
		TypeId:    ast.Oid(n.TypeId),
		TypeMod:   n.TypeMod,
		Collation: ast.Oid(n.Collation),
	}
}

func convertCaseWhen(n *pg.CaseWhen) *ast.CaseWhen {
	if n == nil {
		return nil
	}
	return &ast.CaseWhen{
		Xpr:      convertNode(n.Xpr),
		Expr:     convertNode(n.Expr),
		Result:   convertNode(n.Result),
		Location: int(n.Location),
	}
}

func convertCheckPointStmt(n *pg.CheckPointStmt) *ast.CheckPointStmt {
	if n == nil {
		return nil
	}
	return &ast.CheckPointStmt{}
}

func convertClosePortalStmt(n *pg.ClosePortalStmt) *ast.ClosePortalStmt {
	if n == nil {
		return nil
	}
	return &ast.ClosePortalStmt{
		Portalname: makeString(n.Portalname),
	}
}

func convertClusterStmt(n *pg.ClusterStmt) *ast.ClusterStmt {
	if n == nil {
		return nil
	}
	return &ast.ClusterStmt{
		Relation:  convertRangeVar(n.Relation),
		Indexname: makeString(n.Indexname),
	}
}

func convertCoalesceExpr(n *pg.CoalesceExpr) *ast.CoalesceExpr {
	if n == nil {
		return nil
	}
	return &ast.CoalesceExpr{
		Xpr:            convertNode(n.Xpr),
		Coalescetype:   ast.Oid(n.Coalescetype),
		Coalescecollid: ast.Oid(n.Coalescecollid),
		Args:           convertSlice(n.Args),
		Location:       int(n.Location),
	}
}

func convertCoerceToDomain(n *pg.CoerceToDomain) *ast.CoerceToDomain {
	if n == nil {
		return nil
	}
	return &ast.CoerceToDomain{
		Xpr:            convertNode(n.Xpr),
		Arg:            convertNode(n.Arg),
		Resulttype:     ast.Oid(n.Resulttype),
		Resulttypmod:   n.Resulttypmod,
		Resultcollid:   ast.Oid(n.Resultcollid),
		Coercionformat: ast.CoercionForm(n.Coercionformat),
		Location:       int(n.Location),
	}
}

func convertCoerceToDomainValue(n *pg.CoerceToDomainValue) *ast.CoerceToDomainValue {
	if n == nil {
		return nil
	}
	return &ast.CoerceToDomainValue{
		Xpr:       convertNode(n.Xpr),
		TypeId:    ast.Oid(n.TypeId),
		TypeMod:   n.TypeMod,
		Collation: ast.Oid(n.Collation),
		Location:  int(n.Location),
	}
}

func convertCoerceViaIO(n *pg.CoerceViaIO) *ast.CoerceViaIO {
	if n == nil {
		return nil
	}
	return &ast.CoerceViaIO{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Resulttype:   ast.Oid(n.Resulttype),
		Resultcollid: ast.Oid(n.Resultcollid),
		Coerceformat: ast.CoercionForm(n.Coerceformat),
		Location:     int(n.Location),
	}
}

func convertCollateClause(n *pg.CollateClause) *ast.CollateClause {
	if n == nil {
		return nil
	}
	return &ast.CollateClause{
		Arg:      convertNode(n.Arg),
		Collname: convertSlice(n.Collname),
		Location: int(n.Location),
	}
}

func convertCollateExpr(n *pg.CollateExpr) *ast.CollateExpr {
	if n == nil {
		return nil
	}
	return &ast.CollateExpr{
		Xpr:      convertNode(n.Xpr),
		Arg:      convertNode(n.Arg),
		CollOid:  ast.Oid(n.CollOid),
		Location: int(n.Location),
	}
}

func convertColumnDef(n *pg.ColumnDef) *ast.ColumnDef {
	if n == nil {
		return nil
	}
	return &ast.ColumnDef{
		Colname:       n.Colname,
		TypeName:      convertTypeName(n.TypeName),
		Inhcount:      int(n.Inhcount),
		IsLocal:       n.IsLocal,
		IsNotNull:     n.IsNotNull,
		IsFromType:    n.IsFromType,
		Storage:       makeByte(n.Storage),
		RawDefault:    convertNode(n.RawDefault),
		CookedDefault: convertNode(n.CookedDefault),
		Identity:      makeByte(n.Identity),
		CollClause:    convertCollateClause(n.CollClause),
		CollOid:       ast.Oid(n.CollOid),
		Constraints:   convertSlice(n.Constraints),
		Fdwoptions:    convertSlice(n.Fdwoptions),
		Location:      int(n.Location),
	}
}

func convertColumnRef(n *pg.ColumnRef) *ast.ColumnRef {
	if n == nil {
		return nil
	}
	return &ast.ColumnRef{
		Fields:   convertSlice(n.Fields),
		Location: int(n.Location),
	}
}

func convertCommentStmt(n *pg.CommentStmt) *ast.CommentStmt {
	if n == nil {
		return nil
	}
	return &ast.CommentStmt{
		Objtype: ast.ObjectType(n.Objtype),
		Object:  convertNode(n.Object),
		Comment: makeString(n.Comment),
	}
}

func convertCommonTableExpr(n *pg.CommonTableExpr) *ast.CommonTableExpr {
	if n == nil {
		return nil
	}
	return &ast.CommonTableExpr{
		Ctename:          makeString(n.Ctename),
		Aliascolnames:    convertSlice(n.Aliascolnames),
		Ctequery:         convertNode(n.Ctequery),
		Location:         int(n.Location),
		Cterecursive:     n.Cterecursive,
		Cterefcount:      int(n.Cterefcount),
		Ctecolnames:      convertSlice(n.Ctecolnames),
		Ctecoltypes:      convertSlice(n.Ctecoltypes),
		Ctecoltypmods:    convertSlice(n.Ctecoltypmods),
		Ctecolcollations: convertSlice(n.Ctecolcollations),
	}
}

func convertCompositeTypeStmt(n *pg.CompositeTypeStmt) *ast.CompositeTypeStmt {
	if n == nil {
		return nil
	}
	rel := parseRelationFromRangeVar(n.Typevar)
	return &ast.CompositeTypeStmt{
		TypeName: rel.TypeName(),
	}
}

func convertConstraint(n *pg.Constraint) *ast.Constraint {
	if n == nil {
		return nil
	}
	return &ast.Constraint{
		Contype:        ast.ConstrType(n.Contype),
		Conname:        makeString(n.Conname),
		Deferrable:     n.Deferrable,
		Initdeferred:   n.Initdeferred,
		Location:       int(n.Location),
		IsNoInherit:    n.IsNoInherit,
		RawExpr:        convertNode(n.RawExpr),
		CookedExpr:     makeString(n.CookedExpr),
		GeneratedWhen:  makeByte(n.GeneratedWhen),
		Keys:           convertSlice(n.Keys),
		Exclusions:     convertSlice(n.Exclusions),
		Options:        convertSlice(n.Options),
		Indexname:      makeString(n.Indexname),
		Indexspace:     makeString(n.Indexspace),
		AccessMethod:   makeString(n.AccessMethod),
		WhereClause:    convertNode(n.WhereClause),
		Pktable:        convertRangeVar(n.Pktable),
		FkAttrs:        convertSlice(n.FkAttrs),
		PkAttrs:        convertSlice(n.PkAttrs),
		FkMatchtype:    makeByte(n.FkMatchtype),
		FkUpdAction:    makeByte(n.FkUpdAction),
		FkDelAction:    makeByte(n.FkDelAction),
		OldConpfeqop:   convertSlice(n.OldConpfeqop),
		OldPktableOid:  ast.Oid(n.OldPktableOid),
		SkipValidation: n.SkipValidation,
		InitiallyValid: n.InitiallyValid,
	}
}

func convertConstraintsSetStmt(n *pg.ConstraintsSetStmt) *ast.ConstraintsSetStmt {
	if n == nil {
		return nil
	}
	return &ast.ConstraintsSetStmt{
		Constraints: convertSlice(n.Constraints),
		Deferred:    n.Deferred,
	}
}

func convertConvertRowtypeExpr(n *pg.ConvertRowtypeExpr) *ast.ConvertRowtypeExpr {
	if n == nil {
		return nil
	}
	return &ast.ConvertRowtypeExpr{
		Xpr:           convertNode(n.Xpr),
		Arg:           convertNode(n.Arg),
		Resulttype:    ast.Oid(n.Resulttype),
		Convertformat: ast.CoercionForm(n.Convertformat),
		Location:      int(n.Location),
	}
}

func convertCopyStmt(n *pg.CopyStmt) *ast.CopyStmt {
	if n == nil {
		return nil
	}
	return &ast.CopyStmt{
		Relation:  convertRangeVar(n.Relation),
		Query:     convertNode(n.Query),
		Attlist:   convertSlice(n.Attlist),
		IsFrom:    n.IsFrom,
		IsProgram: n.IsProgram,
		Filename:  makeString(n.Filename),
		Options:   convertSlice(n.Options),
	}
}

func convertCreateAmStmt(n *pg.CreateAmStmt) *ast.CreateAmStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateAmStmt{
		Amname:      makeString(n.Amname),
		HandlerName: convertSlice(n.HandlerName),
		Amtype:      makeByte(n.Amtype),
	}
}

func convertCreateCastStmt(n *pg.CreateCastStmt) *ast.CreateCastStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateCastStmt{
		Sourcetype: convertTypeName(n.Sourcetype),
		Targettype: convertTypeName(n.Targettype),
		Func:       convertObjectWithArgs(n.Func),
		Context:    ast.CoercionContext(n.Context),
		Inout:      n.Inout,
	}
}

func convertCreateConversionStmt(n *pg.CreateConversionStmt) *ast.CreateConversionStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateConversionStmt{
		ConversionName:  convertSlice(n.ConversionName),
		ForEncodingName: makeString(n.ForEncodingName),
		ToEncodingName:  makeString(n.ToEncodingName),
		FuncName:        convertSlice(n.FuncName),
		Def:             n.Def,
	}
}

func convertCreateDomainStmt(n *pg.CreateDomainStmt) *ast.CreateDomainStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateDomainStmt{
		Domainname:  convertSlice(n.Domainname),
		TypeName:    convertTypeName(n.TypeName),
		CollClause:  convertCollateClause(n.CollClause),
		Constraints: convertSlice(n.Constraints),
	}
}

func convertCreateEnumStmt(n *pg.CreateEnumStmt) *ast.CreateEnumStmt {
	if n == nil {
		return nil
	}
	rel, err := parseRelationFromNodes(n.TypeName)
	if err != nil {
		panic(err)
	}
	return &ast.CreateEnumStmt{
		TypeName: rel.TypeName(),
		Vals:     convertSlice(n.Vals),
	}
}

func convertCreateEventTrigStmt(n *pg.CreateEventTrigStmt) *ast.CreateEventTrigStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateEventTrigStmt{
		Trigname:   makeString(n.Trigname),
		Eventname:  makeString(n.Eventname),
		Whenclause: convertSlice(n.Whenclause),
		Funcname:   convertSlice(n.Funcname),
	}
}

func convertCreateExtensionStmt(n *pg.CreateExtensionStmt) *ast.CreateExtensionStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateExtensionStmt{
		Extname:     makeString(n.Extname),
		IfNotExists: n.IfNotExists,
		Options:     convertSlice(n.Options),
	}
}

func convertCreateFdwStmt(n *pg.CreateFdwStmt) *ast.CreateFdwStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateFdwStmt{
		Fdwname:     makeString(n.Fdwname),
		FuncOptions: convertSlice(n.FuncOptions),
		Options:     convertSlice(n.Options),
	}
}

func convertCreateForeignServerStmt(n *pg.CreateForeignServerStmt) *ast.CreateForeignServerStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateForeignServerStmt{
		Servername:  makeString(n.Servername),
		Servertype:  makeString(n.Servertype),
		Version:     makeString(n.Version),
		Fdwname:     makeString(n.Fdwname),
		IfNotExists: n.IfNotExists,
		Options:     convertSlice(n.Options),
	}
}

func convertCreateForeignTableStmt(n *pg.CreateForeignTableStmt) *ast.CreateForeignTableStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateForeignTableStmt{
		Servername: makeString(n.Servername),
		Options:    convertSlice(n.Options),
	}
}

func convertCreateFunctionStmt(n *pg.CreateFunctionStmt) *ast.CreateFunctionStmt {
	if n == nil {
		return nil
	}
	rel, err := parseRelationFromNodes(n.Funcname)
	if err != nil {
		panic(err)
	}
	return &ast.CreateFunctionStmt{
		Replace:    n.Replace,
		Func:       rel.FuncName(),
		Params:     convertSlice(n.Parameters),
		ReturnType: convertTypeName(n.ReturnType),
		Options:    convertSlice(n.Options),
	}
}

func convertCreateOpClassItem(n *pg.CreateOpClassItem) *ast.CreateOpClassItem {
	if n == nil {
		return nil
	}
	return &ast.CreateOpClassItem{
		Itemtype:    int(n.Itemtype),
		Name:        convertObjectWithArgs(n.Name),
		Number:      int(n.Number),
		OrderFamily: convertSlice(n.OrderFamily),
		ClassArgs:   convertSlice(n.ClassArgs),
		Storedtype:  convertTypeName(n.Storedtype),
	}
}

func convertCreateOpClassStmt(n *pg.CreateOpClassStmt) *ast.CreateOpClassStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateOpClassStmt{
		Opclassname:  convertSlice(n.Opclassname),
		Opfamilyname: convertSlice(n.Opfamilyname),
		Amname:       makeString(n.Amname),
		Datatype:     convertTypeName(n.Datatype),
		Items:        convertSlice(n.Items),
		IsDefault:    n.IsDefault,
	}
}

func convertCreateOpFamilyStmt(n *pg.CreateOpFamilyStmt) *ast.CreateOpFamilyStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateOpFamilyStmt{
		Opfamilyname: convertSlice(n.Opfamilyname),
		Amname:       makeString(n.Amname),
	}
}

func convertCreatePLangStmt(n *pg.CreatePLangStmt) *ast.CreatePLangStmt {
	if n == nil {
		return nil
	}
	return &ast.CreatePLangStmt{
		Replace:     n.Replace,
		Plname:      makeString(n.Plname),
		Plhandler:   convertSlice(n.Plhandler),
		Plinline:    convertSlice(n.Plinline),
		Plvalidator: convertSlice(n.Plvalidator),
		Pltrusted:   n.Pltrusted,
	}
}

func convertCreatePolicyStmt(n *pg.CreatePolicyStmt) *ast.CreatePolicyStmt {
	if n == nil {
		return nil
	}
	return &ast.CreatePolicyStmt{
		PolicyName: makeString(n.PolicyName),
		Table:      convertRangeVar(n.Table),
		CmdName:    makeString(n.CmdName),
		Permissive: n.Permissive,
		Roles:      convertSlice(n.Roles),
		Qual:       convertNode(n.Qual),
		WithCheck:  convertNode(n.WithCheck),
	}
}

func convertCreatePublicationStmt(n *pg.CreatePublicationStmt) *ast.CreatePublicationStmt {
	if n == nil {
		return nil
	}
	return &ast.CreatePublicationStmt{
		Pubname:      makeString(n.Pubname),
		Options:      convertSlice(n.Options),
		Tables:       convertSlice(n.Pubobjects),
		ForAllTables: n.ForAllTables,
	}
}

func convertCreateRangeStmt(n *pg.CreateRangeStmt) *ast.CreateRangeStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateRangeStmt{
		TypeName: convertSlice(n.TypeName),
		Params:   convertSlice(n.Params),
	}
}

func convertCreateRoleStmt(n *pg.CreateRoleStmt) *ast.CreateRoleStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateRoleStmt{
		StmtType: ast.RoleStmtType(n.StmtType),
		Role:     makeString(n.Role),
		Options:  convertSlice(n.Options),
	}
}

func convertCreateSchemaStmt(n *pg.CreateSchemaStmt) *ast.CreateSchemaStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateSchemaStmt{
		Name:        makeString(n.Schemaname),
		Authrole:    convertRoleSpec(n.Authrole),
		SchemaElts:  convertSlice(n.SchemaElts),
		IfNotExists: n.IfNotExists,
	}
}

func convertCreateSeqStmt(n *pg.CreateSeqStmt) *ast.CreateSeqStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateSeqStmt{
		Sequence:    convertRangeVar(n.Sequence),
		Options:     convertSlice(n.Options),
		OwnerId:     ast.Oid(n.OwnerId),
		ForIdentity: n.ForIdentity,
		IfNotExists: n.IfNotExists,
	}
}

func convertCreateStatsStmt(n *pg.CreateStatsStmt) *ast.CreateStatsStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateStatsStmt{
		Defnames:    convertSlice(n.Defnames),
		StatTypes:   convertSlice(n.StatTypes),
		Exprs:       convertSlice(n.Exprs),
		Relations:   convertSlice(n.Relations),
		IfNotExists: n.IfNotExists,
	}
}

func convertCreateStmt(n *pg.CreateStmt) *ast.CreateStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateStmt{
		Relation:       convertRangeVar(n.Relation),
		TableElts:      convertSlice(n.TableElts),
		InhRelations:   convertSlice(n.InhRelations),
		Partbound:      convertPartitionBoundSpec(n.Partbound),
		Partspec:       convertPartitionSpec(n.Partspec),
		OfTypename:     convertTypeName(n.OfTypename),
		Constraints:    convertSlice(n.Constraints),
		Options:        convertSlice(n.Options),
		Oncommit:       ast.OnCommitAction(n.Oncommit),
		Tablespacename: makeString(n.Tablespacename),
		IfNotExists:    n.IfNotExists,
	}
}

func convertCreateSubscriptionStmt(n *pg.CreateSubscriptionStmt) *ast.CreateSubscriptionStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateSubscriptionStmt{
		Subname:     makeString(n.Subname),
		Conninfo:    makeString(n.Conninfo),
		Publication: convertSlice(n.Publication),
		Options:     convertSlice(n.Options),
	}
}

func convertCreateTableAsStmt(n *pg.CreateTableAsStmt) *ast.CreateTableAsStmt {
	if n == nil {
		return nil
	}
	res := &ast.CreateTableAsStmt{
		Query:        convertNode(n.Query),
		Into:         convertIntoClause(n.Into),
		Relkind:      ast.ObjectType(n.Objtype),
		IsSelectInto: n.IsSelectInto,
		IfNotExists:  n.IfNotExists,
	}
	return res
}

func convertCreateTableSpaceStmt(n *pg.CreateTableSpaceStmt) *ast.CreateTableSpaceStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateTableSpaceStmt{
		Tablespacename: makeString(n.Tablespacename),
		Owner:          convertRoleSpec(n.Owner),
		Location:       makeString(n.Location),
		Options:        convertSlice(n.Options),
	}
}

func convertCreateTransformStmt(n *pg.CreateTransformStmt) *ast.CreateTransformStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateTransformStmt{
		Replace:  n.Replace,
		TypeName: convertTypeName(n.TypeName),
		Lang:     makeString(n.Lang),
		Fromsql:  convertObjectWithArgs(n.Fromsql),
		Tosql:    convertObjectWithArgs(n.Tosql),
	}
}

func convertCreateTrigStmt(n *pg.CreateTrigStmt) *ast.CreateTrigStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateTrigStmt{
		Trigname:       makeString(n.Trigname),
		Relation:       convertRangeVar(n.Relation),
		Funcname:       convertSlice(n.Funcname),
		Args:           convertSlice(n.Args),
		Row:            n.Row,
		Timing:         int16(n.Timing),
		Events:         int16(n.Events),
		Columns:        convertSlice(n.Columns),
		WhenClause:     convertNode(n.WhenClause),
		Isconstraint:   n.Isconstraint,
		TransitionRels: convertSlice(n.TransitionRels),
		Deferrable:     n.Deferrable,
		Initdeferred:   n.Initdeferred,
		Constrrel:      convertRangeVar(n.Constrrel),
	}
}

func convertCreateUserMappingStmt(n *pg.CreateUserMappingStmt) *ast.CreateUserMappingStmt {
	if n == nil {
		return nil
	}
	return &ast.CreateUserMappingStmt{
		User:        convertRoleSpec(n.User),
		Servername:  makeString(n.Servername),
		IfNotExists: n.IfNotExists,
		Options:     convertSlice(n.Options),
	}
}

func convertCreatedbStmt(n *pg.CreatedbStmt) *ast.CreatedbStmt {
	if n == nil {
		return nil
	}
	return &ast.CreatedbStmt{
		Dbname:  makeString(n.Dbname),
		Options: convertSlice(n.Options),
	}
}

func convertCurrentOfExpr(n *pg.CurrentOfExpr) *ast.CurrentOfExpr {
	if n == nil {
		return nil
	}
	return &ast.CurrentOfExpr{
		Xpr:         convertNode(n.Xpr),
		Cvarno:      ast.Index(n.Cvarno),
		CursorName:  makeString(n.CursorName),
		CursorParam: int(n.CursorParam),
	}
}

func convertDeallocateStmt(n *pg.DeallocateStmt) *ast.DeallocateStmt {
	if n == nil {
		return nil
	}
	return &ast.DeallocateStmt{
		Name: makeString(n.Name),
	}
}

func convertDeclareCursorStmt(n *pg.DeclareCursorStmt) *ast.DeclareCursorStmt {
	if n == nil {
		return nil
	}
	return &ast.DeclareCursorStmt{
		Portalname: makeString(n.Portalname),
		Options:    int(n.Options),
		Query:      convertNode(n.Query),
	}
}

func convertDefElem(n *pg.DefElem) *ast.DefElem {
	if n == nil {
		return nil
	}
	return &ast.DefElem{
		Defnamespace: makeString(n.Defnamespace),
		Defname:      makeString(n.Defname),
		Arg:          convertNode(n.Arg),
		Defaction:    ast.DefElemAction(n.Defaction),
		Location:     int(n.Location),
	}
}

func convertDefineStmt(n *pg.DefineStmt) *ast.DefineStmt {
	if n == nil {
		return nil
	}
	return &ast.DefineStmt{
		Kind:        ast.ObjectType(n.Kind),
		Oldstyle:    n.Oldstyle,
		Defnames:    convertSlice(n.Defnames),
		Args:        convertSlice(n.Args),
		Definition:  convertSlice(n.Definition),
		IfNotExists: n.IfNotExists,
	}
}

func convertDeleteStmt(n *pg.DeleteStmt) *ast.DeleteStmt {
	if n == nil {
		return nil
	}
	return &ast.DeleteStmt{
		Relation:      convertRangeVar(n.Relation),
		UsingClause:   convertSlice(n.UsingClause),
		WhereClause:   convertNode(n.WhereClause),
		ReturningList: convertSlice(n.ReturningList),
		WithClause:    convertWithClause(n.WithClause),
	}
}

func convertDiscardStmt(n *pg.DiscardStmt) *ast.DiscardStmt {
	if n == nil {
		return nil
	}
	return &ast.DiscardStmt{
		Target: ast.DiscardMode(n.Target),
	}
}

func convertDoStmt(n *pg.DoStmt) *ast.DoStmt {
	if n == nil {
		return nil
	}
	return &ast.DoStmt{
		Args: convertSlice(n.Args),
	}
}

func convertDropOwnedStmt(n *pg.DropOwnedStmt) *ast.DropOwnedStmt {
	if n == nil {
		return nil
	}
	return &ast.DropOwnedStmt{
		Roles:    convertSlice(n.Roles),
		Behavior: ast.DropBehavior(n.Behavior),
	}
}

func convertDropRoleStmt(n *pg.DropRoleStmt) *ast.DropRoleStmt {
	if n == nil {
		return nil
	}
	return &ast.DropRoleStmt{
		Roles:     convertSlice(n.Roles),
		MissingOk: n.MissingOk,
	}
}

func convertDropStmt(n *pg.DropStmt) *ast.DropStmt {
	if n == nil {
		return nil
	}
	return &ast.DropStmt{
		Objects:    convertSlice(n.Objects),
		RemoveType: ast.ObjectType(n.RemoveType),
		Behavior:   ast.DropBehavior(n.Behavior),
		MissingOk:  n.MissingOk,
		Concurrent: n.Concurrent,
	}
}

func convertDropSubscriptionStmt(n *pg.DropSubscriptionStmt) *ast.DropSubscriptionStmt {
	if n == nil {
		return nil
	}
	return &ast.DropSubscriptionStmt{
		Subname:   makeString(n.Subname),
		MissingOk: n.MissingOk,
		Behavior:  ast.DropBehavior(n.Behavior),
	}
}

func convertDropTableSpaceStmt(n *pg.DropTableSpaceStmt) *ast.DropTableSpaceStmt {
	if n == nil {
		return nil
	}
	return &ast.DropTableSpaceStmt{
		Tablespacename: makeString(n.Tablespacename),
		MissingOk:      n.MissingOk,
	}
}

func convertDropUserMappingStmt(n *pg.DropUserMappingStmt) *ast.DropUserMappingStmt {
	if n == nil {
		return nil
	}
	return &ast.DropUserMappingStmt{
		User:       convertRoleSpec(n.User),
		Servername: makeString(n.Servername),
		MissingOk:  n.MissingOk,
	}
}

func convertDropdbStmt(n *pg.DropdbStmt) *ast.DropdbStmt {
	if n == nil {
		return nil
	}
	return &ast.DropdbStmt{
		Dbname:    makeString(n.Dbname),
		MissingOk: n.MissingOk,
	}
}

func convertExecuteStmt(n *pg.ExecuteStmt) *ast.ExecuteStmt {
	if n == nil {
		return nil
	}
	return &ast.ExecuteStmt{
		Name:   makeString(n.Name),
		Params: convertSlice(n.Params),
	}
}

func convertExplainStmt(n *pg.ExplainStmt) *ast.ExplainStmt {
	if n == nil {
		return nil
	}
	return &ast.ExplainStmt{
		Query:   convertNode(n.Query),
		Options: convertSlice(n.Options),
	}
}

func convertFetchStmt(n *pg.FetchStmt) *ast.FetchStmt {
	if n == nil {
		return nil
	}
	return &ast.FetchStmt{
		Direction:  ast.FetchDirection(n.Direction),
		HowMany:    n.HowMany,
		Portalname: makeString(n.Portalname),
		Ismove:     n.Ismove,
	}
}

func convertFieldSelect(n *pg.FieldSelect) *ast.FieldSelect {
	if n == nil {
		return nil
	}
	return &ast.FieldSelect{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Fieldnum:     ast.AttrNumber(n.Fieldnum),
		Resulttype:   ast.Oid(n.Resulttype),
		Resulttypmod: n.Resulttypmod,
		Resultcollid: ast.Oid(n.Resultcollid),
	}
}

func convertFieldStore(n *pg.FieldStore) *ast.FieldStore {
	if n == nil {
		return nil
	}
	return &ast.FieldStore{
		Xpr:        convertNode(n.Xpr),
		Arg:        convertNode(n.Arg),
		Newvals:    convertSlice(n.Newvals),
		Fieldnums:  convertSlice(n.Fieldnums),
		Resulttype: ast.Oid(n.Resulttype),
	}
}

func convertFloat(n *pg.Float) *ast.Float {
	if n == nil {
		return nil
	}
	return &ast.Float{
		Str: n.Fval,
	}
}

func convertFromExpr(n *pg.FromExpr) *ast.FromExpr {
	if n == nil {
		return nil
	}
	return &ast.FromExpr{
		Fromlist: convertSlice(n.Fromlist),
		Quals:    convertNode(n.Quals),
	}
}

func convertFuncCall(n *pg.FuncCall) *ast.FuncCall {
	if n == nil {
		return nil
	}
	rel, err := parseRelationFromNodes(n.Funcname)
	if err != nil {
		// TODO: How should we handle errors?
		panic(err)
	}
	return &ast.FuncCall{
		Func:           rel.FuncName(),
		Funcname:       convertSlice(n.Funcname),
		Args:           convertSlice(n.Args),
		AggOrder:       convertSlice(n.AggOrder),
		AggFilter:      convertNode(n.AggFilter),
		AggWithinGroup: n.AggWithinGroup,
		AggStar:        n.AggStar,
		AggDistinct:    n.AggDistinct,
		FuncVariadic:   n.FuncVariadic,
		Over:           convertWindowDef(n.Over),
		Location:       int(n.Location),
	}
}

func convertFuncExpr(n *pg.FuncExpr) *ast.FuncExpr {
	if n == nil {
		return nil
	}
	return &ast.FuncExpr{
		Xpr:            convertNode(n.Xpr),
		Funcid:         ast.Oid(n.Funcid),
		Funcresulttype: ast.Oid(n.Funcresulttype),
		Funcretset:     n.Funcretset,
		Funcvariadic:   n.Funcvariadic,
		Funcformat:     ast.CoercionForm(n.Funcformat),
		Funccollid:     ast.Oid(n.Funccollid),
		Inputcollid:    ast.Oid(n.Inputcollid),
		Args:           convertSlice(n.Args),
		Location:       int(n.Location),
	}
}

func convertFunctionParameter(n *pg.FunctionParameter) *ast.FunctionParameter {
	if n == nil {
		return nil
	}
	return &ast.FunctionParameter{
		Name:    makeString(n.Name),
		ArgType: convertTypeName(n.ArgType),
		Mode:    ast.FunctionParameterMode(n.Mode),
		Defexpr: convertNode(n.Defexpr),
	}
}

func convertGrantRoleStmt(n *pg.GrantRoleStmt) *ast.GrantRoleStmt {
	if n == nil {
		return nil
	}
	return &ast.GrantRoleStmt{
		GrantedRoles: convertSlice(n.GrantedRoles),
		GranteeRoles: convertSlice(n.GranteeRoles),
		IsGrant:      n.IsGrant,
		AdminOpt:     n.AdminOpt,
		Grantor:      convertRoleSpec(n.Grantor),
		Behavior:     ast.DropBehavior(n.Behavior),
	}
}

func convertGrantStmt(n *pg.GrantStmt) *ast.GrantStmt {
	if n == nil {
		return nil
	}
	return &ast.GrantStmt{
		IsGrant:     n.IsGrant,
		Targtype:    ast.GrantTargetType(n.Targtype),
		Objtype:     ast.GrantObjectType(n.Objtype),
		Objects:     convertSlice(n.Objects),
		Privileges:  convertSlice(n.Privileges),
		Grantees:    convertSlice(n.Grantees),
		GrantOption: n.GrantOption,
		Behavior:    ast.DropBehavior(n.Behavior),
	}
}

func convertGroupingFunc(n *pg.GroupingFunc) *ast.GroupingFunc {
	if n == nil {
		return nil
	}
	return &ast.GroupingFunc{
		Xpr:         convertNode(n.Xpr),
		Args:        convertSlice(n.Args),
		Refs:        convertSlice(n.Refs),
		Cols:        convertSlice(n.Cols),
		Agglevelsup: ast.Index(n.Agglevelsup),
		Location:    int(n.Location),
	}
}

func convertGroupingSet(n *pg.GroupingSet) *ast.GroupingSet {
	if n == nil {
		return nil
	}
	return &ast.GroupingSet{
		Kind:     ast.GroupingSetKind(n.Kind),
		Content:  convertSlice(n.Content),
		Location: int(n.Location),
	}
}

func convertImportForeignSchemaStmt(n *pg.ImportForeignSchemaStmt) *ast.ImportForeignSchemaStmt {
	if n == nil {
		return nil
	}
	return &ast.ImportForeignSchemaStmt{
		ServerName:   makeString(n.ServerName),
		RemoteSchema: makeString(n.RemoteSchema),
		LocalSchema:  makeString(n.LocalSchema),
		ListType:     ast.ImportForeignSchemaType(n.ListType),
		TableList:    convertSlice(n.TableList),
		Options:      convertSlice(n.Options),
	}
}

func convertIndexElem(n *pg.IndexElem) *ast.IndexElem {
	if n == nil {
		return nil
	}
	return &ast.IndexElem{
		Name:          makeString(n.Name),
		Expr:          convertNode(n.Expr),
		Indexcolname:  makeString(n.Indexcolname),
		Collation:     convertSlice(n.Collation),
		Opclass:       convertSlice(n.Opclass),
		Ordering:      ast.SortByDir(n.Ordering),
		NullsOrdering: ast.SortByNulls(n.NullsOrdering),
	}
}

func convertIndexStmt(n *pg.IndexStmt) *ast.IndexStmt {
	if n == nil {
		return nil
	}
	return &ast.IndexStmt{
		Idxname:        makeString(n.Idxname),
		Relation:       convertRangeVar(n.Relation),
		AccessMethod:   makeString(n.AccessMethod),
		TableSpace:     makeString(n.TableSpace),
		IndexParams:    convertSlice(n.IndexParams),
		Options:        convertSlice(n.Options),
		WhereClause:    convertNode(n.WhereClause),
		ExcludeOpNames: convertSlice(n.ExcludeOpNames),
		Idxcomment:     makeString(n.Idxcomment),
		IndexOid:       ast.Oid(n.IndexOid),
		OldNode:        ast.Oid(n.OldNode),
		Unique:         n.Unique,
		Primary:        n.Primary,
		Isconstraint:   n.Isconstraint,
		Deferrable:     n.Deferrable,
		Initdeferred:   n.Initdeferred,
		Transformed:    n.Transformed,
		Concurrent:     n.Concurrent,
		IfNotExists:    n.IfNotExists,
	}
}

func convertInferClause(n *pg.InferClause) *ast.InferClause {
	if n == nil {
		return nil
	}
	return &ast.InferClause{
		IndexElems:  convertSlice(n.IndexElems),
		WhereClause: convertNode(n.WhereClause),
		Conname:     makeString(n.Conname),
		Location:    int(n.Location),
	}
}

func convertInferenceElem(n *pg.InferenceElem) *ast.InferenceElem {
	if n == nil {
		return nil
	}
	return &ast.InferenceElem{
		Xpr:          convertNode(n.Xpr),
		Expr:         convertNode(n.Expr),
		Infercollid:  ast.Oid(n.Infercollid),
		Inferopclass: ast.Oid(n.Inferopclass),
	}
}

func convertInlineCodeBlock(n *pg.InlineCodeBlock) *ast.InlineCodeBlock {
	if n == nil {
		return nil
	}
	return &ast.InlineCodeBlock{
		SourceText:    makeString(n.SourceText),
		LangOid:       ast.Oid(n.LangOid),
		LangIsTrusted: n.LangIsTrusted,
	}
}

func convertInsertStmt(n *pg.InsertStmt) *ast.InsertStmt {
	if n == nil {
		return nil
	}
	return &ast.InsertStmt{
		Relation:         convertRangeVar(n.Relation),
		Cols:             convertSlice(n.Cols),
		SelectStmt:       convertNode(n.SelectStmt),
		OnConflictClause: convertOnConflictClause(n.OnConflictClause),
		ReturningList:    convertSlice(n.ReturningList),
		WithClause:       convertWithClause(n.WithClause),
		Override:         ast.OverridingKind(n.Override),
	}
}

func convertInteger(n *pg.Integer) *ast.Integer {
	if n == nil {
		return nil
	}
	return &ast.Integer{
		Ival: int64(n.Ival),
	}
}

func convertIntoClause(n *pg.IntoClause) *ast.IntoClause {
	if n == nil {
		return nil
	}
	return &ast.IntoClause{
		Rel:            convertRangeVar(n.Rel),
		ColNames:       convertSlice(n.ColNames),
		Options:        convertSlice(n.Options),
		OnCommit:       ast.OnCommitAction(n.OnCommit),
		TableSpaceName: makeString(n.TableSpaceName),
		ViewQuery:      convertNode(n.ViewQuery),
		SkipData:       n.SkipData,
	}
}

func convertJoinExpr(n *pg.JoinExpr) *ast.JoinExpr {
	if n == nil {
		return nil
	}
	return &ast.JoinExpr{
		Jointype:    ast.JoinType(n.Jointype),
		IsNatural:   n.IsNatural,
		Larg:        convertNode(n.Larg),
		Rarg:        convertNode(n.Rarg),
		UsingClause: convertSlice(n.UsingClause),
		Quals:       convertNode(n.Quals),
		Alias:       convertAlias(n.Alias),
		Rtindex:     int(n.Rtindex),
	}
}

func convertListenStmt(n *pg.ListenStmt) *ast.ListenStmt {
	if n == nil {
		return nil
	}
	return &ast.ListenStmt{
		Conditionname: makeString(n.Conditionname),
	}
}

func convertLoadStmt(n *pg.LoadStmt) *ast.LoadStmt {
	if n == nil {
		return nil
	}
	return &ast.LoadStmt{
		Filename: makeString(n.Filename),
	}
}

func convertLockStmt(n *pg.LockStmt) *ast.LockStmt {
	if n == nil {
		return nil
	}
	return &ast.LockStmt{
		Relations: convertSlice(n.Relations),
		Mode:      int(n.Mode),
		Nowait:    n.Nowait,
	}
}

func convertLockingClause(n *pg.LockingClause) *ast.LockingClause {
	if n == nil {
		return nil
	}
	return &ast.LockingClause{
		LockedRels: convertSlice(n.LockedRels),
		Strength:   ast.LockClauseStrength(n.Strength),
		WaitPolicy: ast.LockWaitPolicy(n.WaitPolicy),
	}
}

func convertMinMaxExpr(n *pg.MinMaxExpr) *ast.MinMaxExpr {
	if n == nil {
		return nil
	}
	return &ast.MinMaxExpr{
		Xpr:          convertNode(n.Xpr),
		Minmaxtype:   ast.Oid(n.Minmaxtype),
		Minmaxcollid: ast.Oid(n.Minmaxcollid),
		Inputcollid:  ast.Oid(n.Inputcollid),
		Op:           ast.MinMaxOp(n.Op),
		Args:         convertSlice(n.Args),
		Location:     int(n.Location),
	}
}

func convertMultiAssignRef(n *pg.MultiAssignRef) *ast.MultiAssignRef {
	if n == nil {
		return nil
	}
	return &ast.MultiAssignRef{
		Source:   convertNode(n.Source),
		Colno:    int(n.Colno),
		Ncolumns: int(n.Ncolumns),
	}
}

func convertNamedArgExpr(n *pg.NamedArgExpr) *ast.NamedArgExpr {
	if n == nil {
		return nil
	}
	return &ast.NamedArgExpr{
		Xpr:       convertNode(n.Xpr),
		Arg:       convertNode(n.Arg),
		Name:      makeString(n.Name),
		Argnumber: int(n.Argnumber),
		Location:  int(n.Location),
	}
}

func convertNextValueExpr(n *pg.NextValueExpr) *ast.NextValueExpr {
	if n == nil {
		return nil
	}
	return &ast.NextValueExpr{
		Xpr:    convertNode(n.Xpr),
		Seqid:  ast.Oid(n.Seqid),
		TypeId: ast.Oid(n.TypeId),
	}
}

func convertNotifyStmt(n *pg.NotifyStmt) *ast.NotifyStmt {
	if n == nil {
		return nil
	}
	return &ast.NotifyStmt{
		Conditionname: makeString(n.Conditionname),
		Payload:       makeString(n.Payload),
	}
}

func convertNullTest(n *pg.NullTest) *ast.NullTest {
	if n == nil {
		return nil
	}
	return &ast.NullTest{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Nulltesttype: ast.NullTestType(n.Nulltesttype),
		Argisrow:     n.Argisrow,
		Location:     int(n.Location),
	}
}

func convertObjectWithArgs(n *pg.ObjectWithArgs) *ast.ObjectWithArgs {
	if n == nil {
		return nil
	}
	return &ast.ObjectWithArgs{
		Objname:         convertSlice(n.Objname),
		Objargs:         convertSlice(n.Objargs),
		ArgsUnspecified: n.ArgsUnspecified,
	}
}

func convertOnConflictClause(n *pg.OnConflictClause) *ast.OnConflictClause {
	if n == nil {
		return nil
	}
	return &ast.OnConflictClause{
		Action:      ast.OnConflictAction(n.Action),
		Infer:       convertInferClause(n.Infer),
		TargetList:  convertSlice(n.TargetList),
		WhereClause: convertNode(n.WhereClause),
		Location:    int(n.Location),
	}
}

func convertOnConflictExpr(n *pg.OnConflictExpr) *ast.OnConflictExpr {
	if n == nil {
		return nil
	}
	return &ast.OnConflictExpr{
		Action:          ast.OnConflictAction(n.Action),
		ArbiterElems:    convertSlice(n.ArbiterElems),
		ArbiterWhere:    convertNode(n.ArbiterWhere),
		Constraint:      ast.Oid(n.Constraint),
		OnConflictSet:   convertSlice(n.OnConflictSet),
		OnConflictWhere: convertNode(n.OnConflictWhere),
		ExclRelIndex:    int(n.ExclRelIndex),
		ExclRelTlist:    convertSlice(n.ExclRelTlist),
	}
}

func convertOpExpr(n *pg.OpExpr) *ast.OpExpr {
	if n == nil {
		return nil
	}
	return &ast.OpExpr{
		Xpr:          convertNode(n.Xpr),
		Opno:         ast.Oid(n.Opno),
		Opfuncid:     ast.Oid(n.Opfuncid),
		Opresulttype: ast.Oid(n.Opresulttype),
		Opretset:     n.Opretset,
		Opcollid:     ast.Oid(n.Opcollid),
		Inputcollid:  ast.Oid(n.Inputcollid),
		Args:         convertSlice(n.Args),
		Location:     int(n.Location),
	}
}

func convertParam(n *pg.Param) *ast.Param {
	if n == nil {
		return nil
	}
	return &ast.Param{
		Xpr:         convertNode(n.Xpr),
		Paramkind:   ast.ParamKind(n.Paramkind),
		Paramid:     int(n.Paramid),
		Paramtype:   ast.Oid(n.Paramtype),
		Paramtypmod: n.Paramtypmod,
		Paramcollid: ast.Oid(n.Paramcollid),
		Location:    int(n.Location),
	}
}

func convertParamRef(n *pg.ParamRef) *ast.ParamRef {
	if n == nil {
		return nil
	}
	var dollar bool
	if n.Number != 0 {
		dollar = true
	}
	return &ast.ParamRef{
		Dollar:   dollar,
		Number:   int(n.Number),
		Location: int(n.Location),
	}
}

func convertPartitionBoundSpec(n *pg.PartitionBoundSpec) *ast.PartitionBoundSpec {
	if n == nil {
		return nil
	}
	return &ast.PartitionBoundSpec{
		Strategy:    makeByte(n.Strategy),
		Listdatums:  convertSlice(n.Listdatums),
		Lowerdatums: convertSlice(n.Lowerdatums),
		Upperdatums: convertSlice(n.Upperdatums),
		Location:    int(n.Location),
	}
}

func convertPartitionCmd(n *pg.PartitionCmd) *ast.PartitionCmd {
	if n == nil {
		return nil
	}
	return &ast.PartitionCmd{
		Name:  convertRangeVar(n.Name),
		Bound: convertPartitionBoundSpec(n.Bound),
	}
}

func convertPartitionElem(n *pg.PartitionElem) *ast.PartitionElem {
	if n == nil {
		return nil
	}
	return &ast.PartitionElem{
		Name:      makeString(n.Name),
		Expr:      convertNode(n.Expr),
		Collation: convertSlice(n.Collation),
		Opclass:   convertSlice(n.Opclass),
		Location:  int(n.Location),
	}
}

func convertPartitionRangeDatum(n *pg.PartitionRangeDatum) *ast.PartitionRangeDatum {
	if n == nil {
		return nil
	}
	return &ast.PartitionRangeDatum{
		Kind:     ast.PartitionRangeDatumKind(n.Kind),
		Value:    convertNode(n.Value),
		Location: int(n.Location),
	}
}

func convertPartitionSpec(n *pg.PartitionSpec) *ast.PartitionSpec {
	if n == nil {
		return nil
	}
	return &ast.PartitionSpec{
		Strategy:   makeString(n.Strategy),
		PartParams: convertSlice(n.PartParams),
		Location:   int(n.Location),
	}
}

func convertPrepareStmt(n *pg.PrepareStmt) *ast.PrepareStmt {
	if n == nil {
		return nil
	}
	return &ast.PrepareStmt{
		Name:     makeString(n.Name),
		Argtypes: convertSlice(n.Argtypes),
		Query:    convertNode(n.Query),
	}
}

func convertQuery(n *pg.Query) *ast.Query {
	if n == nil {
		return nil
	}
	return &ast.Query{
		CommandType:      ast.CmdType(n.CommandType),
		QuerySource:      ast.QuerySource(n.QuerySource),
		CanSetTag:        n.CanSetTag,
		UtilityStmt:      convertNode(n.UtilityStmt),
		ResultRelation:   int(n.ResultRelation),
		HasAggs:          n.HasAggs,
		HasWindowFuncs:   n.HasWindowFuncs,
		HasTargetSrfs:    n.HasTargetSrfs,
		HasSubLinks:      n.HasSubLinks,
		HasDistinctOn:    n.HasDistinctOn,
		HasRecursive:     n.HasRecursive,
		HasModifyingCte:  n.HasModifyingCte,
		HasForUpdate:     n.HasForUpdate,
		HasRowSecurity:   n.HasRowSecurity,
		CteList:          convertSlice(n.CteList),
		Rtable:           convertSlice(n.Rtable),
		Jointree:         convertFromExpr(n.Jointree),
		TargetList:       convertSlice(n.TargetList),
		Override:         ast.OverridingKind(n.Override),
		OnConflict:       convertOnConflictExpr(n.OnConflict),
		ReturningList:    convertSlice(n.ReturningList),
		GroupClause:      convertSlice(n.GroupClause),
		GroupingSets:     convertSlice(n.GroupingSets),
		HavingQual:       convertNode(n.HavingQual),
		WindowClause:     convertSlice(n.WindowClause),
		DistinctClause:   convertSlice(n.DistinctClause),
		SortClause:       convertSlice(n.SortClause),
		LimitOffset:      convertNode(n.LimitOffset),
		LimitCount:       convertNode(n.LimitCount),
		RowMarks:         convertSlice(n.RowMarks),
		SetOperations:    convertNode(n.SetOperations),
		ConstraintDeps:   convertSlice(n.ConstraintDeps),
		WithCheckOptions: convertSlice(n.WithCheckOptions),
		StmtLocation:     int(n.StmtLocation),
		StmtLen:          int(n.StmtLen),
	}
}

func convertRangeFunction(n *pg.RangeFunction) *ast.RangeFunction {
	if n == nil {
		return nil
	}
	return &ast.RangeFunction{
		Lateral:    n.Lateral,
		Ordinality: n.Ordinality,
		IsRowsfrom: n.IsRowsfrom,
		Functions:  convertSlice(n.Functions),
		Alias:      convertAlias(n.Alias),
		Coldeflist: convertSlice(n.Coldeflist),
	}
}

func convertRangeSubselect(n *pg.RangeSubselect) *ast.RangeSubselect {
	if n == nil {
		return nil
	}
	return &ast.RangeSubselect{
		Lateral:  n.Lateral,
		Subquery: convertNode(n.Subquery),
		Alias:    convertAlias(n.Alias),
	}
}

func convertRangeTableFunc(n *pg.RangeTableFunc) *ast.RangeTableFunc {
	if n == nil {
		return nil
	}
	return &ast.RangeTableFunc{
		Lateral:    n.Lateral,
		Docexpr:    convertNode(n.Docexpr),
		Rowexpr:    convertNode(n.Rowexpr),
		Namespaces: convertSlice(n.Namespaces),
		Columns:    convertSlice(n.Columns),
		Alias:      convertAlias(n.Alias),
		Location:   int(n.Location),
	}
}

func convertRangeTableFuncCol(n *pg.RangeTableFuncCol) *ast.RangeTableFuncCol {
	if n == nil {
		return nil
	}
	return &ast.RangeTableFuncCol{
		Colname:       makeString(n.Colname),
		TypeName:      convertTypeName(n.TypeName),
		ForOrdinality: n.ForOrdinality,
		IsNotNull:     n.IsNotNull,
		Colexpr:       convertNode(n.Colexpr),
		Coldefexpr:    convertNode(n.Coldefexpr),
		Location:      int(n.Location),
	}
}

func convertRangeTableSample(n *pg.RangeTableSample) *ast.RangeTableSample {
	if n == nil {
		return nil
	}
	return &ast.RangeTableSample{
		Relation:   convertNode(n.Relation),
		Method:     convertSlice(n.Method),
		Args:       convertSlice(n.Args),
		Repeatable: convertNode(n.Repeatable),
		Location:   int(n.Location),
	}
}

func convertRangeTblEntry(n *pg.RangeTblEntry) *ast.RangeTblEntry {
	if n == nil {
		return nil
	}
	return &ast.RangeTblEntry{
		Rtekind:         ast.RTEKind(n.Rtekind),
		Relid:           ast.Oid(n.Relid),
		Relkind:         makeByte(n.Relkind),
		Tablesample:     convertTableSampleClause(n.Tablesample),
		Subquery:        convertQuery(n.Subquery),
		SecurityBarrier: n.SecurityBarrier,
		Jointype:        ast.JoinType(n.Jointype),
		Joinaliasvars:   convertSlice(n.Joinaliasvars),
		Functions:       convertSlice(n.Functions),
		Funcordinality:  n.Funcordinality,
		Tablefunc:       convertTableFunc(n.Tablefunc),
		ValuesLists:     convertSlice(n.ValuesLists),
		Ctename:         makeString(n.Ctename),
		Ctelevelsup:     ast.Index(n.Ctelevelsup),
		SelfReference:   n.SelfReference,
		Coltypes:        convertSlice(n.Coltypes),
		Coltypmods:      convertSlice(n.Coltypmods),
		Colcollations:   convertSlice(n.Colcollations),
		Enrname:         makeString(n.Enrname),
		Enrtuples:       n.Enrtuples,
		Alias:           convertAlias(n.Alias),
		Eref:            convertAlias(n.Eref),
		Lateral:         n.Lateral,
		Inh:             n.Inh,
		InFromCl:        n.InFromCl,
		RequiredPerms:   ast.AclMode(n.RequiredPerms),
		CheckAsUser:     ast.Oid(n.CheckAsUser),
		SelectedCols:    makeUint32Slice(n.SelectedCols),
		InsertedCols:    makeUint32Slice(n.InsertedCols),
		UpdatedCols:     makeUint32Slice(n.UpdatedCols),
		SecurityQuals:   convertSlice(n.SecurityQuals),
	}
}

func convertRangeTblFunction(n *pg.RangeTblFunction) *ast.RangeTblFunction {
	if n == nil {
		return nil
	}
	return &ast.RangeTblFunction{
		Funcexpr:          convertNode(n.Funcexpr),
		Funccolcount:      int(n.Funccolcount),
		Funccolnames:      convertSlice(n.Funccolnames),
		Funccoltypes:      convertSlice(n.Funccoltypes),
		Funccoltypmods:    convertSlice(n.Funccoltypmods),
		Funccolcollations: convertSlice(n.Funccolcollations),
		Funcparams:        makeUint32Slice(n.Funcparams),
	}
}

func convertRangeTblRef(n *pg.RangeTblRef) *ast.RangeTblRef {
	if n == nil {
		return nil
	}
	return &ast.RangeTblRef{
		Rtindex: int(n.Rtindex),
	}
}

func convertRangeVar(n *pg.RangeVar) *ast.RangeVar {
	if n == nil {
		return nil
	}
	return &ast.RangeVar{
		Catalogname:    makeString(n.Catalogname),
		Schemaname:     makeString(n.Schemaname),
		Relname:        makeString(n.Relname),
		Inh:            n.Inh,
		Relpersistence: makeByte(n.Relpersistence),
		Alias:          convertAlias(n.Alias),
		Location:       int(n.Location),
	}
}

func convertRawStmt(n *pg.RawStmt) *ast.RawStmt {
	if n == nil {
		return nil
	}
	return &ast.RawStmt{
		Stmt:         convertNode(n.Stmt),
		StmtLocation: int(n.StmtLocation),
		StmtLen:      int(n.StmtLen),
	}
}

func convertReassignOwnedStmt(n *pg.ReassignOwnedStmt) *ast.ReassignOwnedStmt {
	if n == nil {
		return nil
	}
	return &ast.ReassignOwnedStmt{
		Roles:   convertSlice(n.Roles),
		Newrole: convertRoleSpec(n.Newrole),
	}
}

func convertRefreshMatViewStmt(n *pg.RefreshMatViewStmt) *ast.RefreshMatViewStmt {
	if n == nil {
		return nil
	}
	return &ast.RefreshMatViewStmt{
		Concurrent: n.Concurrent,
		SkipData:   n.SkipData,
		Relation:   convertRangeVar(n.Relation),
	}
}

func convertReindexStmt(n *pg.ReindexStmt) *ast.ReindexStmt {
	if n == nil {
		return nil
	}
	return &ast.ReindexStmt{
		Kind:     ast.ReindexObjectType(n.Kind),
		Relation: convertRangeVar(n.Relation),
		Name:     makeString(n.Name),
		// Options:  int(n.Options), TODO: Support params
	}
}
