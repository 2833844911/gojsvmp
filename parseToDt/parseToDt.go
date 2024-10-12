package parseToDt

import (
	"encoding/json"
	"fmt"
	"myvmp/ast"
	"myvmp/token"
)

func PrushData(dt ast.Statement) string {
	jsonString, _ := json.Marshal(dt)
	return string(jsonString)

}

func parseAST(data map[string]interface{}, astInfo *ast.Statement) ast.Statement {
	typee := data["TypeInfo"].(string)
	switch typee {
	case token.Prog:
		dt := &ast.Program{Body: make([]*ast.Statement, 0)}
		kss := data["Body"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Body = append(dt.Body, &lpm2)
		}
		return dt
	case token.IDENT:
		kss := data["Name"].(string)
		dt := &ast.Identifier{Name: kss}
		return dt
	case token.NULL:
		kss := data["Value"].(string)
		dt := &ast.NullIdentifier{Value: kss}
		return dt
	case token.INT:
		kss := data["Value"].(float64)
		dt := &ast.NumericLiteral{Value: kss}
		return dt
	case token.VAR:
		dt := &ast.VariableDeclaration{Declarations: make([]*ast.Statement, 0)}
		kss := data["Declarations"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Declarations = append(dt.Declarations, &lpm2)
		}
		dt.Init = parseAST(data["Init"].(map[string]interface{}), nil)
		return dt
	case token.Bin:
		left := parseAST(data["Left"].(map[string]interface{}), nil)
		right := parseAST(data["Right"].(map[string]interface{}), nil)
		Operator := data["Operator"].(string)
		dt := &ast.BinaryExpression{Operator: Operator, Left: left, Right: right}
		return dt
	case token.NOP:
		return &ast.NOP{}
	case token.OVER:
		return &ast.OVER{}
	case token.Ass:
		left := parseAST(data["Left"].(map[string]interface{}), nil)
		right := parseAST(data["Right"].(map[string]interface{}), nil)
		Operator := data["Operator"].(string)
		dt := &ast.AssignmentExpression{Operator: Operator, Left: left, Right: right}
		return dt
	case token.Call:
		dt := &ast.CallExpression{Arguments: make([]*ast.Statement, 0)}
		kss := data["Arguments"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Arguments = append(dt.Arguments, &lpm2)
		}
		Caller := parseAST(data["Caller"].(map[string]interface{}), nil)
		dt.Caller = Caller
		return dt
	case token.IfStat:
		Test := parseAST(data["Test"].(map[string]interface{}), nil)
		Consequent := parseAST(data["Consequent"].(map[string]interface{}), nil)
		var Alternate ast.Statement
		if data["Alternate"] != nil {
			Alternate = parseAST(data["Alternate"].(map[string]interface{}), nil)

		} else {
			Alternate = nil
		}
		dt := &ast.IfStatement{Test: Test, Consequent: Consequent, Alternate: Alternate}
		return dt
	case token.Block:
		dt := &ast.BlockStatement{Body: make([]*ast.Statement, 0)}
		kss := data["Body"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Body = append(dt.Body, &lpm2)
		}
		return dt
	case token.Unary:
		Argument := parseAST(data["Argument"].(map[string]interface{}), nil)
		Prefix := data["Prefix"].(bool)
		Operator := data["Operator"].(string)
		dt := &ast.UnaryExpression{
			Argument: Argument,
			Operator: Operator,
			Prefix:   Prefix,
		}
		return dt
	case token.FuncD:
		Id := parseAST(data["Id"].(map[string]interface{}), nil)
		Body := parseAST(data["Body"].(map[string]interface{}), nil)
		dt := &ast.FunctionDeclaration{Id: Id, Body: Body, Params: make([]*ast.Statement, 0)}
		kss := data["Params"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Params = append(dt.Params, &lpm2)
		}
		return dt
	case token.FuncE:
		var Id ast.Statement
		if data["Id"] != nil {
			Id = parseAST(data["Id"].(map[string]interface{}), nil)
		} else {
			Id = nil
		}
		Body := parseAST(data["Body"].(map[string]interface{}), nil)
		dt := &ast.FunctionExpression{Id: Id, Body: Body, Params: make([]*ast.Statement, 0)}
		kss := data["Params"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Params = append(dt.Params, &lpm2)
		}
		return dt
	case token.Member:
		Object := parseAST(data["Object"].(map[string]interface{}), nil)
		Property := parseAST(data["Property"].(map[string]interface{}), nil)
		dt := &ast.MemberExpression{
			Object:   Object,
			Property: Property,
		}
		return dt
	case token.Stri:
		kss := data["Value"].(string)
		dt := &ast.StringLiteral{
			Value: kss,
		}
		return dt
	case token.THIS:
		tthis := &ast.ThisExpression{}
		return tthis
	case token.BREAK:
		dt := &ast.BreakStatement{}
		return dt
	case token.CONTINUE:
		dt := &ast.ContinueStatement{}
		return dt
	case token.ForS:
		var Init ast.Statement
		var Test ast.Statement
		var Updata ast.Statement
		if data["Init"] != nil {
			Init = parseAST(data["Init"].(map[string]interface{}), nil)
		}

		if data["Test"] != nil {
			Test = parseAST(data["Test"].(map[string]interface{}), nil)
		}

		if data["Updata"] != nil {
			Updata = parseAST(data["Updata"].(map[string]interface{}), nil)
		}
		Body := parseAST(data["Body"].(map[string]interface{}), nil)
		dt := &ast.ForStatement{
			Init:   Init,
			Body:   Body,
			Updata: Updata,
			Test:   Test,
		}
		return dt
	case token.ForI:
		var Init ast.Statement
		var Test ast.Statement
		if data["Left"] != nil {
			Init = parseAST(data["Left"].(map[string]interface{}), nil)
		}

		if data["Right"] != nil {
			Test = parseAST(data["Right"].(map[string]interface{}), nil)
		}

		Body := parseAST(data["Body"].(map[string]interface{}), nil)
		dt := &ast.ForInStatement{
			Left:  Init,
			Body:  Body,
			Right: Test,
		}
		return dt
	case token.ArrayE:
		dt := &ast.ArrayExpression{Elements: make([]*ast.Statement, 0)}
		kss := data["Elements"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Elements = append(dt.Elements, &lpm2)
		}
		return dt
	case token.Object:
		dt := &ast.ObjectExpression{Properties: make([]*ast.Statement, 0)}
		kss := data["Properties"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Properties = append(dt.Properties, &lpm2)
		}
		return dt
	case token.Prop:
		Key := parseAST(data["Key"].(map[string]interface{}), nil)
		Value := parseAST(data["Value"].(map[string]interface{}), nil)
		dt := &ast.Property{Key: Key, Value: Value}
		return dt
	case token.TRY:
		Block := parseAST(data["Block"].(map[string]interface{}), nil)
		Handler := parseAST(data["Handler"].(map[string]interface{}), nil)
		dt := &ast.TryStatement{Block: Block, Handler: Handler}
		return dt
	case token.CATCH:
		var Param ast.Statement
		if Param != nil {
			Param = parseAST(data["Param"].(map[string]interface{}), nil)
		} else {
			Param = nil
		}

		Body := parseAST(data["Body"].(map[string]interface{}), nil)
		dt := &ast.CatchClause{Param: Param, Body: Body}
		return dt

	case token.RETURN:
		Argument := parseAST(data["Argument"].(map[string]interface{}), nil)
		dt := &ast.ReturnStatement{Argument: Argument}
		return dt
	case token.Debug:
		dt := &ast.DebugStatement{}
		return dt
	case token.NEW:
		Callee := parseAST(data["Callee"].(map[string]interface{}), nil)

		dt := &ast.NewExpression{Callee: Callee, Arguments: make([]*ast.Statement, 0)}
		kss := data["Arguments"].([]interface{})
		for _, value := range kss {
			lpp := value.(map[string]interface{})
			lpm2 := parseAST(lpp, nil)

			dt.Arguments = append(dt.Arguments, &lpm2)
		}
		return dt
	}
	return nil
}

func LoadStr(dd string) ast.Statement {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(dd), &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return parseAST(data, nil)
}
