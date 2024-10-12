package parse

import (
	"fmt"
	"myvmp/ast"
	"myvmp/token"
	"os"
	"strconv"
)

type allDtInfo struct {
	Start    int
	ShangOne ast.Statement
	Alldtd   []*token.TokenType
}

func (oo *allDtInfo) getDt() *token.TokenType {
	v := oo.Alldtd[oo.Start]
	oo.Start++
	return v
}

func (oo *allDtInfo) getNextkey() *token.TokenType {
	v := oo.Alldtd[oo.Start]
	return v
}

func newdt(vhh []*token.TokenType) *allDtInfo {
	data := &allDtInfo{Start: 0, Alldtd: vhh}
	return data

}

func (oo *allDtInfo) isover() bool {
	if oo.Start >= len(oo.Alldtd) {
		return true
	}
	return false
}

func parseBinaryExpression(alldata *allDtInfo, opert *token.TokenType) ast.Statement {
	huu := alldata.ShangOne
	alldata.ShangOne = nil
	Right := parseData(alldata)
	alldata.ShangOne = huu
	hds := alldata.getNextkey()
	if opert.PAXU <= hds.PAXU && alldata.getNextkey().TypeInfo != token.OVER && alldata.getNextkey().TypeInfo != token.YOUOK && alldata.getNextkey().TypeInfo != token.ZUOZ && alldata.getNextkey().TypeInfo != token.DH && alldata.getNextkey().TypeInfo != token.HUOHUO && alldata.getNextkey().TypeInfo != token.YUYU {
		aldt := &ast.BinaryExpression{Left: alldata.ShangOne}
		aldt.Operator = opert.TypeInfo
		alldata.ShangOne = Right
		mytkdd := alldata.getDt()
		Left := parseBinaryExpression(alldata, mytkdd)
		aldt.Right = Left
		return aldt
	} else {
		aldt := &ast.BinaryExpression{Left: alldata.ShangOne}
		aldt.Operator = opert.TypeInfo
		aldt.Right = Right
		alldata.ShangOne = aldt
		if alldata.getNextkey().TypeInfo == token.OVER || alldata.getNextkey().TypeInfo == token.YOUOK || alldata.getNextkey().TypeInfo == token.ZUOZ || alldata.getNextkey().TypeInfo == token.DH || alldata.getNextkey().TypeInfo == token.YUYU || alldata.getNextkey().TypeInfo == token.HUOHUO {
			sseee := alldata.ShangOne
			return sseee
		}
		hu := parseData(alldata)

		var dsadasd ast.Statement
		if hu == nil {
			dsadasd = aldt
		} else {
			dsadasd = hu
		}
		return dsadasd
	}
}

func parseYpuxianKuoHao(alldata *allDtInfo) ast.Statement {
	alldata.ShangOne = nil
	for {
		dtt := parseData(alldata)

		if alldata.getNextkey().TypeInfo == token.YOUOK {
			alldata.Start++
			zbrt := dtt
			alldata.ShangOne = dtt
			var dtife ast.Statement
			var dtif ast.Statement
			dtife = zbrt
			if alldata.getNextkey().TypeInfo == token.ZHUOK {
				alldata.ShangOne = dtife
				alldata.Start++
				dtif = parseCallE(alldata)
			} else if alldata.getNextkey().TypeInfo == token.Dian {
				alldata.ShangOne = dtife
				alldata.Start++
				dtif = parseMemberExpression(alldata, false)
			} else if alldata.getNextkey().TypeInfo == token.YOUZ {
				alldata.ShangOne = dtife
				alldata.Start++
				dtif = parseMemberExpression(alldata, true)
			} else if alldata.getNextkey().TypeInfo == token.DENYU {
				alldata.ShangOne = nil
				alldata.Start++
				dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.DENYU}
				dds := parseAss(alldata)
				dshhh.Right = dds
				dtif = dshhh
			} else if alldata.getNextkey().TypeInfo == token.JIADEN {
				alldata.ShangOne = nil
				alldata.Start++
				dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.JIADEN}
				dds := parseAss(alldata)
				dshhh.Right = dds
				return dshhh
			} else if alldata.getNextkey().TypeInfo == token.JANDEN {
				alldata.ShangOne = nil
				alldata.Start++
				dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.JANDEN}
				dds := parseAss(alldata)
				dshhh.Right = dds
				return dshhh
			} else {
				dtif = dtife
			}
			alldata.ShangOne = dtif
			//huff := parseBin(alldata)
			//if huff == nil {
			//	dtif = dtif
			//} else {
			//	dtif = huff
			//}
			return dtif

		}
		alldata.ShangOne = dtt
	}

}
func parseCallE(alldata *allDtInfo) ast.Statement {
	caleer := alldata.ShangOne
	alldata.ShangOne = nil
	Calle := &ast.CallExpression{Caller: caleer, Arguments: make([]*ast.Statement, 0)}
	if alldata.getNextkey().TypeInfo != token.YOUOK {

		for {
			dtt := parseData(alldata)

			if alldata.getNextkey().TypeInfo == token.DH {
				alldata.Start++
				kp := dtt
				Calle.Arguments = append(Calle.Arguments, &kp)

				alldata.ShangOne = nil
				continue
			}
			if alldata.getNextkey().TypeInfo == token.YOUOK {

				alldata.Start++
				kp := dtt
				alldata.ShangOne = nil
				Calle.Arguments = append(Calle.Arguments, &kp)
				break
			}
			alldata.ShangOne = dtt

		}
	} else {
		alldata.Start++
	}

	//alldata.ShangOne = Calle
	dtife := Calle
	var dtif ast.Statement
	if alldata.getNextkey().TypeInfo == token.ZHUOK {
		alldata.ShangOne = dtife
		alldata.Start++
		dtif = parseCallE(alldata)
	} else if alldata.getNextkey().TypeInfo == token.Dian {
		alldata.ShangOne = dtife
		alldata.Start++
		dtif = parseMemberExpression(alldata, false)
	} else if alldata.getNextkey().TypeInfo == token.YOUZ {
		alldata.ShangOne = dtife
		alldata.Start++
		dtif = parseMemberExpression(alldata, true)
	} else if alldata.getNextkey().TypeInfo == token.DENYU {
		alldata.ShangOne = nil
		alldata.Start++
		dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.DENYU}
		dds := parseAss(alldata)
		dshhh.Right = dds
		dtif = dshhh
	} else if alldata.getNextkey().TypeInfo == token.JIADEN {
		alldata.ShangOne = nil
		alldata.Start++
		dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.JIADEN}
		dds := parseAss(alldata)
		dshhh.Right = dds
		return dshhh
	} else if alldata.getNextkey().TypeInfo == token.JANDEN {
		alldata.ShangOne = nil
		alldata.Start++
		dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.JANDEN}
		dds := parseAss(alldata)
		dshhh.Right = dds
		return dshhh
	} else {
		dtif = dtife
	}
	alldata.ShangOne = dtif
	//hu := parseBin(alldata)
	//if hu == nil {
	//	dtif = dtif
	//} else {
	//	dtif = hu
	//}

	return dtif
}

func parseNEW(alldata *allDtInfo) ast.Statement {
	caleer := parseData(alldata)
	dasd := caleer.(*ast.CallExpression)
	alldata.ShangOne = nil
	Calle := &ast.NewExpression{Callee: dasd.Caller, Arguments: dasd.Arguments}
	return Calle
}

func parseMemberExpression(alldata *allDtInfo, computed bool) ast.Statement {
	allh := ast.MemberExpression{Object: alldata.ShangOne}
	if computed == false {
		dtt := alldata.getDt()
		allh.Property = &ast.StringLiteral{Value: dtt.Value}
	} else {

		dtt := overDD(alldata)
		alldata.Start++
		allh.Property = dtt
	}
	alldata.ShangOne = &allh
	if alldata.getNextkey().TypeInfo == token.ZHUOK {
		alldata.Start++
		dtif := parseCallE(alldata)
		return dtif
	} else if alldata.getNextkey().TypeInfo == token.DENYU {
		alldata.ShangOne = nil
		alldata.Start++
		dshhh := &ast.AssignmentExpression{Left: &allh, Operator: token.DENYU}
		dds := parseAss(alldata)
		dshhh.Right = dds
		return dshhh
	} else if alldata.getNextkey().TypeInfo == token.JIADEN {
		alldata.ShangOne = nil
		alldata.Start++
		dshhh := &ast.AssignmentExpression{Left: &allh, Operator: token.JIADEN}
		dds := parseAss(alldata)
		dshhh.Right = dds
		return dshhh
	} else if alldata.getNextkey().TypeInfo == token.JANDEN {
		alldata.ShangOne = nil
		alldata.Start++
		dshhh := &ast.AssignmentExpression{Left: &allh, Operator: token.JANDEN}
		dds := parseAss(alldata)
		dshhh.Right = dds
		return dshhh
	}
	if alldata.getNextkey().TypeInfo != token.Dian && alldata.getNextkey().TypeInfo != token.YOUZ {
		alldata.ShangOne = &allh
		//hu := parseBin(alldata)
		//var dsadasd ast.Statement
		//if hu == nil {
		//	dsadasd = &allh
		//} else {
		//	dsadasd = hu
		//}
		dsadasd := &allh
		return dsadasd
	}
	if alldata.getNextkey().TypeInfo == token.Dian {
		alldata.Start++

		return parseMemberExpression(alldata, false)
	} else {
		alldata.Start++

		return parseMemberExpression(alldata, true)
	}

}
func parseBlack(alldata *allDtInfo) ast.Statement {
	dataBody := []*ast.Statement{}
	djjj := &ast.BlockStatement{Body: dataBody}

	for {
		ji := parseData(alldata)

		if ji.StatementNode() == token.ZUOKH {
			break
		}
		alldata.ShangOne = nil
		if ji.StatementNode() == token.OVER {
			continue
		}
		dashdk := ji
		djjj.Body = append(djjj.Body, &dashdk)

	}
	alldata.ShangOne = nil
	return djjj
}

func parseIf(alldata *allDtInfo) ast.Statement {
	alldata.Start++
	for {
		dtt := parseData(alldata)
		alldata.ShangOne = dtt

		if alldata.getNextkey().TypeInfo == token.YOUOK {
			//alldata.ShangOne = nil
			alldata.Start++
			break
		}
		if alldata.isover() {
			fmt.Println("缺少 ')'")
			return nil
		}
	}
	test := alldata.ShangOne
	alldata.ShangOne = nil
	dtt := ast.IfStatement{Test: test}
	alldata.Start++
	Cope := parseBlack(alldata)
	dtt.Consequent = Cope
	if alldata.getNextkey().TypeInfo == token.ELSE {
		alldata.Start++
		if alldata.getNextkey().TypeInfo == token.YOUKH {
			alldata.Start++
			dtt.Alternate = parseBlack(alldata)
		} else {
			dtt.Alternate = parseData(alldata)
		}

	}
	return &dtt
}

func parseForIn(alldata *allDtInfo, intt ast.Statement) ast.Statement {
	fori := &ast.ForInStatement{
		Left: intt,
	}
	for {
		dtt := parseData(alldata)
		alldata.ShangOne = dtt

		if alldata.getNextkey().TypeInfo == token.YOUOK {
			fori.Right = dtt
			alldata.Start++
			break
		}
	}
	alldata.Start++
	Cope := parseBlack(alldata)
	fori.Body = Cope

	return fori
}

func parseFor(alldata *allDtInfo) ast.Statement {
	alldata.Start++
	alldata.ShangOne = nil
	dat := &ast.ForStatement{}
	if alldata.getNextkey().TypeInfo == token.OVER {
	} else {
		for {
			dtt := parseData(alldata)
			alldata.ShangOne = dtt

			if alldata.getNextkey().TypeInfo == token.OVER {
				dat.Init = dtt
				alldata.Start++
				break
			} else if alldata.getNextkey().TypeInfo == token.IN {
				alldata.Start++
				return parseForIn(alldata, dtt)
			}
		}
	}
	if alldata.getNextkey().TypeInfo == token.OVER {
		alldata.Start++
	} else {
		for {
			dtt := parseData(alldata)
			alldata.ShangOne = dtt

			if alldata.getNextkey().TypeInfo == token.OVER {
				dat.Test = dtt
				alldata.Start++
				break
			}
		}
	}
	if alldata.getNextkey().TypeInfo == token.YOUOK {
		alldata.Start++
	} else {
		for {
			dtt := parseData(alldata)
			alldata.ShangOne = dtt

			if alldata.getNextkey().TypeInfo == token.YOUOK {
				dat.Updata = dtt
				alldata.Start++
				break
			}
		}
	}
	alldata.Start++
	Cope := parseBlack(alldata)
	dat.Body = Cope
	alldata.ShangOne = nil
	return dat
}
func parseAss(alldata *allDtInfo) ast.Statement {
	for {
		dtt := parseData(alldata)
		alldata.ShangOne = dtt

		if alldata.getNextkey().TypeInfo == token.OVER {

			//alldata.Start++
			alldata.ShangOne = nil
			return dtt
		} else if alldata.getNextkey().TypeInfo == token.YOUOK {

			alldata.ShangOne = nil
			return dtt
		} else if alldata.getNextkey().TypeInfo == token.ZUOKH {

			alldata.ShangOne = nil
			return dtt
		}
	}

}

func parseFun(alldata *allDtInfo) ast.Statement {

	if alldata.getNextkey().TypeInfo == token.ZHUOK {
		dayy := &ast.FunctionExpression{Params: make([]*ast.Statement, 0)}
		alldata.Start++
		if alldata.getNextkey().TypeInfo != token.YOUOK {
			for {
				dtt := parseData(alldata)

				if alldata.getNextkey().TypeInfo == token.DH {
					alldata.Start++
					kp := dtt
					dayy.Params = append(dayy.Params, &kp)

					alldata.ShangOne = nil
					continue
				}
				if alldata.getNextkey().TypeInfo == token.YOUOK {
					alldata.Start++
					dayy.Params = append(dayy.Params, &dtt)

					alldata.ShangOne = nil
					break
				}
			}
		} else {
			alldata.Start++
		}
		alldata.Start++
		dayy.Body = parseBlack(alldata)
		return dayy
	} else {
		dayy := &ast.FunctionDeclaration{Params: make([]*ast.Statement, 0)}
		dsd := &ast.Identifier{Name: alldata.getDt().Value}
		dayy.Id = dsd
		alldata.Start++
		if alldata.getNextkey().TypeInfo != token.YOUOK {
			for {
				dtt := parseData(alldata)

				if alldata.getNextkey().TypeInfo == token.DH {
					alldata.Start++
					kp := dtt
					dayy.Params = append(dayy.Params, &kp)

					alldata.ShangOne = nil
					continue
				}
				if alldata.getNextkey().TypeInfo == token.YOUOK {
					alldata.Start++
					dayy.Params = append(dayy.Params, &dtt)

					alldata.ShangOne = nil
					break
				}
			}
		} else {
			alldata.Start++
		}
		alldata.Start++
		dayy.Body = parseBlack(alldata)
		return dayy
	}

}

func parseVAR(alldata *allDtInfo) ast.Statement {
	das := ast.VariableDeclaration{Declarations: make([]*ast.Statement, 0)}
	for {
		gu := alldata.getDt()
		dttw := ast.Identifier{Name: gu.Value}
		var dtt ast.Statement = &dttw
		das.Declarations = append(das.Declarations, &dtt)
		if alldata.getNextkey().TypeInfo == token.DH {
			alldata.Start++
			alldata.ShangOne = nil
			continue
		}
		if alldata.getNextkey().TypeInfo == token.DENYU {

			alldata.Start++
			alldata.ShangOne = nil
			break
		} else if alldata.getNextkey().TypeInfo == token.OVER {

			alldata.ShangOne = nil
			return &das
		} else if alldata.getNextkey().TypeInfo == token.IN {

			alldata.ShangOne = nil
			return &das
		}
	}
	das.Init = overYH(alldata)
	return &das
}

func overYH(alldata *allDtInfo) ast.Statement {
	if alldata.getNextkey().TypeInfo == token.OVER {
		alldata.Start++
		alldata.ShangOne = nil
		return &ast.OVER{TypeInfo: token.OVER}
	}
	for {
		dtt := parseData(alldata)
		alldata.ShangOne = dtt

		if alldata.getNextkey().TypeInfo == token.OVER {

			alldata.ShangOne = nil
			return dtt
		} else if alldata.getNextkey().TypeInfo == token.YOUOK {

			alldata.ShangOne = nil
			return dtt
		} else if alldata.getNextkey().TypeInfo == token.ZUOKH {
			alldata.ShangOne = nil
			return dtt
		}
	}
}

func overDD(alldata *allDtInfo) ast.Statement {
	if alldata.getNextkey().TypeInfo == token.OVER {
		alldata.Start++
		alldata.ShangOne = nil
		return &ast.OVER{TypeInfo: token.OVER}
	}
	for {
		dtt := parseData(alldata)
		alldata.ShangOne = dtt

		if alldata.getNextkey().TypeInfo == token.ZUOZ {

			alldata.ShangOne = nil
			return dtt
		}
	}
}

func parseTRY(alldata *allDtInfo) ast.Statement {
	das := ast.TryStatement{}
	alldata.Start++
	das.Block = parseBlack(alldata)
	alldata.Start++
	dasghj := ast.CatchClause{}

	if alldata.getNextkey().TypeInfo == token.ZHUOK {
		alldata.Start++
		dasd := alldata.getDt()
		bh := ast.Identifier{Name: dasd.Value}
		dasghj.Param = &bh
		alldata.Start++
		alldata.Start++

	} else {
		alldata.Start++
	}
	dasghj.Body = parseBlack(alldata)
	das.Handler = &dasghj
	alldata.ShangOne = nil
	return &das
}

func parseUnary(alldata *allDtInfo, sdqm token.TokenType) ast.Statement {
	das := ast.UnaryExpression{Operator: sdqm.Value}
	das.Argument = parseData(alldata)
	return &das
}
func parseArray(alldata *allDtInfo) ast.Statement {
	das := ast.ArrayExpression{Elements: make([]*ast.Statement, 0)}
	if alldata.getNextkey().TypeInfo != token.ZUOZ {

		for {
			dtt := parseData(alldata)

			if alldata.getNextkey().TypeInfo == token.DH {
				alldata.Start++
				kp := dtt
				das.Elements = append(das.Elements, &kp)

				alldata.ShangOne = nil
				continue
			}
			if alldata.getNextkey().TypeInfo == token.ZUOZ {

				alldata.Start++
				kp := dtt
				das.Elements = append(das.Elements, &kp)
				break
			}

		}
	} else {
		alldata.Start++
	}
	if alldata.getNextkey().TypeInfo != token.Dian && alldata.getNextkey().TypeInfo != token.YOUZ {
		//hu := parseBin(alldata)
		//var dsadasd ast.Statement
		//if hu == nil {
		//	dsadasd = &allh
		//} else {
		//	dsadasd = hu
		//}
		dsadasd := &das
		return dsadasd
	}
	if alldata.getNextkey().TypeInfo == token.Dian {
		alldata.Start++
		alldata.ShangOne = &das

		return parseMemberExpression(alldata, false)
	} else {
		alldata.Start++
		alldata.ShangOne = &das
		return parseMemberExpression(alldata, true)
	}
}

func parseObject(alldata *allDtInfo) ast.Statement {
	das := ast.ObjectExpression{Properties: make([]*ast.Statement, 0)}
	var dsfhc ast.Statement
	if alldata.getNextkey().TypeInfo != token.ZUOKH {

		for {
			jklads := ast.Property{}
			dtt := parseData(alldata)
			if dtt.StatementNode() == token.OVER {
				continue
			}
			if dtt.StatementNode() == token.ZUOKH {
				if dsfhc != nil {
					das.Properties = append(das.Properties, &dsfhc)
				}
				break
			}

			alldata.Start++
			jklads.Key = dtt
			value := parseData(alldata)
			jklads.Value = value
			var dsf ast.Statement = &jklads
			dsfhc = dsf
			if alldata.getNextkey().TypeInfo == token.DH {
				alldata.Start++
				das.Properties = append(das.Properties, &dsf)
				dsfhc = nil
				alldata.ShangOne = nil
				continue
			}
			if alldata.getNextkey().TypeInfo == token.ZUOKH {

				alldata.Start++
				das.Properties = append(das.Properties, &dsf)
				break
			}

		}
	} else {
		alldata.Start++
	}
	return &das
}

func parseRet(alldata *allDtInfo) ast.Statement {
	das := ast.ReturnStatement{}
	das.Argument = parseData(alldata)
	return &das
}
func parseUpdate(alldata *allDtInfo, opert *token.TokenType, ty bool) ast.Statement {
	jhi := alldata.ShangOne
	das := ast.UnaryExpression{Argument: jhi, Operator: opert.Value}
	das.Prefix = ty
	alldata.ShangOne = &das
	//hu := parseBin(alldata)
	//var dsadasd ast.Statement
	//if hu == nil {
	//	dsadasd = &das
	//} else {
	//	dsadasd = hu
	//}
	dsadasd := &das
	return dsadasd

}

func parseBin(alldata *allDtInfo, mytk *token.TokenType) ast.Statement {
	//mytk := alldata.getNextkey()
	var dtif ast.Statement
	//alldata.Start++

	switch mytk.TypeInfo {
	case token.ADD:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.CHEN:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.SDD:
		if alldata.ShangOne == nil {
			dsdh := &ast.NumericLiteral{Value: 0}
			var dsdoooo ast.Statement = dsdh
			alldata.ShangOne = dsdoooo
		}
		dtif = parseBinaryExpression(alldata, mytk)
	case token.DXIAND:
		alldata.Start++
		dtif = parseBinaryExpression(alldata, mytk)
	case token.CHU:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.XIAND:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.YIHUO:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.HUO:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.HUOHUO:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.YU:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.YUYU:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.XIAOYH:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.XIAOYHYH:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.XIAOYHYHYH:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.QUYU:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.XIAOYHDY:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.DAYH:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.DAYHDY:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.DAYHYH:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.DAYHYHYU:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.BUDY:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.KAIFAN:
		dtif = parseBinaryExpression(alldata, mytk)
	case token.BUDYDY:
		dtif = parseBinaryExpression(alldata, mytk)

	case token.UPADD:
		dtif = parseUpdate(alldata, mytk, false)
	case token.UPASD:
		dtif = parseUpdate(alldata, mytk, false)

	default:
		fmt.Println("js 语法错误", mytk.Value)
		os.Exit(0)
		alldata.Start--

	}

	return dtif
}

func parseData(alldata *allDtInfo) ast.Statement {
	mytk := alldata.getDt()
	var dtif ast.Statement
	switch mytk.TypeInfo {
	case token.IF:
		dtif = parseIf(alldata)
	case token.YOUKH:

		dtif = parseObject(alldata)
	case token.TYPEOF:
		dtif = parseUnary(alldata, *mytk)
	//case token.SDD:
	//	dtif = parseUnary(alldata, *mytk)
	case token.UPADD:
		if alldata.ShangOne == nil {
			alldata.ShangOne = parseData(alldata)
			dtif = parseUpdate(alldata, mytk, true)
		} else {
			dsada := alldata.ShangOne
			dtif = &ast.UnaryExpression{Argument: dsada, Operator: mytk.TypeInfo, Prefix: false}
		}

	case token.UPASD:
		if alldata.ShangOne == nil {
			alldata.ShangOne = parseData(alldata)
			dtif = parseUpdate(alldata, mytk, true)
		} else {
			dsada := alldata.ShangOne
			dtif = &ast.UnaryExpression{Argument: dsada, Operator: mytk.TypeInfo, Prefix: false}
		}

	case token.QUFAN:
		dtif = parseUnary(alldata, *mytk)
	case token.FOR:
		dtif = parseFor(alldata)
	case token.TRY:
		dtif = parseTRY(alldata)
	case token.NEW:
		dtif = parseNEW(alldata)

	case token.RETURN:
		dtif = parseRet(alldata)
	case token.VAR:
		dtif = parseVAR(alldata)
	case token.Debug:
		dtif = &ast.DebugStatement{}
	case token.INT:
		zf, _ := strconv.ParseFloat(mytk.Value, 64)
		dtif = &ast.NumericLiteral{Value: zf}
		alldata.ShangOne = dtif
		//hu := parseBin(alldata)
		//if hu == nil {
		//	dtif = dtif
		//} else {
		//	dtif = hu
		//}
	case token.NULL:
		dtif = &ast.NullIdentifier{}
		alldata.ShangOne = dtif
		//hu := parseBin(alldata)
		//if hu == nil {
		//	dtif = dtif
		//} else {
		//	dtif = hu
		//}

	case token.Str:
		dtifed := &ast.StringLiteral{Value: mytk.Value}
		var dtife ast.Statement = dtifed
		if alldata.getNextkey().TypeInfo == token.ZHUOK {
			alldata.ShangOne = dtife
			alldata.Start++
			dtif = parseCallE(alldata)

		} else if alldata.getNextkey().TypeInfo == token.Dian {
			alldata.ShangOne = dtife
			alldata.Start++
			dtif = parseMemberExpression(alldata, false)
		} else if alldata.getNextkey().TypeInfo == token.YOUZ {
			alldata.ShangOne = dtife
			alldata.Start++
			dtif = parseMemberExpression(alldata, true)
		} else if alldata.getNextkey().TypeInfo == token.DENYU {
			alldata.ShangOne = nil
			alldata.Start++
			dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.DENYU}
			dds := parseAss(alldata)
			dshhh.Right = dds
			dtif = dshhh
		} else if alldata.getNextkey().TypeInfo == token.JIADEN {
			alldata.ShangOne = nil
			alldata.Start++
			dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.JIADEN}
			dds := parseAss(alldata)
			dshhh.Right = dds
			return dshhh
		} else if alldata.getNextkey().TypeInfo == token.JANDEN {
			alldata.ShangOne = nil
			alldata.Start++
			dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.JANDEN}
			dds := parseAss(alldata)
			dshhh.Right = dds
			return dshhh
		} else {
			alldata.ShangOne = dtife
			//hu := parseBin(alldata)
			//if hu == nil {
			//	dtif = dtife
			//} else {
			//	dtif = hu
			//}
			dtif = dtife

		}
	case token.THIS:
		fallthrough
	case token.IDENT:
		var dtife ast.Statement
		if mytk.TypeInfo == token.THIS {
			dtife = &ast.ThisExpression{}
		} else {
			dtife = &ast.Identifier{Name: mytk.Value}
		}

		if alldata.getNextkey().TypeInfo == token.ZHUOK {
			alldata.ShangOne = dtife
			alldata.Start++
			dtif = parseCallE(alldata)

		} else if alldata.getNextkey().TypeInfo == token.Dian {
			alldata.ShangOne = dtife
			alldata.Start++
			dtif = parseMemberExpression(alldata, false)
		} else if alldata.getNextkey().TypeInfo == token.YOUZ {
			alldata.ShangOne = dtife
			alldata.Start++
			dtif = parseMemberExpression(alldata, true)
		} else if alldata.getNextkey().TypeInfo == token.DENYU {
			alldata.ShangOne = nil
			alldata.Start++
			dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.DENYU}
			dds := parseAss(alldata)
			dshhh.Right = dds
			dtif = dshhh
		} else if alldata.getNextkey().TypeInfo == token.JIADEN {
			alldata.ShangOne = nil
			alldata.Start++
			dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.JIADEN}
			dds := parseAss(alldata)
			dshhh.Right = dds
			return dshhh
		} else if alldata.getNextkey().TypeInfo == token.JANDEN {
			alldata.ShangOne = nil
			alldata.Start++
			dshhh := &ast.AssignmentExpression{Left: dtife, Operator: token.JANDEN}
			dds := parseAss(alldata)
			dshhh.Right = dds
			return dshhh
		} else {
			alldata.ShangOne = dtife
			//hu := parseBin(alldata)
			//if hu == nil {
			//	dtif = dtife
			//} else {
			//	dtif = hu
			//}
			dtif = dtife

		}
	case token.YOUZ:
		dtif = parseArray(alldata)

	case token.OVER:
		dtif = &ast.OVER{TypeInfo: token.OVER}
	case token.ZHUOK:
		dtif = parseYpuxianKuoHao(alldata)
	case token.CONTINUE:
		dtif = &ast.ContinueStatement{}
	case token.BREAK:
		dtif = &ast.BreakStatement{}
	case token.FUN:
		dtif = parseFun(alldata)

	case token.YOUOK:
		dtif = &ast.OVER{TypeInfo: token.YOUOK}
	case token.ZUOKH:
		dtif = &ast.OVER{TypeInfo: token.ZUOKH}
	case token.ZUOZ:
		dtif = &ast.OVER{TypeInfo: token.ZUOZ}
	case token.MAOHAO:
		dtif = &ast.OVER{TypeInfo: token.MAOHAO}
	case token.DH:
		dtif = &ast.OVER{TypeInfo: token.DH}
	default:
		dtif = parseBin(alldata, mytk)
	}
	//dtif.StatementNode()
	return dtif
}

func stardFun(alldata *allDtInfo) []*ast.Statement {
	dataBody := []*ast.Statement{}

	for {
		ji := overYH(alldata)
		if ji.StatementNode() == token.OVER {
			if alldata.isover() {
				break
			}
			continue
		}
		jkdas := ji
		dataBody = append(dataBody, &jkdas)
		alldata.ShangOne = nil

		if alldata.isover() {
			break
		}
	}
	return dataBody
}

func NewParse(fd []*token.TokenType) []*ast.Statement {
	dt := newdt(fd)
	dsd := stardFun(dt)
	return dsd
}
