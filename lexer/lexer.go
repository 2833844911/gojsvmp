package lexer

import (
	"myvmp/token"
	"regexp"
)

type LexerParse struct {
	keyList   []string
	value     string
	start     int
	maxlength int
}
type CardList struct {
	value []*token.TokenType
}

func New(dt string) *LexerParse {
	//dt = strings.Replace(dt, "\n", token.Kong, -1)
	//dt = strings.Replace(dt, "\n", token.OVER, -1)
	//dt = strings.Replace(dt, "\t", token.Kong, -1)
	//dt = strings.Replace(dt, "\r", token.Kong, -1)
	re := regexp.MustCompile(`[;]+\s*[;]+`)
	dt = re.ReplaceAllString(dt, ";")
	ree := regexp.MustCompile(`[{]+\s*[;]+`)
	dt = ree.ReplaceAllString(dt, "{")
	return &LexerParse{value: dt, start: -1, maxlength: len(dt)}
}
func (lp *LexerParse) clearNop() {
	for {
		lp.start++
		if lp.start >= lp.maxlength {
			return
		}

		chardt := lp.value[lp.start : lp.start+1]
		if chardt != token.NOP {
			lp.start--
			return
		}

	}
}

func (lp *LexerParse) clearZhus() {
	for {
		lp.start++
		if lp.start >= lp.maxlength {
			return
		}

		chardt := lp.value[lp.start : lp.start+1]
		if chardt == token.OVER || chardt == token.HUANH {
			lp.start--
			return
		}

	}
}

func (lp *LexerParse) readCard() string {
	carddt := token.Kong
	for {
		lp.start++
		if lp.start >= lp.maxlength {
			//fmt.Println("结束")
			return token.NNNN
		}
		chardt := lp.value[lp.start : lp.start+1]
		if chardt == token.NOP || chardt == token.TB || chardt == token.TR || chardt == token.YIHUO || chardt == token.HUANH || chardt == token.DUOZF || chardt == token.Dian || chardt == token.QUYU || chardt == token.Str || chardt == token.Str2 || chardt == token.QUFAN || chardt == token.MAOHAO || chardt == token.YU || chardt == token.HUO || chardt == token.DENYU || chardt == token.XIAOYH || chardt == token.DAYH || chardt == token.ZUOZ || chardt == token.YOUZ || chardt == token.ZUOKH || chardt == token.YOUKH || chardt == token.ZHUOK || chardt == token.YOUOK || chardt == token.OVER || chardt == token.ADD || chardt == token.SDD || chardt == token.DH || chardt == token.CHEN || chardt == token.CHU {
			if chardt == token.NOP {
				lp.clearNop()
				if carddt == token.NNNN {
					continue
				}

			}

			if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.CHU+token.CHU {
				lp.clearZhus()
				continue
			}

			if carddt != token.NNNN {
				if chardt != token.NOP {
					lp.start--
				}

			} else {

				if lp.start+3 < lp.maxlength && lp.value[lp.start:lp.start+3] == token.DXIAND {
					lp.start = lp.start + 2
					return lp.value[lp.start-2 : lp.start+1]
				}
				if lp.start+3 < lp.maxlength && lp.value[lp.start:lp.start+3] == token.DAYHYHYU {
					lp.start = lp.start + 2
					return lp.value[lp.start-2 : lp.start+1]
				}
				if lp.start+3 < lp.maxlength && lp.value[lp.start:lp.start+3] == token.XIAOYHYHYH {
					lp.start = lp.start + 2
					return lp.value[lp.start-2 : lp.start+1]
				}
				if lp.start+3 < lp.maxlength && lp.value[lp.start:lp.start+3] == token.BUDYDY {
					lp.start = lp.start + 2
					return lp.value[lp.start-2 : lp.start+1]
				}

				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.XIAND {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.KAIFAN {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.XIAOYHDY {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.DAYHDY {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.BUDY {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.DAYHYH {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.UPADD {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.UPASD {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.JIADEN {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.JANDEN {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.XIAOYHYH {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.HUOHUO {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				if lp.start+2 < lp.maxlength && lp.value[lp.start:lp.start+2] == token.YUYU {
					lp.start = lp.start + 1
					return lp.value[lp.start-1 : lp.start+1]
				}
				carddt = chardt
			}
			if chardt == token.Dian && isDigit(lp.value[lp.start+2]) {
				lp.start++
			} else {
				return carddt
			}

		}
		carddt += chardt
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func (lp *LexerParse) readStr() string {
	carddt := ""
	qian := ""
	for {
		lp.start++
		cff := lp.value[lp.start : lp.start+1]
		if qian == "\\" && cff == "n" {
			cff = "\n"
		} else if qian == "\\" && cff == "t" {
			cff = "\t"
		} else if qian == "\\" && cff == "\"" {
			qian = cff
			carddt += "\""
			continue
		} else if qian == "\\" && cff == "\\" {
			qian = "\\\\"
			cff = "\\"
			carddt += cff
			continue
		}
		qian = cff
		if cff == token.Str {
			break
		}
		if qian == "\\" {
			continue
		}
		carddt += cff

	}
	return carddt

}
func (lp *LexerParse) readStr3() string {
	carddt := ""
	qian := ""
	for {
		lp.start++
		cff := lp.value[lp.start : lp.start+1]
		if qian == "\\" && cff == "n" {
			cff = "\n"
		} else if qian == "\\" && cff == "t" {
			cff = "\t"
		} else if qian == "\\" && cff == "\"" {
			qian = cff
			carddt += "\""
			continue
		} else if qian == "\\" && cff == "\\" {
			qian = "\\\\"
			cff = "\\"
			carddt += cff
			continue
		}
		qian = cff
		if cff == token.DUOZF {
			break
		}
		if qian == "\\" {
			continue
		}
		carddt += cff

	}
	return carddt

}
func (lp *LexerParse) readStr2() string {
	carddt := token.Kong
	qian := ""
	for {
		lp.start++
		cff := lp.value[lp.start : lp.start+1]
		if qian == "\\" && cff == "n" {
			cff = "\n"
		} else if qian == "\\" && cff == "t" {
			cff = "\t"
		} else if qian == "\\" && cff == "'" {
			qian = cff
			carddt += "\""
			continue
		} else if qian == "\\" && cff == "\\" {
			qian = "\\\\"
			cff = "\\"
			carddt += cff
			continue
		}
		qian = cff
		if cff == token.Str2 {
			break
		}
		if qian == "\\" {
			continue
		}
		carddt += cff

	}
	return carddt

}
func (lp *LexerParse) nextCard() *token.TokenType {

	for {
		if lp.start >= lp.maxlength {
			//fmt.Println("结束")
			return &token.TokenType{
				TypeInfo: token.DONT,
			}
		}
		card := lp.readCard()
		switch card {
		case token.VAR:
			return &token.TokenType{
				TypeInfo: token.VAR,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.UPADD:
			return &token.TokenType{
				TypeInfo: token.UPADD,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.UPASD:
			return &token.TokenType{
				TypeInfo: token.UPASD,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.IN:
			return &token.TokenType{
				TypeInfo: token.IN,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.DAYHDY:
			return &token.TokenType{
				TypeInfo: token.DAYHDY,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.XIAOYHDY:
			return &token.TokenType{
				TypeInfo: token.XIAOYHDY,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.MAOHAO:
			return &token.TokenType{
				TypeInfo: token.MAOHAO,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.DAYHYH:
			return &token.TokenType{
				TypeInfo: token.DAYHYH,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.DAYHYHYU:
			return &token.TokenType{
				TypeInfo: token.DAYHYHYU,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.XIAOYHYH:
			return &token.TokenType{
				TypeInfo: token.XIAOYHYH,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.HUOHUO:
			return &token.TokenType{
				TypeInfo: token.HUOHUO,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.YUYU:
			return &token.TokenType{
				TypeInfo: token.YUYU,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.HUO:
			return &token.TokenType{
				TypeInfo: token.HUO,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.YU:
			return &token.TokenType{
				TypeInfo: token.YU,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.XIAOYHYHYH:
			return &token.TokenType{
				TypeInfo: token.XIAOYHYHYH,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.DENYU:

			return &token.TokenType{
				TypeInfo: token.DENYU,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.JIADEN:

			return &token.TokenType{
				TypeInfo: token.JIADEN,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.JANDEN:

			return &token.TokenType{
				TypeInfo: token.JANDEN,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.SDD:

			return &token.TokenType{
				TypeInfo: token.SDD,
				Value:    card,
				PAXU:     token.TWO,
			}
		case token.IF:

			return &token.TokenType{
				TypeInfo: token.IF,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.NULL:

			return &token.TokenType{
				TypeInfo: token.NULL,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.CHU:

			return &token.TokenType{
				TypeInfo: token.CHU,
				Value:    card,
				PAXU:     token.SCRRR,
			}
		case token.YIHUO:

			return &token.TokenType{
				TypeInfo: token.YIHUO,
				Value:    card,
				PAXU:     token.SCRRR,
			}
		case token.TYPEOF:

			return &token.TokenType{
				TypeInfo: token.TYPEOF,
				Value:    card,
				PAXU:     token.SCRRR,
			}
		case token.OVER:
			return &token.TokenType{
				TypeInfo: token.OVER,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.HUANH:
			return &token.TokenType{
				TypeInfo: token.OVER,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.Str:
			ddd := lp.readStr()
			return &token.TokenType{
				TypeInfo: token.Str,
				Value:    ddd,
				PAXU:     token.FOTT,
			}
		case token.TB:
			continue
		case token.TR:
			continue
		case token.Str2:
			ddd := lp.readStr2()
			return &token.TokenType{
				TypeInfo: token.Str,
				Value:    ddd,
				PAXU:     token.FOTT,
			}
		case token.DUOZF:
			ddd := lp.readStr3()
			return &token.TokenType{
				TypeInfo: token.Str,
				Value:    ddd,
				PAXU:     token.FOTT,
			}
		case token.ADD:
			return &token.TokenType{
				TypeInfo: token.ADD,
				Value:    card,
				PAXU:     token.TWO,
			}
		case token.TRY:
			return &token.TokenType{
				TypeInfo: token.TRY,
				Value:    card,
				PAXU:     token.TWO,
			}
		case token.CATCH:
			return &token.TokenType{
				TypeInfo: token.CATCH,
				Value:    card,
				PAXU:     token.TWO,
			}
		case token.FOR:
			return &token.TokenType{
				TypeInfo: token.FOR,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.XIAOYH:
			return &token.TokenType{
				TypeInfo: token.XIAOYH,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.Debug:
			return &token.TokenType{
				TypeInfo: token.Debug,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.DAYH:
			return &token.TokenType{
				TypeInfo: token.DAYH,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.BREAK:
			return &token.TokenType{
				TypeInfo: token.BREAK,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.CONTINUE:
			return &token.TokenType{
				TypeInfo: token.CONTINUE,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.THIS:
			return &token.TokenType{
				TypeInfo: token.THIS,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.DH:
			return &token.TokenType{
				TypeInfo: token.DH,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.FUN:
			return &token.TokenType{
				TypeInfo: token.FUN,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.ELSE:
			return &token.TokenType{
				TypeInfo: token.ELSE,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.YOUOK:
			return &token.TokenType{
				TypeInfo: token.YOUOK,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.YOUKH:
			return &token.TokenType{
				TypeInfo: token.YOUKH,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.ZUOKH:
			return &token.TokenType{
				TypeInfo: token.ZUOKH,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.ZHUOK:
			return &token.TokenType{
				TypeInfo: token.ZHUOK,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.KAIFAN:
			return &token.TokenType{
				TypeInfo: token.KAIFAN,
				Value:    card,
				PAXU:     token.SCRRR,
			}
		case token.XIAND:
			return &token.TokenType{
				TypeInfo: token.XIAND,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.BUDY:
			return &token.TokenType{
				TypeInfo: token.BUDY,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.BUDYDY:
			return &token.TokenType{
				TypeInfo: token.BUDYDY,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.QUFAN:
			return &token.TokenType{
				TypeInfo: token.QUFAN,
				Value:    card,
				PAXU:     token.SIX,
			}
		case token.DXIAND:
			return &token.TokenType{
				TypeInfo: token.DXIAND,
				Value:    card,
				PAXU:     token.ZNO,
			}
		case token.CHEN:
			return &token.TokenType{
				TypeInfo: token.CHEN,
				Value:    card,
				PAXU:     token.SCRRR,
			}
		case token.QUYU:
			return &token.TokenType{
				TypeInfo: token.QUYU,
				Value:    card,
				PAXU:     token.SCRRR,
			}
		case token.END:
			return &token.TokenType{
				TypeInfo: token.END,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.NEW:
			return &token.TokenType{
				TypeInfo: token.NEW,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.RETURN:
			return &token.TokenType{
				TypeInfo: token.RETURN,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.YOUZ:
			return &token.TokenType{
				TypeInfo: token.YOUZ,
				Value:    card,
				PAXU:     token.SCRRR,
			}
		case token.ZUOZ:
			return &token.TokenType{
				TypeInfo: token.ZUOZ,
				Value:    card,
				PAXU:     token.ONE,
			}
		case token.Dian:
			return &token.TokenType{
				TypeInfo: token.Dian,
				Value:    card,
				PAXU:     token.ONE,
			}

		default:
			if isDigit(card[0]) {
				return &token.TokenType{
					TypeInfo: token.INT,
					Value:    card,
					PAXU:     token.ONE,
				}
			} else {
				return &token.TokenType{
					TypeInfo: token.IDENT,
					Value:    card,
					PAXU:     token.ONE,
				}
			}

		}

	}
}

func (lp *LexerParse) Input() []*token.TokenType {
	dataList := []*token.TokenType{}
	for {
		card := lp.nextCard()
		if (*card).TypeInfo == token.END {
			return dataList
		}
		dataList = append(dataList, card)
	}
}
