package ast

import (
	"myvmp/token"
)

// The base Node interface
type Node interface {
	toString()
}

type Statement interface {
	Node
	StatementNode() string
}

type Program struct {
	Body     []*Statement
	TypeInfo string
}

func (nn *Program) StatementNode() string {
	nn.TypeInfo = token.Prog
	return token.Prog
}

func (nn *Program) toString() {

}

type Identifier struct {
	Name     string
	PAIX     int
	TypeInfo string
}

func (nn *Identifier) StatementNode() string {
	nn.TypeInfo = token.IDENT
	return token.IDENT
}

func (nn *Identifier) toString() {

}

type NullIdentifier struct {
	Value    string
	PAIX     int
	TypeInfo string
}

func (nn *NullIdentifier) StatementNode() string {
	nn.TypeInfo = token.NULL
	return token.NULL
}

func (nn *NullIdentifier) toString() {

}

type VariableDeclaration struct {
	Token        string
	Declarations []*Statement
	Init         Statement
	PAIX         int
	TypeInfo     string
}

func (nn *VariableDeclaration) StatementNode() string {
	nn.TypeInfo = token.VAR
	return token.VAR
}
func (nn *VariableDeclaration) toString() {

}

type NumericLiteral struct {
	Value    float64
	PAIX     int
	TypeInfo string
}

func (nn *NumericLiteral) StatementNode() string {
	nn.TypeInfo = token.INT
	return token.INT
}
func (nn *NumericLiteral) toString() {

}

type BinaryExpression struct {
	Left     Statement
	Right    Statement
	Operator string
	PAIX     int
	TypeInfo string
}

func (nn *BinaryExpression) StatementNode() string {
	nn.TypeInfo = token.Bin
	return token.Bin
}
func (nn *BinaryExpression) toString() {

}

type NOP struct {
	TypeInfo string
}

func (nn *NOP) StatementNode() string {
	nn.TypeInfo = token.NOP
	return token.NOP
}
func (nn *NOP) toString() {

}

type OVER struct {
	TypeInfo string
}

func (nn *OVER) StatementNode() string {

	//nn.TypeInfo = token.OVER
	return nn.TypeInfo
}
func (nn *OVER) toString() {

}

type AssignmentExpression struct {
	Left     Statement
	Right    Statement
	Operator string
	PAIX     int
	TypeInfo string
}

func (nn *AssignmentExpression) StatementNode() string {
	nn.TypeInfo = token.Ass
	return token.Ass
}
func (nn *AssignmentExpression) toString() {

}

type CallExpression struct {
	Caller    Statement
	Arguments []*Statement
	PAIX      int
	TypeInfo  string
}

func (nn *CallExpression) StatementNode() string {
	nn.TypeInfo = token.Call
	return token.Call
}
func (nn *CallExpression) toString() {

}

type IfStatement struct {
	Test       Statement
	Consequent Statement
	Alternate  Statement
	PAIX       int
	TypeInfo   string
}

func (nn *IfStatement) StatementNode() string {
	nn.TypeInfo = token.IfStat
	return token.IfStat
}
func (nn *IfStatement) toString() {

}

type BlockStatement struct {
	Body     []*Statement
	PAIX     int
	TypeInfo string
}

func (nn *BlockStatement) StatementNode() string {
	nn.TypeInfo = token.Block
	return token.Block
}
func (nn *BlockStatement) toString() {

}

type UnaryExpression struct {
	Argument Statement
	Prefix   bool
	Operator string
	PAIX     int
	TypeInfo string
}

func (nn *UnaryExpression) StatementNode() string {
	nn.TypeInfo = token.Unary
	return token.Unary
}
func (nn *UnaryExpression) toString() {

}

type FunctionDeclaration struct {
	Id       Statement
	Params   []*Statement
	Body     Statement
	PAIX     int
	TypeInfo string
}

func (nn *FunctionDeclaration) StatementNode() string {
	nn.TypeInfo = token.FuncD
	return token.FuncD
}
func (nn *FunctionDeclaration) toString() {

}

type FunctionExpression struct {
	Id       Statement
	Params   []*Statement
	Body     Statement
	PAIX     int
	TypeInfo string
}

func (nn *FunctionExpression) StatementNode() string {
	nn.TypeInfo = token.FuncE
	return token.FuncE
}
func (nn *FunctionExpression) toString() {

}

type MemberExpression struct {
	Object   Statement
	Property Statement
	PAIX     int
	TypeInfo string
}

func (nn *MemberExpression) StatementNode() string {
	nn.TypeInfo = token.Member
	return token.Member
}
func (nn *MemberExpression) toString() {

}

type StringLiteral struct {
	Value    string
	PAIX     int
	TypeInfo string
}

func (nn *StringLiteral) StatementNode() string {
	nn.TypeInfo = token.Stri
	return token.Stri
}
func (nn *StringLiteral) toString() {

}

type ThisExpression struct {
	PAIX     int
	TypeInfo string
}

func (nn *ThisExpression) StatementNode() string {
	nn.TypeInfo = token.THIS
	return token.THIS
}
func (nn *ThisExpression) toString() {

}

type BreakStatement struct {
	PAIX     int
	TypeInfo string
}

func (nn *BreakStatement) StatementNode() string {
	nn.TypeInfo = token.BREAK
	return token.BREAK
}
func (nn *BreakStatement) toString() {

}

type ContinueStatement struct {
	PAIX     int
	TypeInfo string
}

func (nn *ContinueStatement) StatementNode() string {
	nn.TypeInfo = token.CONTINUE
	return token.CONTINUE
}
func (nn *ContinueStatement) toString() {

}

type ForStatement struct {
	Init     Statement
	Test     Statement
	Updata   Statement
	Body     Statement
	PAIX     int
	TypeInfo string
}

func (nn *ForStatement) StatementNode() string {
	nn.TypeInfo = token.ForS
	return token.ForS
}

func (nn *ForStatement) toString() {

}

type ForInStatement struct {
	Left     Statement
	Right    Statement
	Body     Statement
	PAIX     int
	TypeInfo string
}

func (nn *ForInStatement) StatementNode() string {
	nn.TypeInfo = token.ForI
	return token.ForI
}

func (nn *ForInStatement) toString() {

}

type ArrayExpression struct {
	Elements []*Statement
	PAIX     int
	TypeInfo string
}

func (nn *ArrayExpression) StatementNode() string {
	nn.TypeInfo = token.ArrayE
	return token.ArrayE
}

func (nn *ArrayExpression) toString() {

}

type ObjectExpression struct {
	Properties []*Statement
	PAIX       int
	TypeInfo   string
}

func (nn *ObjectExpression) StatementNode() string {
	nn.TypeInfo = token.Object
	return token.Object
}

func (nn *ObjectExpression) toString() {

}

type Property struct {
	Key      Statement
	Value    Statement
	PAIX     int
	TypeInfo string
}

func (nn *Property) StatementNode() string {
	nn.TypeInfo = token.Prop
	return token.Prop
}

func (nn *Property) toString() {

}

type NewExpression struct {
	Callee    Statement
	Arguments []*Statement
	PAIX      int
	TypeInfo  string
}

func (nn *NewExpression) StatementNode() string {
	nn.TypeInfo = token.NEW
	return token.NEW
}

func (nn *NewExpression) toString() {

}

type ReturnStatement struct {
	Argument Statement
	PAIX     int
	TypeInfo string
}

func (nn *ReturnStatement) StatementNode() string {
	nn.TypeInfo = token.RETURN
	return token.RETURN
}

func (nn *ReturnStatement) toString() {

}

type DebugStatement struct {
	PAIX     int
	TypeInfo string
}

func (nn *DebugStatement) StatementNode() string {
	nn.TypeInfo = token.Debug
	return token.Debug
}

func (nn *DebugStatement) toString() {

}

type TryStatement struct {
	Block    Statement
	Handler  Statement
	PAIX     int
	TypeInfo string
}

func (nn *TryStatement) StatementNode() string {
	nn.TypeInfo = token.TRY
	return token.TRY
}

func (nn *TryStatement) toString() {

}

type CatchClause struct {
	Param    Statement
	Body     Statement
	PAIX     int
	TypeInfo string
}

func (nn *CatchClause) StatementNode() string {
	nn.TypeInfo = token.CATCH
	return token.CATCH
}

func (nn *CatchClause) toString() {

}
