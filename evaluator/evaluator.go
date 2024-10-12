package evaluator

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"math"
	"myvmp/ast"
	"myvmp/banding"
	"myvmp/lexer"
	"myvmp/object"
	"myvmp/parse"
	"myvmp/parseToDt"
	"myvmp/promise"
	"myvmp/require"
	"myvmp/token"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

var Zhidy map[string]*func([]*object.Object) object.Object

func binString(dtt string, dttr string, typee string) object.Object {
	switch typee {
	case token.ADD:
		return &object.StringObject{Value: dtt + dttr}
	case token.SDD: // 这里可能是返回数字
		return &object.NanObject{}
	case token.CHEN: // 这里可能是返回数字
		return &object.NanObject{}
	case token.CHU: // 这里可能是返回数字
		return &object.NanObject{}
	case token.XIAOYHYH:
		return &object.NumericObject{Value: 0}
	case token.XIAOYHYHYH:
		return &object.NumericObject{Value: 0}
	case token.DAYHYH:
		return &object.NumericObject{Value: 0}
	case token.YIHUO:
		return &object.NumericObject{Value: 0}
	case token.DAYHYHYU:
		return &object.NumericObject{Value: 0}
	case token.XIAND:
		return &object.BoolObject{Value: dtt == dttr}
	case token.XIAOYH:
		return &object.BoolObject{Value: dtt < dttr}
	case token.DAYH:
		return &object.BoolObject{Value: dtt > dttr}
	case token.DAYHDY:
		return &object.BoolObject{Value: dtt >= dttr}
	case token.XIAOYHDY:
		return &object.BoolObject{Value: dtt <= dttr}
	case token.BUDY:
		return &object.BoolObject{Value: dtt != dttr}
	case token.BUDYDY:
		return &object.BoolObject{Value: dtt != dttr}

	}
	return &object.NanObject{}
}

func binInt(dtt float64, dttr float64, typee string) object.Object {
	switch typee {
	case token.ADD:
		return &object.NumericObject{Value: dtt + dttr}
	case token.SDD: // 这里可能是返回数字
		return &object.NumericObject{Value: dtt - dttr}
	case token.KAIFAN: // 这里可能是返回数字
		num := math.Pow(dtt, dttr)
		return &object.NumericObject{Value: num}
	case token.CHEN: // 这里可能是返回数字
		return &object.NumericObject{Value: dtt * dttr}
	case token.CHU: // 这里可能是返回数字
		return &object.NumericObject{Value: dtt / dttr}
	case token.XIAOYHYH:
		return &object.NumericObject{Value: float64(int64(dtt) << int64(dttr))}
	case token.XIAOYHYHYH:
		dsd := float64(int64(dtt)<<int64(dttr) + 100)
		return &object.NumericObject{Value: dsd}
	case token.DAYHYH:
		return &object.NumericObject{Value: float64(int64(dtt) >> int64(dttr))}
	case token.QUYU:
		return &object.NumericObject{Value: float64(int64(dtt) % int64(dttr))}
	case token.DAYHYHYU:
		return &object.NumericObject{Value: float64(int64(dtt) >> int64(dttr))}
	case token.XIAND:
		return &object.BoolObject{Value: dtt == dttr}
	case token.XIAOYH:
		return &object.BoolObject{Value: dtt < dttr}
	case token.DAYH:
		return &object.BoolObject{Value: dtt > dttr}
	case token.DAYHDY:
		return &object.BoolObject{Value: dtt >= dttr}
	case token.YIHUO:
		return &object.NumericObject{Value: float64(int(dtt) ^ int(dttr))}
	case token.XIAOYHDY:
		return &object.BoolObject{Value: dtt <= dttr}
	case token.BUDY:
		return &object.BoolObject{Value: dtt != dttr}
	case token.BUDYDY:
		return &object.BoolObject{Value: dtt != dttr}
	case token.HUO:
		return &object.BoolObject{Value: dtt != 0 || dttr != 0}
	case token.HUOHUO:
		ds := &object.NumericObject{Value: dtt}
		var dsj object.Object = ds
		dsr := &object.NumericObject{Value: dttr}
		var dsjr object.Object = dsr
		if getBoolInfo(&dsj).Value {
			return ds
		} else if getBoolInfo(&dsjr).Value {
			return dsr
		}

		return &object.BoolObject{Value: false}
	case token.YU:
		return &object.BoolObject{Value: dtt != 0 && dttr != 0}
	case token.YUYU:
		return &object.BoolObject{Value: dtt != 0 && dttr != 0}

	}
	return nil
}

func yunxBin(dtt *ast.Statement, env *object.Environment) object.Object {
	ddd := (*dtt).(*ast.BinaryExpression)

	leftData := ParseData(&ddd.Left, env)
	if ddd.Operator == token.YUYU {
		jj := getBoolInfo(&leftData)
		if !jj.Value {
			return &jj
		}

	} else if ddd.Operator == token.HUOHUO {
		jj := getBoolInfo(&leftData)
		if jj.Value {
			return leftData
		}
	}
	rightData := ParseData(&ddd.Right, env)

	if leftData.Type() == token.TYNUM && rightData.Type() == token.TYNUM {
		return binInt(leftData.(*object.NumericObject).Value, rightData.(*object.NumericObject).Value, ddd.Operator)
	} else if leftData.Type() == token.BOOL && rightData.Type() == token.BOOL {
		var zuo float64
		if leftData.(*object.BoolObject).Value {
			zuo = 1
		} else {
			zuo = 0
		}
		var yuo float64
		if rightData.(*object.BoolObject).Value {
			yuo = 1
		} else {
			yuo = 0
		}

		return binInt(zuo, yuo, ddd.Operator)
	} else if leftData.Type() == token.TYNUM && rightData.Type() == token.BOOL {
		var zuo float64
		zuo = leftData.(*object.NumericObject).Value
		var yuo float64
		if rightData.(*object.BoolObject).Value {
			yuo = 1
		} else {
			yuo = 0
		}

		return binInt(zuo, yuo, ddd.Operator)
	} else if leftData.Type() == token.BOOL && rightData.Type() == token.TYNUM {
		var zuo float64
		if leftData.(*object.BoolObject).Value {
			zuo = 1
		} else {
			zuo = 0
		}
		var yuo float64
		yuo = rightData.(*object.NumericObject).Value

		return binInt(zuo, yuo, ddd.Operator)
	} else {
		if rightData.Type() == token.ArrayE && leftData.Type() == token.ArrayE && (ddd.Operator == token.DXIAND || ddd.Operator == token.XIAND || ddd.Operator == token.BUDY || ddd.Operator == token.BUDYDY) {
			nkl := rightData.(*object.Environment)
			nkle := leftData.(*object.Environment)
			switch ddd.Operator {
			case token.DXIAND:
				return &object.BoolObject{Value: nkl == nkle}
			case token.XIAND:
				return &object.BoolObject{Value: nkl == nkle}
			case token.BUDY:
				return &object.BoolObject{Value: nkl != nkle}
			case token.BUDYDY:
				return &object.BoolObject{Value: nkl != nkle}

			}

		} else if rightData.Type() == token.Object && leftData.Type() == token.Object && (ddd.Operator == token.DXIAND || ddd.Operator == token.XIAND || ddd.Operator == token.BUDY || ddd.Operator == token.BUDYDY) {
			nkl := rightData.(*object.Environment)
			nkle := leftData.(*object.Environment)
			switch ddd.Operator {
			case token.DXIAND:
				return &object.BoolObject{Value: nkl == nkle}
			case token.XIAND:
				return &object.BoolObject{Value: nkl == nkle}
			case token.BUDY:
				return &object.BoolObject{Value: nkl != nkle}
			case token.BUDYDY:
				return &object.BoolObject{Value: nkl != nkle}

			}
		} else if (rightData.Type() == token.ENV || rightData.Type() == token.THIS) && (leftData.Type() == token.ENV || rightData.Type() == token.THIS) && (ddd.Operator == token.DXIAND || ddd.Operator == token.XIAND || ddd.Operator == token.BUDY || ddd.Operator == token.BUDYDY) {
			nkl := rightData.(*object.Environment)
			nkle := leftData.(*object.Environment)
			switch ddd.Operator {
			case token.DXIAND:
				return &object.BoolObject{Value: nkl == nkle}
			case token.XIAND:
				return &object.BoolObject{Value: nkl == nkle}
			case token.BUDY:
				return &object.BoolObject{Value: nkl != nkle}
			case token.BUDYDY:
				return &object.BoolObject{Value: nkl != nkle}

			}
		}
		rightStr := rightData.ToString()
		leftStr := leftData.ToString()
		return binString(leftStr, rightStr, ddd.Operator)
	}
	return rightData
}
func yunxINT(dtt *ast.Statement) object.Object {
	ddd := (*dtt).(*ast.NumericLiteral).Value
	return &object.NumericObject{
		Value: ddd,
	}
}

func yunxStr(dtt *ast.Statement) object.Object {
	ddd := (*dtt).(*ast.StringLiteral).Value
	return &object.StringObject{
		Value: ddd,
	}
}

func yunxVar(dtt *ast.Statement, env *object.Environment) object.Object {
	ddd := (*dtt).(*ast.VariableDeclaration)
	qianmdy := ddd.Declarations
	hmd := ParseData(&ddd.Init, env)
	for _, key := range qianmdy {
		//(*env).Store[key.Name] = hmd
		ddd := (*key).(*ast.Identifier)
		(*env).Store.Set(ddd.Name, hmd)
	}
	return nil
}
func yunxIDENT(dtt *ast.Statement, env *object.Environment) object.Object {
	ddd := (*dtt).(*ast.Identifier)
	data := findKey(ddd.Name, env)
	return data
}

func getBoolInfo(dt *object.Object) object.BoolObject {
	if *dt == nil {
		return object.BoolObject{Value: false}
	}
	tb := (*dt).Type()
	switch tb {
	case token.BOOL:
		return *((*dt).(*object.BoolObject))
	case token.TYNUM:
		dtee := (*dt).(*object.NumericObject)
		if dtee.Value == 0 {
			return object.BoolObject{Value: false}
		} else {
			return object.BoolObject{Value: true}
		}
	case token.NULL:
		return object.BoolObject{Value: false}
	case token.NANINFO:
		return object.BoolObject{Value: false}

	}
	return object.BoolObject{Value: true}
}

func hebing(as ast.Statement, rightData object.Object, optt string, env *object.Environment) object.Object {
	switch optt {
	case token.DENYU:
		return rightData
	case token.JIADEN:
		leftData := ParseData(&as, env)
		if leftData.Type() == token.TYNUM && rightData.Type() == token.TYNUM {
			return binInt(leftData.(*object.NumericObject).Value, rightData.(*object.NumericObject).Value, token.ADD)
		} else {
			rightStr := rightData.ToString()
			leftStr := leftData.ToString()
			return binString(leftStr, rightStr, token.ADD)
		}
	case token.JANDEN:
		leftData := ParseData(&as, env)
		if leftData.Type() == token.TYNUM && rightData.Type() == token.TYNUM {
			return binInt(leftData.(*object.NumericObject).Value, rightData.(*object.NumericObject).Value, token.SDD)
		} else {
			rightStr := rightData.ToString()
			leftStr := leftData.ToString()
			return binString(leftStr, rightStr, token.SDD)
		}

	}
	return nil

}
func yunxAss(dtt *ast.Statement, env *object.Environment) object.Object {
	ddd := (*dtt).(*ast.AssignmentExpression)
	Valuee := ParseData(&ddd.Right, env)
	Value := hebing(ddd.Left, Valuee, ddd.Operator, env)
	key := ddd.Left //之后需要判断是否对象
	if key.StatementNode() == token.IDENT {
		data := setKey(key.(*ast.Identifier).Name, Value, env)
		return data
	}
	if key.StatementNode() == token.Member {
		keyValue := key.(*ast.MemberExpression)
		mnnn := ParseData(&keyValue.Object, env)
		zkey := ParseData(&keyValue.Property, env)

		if mnnn.Type() == token.BYTE {
			zkey2 := zkey.(*object.NumericObject)
			Value2 := Value.(*object.NumericObject)
			dmmm := mnnn.(*object.ByteObject)
			dmmm.Value[int(zkey2.Value)] = byte(Value2.Value)
		} else {
			nj := mnnn.(*object.Environment)
			if nj.Type() == token.ArrayE && zkey.Type() == token.TYNUM {
				zkey2 := zkey.(*object.NumericObject)
				if int(zkey2.Value) >= len(nj.Value) {
					// 扩展切片到足够的长度
					newSlice := make([]*object.Object, int(zkey2.Value)+1)
					copy(newSlice, nj.Value)
					nj.Value = newSlice
				}
				nj.Value[int(zkey2.Value)] = &Value
			} else {
				var strkey string
				if zkey.Type() == token.TYNUM {
					zkey2 := zkey.(*object.NumericObject)
					strkey = strconv.Itoa(int(zkey2.Value))
				} else {
					zkey2 := zkey.(*object.StringObject)
					strkey = zkey2.Value
				}
				//nj.Store[strkey] = Value
				nj.Store.Set(strkey, Value)
				//setKey(strkey, Value, nj)
			}
		}

	}
	return Value
}

func yunxUpINfo(key ast.Statement, env *object.Environment, Value object.Object) object.Object {
	if key.StatementNode() == token.IDENT {
		data := setKey(key.(*ast.Identifier).Name, Value, env)
		return data
	}
	if key.StatementNode() == token.Member {
		keyValue := key.(*ast.MemberExpression)
		mnnn := ParseData(&keyValue.Object, env)
		zkey := ParseData(&keyValue.Property, env)

		nj := mnnn.(*object.Environment)
		if nj.Type() == token.ArrayE && zkey.Type() == token.TYNUM {
			zkey2 := zkey.(*object.NumericObject)
			if int(zkey2.Value) >= len(nj.Value) {
				// 扩展切片到足够的长度
				newSlice := make([]*object.Object, int(zkey2.Value)+1)
				copy(newSlice, nj.Value)
				nj.Value = newSlice
			}
			nj.Value[int(zkey2.Value)] = &Value
		} else {
			var strkey string
			if zkey.Type() == token.TYNUM {
				zkey2 := zkey.(*object.NumericObject)
				strkey = strconv.Itoa(int(zkey2.Value))
			} else {
				zkey2 := zkey.(*object.StringObject)
				strkey = zkey2.Value
			}
			//nj.Store[strkey] = Value
			nj.Store.Set(strkey, Value)
			//setKey(strkey, Value, nj)
		}

	}
	return Value
}

func yunxIfStat(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.IfStatement)
	test := ParseData(&dt.Test, env)
	boInfo := getBoolInfo(&test)
	if boInfo.Value {
		ds := ParseData(&dt.Consequent, env)
		if ds != nil && ds.Type() == token.BREAK {
			return ds
		} else if ds != nil && ds.Type() == token.CONTINUE {
			return ds
		} else if ds != nil && ds.Type() == token.RETURN {
			return ds
		}
	} else {

		ds := ParseData(&dt.Alternate, env)
		if ds != nil && ds.Type() == token.BREAK {
			return ds
		} else if ds != nil && ds.Type() == token.CONTINUE {
			return ds
		} else if ds != nil && ds.Type() == token.RETURN {
			return ds
		}
	}

	return nil
}

func yunxBlock(dtt *ast.Statement, envup *object.Environment) object.Object {
	env := object.NewEnv(envup)
	blockList := (*dtt).(*ast.BlockStatement).Body
	for _, dt := range blockList {
		dg := ParseData(dt, env)
		if dg != nil && dg.Type() == token.BREAK {
			return dg
		} else if dg != nil && dg.Type() == token.CONTINUE {
			return dg
		} else if dg != nil && dg.Type() == token.RETURN {
			return dg
		}
	}
	return nil
}

func yunxBleak(dtt *ast.Statement, envup *object.Environment) object.Object {

	return &object.BreakObject{}
}
func yunxCONTINUE(dtt *ast.Statement, envup *object.Environment) object.Object {

	return &object.ContinueObject{}
}

func yunxUnary(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.UnaryExpression)
	qiand := dt.Operator
	switch qiand {
	case token.TYPEOF:
		dgg := ParseData(&dt.Argument, env)
		return &object.StringObject{Value: dgg.Type()}
	case token.QUFAN:
		dgg := ParseData(&dt.Argument, env)
		boolinfo := getBoolInfo(&dgg)
		return &object.BoolObject{Value: !boolinfo.Value}
	case token.UPADD:
		dgg := ParseData(&dt.Argument, env)
		ds := dgg.Type()
		switch ds {
		case token.TYNUM:
			ks := dgg.(*object.NumericObject)

			//ji := ks.Value
			var dada float64
			if dt.Prefix {
				dada = ks.Value + 1
			} else {
				dada = ks.Value
			}

			yunxUpINfo(dt.Argument, env, &object.NumericObject{Value: ks.Value + 1})
			return &object.NumericObject{Value: dada}
		default:
			return &object.NanObject{}

		}
	case token.UPASD:
		dgg := ParseData(&dt.Argument, env)
		ds := dgg.Type()
		switch ds {
		case token.TYNUM:
			ks := dgg.(*object.NumericObject)

			var dada float64
			fmt.Println(dt.Prefix)
			if dt.Prefix {
				dada = ks.Value - 1
			} else {
				dada = ks.Value
			}

			yunxUpINfo(dt.Argument, env, &object.NumericObject{Value: ks.Value - 1})
			return &object.NumericObject{Value: dada}
		default:
			return &object.NanObject{}

		}
	case token.SDD:
		dgg := ParseData(&dt.Argument, env)
		ddddd := dgg.(*object.NumericObject)
		return &object.NumericObject{Value: -ddddd.Value}

	}

	return nil
}

func yunxArrayE(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.ArrayExpression)
	dtte := object.NewArray()
	for _, cle := range dt.Elements {
		lppem := ParseData(cle, env)
		dtte.Value = append(dtte.Value, &lppem)
	}

	return &dtte
}

func yunxFuncD(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.FunctionDeclaration)
	idN := dt.Id.(*ast.Identifier)
	dddfff := object.FunctionDeclarationObject{Params: dt.Params, Body: dt.Body, Env: env}

	//env.Store[idN.Name] = &dddfff
	env.Store.Set(idN.Name, &dddfff)

	return nil
}

func yunxFuncE(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.FunctionExpression)

	dddfff := object.FunctionDeclarationObject{Params: dt.Params, Body: dt.Body, Env: env}

	return &dddfff
}

func yunxRETURN(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.ReturnStatement)
	retuVal := ParseData(&dt.Argument, env)
	dddfff := object.ReturnStatementObject{}
	dddfff.Value = retuVal

	return &dddfff
}
func yunxCall(dtt *ast.Statement, env *object.Environment) object.Object {

	dt := (*dtt).(*ast.CallExpression)
	//if dt.Caller.StatementNode() == token.IDENT {
	//	bbb := dt.Caller.(*ast.Identifier).Name
	//
	//}

	funct := ParseData(&dt.Caller, env)
	dbb := []*object.Object{}

	for _, hh := range dt.Arguments {
		arg := ParseData(hh, env)
		dbb = append(dbb, &arg)
	}
	if funct.Type() == token.BULLE && dt.Caller.StatementNode() == token.IDENT {
		jiii := dt.Caller.(*ast.Identifier)

		vay, ok := Zhidy[jiii.Name]
		if ok {
			return (*vay)(dbb)
		}
	}

	if funct.Type() == token.TYSTR {
		kppp := funct.(*object.StringObject)
		if kppp.Key == token.Slice {
			ldasd0 := (*dbb[0]).(*object.NumericObject)
			ldasd1 := (*dbb[1]).(*object.NumericObject)
			return &object.StringObject{
				Value: kppp.Slice(int(ldasd0.Value), int(ldasd1.Value)),
			}
		}
	} else if funct.Type() == token.ArrayE {
		kppp := funct.(*object.Environment)
		if kppp.Key == token.Slice {
			ldasd0 := (*dbb[0]).(*object.NumericObject)
			ldasd1 := (*dbb[1]).(*object.NumericObject)
			kkoo := kppp.Slice(int(ldasd0.Value), int(ldasd1.Value))
			nkk := object.NewArray()
			nkk.Value = kkoo
			return &nkk
		}
	}

	//异步太快问题
	dasd := 0
	for {
		if funct.Type() != token.NULL {
			break
		}
		dasd += 1
		if dasd > 100000000 {
			fmt.Println("没有找到可以执行的函数")
			os.Exit(0)
		}

	}

	kfunob := (funct).(*object.FunctionDeclarationObject)
	kfunob.Args = dbb
	var fuenv *object.Environment
	if dt.Caller.StatementNode() == token.Member {

		//if kfunob.Env == nil {
		//	dgg := dt.Caller.(*ast.MemberExpression)
		//	kff := ParseData(&dgg.Object, env)
		//	if kff.Type() == token.THIS {
		//		bjjj := kff.(*object.Environment)
		//		fuenv = object.NewEnv(bjjj)
		//	}
		//} else {
		//	fuenv = kfunob.Env
		//}
		fuenv = kfunob.Env

	} else {
		fuenv = object.NewEnv(env)
	}
	if kfunob.Callthis == 1 {
		dadshkj := (*dbb[0]).(*object.Environment)

		fuenv = dadshkj
		dbb = dbb[1:]
	} else if kfunob.Callthis == 2 {
		dadshkj := (*dbb[0]).(*object.Environment)

		fuenv = dadshkj
		dasd := (*dbb[1]).(*object.Environment)

		dbb = dasd.Value
	}

	for idx, vkey := range kfunob.Params {
		zzkey := (*vkey).(*ast.Identifier).Name
		if idx >= len(dbb) {
			//fuenv.Store[zzkey] = &object.NULLObject{}
			fuenv.Store.Set(zzkey, &object.NULLObject{})
			continue
		}
		//fuenv.Store[zzkey] = *dbb[idx]
		fuenv.Store.Set(zzkey, *dbb[idx])

	}
	knn := object.NewArray()
	knn.Value = dbb
	fuenv.Store.Set(token.Arguments, &knn)

	if kfunob.IsNative == 1 {
		return (*kfunob.NativeBody)(kfunob)
	}

	djii := ParseData(&kfunob.Body, fuenv)
	if djii == nil {
		return &object.NULLObject{}
	}
	if djii.Type() == token.RETURN {
		dd := djii.(*object.ReturnStatementObject)

		return dd.Value
	}

	return djii
}
func yunxMember(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.MemberExpression)
	left := ParseData(&dt.Object, env)
	right := ParseData(&dt.Property, env)
	//fmt.Println(left)
	//if left == nil {
	//	fmt.Println()
	//}

	if left.Type() == token.FUN && right.Type() == token.TYSTR {
		dasda := left.(*object.FunctionDeclarationObject)
		dasd := right.(*object.StringObject).Value
		if dasd == token.CALL {

			dasda.Callthis = 1
			return dasda
		} else if dasd == token.APPLY {
			dasda.Callthis = 2
			return dasda
		}

	}

	if left.Type() == token.ArrayE {
		dsg := left.(*object.Environment)
		if right.Type() == token.TYNUM {
			hjj := right.(*object.NumericObject)
			lpINt := int(hjj.Value)
			if len(dsg.Value) <= lpINt {
				return &object.NULLObject{}
			}
			lpffffp := dsg.Value[lpINt]
			if lpffffp == nil {
				return &object.NULLObject{}
			}
			return *lpffffp
		} else {

			hjj := right.(*object.StringObject)
			lpffffp, _ := dsg.Store.Get(hjj.Value)

			if lpffffp == nil {

				if right.ToString() == token.Slice {
					lkkk := left.(*object.Environment)
					lkkk.Key = token.Slice
					return left
				} else if right.ToString() == token.Length {
					lkkk := left.(*object.Environment)
					ddd := &object.NumericObject{Value: float64(len(lkkk.Value))}
					return ddd
				}
				return &object.NULLObject{}
			}
			if lpffffp.Type() == token.FUN {
				jj := left.(*object.Environment)
				lpffffp.(*object.FunctionDeclarationObject).Env = jj
			}
			return lpffffp
		}
	} else {

		if left.Type() == token.BYTE && right.Type() == token.TYSTR && right.ToString() == token.Length {
			dd := left.(*object.ByteObject)
			return &object.NumericObject{Value: float64(len(dd.Value))}
		}
		if left.Type() == token.BYTE && right.Type() == token.TYNUM {
			dd := left.(*object.ByteObject)
			ddIdx := right.(*object.NumericObject).Value
			return &object.NumericObject{Value: float64(dd.Value[int(ddIdx)])}
		}

		var key string
		if right.Type() == token.TYNUM {
			if left.Type() == token.TYSTR {
				hjj := right.(*object.NumericObject)
				lkkk := left.(*object.StringObject)
				// 将字符串转换为 rune 切片
				runes := []rune(lkkk.Value)
				return &object.StringObject{Value: string(runes[int64(hjj.Value):int64(hjj.Value+1)])}
			} else {
				hjj := right.(*object.NumericObject)
				key = strconv.Itoa(int(hjj.Value))
			}

		} else {

			if left.Type() == token.TYSTR {
				if right.ToString() == token.Length {
					lkkk := left.(*object.StringObject)
					ddd := &object.NumericObject{Value: float64(utf8.RuneCountInString(lkkk.Value))}
					return ddd
				}
				//hjj := right.(*object.StringObject)
				lkkk := left.(*object.StringObject)
				lkkk.Key = token.Slice
				return left

			} else {
				hjj := right.(*object.StringObject)
				key = hjj.Value
			}

		}
		jh := findKey(key, left.(*object.Environment))
		if jh == nil || jh.Type() == token.NULL {
			if left.Type() == token.Object && right.ToString() == token.Length {
				lkkk := left.(*object.Environment)
				ddd := &object.NumericObject{Value: float64(len(lkkk.Store.M))}
				return ddd
			}
			return &object.NULLObject{}
		}

		if jh.Type() == token.FUN {
			jj := left.(*object.Environment)
			jh.(*object.FunctionDeclarationObject).Env = jj
		}

		return jh
	}
}

func yunxFOR(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.ForStatement)
	init := dt.Init

	test := dt.Test
	updata := dt.Updata
	body := dt.Body

	if init != nil {
		ParseData(&init, env)
	}
	for {
		if test != nil {
			lppp := ParseData(&test, env)
			if !getBoolInfo(&lppp).Value {
				break
			}
		}
		kf := ParseData(&body, env)
		if updata != nil {
			ParseData(&updata, env)
		}
		if kf != nil && kf.Type() == token.BREAK {
			break
		} else if kf != nil && kf.Type() == token.CONTINUE {

			continue
		} else if kf != nil && kf.Type() == token.RETURN {
			return kf
		}
	}

	return nil
}

func yunxFOI(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.ForInStatement)
	init := dt.Left

	right := ParseData(&dt.Right, env)
	dtall := make([]object.Object, 0)
	if right.Type() == token.Object || right.Type() == token.ENV || right.Type() == token.THIS {
		dsgre := right.(*object.Environment)
		for key, _ := range dsgre.Store.M {
			dtall = append(dtall, &object.StringObject{Value: key})
		}

	} else if token.ArrayE == right.Type() {
		dsgre := right.(*object.Environment)
		for i := 0; i < len(dsgre.Value); i++ {
			dtall = append(dtall, &object.NumericObject{Value: float64(i)})
		}
	} else {
		fmt.Println("can not for in the object !")
		os.Exit(0)
	}

	body := dt.Body

	if init != nil {
		ParseData(&init, env)
	}
	needgbkey := (*init.(*ast.VariableDeclaration).Declarations[0]).(*ast.Identifier).Name

	for u := 0; u < len(dtall); u++ {
		env.Store.Set(needgbkey, dtall[u])
		kf := ParseData(&body, env)

		if kf != nil && kf.Type() == token.BREAK {
			break
		} else if kf != nil && kf.Type() == token.CONTINUE {

			continue
		} else if kf != nil && kf.Type() == token.RETURN {
			return kf
		}
	}

	return nil
}

func yunxObject(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.ObjectExpression)
	dtte := object.NewObject()
	dtte.Outer = env
	for _, cle := range dt.Properties {
		das := (*cle).(*ast.Property)
		key := das.Key.StatementNode()
		var keyv string
		if key == token.Stri {
			lppp := das.Key
			keyv = lppp.(*ast.StringLiteral).Value
		} else if key == token.IDENT {
			lppp := das.Key
			keyv = lppp.(*ast.Identifier).Name
		}
		lppem := ParseData(&das.Value, env)
		dtte.Store.Set(keyv, lppem)
	}

	return &dtte
}

func Try(fn func(*ast.Statement, *object.Environment) object.Object, ass *ast.Statement, dddd *object.Environment) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("panic: %v", r)
			}
		}
	}()
	fn(ass, dddd)
	return nil
}

func Catch(err error, handler func(error)) {
	if err != nil {
		handler(err)
	}
}

func yunxTRY(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.TryStatement)
	dtte := object.NewEnv(env)

	err := Try(ParseData, &dt.Block, dtte)

	Catch(err, func(e error) {
		// This is the catch block to handle errors
		cadtte := object.NewEnv(env)
		dadlcatch := dt.Handler.(*ast.CatchClause)
		if dadlcatch.Param != nil {
			cadtte.Store.Set(dadlcatch.Param.(*ast.Identifier).Name, &object.StringObject{Value: e.Error()})
		}
		ParseData(&dadlcatch.Body, cadtte)
	})
	return &object.NULLObject{}
}
func yunxTHIS(dtt *ast.Statement, env *object.Environment) object.Object {
	if env.Type() == token.THIS {
		return env
	}
	if env.Outer != nil {
		return yunxTHIS(dtt, env.Outer)
	}
	return env
}
func yunxNEW(dtt *ast.Statement, env *object.Environment) object.Object {
	dt := (*dtt).(*ast.NewExpression)
	funct := ParseData(&dt.Callee, env)
	dbb := []*object.Object{}

	for _, hh := range dt.Arguments {
		arg := ParseData(hh, env)
		dbb = append(dbb, &arg)
	}

	//if funct.Type() == token.BULLE && dt.Callee.StatementNode() == token.IDENT {
	//	dasd := dt.Callee.(*ast.Identifier)
	//	if dasd.Name == token.Promise {
	//		ji := ParseData
	//		return promise.Init(&ji, dbb[0], env)
	//	}
	//
	//}

	kfunob := (funct).(*object.FunctionDeclarationObject)
	kfunob.Args = dbb
	fuenv := object.NewEnv(env)
	fuenv.TypeInfo = token.THIS
	for idx, vkey := range kfunob.Params {
		zzkey := (*vkey).(*ast.Identifier).Name
		if idx >= len(dbb) {
			fuenv.Store.Set(zzkey, &object.NULLObject{})
			continue
		}
		fuenv.Store.Set(zzkey, *dbb[idx])
	}
	if kfunob.IsNative == 1 {
		kfunob.Env = fuenv
		return (*kfunob.NativeBody)(kfunob)
	} else {
		ParseData(&kfunob.Body, fuenv)
	}

	return fuenv
}
func Promise(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	dbb := dt[0]
	ji := ParseData
	env := myfun.Env
	return promise.Init(&ji, dbb, env)
}
func Wait(myfun *object.FunctionDeclarationObject) object.Object {
	promise.Done()
	return &object.NULLObject{}

}
func ParseData(dtt *ast.Statement, env *object.Environment) object.Object {
	if *dtt == nil {
		return nil
	}
	Leix := (*dtt).StatementNode()
	//fmt.Println(Leix)
	switch Leix {
	case token.Bin:
		return yunxBin(dtt, env)
	case token.Unary:
		return yunxUnary(dtt, env)
	case token.INT:
		return yunxINT(dtt)
	case token.Stri:
		return yunxStr(dtt)
	case token.VAR:
		return yunxVar(dtt, env)
	case token.IDENT:
		return yunxIDENT(dtt, env)
	case token.Ass:
		return yunxAss(dtt, env)
	case token.IfStat:
		return yunxIfStat(dtt, env)
	case token.ForS:
		return yunxFOR(dtt, env)
	case token.ForI:
		return yunxFOI(dtt, env)
	case token.Block:
		return yunxBlock(dtt, env)
	case token.BREAK:
		return yunxBleak(dtt, env)
	case token.CONTINUE:
		return yunxCONTINUE(dtt, env)
	case token.ArrayE:
		return yunxArrayE(dtt, env)
	case token.TRY:
		return yunxTRY(dtt, env)
	case token.Debug:
		for {
			fmt.Printf("断点中:")
			scanner := bufio.NewScanner(os.Stdin)

			// 读取输入
			scanner.Scan()
			input := scanner.Text()
			input = strings.TrimSpace(input)
			// 检查是否有错误
			if err := scanner.Err(); err != nil {
				fmt.Println("读取输入时发生错误:", err)
				return nil
			}
			zl := ""
			if len(input) >= 2 && input[1:2] == ":" {
				zl = input[0:1]
				input = input[1:]
			} else if len(input) > 0 {
				zl = "w"
			}
			if zl == "c" {
				return nil
			} else if zl == "w" {
				sz := EvalTG(input+";", env)
				fmt.Println(sz)
			}
		}

		return nil
	case token.Member:
		return yunxMember(dtt, env)
	case token.Object:
		return yunxObject(dtt, env)
	case token.FuncD:
		return yunxFuncD(dtt, env)
	case token.FuncE:
		return yunxFuncE(dtt, env)
	case token.Call:
		return yunxCall(dtt, env)
	case token.RETURN:
		return yunxRETURN(dtt, env)
	case token.THIS:
		return yunxTHIS(dtt, env)
	case token.NEW:
		return yunxNEW(dtt, env)
	}
	return nil
}
func findKey(key string, env *object.Environment) object.Object {
	value, ok := env.Store.Get(key)
	if ok {

		return value
	}
	shanc := env.Outer
	if shanc == nil {

		return &object.NULLObject{}
	}
	return findKey(key, shanc)
}

func setKey(key string, value object.Object, env *object.Environment) object.Object {

	_, ok := env.Store.Get(key)
	if ok {
		env.Store.Set(key, value)
		return value

	}

	shanc := env.Outer
	if shanc == nil {
		env.Store.Set(key, value)
		return value
	}
	return setKey(key, value, shanc)
}

func EvalDDD(myfun *object.FunctionDeclarationObject) object.Object {
	// 参数一： js代码
	// 参数二： this环境
	dt := myfun.Args
	code := (*dt[0]).(*object.StringObject).Value + ";"
	dsf := (*dt[1]).(*object.Environment)
	dtfg := lexer.New(code)
	kk := (*dtfg).Input()
	fff := parse.NewParse(kk)
	var ff object.Object
	for _, dtee := range fff {
		ff = ParseData(dtee, dsf)
	}
	if ff == nil {
		return &object.NULLObject{}
	}
	return ff
}

//func getCyDt(this js.Value, inputs []js.Value) interface{} {
//	jii := inputs[0].String()
//	alldt, _ := allenv.Store.Get("startFun")
//	kd := alldt.(*object.FunctionDeclarationObject)
//	env := object.NewEnv(allenv)
//	env.Store.Set("cbu", &object.StringObject{Value: jii})
//	allenv.Store.Set("l", &object.StringObject{Value: jii})
//
//	ParseData(&kd.Body, env)
//
//	return js.ValueOf("data in window.cydata")
//}
//
//func registerCallbacks() {
//	js.Global().Set("getCyDt", js.FuncOf(getCyDt))
//}

func Require(myfun *object.FunctionDeclarationObject) object.Object {
	dte := myfun.Args
	daoirfile := (*dte[0]).ToString()

	if daoirfile[len(daoirfile)-3:] != ".js" {
		daoirfile += ".js"
	}

	dsdaskj, okk := requireall.DtInfo[daoirfile]
	if okk {
		return dsdaskj
	}

	requireall.IsDo = daoirfile
	code := require.ReadFile(daoirfile)
	env := object.NewEnv(nil)

	dt := lexer.New(code)
	kk := (*dt).Input()
	fff := parse.NewParse(kk)
	StartEval(fff, env)
	requireall.IsDo = ""
	requireall.DtInfo[daoirfile] = env
	return env
}

var allenv *object.Environment
var requireall *require.RequireInfo

func new_Func(ddd *func(*object.FunctionDeclarationObject) object.Object) object.Object {
	d := &object.FunctionDeclarationObject{IsNative: 1, NativeBody: ddd}
	return d
}
func StartEval(data []*ast.Statement, env *object.Environment) object.Object {
	ghsa, qbhj := banding.Init()
	for key, dsd := range ghsa {
		env.Store.Set(key, dsd)
	}
	eval := EvalDDD
	require := Require
	promise := Promise
	wait := Wait

	env.Store.Set(token.Eval, new_Func(&eval))
	env.Store.Set(token.Require, new_Func(&require))
	env.Store.Set(token.Promise, new_Func(&promise))
	env.Store.Set(token.Wait, new_Func(&wait))
	var dddcbbbb object.Object
	allenv = env
	env.TypeInfo = token.THIS
	for key, vlu := range qbhj {
		//env.Store[key] = vlu
		env.Store.Set(key, vlu)
	}
	for _, dt := range data {
		dddcbbbb = ParseData(dt, env)
	}
	//registerCallbacks()
	if dddcbbbb != nil {
		fmt.Println(dddcbbbb.ToString())
	}

	return dddcbbbb
}

func Eval(code string) object.Object {
	if requireall == nil {

		requireall = &require.RequireInfo{DtInfo: map[string]*object.Environment{}}
	}
	env := object.NewEnv(nil)
	promise.CyJSInit()
	dt := lexer.New(code)
	kk := (*dt).Input()
	fff := parse.NewParse(kk)
	//registerCallbacks()
	gg := StartEval(fff, env)
	promise.Done()

	return gg
}

func EvalTG(code string, env *object.Environment) object.Object {

	dt := lexer.New(code)
	kk := (*dt).Input()
	fff := parse.NewParse(kk)
	//registerCallbacks()
	gg := StartEval(fff, env)

	return gg
}

func EvalDt() object.Object {

	encodedText := `eyJCb2R5IjpbeyJJZCI6eyJOYW1lIjoiZGF0YSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlBhcmFtcyI6W3siTmFtZSI6InN0ckluZm8iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiQm9keSI6eyJCb2R5IjpbeyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6InN0ckxpc3QiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkVsZW1lbnRzIjpbXSwiUEFJWCI6MCwiVHlwZUluZm8iOiJwcHBwcCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiSW5pdCI6eyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImkiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LCJUZXN0Ijp7IkxlZnQiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJDYWxsZXIiOnsiTmFtZSI6ImxlbiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIkFyZ3VtZW50cyI6W3siTmFtZSI6InN0ckluZm8iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIk9wZXJhdG9yIjoiXHUwMDNjIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiVXBkYXRhIjp7IkxlZnQiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjQsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiKyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIk9wZXJhdG9yIjoiPSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidXV1In0sIkJvZHkiOnsiQm9keSI6W3siVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJjZCIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiQ2FsbGVyIjp7Ik5hbWUiOiJsZW4iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJzdHJJbmZvIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJ2aGgiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkxlZnQiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6NCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IlRlc3QiOnsiTGVmdCI6eyJOYW1lIjoidmhoIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTmFtZSI6ImNkIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiT3BlcmF0b3IiOiJcdTAwM2U9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiQ29uc2VxdWVudCI6eyJCb2R5IjpbeyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6InQxIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJDYWxsZXIiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJzdHJJbmZvIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOiJzbGljZSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vb28ifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIkFyZ3VtZW50cyI6W3siTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LHsiQ2FsbGVyIjp7Ik5hbWUiOiJsZW4iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJzdHJJbmZvIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJJbml0Ijp7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiaTIiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LCJUZXN0Ijp7IkxlZnQiOnsiTmFtZSI6ImkyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTGVmdCI6eyJMZWZ0Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjQsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiKyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlJpZ2h0Ijp7Ik5hbWUiOiJjZCIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIk9wZXJhdG9yIjoiLSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIk9wZXJhdG9yIjoiXHUwMDNjIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiVXBkYXRhIjp7IkxlZnQiOnsiTmFtZSI6ImkyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTGVmdCI6eyJOYW1lIjoiaTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifSwiQm9keSI6eyJCb2R5IjpbeyJMZWZ0Ijp7Ik5hbWUiOiJ0MSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IkxlZnQiOnsiTmFtZSI6InQxIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOiIgIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vbyJ9LCJPcGVyYXRvciI6IisiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJPcGVyYXRvciI6Ij0iLCJQQUlYIjowLCJUeXBlSW5mbyI6InV1dSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJwcHBwIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImExIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJDYWxsZXIiOnsiTmFtZSI6ImN5Y2hhciIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIkFyZ3VtZW50cyI6W3siT2JqZWN0Ijp7Ik5hbWUiOiJ0MSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJhMiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiQ2FsbGVyIjp7Ik5hbWUiOiJjeWNoYXIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik9iamVjdCI6eyJOYW1lIjoidDEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiYTMiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkNhbGxlciI6eyJOYW1lIjoiY3ljaGFyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJPYmplY3QiOnsiTmFtZSI6InQxIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOjIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImE0IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJDYWxsZXIiOnsiTmFtZSI6ImN5Y2hhciIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIkFyZ3VtZW50cyI6W3siT2JqZWN0Ijp7Ik5hbWUiOiJ0MSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjozLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJoaiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiRWxlbWVudHMiOlt7Ik5hbWUiOiJhMSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0seyJOYW1lIjoiYTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LHsiTmFtZSI6ImEzIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSx7Ik5hbWUiOiJhNCIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6InBwcHBwIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJDYWxsZXIiOnsiTmFtZSI6ImN5YXBwZW5kIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoic3RyTGlzdCIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0seyJOYW1lIjoiaGoiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0seyJQQUlYIjowLCJUeXBlSW5mbyI6ImJyZWFrIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpaSJ9LCJBbHRlcm5hdGUiOm51bGwsIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWkifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoidDEiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkNhbGxlciI6eyJPYmplY3QiOnsiTmFtZSI6InN0ckluZm8iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6InNsaWNlIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vbyJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0seyJMZWZ0Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjQsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiKyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiYTEiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkNhbGxlciI6eyJOYW1lIjoiY3ljaGFyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJPYmplY3QiOnsiTmFtZSI6InQxIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOjAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImEyIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJDYWxsZXIiOnsiTmFtZSI6ImN5Y2hhciIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIkFyZ3VtZW50cyI6W3siT2JqZWN0Ijp7Ik5hbWUiOiJ0MSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjoxLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJhMyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiQ2FsbGVyIjp7Ik5hbWUiOiJjeWNoYXIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik9iamVjdCI6eyJOYW1lIjoidDEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiYTQiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkNhbGxlciI6eyJOYW1lIjoiY3ljaGFyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJPYmplY3QiOnsiTmFtZSI6InQxIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOjMsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImhqIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJFbGVtZW50cyI6W3siTmFtZSI6ImExIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSx7Ik5hbWUiOiJhMiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0seyJOYW1lIjoiYTMiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LHsiTmFtZSI6ImE0IiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcHAifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IkNhbGxlciI6eyJOYW1lIjoiY3lhcHBlbmQiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJzdHJMaXN0IiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSx7Ik5hbWUiOiJoaiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcCJ9LHsiQXJndW1lbnQiOnsiTmFtZSI6InN0ckxpc3QiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InJldHVybiJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWlpaSJ9LHsiSWQiOnsiTmFtZSI6InN0YXJ0RnVuIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUGFyYW1zIjpbeyJOYW1lIjoiY2J1IiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkJvZHkiOnsiQm9keSI6W3siTGVmdCI6eyJOYW1lIjoibCIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7Ik5hbWUiOiJjYnUiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJPcGVyYXRvciI6Ij0iLCJQQUlYIjowLCJUeXBlSW5mbyI6InV1dSJ9LHsiTGVmdCI6eyJOYW1lIjoibCIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IkxlZnQiOnsiTmFtZSI6ImwiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6ImNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYmNiYiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vb28ifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoibHNibDEiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IklkIjpudWxsLCJQYXJhbXMiOltdLCJCb2R5Ijp7IkJvZHkiOlt7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoia2YiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkNhbGxlciI6eyJOYW1lIjoiZGF0YSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIkFyZ3VtZW50cyI6W3siTmFtZSI6ImwiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJDYWxsZXIiOnsiTmFtZSI6ImNiYl9hIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoia2YiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vbyJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJrcyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiQ2FsbGVlIjp7Ik5hbWUiOiJQcm9taXNlIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoibHNibDEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJuZXcifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoia3MyIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJDYWxsZXIiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJrcyIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjoidGhlbiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vb28ifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIkFyZ3VtZW50cyI6W3siSWQiOm51bGwsIlBhcmFtcyI6W3siTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiQm9keSI6eyJCb2R5IjpbeyJJbml0Ijp7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiaSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiVmFsdWUiOjAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0sIlRlc3QiOnsiTGVmdCI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IkNhbGxlciI6eyJOYW1lIjoibGVuIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSwiT3BlcmF0b3IiOiJcdTAwM2MiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJVcGRhdGEiOnsiTGVmdCI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IkxlZnQiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifSwiQm9keSI6eyJCb2R5IjpbeyJJbml0Ijp7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiaTIiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LCJUZXN0Ijp7IkxlZnQiOnsiTmFtZSI6ImkyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiQ2FsbGVyIjp7Ik5hbWUiOiJsZW4iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJPcGVyYXRvciI6Ilx1MDAzYyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlVwZGF0YSI6eyJMZWZ0Ijp7Ik5hbWUiOiJpMiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IkxlZnQiOnsiTmFtZSI6ImkyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjEsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiKyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIk9wZXJhdG9yIjoiPSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidXV1In0sIkJvZHkiOnsiQm9keSI6W3siVGVzdCI6eyJMZWZ0Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTmFtZSI6ImkyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiT3BlcmF0b3IiOiI9PSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIkNvbnNlcXVlbnQiOnsiQm9keSI6W3siUEFJWCI6MCwiVHlwZUluZm8iOiJjb250aW51ZSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiQWx0ZXJuYXRlIjpudWxsLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImFsbGNkIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUHJvcGVydHkiOnsiVmFsdWUiOjAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUHJvcGVydHkiOnsiVmFsdWUiOjEsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUHJvcGVydHkiOnsiVmFsdWUiOjIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJPYmplY3QiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiTmFtZSI6ImkyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjozLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IkxlZnQiOnsiT2JqZWN0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUmlnaHQiOnsiTGVmdCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7Ik5hbWUiOiJhbGxjZCIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIk9wZXJhdG9yIjoiXiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlJpZ2h0Ijp7IlZhbHVlIjoyMCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifSx7IkxlZnQiOnsiT2JqZWN0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjoxLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUmlnaHQiOnsiTGVmdCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7Ik5hbWUiOiJhbGxjZCIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIk9wZXJhdG9yIjoiXiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlJpZ2h0Ijp7IlZhbHVlIjoxMCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifSx7IlRlc3QiOnsiTGVmdCI6eyJOYW1lIjoiYWxsY2QiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MzQwLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJPcGVyYXRvciI6Ilx1MDAzZSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIkNvbnNlcXVlbnQiOnsiQm9keSI6W3siTGVmdCI6eyJPYmplY3QiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUHJvcGVydHkiOnsiVmFsdWUiOjEsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7IlZhbHVlIjo4MCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIkFsdGVybmF0ZSI6bnVsbCwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaSJ9LHsiTGVmdCI6eyJPYmplY3QiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUHJvcGVydHkiOnsiVmFsdWUiOjIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJMZWZ0Ijp7IkxlZnQiOnsiT2JqZWN0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjoyLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUmlnaHQiOnsiTmFtZSI6ImFsbGNkIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiT3BlcmF0b3IiOiJeIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiUmlnaHQiOnsiVmFsdWUiOjEsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiKyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIk9wZXJhdG9yIjoiPSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidXV1In0seyJMZWZ0Ijp7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MywiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7IkxlZnQiOnsiTGVmdCI6eyJPYmplY3QiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUHJvcGVydHkiOnsiVmFsdWUiOjMsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJOYW1lIjoiYWxsY2QiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJPcGVyYXRvciI6Il4iLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJSaWdodCI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcCJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJwcHBwIn0seyJDYWxsZXIiOnsiTmFtZSI6ImNiYl9iIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiZHNkZCIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiQ2FsbGVyIjp7Ik9iamVjdCI6eyJOYW1lIjoia3MyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOiJ0aGVuIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vbyJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiQXJndW1lbnRzIjpbeyJJZCI6bnVsbCwiUGFyYW1zIjpbeyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJCb2R5Ijp7IkJvZHkiOlt7IkNhbGxlciI6eyJOYW1lIjoiY3lvdXQiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSx7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSx7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb28ifSx7IklkIjpudWxsLCJQYXJhbXMiOlt7Ik5hbWUiOiJhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkJvZHkiOnsiQm9keSI6W3siVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJhbGxkYXRhIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJFbGVtZW50cyI6W10sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcHAifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IkluaXQiOnsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJpIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJWYWx1ZSI6MCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSwiVGVzdCI6eyJMZWZ0Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiQ2FsbGVyIjp7Ik5hbWUiOiJsZW4iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJPcGVyYXRvciI6Ilx1MDAzYyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlVwZGF0YSI6eyJMZWZ0Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTGVmdCI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IlZhbHVlIjoxLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJPcGVyYXRvciI6IisiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJPcGVyYXRvciI6Ij0iLCJQQUlYIjowLCJUeXBlSW5mbyI6InV1dSJ9LCJCb2R5Ijp7IkJvZHkiOlt7IkluaXQiOnsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJpMiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiVmFsdWUiOjAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0sIlRlc3QiOnsiTGVmdCI6eyJOYW1lIjoiaTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6NCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiJcdTAwM2MiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJVcGRhdGEiOnsiTGVmdCI6eyJOYW1lIjoiaTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik5hbWUiOiJpMiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IlZhbHVlIjoxLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJPcGVyYXRvciI6IisiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJPcGVyYXRvciI6Ij0iLCJQQUlYIjowLCJUeXBlSW5mbyI6InV1dSJ9LCJCb2R5Ijp7IkJvZHkiOlt7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiZHBwcHAiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkxlZnQiOnsiT2JqZWN0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlByb3BlcnR5Ijp7Ik5hbWUiOiJpMiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJWYWx1ZSI6NSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIlIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiamlpIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJWYWx1ZSI6MTAwLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiQmxvY2siOnsiQm9keSI6W3siTGVmdCI6eyJOYW1lIjoiamlpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTGVmdCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJPYmplY3QiOnsiTmFtZSI6ImEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUmlnaHQiOnsiTmFtZSI6ImRwcHBwIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiT3BlcmF0b3IiOiJcdTAwM2VcdTAwM2UiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJSaWdodCI6eyJWYWx1ZSI6MywiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIqIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIkhhbmRsZXIiOnsiUGFyYW0iOm51bGwsIkJvZHkiOnsiQm9keSI6W3siTGVmdCI6eyJOYW1lIjoiamlpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTGVmdCI6eyJPYmplY3QiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUHJvcGVydHkiOnsiTmFtZSI6ImkyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7IlZhbHVlIjo0MDAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiJSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIk9wZXJhdG9yIjoiPSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidXV1In1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6ImNhdGNoIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidHJ5In0seyJDYWxsZXIiOnsiTmFtZSI6ImN5YXBwZW5kIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiYWxsZGF0YSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0seyJMZWZ0Ijp7Ik5hbWUiOiJqaWkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MzAwMCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIlIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LHsiVGVzdCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYWxsZGF0YSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IkxlZnQiOnsiQ2FsbGVyIjp7Ik5hbWUiOiJsZW4iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJhbGxkYXRhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJSaWdodCI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiItIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7IlZhbHVlIjoxMDAwLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJPcGVyYXRvciI6Ilx1MDAzZSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIkNvbnNlcXVlbnQiOnsiQm9keSI6W3siTGVmdCI6eyJPYmplY3QiOnsiTmFtZSI6ImFsbGRhdGEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJMZWZ0Ijp7IkNhbGxlciI6eyJOYW1lIjoibGVuIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiYWxsZGF0YSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSwiUmlnaHQiOnsiVmFsdWUiOjEsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiLSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYWxsZGF0YSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IkxlZnQiOnsiQ2FsbGVyIjp7Ik5hbWUiOiJsZW4iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJhbGxkYXRhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJSaWdodCI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiItIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7IlZhbHVlIjoxLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJPcGVyYXRvciI6IisiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJPcGVyYXRvciI6Ij0iLCJQQUlYIjowLCJUeXBlSW5mbyI6InV1dSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiQWx0ZXJuYXRlIjpudWxsLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InBwcHAifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcCJ9LHsiTGVmdCI6eyJPYmplY3QiOnsiTmFtZSI6ImFsbGRhdGEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6MCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7IkxlZnQiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJhbGxkYXRhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOjAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJWYWx1ZSI6MjU1LCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJPcGVyYXRvciI6IiUiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJPcGVyYXRvciI6Ij0iLCJQQUlYIjowLCJUeXBlSW5mbyI6InV1dSJ9LHsiSWQiOnsiTmFtZSI6ImdpZCIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlBhcmFtcyI6W3siTmFtZSI6ImlyciIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJCb2R5Ijp7IkJvZHkiOlt7IklkIjp7Ik5hbWUiOiJkamtpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUGFyYW1zIjpbXSwiQm9keSI6eyJCb2R5IjpbeyJUZXN0Ijp7IkxlZnQiOnsiTmFtZSI6ImlyciIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYWxsZGF0YSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiT3BlcmF0b3IiOiIhPSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIkNvbnNlcXVlbnQiOnsiQm9keSI6W3siQXJndW1lbnQiOnsiVmFsdWUiOjAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicmV0dXJuIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpaSJ9LCJBbHRlcm5hdGUiOm51bGwsIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWkifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiZHNodSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiRWxlbWVudHMiOltdLCJQQUlYIjowLCJUeXBlSW5mbyI6InBwcHBwIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJJbml0Ijp7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiaSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiVmFsdWUiOjAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0sIlRlc3QiOnsiTGVmdCI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IlZhbHVlIjozMiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiJcdTAwM2MiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJVcGRhdGEiOnsiTGVmdCI6eyJOYW1lIjoiaSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IkxlZnQiOnsiTmFtZSI6ImkiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifSwiQm9keSI6eyJCb2R5IjpbeyJMZWZ0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiZHNodSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7Ik5hbWUiOiJpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlJpZ2h0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYWxsZGF0YSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcCJ9LHsiSW5pdCI6eyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6Iml2dnYiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IlZhbHVlIjowLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LCJUZXN0Ijp7IkxlZnQiOnsiTmFtZSI6Iml2dnYiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJDYWxsZXIiOnsiTmFtZSI6ImxlbiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIkFyZ3VtZW50cyI6W3siTmFtZSI6ImFsbGRhdGEiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIk9wZXJhdG9yIjoiXHUwMDNjIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiVXBkYXRhIjp7IkxlZnQiOnsiTmFtZSI6Iml2dnYiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik5hbWUiOiJpdnZ2IiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjEsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiKyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIk9wZXJhdG9yIjoiPSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidXV1In0sIkJvZHkiOnsiQm9keSI6W3siVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJoc2lvYWQiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkxlZnQiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJhbGxkYXRhIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiTmFtZSI6Iml2dnYiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUmlnaHQiOnsiVmFsdWUiOjIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiKiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJMZWZ0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiZHNodSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IkxlZnQiOnsiTmFtZSI6Iml2dnYiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MzIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiJSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik5hbWUiOiJoc2lvYWQiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJMZWZ0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiZHNodSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IkxlZnQiOnsiTmFtZSI6Iml2dnYiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MzIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiJSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vbyJ9LCJSaWdodCI6eyJWYWx1ZSI6MiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiJcdTAwM2VcdTAwM2UiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJPcGVyYXRvciI6IisiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJPcGVyYXRvciI6Ij0iLCJQQUlYIjowLCJUeXBlSW5mbyI6InV1dSJ9LHsiTGVmdCI6eyJPYmplY3QiOnsiTmFtZSI6ImRzaHUiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJMZWZ0Ijp7Ik5hbWUiOiJpdnZ2IiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjMyLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJPcGVyYXRvciI6IiUiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUmlnaHQiOnsiTGVmdCI6eyJPYmplY3QiOnsiTmFtZSI6ImRzaHUiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJMZWZ0Ijp7Ik5hbWUiOiJpdnZ2IiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjMyLCJQQUlYIjowLCJUeXBlSW5mbyI6InR0dCJ9LCJPcGVyYXRvciI6IiUiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUmlnaHQiOnsiVmFsdWUiOjI1NSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIlIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcCJ9LHsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJhbGxkdCIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiUHJvcGVydGllcyI6W3siS2V5Ijp7IlZhbHVlIjoiZGF0YSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vb28ifSwiVmFsdWUiOnsiTmFtZSI6ImRzaHUiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6ImxsbGwifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcHBwIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6InZodWkiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkNhbGxlciI6eyJPYmplY3QiOnsiTmFtZSI6IkpTT04iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJWYWx1ZSI6InN0cmluZ2lmeSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vb28ifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIkFyZ3VtZW50cyI6W3siTmFtZSI6ImFsbGR0IiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiQ2FsbGVyIjp7Ik5hbWUiOiJjYmJfYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIkFyZ3VtZW50cyI6W3siTmFtZSI6ImFsbGR0IiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSx7Ik5hbWUiOiJsIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWlpaSJ9LHsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJrZmMiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkNhbGxlZSI6eyJOYW1lIjoiUHJvbWlzZSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIkFyZ3VtZW50cyI6W3siTmFtZSI6ImRqa2kiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJuZXcifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiaGlobyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiQ2FsbGVyIjp7Ik9iamVjdCI6eyJOYW1lIjoia2ZjIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOiJ0aGVuIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vbyJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiQXJndW1lbnRzIjpbeyJJZCI6bnVsbCwiUGFyYW1zIjpbeyJOYW1lIjoiYWxsZHQyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSx7Ik5hbWUiOiJsMiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJCb2R5Ijp7IkJvZHkiOlt7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoiZHNodTIiLCJQQUlYIjowLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJJbml0Ijp7IkVsZW1lbnRzIjpbXSwiUEFJWCI6MCwiVHlwZUluZm8iOiJwcHBwcCJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiSW5pdCI6eyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6Iml2dnYyIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJWYWx1ZSI6MCwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSwiVGVzdCI6eyJMZWZ0Ijp7Ik5hbWUiOiJpdnZ2MiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlJpZ2h0Ijp7IkNhbGxlciI6eyJOYW1lIjoibGVuIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoibDIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIk9wZXJhdG9yIjoiXHUwMDNjIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiVXBkYXRhIjp7IkxlZnQiOnsiTmFtZSI6Iml2dnYyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTGVmdCI6eyJOYW1lIjoiaXZ2djIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiIrIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiT3BlcmF0b3IiOiI9IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ1dXUifSwiQm9keSI6eyJCb2R5IjpbeyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImhzaW9hZCIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiQ2FsbGVyIjp7Ik5hbWUiOiJjeWNoYXIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik9iamVjdCI6eyJOYW1lIjoibDIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQcm9wZXJ0eSI6eyJOYW1lIjoiaXZ2djIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InZhciJ9LHsiVG9rZW4iOiJ2YXIiLCJEZWNsYXJhdGlvbnMiOlt7Ik5hbWUiOiJoZGFrZGFzaiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiTGVmdCI6eyJOYW1lIjoiaXZ2djIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MzIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiJSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6Imd5cyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiT2JqZWN0Ijp7Ik9iamVjdCI6eyJOYW1lIjoiYWxsZHQyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOiJkYXRhIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vbyJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiUHJvcGVydHkiOnsiTmFtZSI6ImhkYWtkYXNqIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImZzZyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiTGVmdCI6eyJOYW1lIjoiaHNpb2FkIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTmFtZSI6Imd5cyIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIk9wZXJhdG9yIjoiXiIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJDYWxsZXIiOnsiTmFtZSI6ImN5YXBwZW5kIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiZHNodTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LHsiTmFtZSI6ImZzZyIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSx7IlRlc3QiOnsiTGVmdCI6eyJOYW1lIjoiaXZ2djIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJSaWdodCI6eyJWYWx1ZSI6MzIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiXHUwMDNjIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJ5eXkifSwiQ29uc2VxdWVudCI6eyJCb2R5IjpbeyJDYWxsZXIiOnsiTmFtZSI6ImN5YXBwZW5kIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiZHNodTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LHsiTmFtZSI6Imd5cyIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIkFsdGVybmF0ZSI6bnVsbCwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJwcHBwIn0seyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImR0IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiSW5pdCI6eyJQcm9wZXJ0aWVzIjpbeyJLZXkiOnsiVmFsdWUiOiJkYXRhIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vbyJ9LCJWYWx1ZSI6eyJOYW1lIjoiZHNodTIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6ImxsbGwifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoicHBwcHBwIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJDYWxsZXIiOnsiTmFtZSI6ImNiYl9hIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiZHQiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vbyJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0seyJDYWxsZXIiOnsiT2JqZWN0Ijp7Ik5hbWUiOiJoaWhvIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUHJvcGVydHkiOnsiVmFsdWUiOiJ0aGVuIiwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vbyJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiQXJndW1lbnRzIjpbeyJJZCI6bnVsbCwiUGFyYW1zIjpbeyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJCb2R5Ijp7IkJvZHkiOlt7IlRva2VuIjoidmFyIiwiRGVjbGFyYXRpb25zIjpbeyJOYW1lIjoidmh1aSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiQ2FsbGVyIjp7Ik9iamVjdCI6eyJOYW1lIjoiSlNPTiIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn0sIlByb3BlcnR5Ijp7IlZhbHVlIjoic3RyaW5naWZ5IiwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb29vbyJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6Im9vb28ifSwiQXJndW1lbnRzIjpbeyJOYW1lIjoiYSIsIlBBSVgiOi0xLCJUeXBlSW5mbyI6ImFhYWFhIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifSx7IkNhbGxlciI6eyJOYW1lIjoiZ2V0VG9rZW4iLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJ2aHVpIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJvb28ifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpaWlpaSJ9LHsiSW5pdCI6eyJUb2tlbiI6InZhciIsIkRlY2xhcmF0aW9ucyI6W3siTmFtZSI6ImlyciIsIlBBSVgiOjAsIlR5cGVJbmZvIjoiYWFhYWEifV0sIkluaXQiOnsiVmFsdWUiOjAsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIlBBSVgiOjAsIlR5cGVJbmZvIjoidmFyIn0sIlRlc3QiOnsiTGVmdCI6eyJOYW1lIjoiaXJyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjI1NSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ0dHQifSwiT3BlcmF0b3IiOiJcdTAwM2MiLCJQQUlYIjowLCJUeXBlSW5mbyI6Inl5eSJ9LCJVcGRhdGEiOnsiTGVmdCI6eyJOYW1lIjoiaXJyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiTGVmdCI6eyJOYW1lIjoiaXJyIiwiUEFJWCI6LTEsIlR5cGVJbmZvIjoiYWFhYWEifSwiUmlnaHQiOnsiVmFsdWUiOjEsIlBBSVgiOjAsIlR5cGVJbmZvIjoidHR0In0sIk9wZXJhdG9yIjoiKyIsIlBBSVgiOjAsIlR5cGVJbmZvIjoieXl5In0sIk9wZXJhdG9yIjoiPSIsIlBBSVgiOjAsIlR5cGVJbmZvIjoidXV1In0sIkJvZHkiOnsiQm9keSI6W3siQ2FsbGVyIjp7Ik5hbWUiOiJnaWQiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9LCJBcmd1bWVudHMiOlt7Ik5hbWUiOiJpcnIiLCJQQUlYIjotMSwiVHlwZUluZm8iOiJhYWFhYSJ9XSwiUEFJWCI6MCwiVHlwZUluZm8iOiJpaWlpIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWlpaSJ9LCJQQUlYIjowLCJUeXBlSW5mbyI6InBwcHAifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoib29vIn1dLCJQQUlYIjowLCJUeXBlSW5mbyI6ImlpaWkifSwiUEFJWCI6MCwiVHlwZUluZm8iOiJ2YXIifV0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpIn0sIlBBSVgiOjAsIlR5cGVJbmZvIjoiaWlpaWlpaWkifV0sIlR5cGVJbmZvIjoiYWFhYSJ9`
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		fmt.Println("Decode error:", err)
		return nil
	}
	dd := string(decodedBytes)
	sdsgh := sync.WaitGroup{}
	sdsgh.Add(1)
	go func() {
		defer sdsgh.Done()
		promise.CyJSInit()
		ss := parseToDt.LoadStr(dd)
		dyy := ss.(*ast.Program)
		env := object.NewEnv(nil)
		StartEval(dyy.Body, env)

		promise.Done()

	}()
	sdsgh.Wait()

	return nil
}
