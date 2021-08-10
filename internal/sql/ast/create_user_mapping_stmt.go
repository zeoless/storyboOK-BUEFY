package ast

type CreateUserMappingStmt struct {
	User        *RoleSpec
	Servername  *string
	IfNotE