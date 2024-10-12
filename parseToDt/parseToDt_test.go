package parseToDt

import (
	"fmt"
	"myvmp/ast"
	"myvmp/evaluator"
	"myvmp/lexer"
	"myvmp/parse"
	"myvmp/promise"
	"testing"
)

func TestNextTokenee(t *testing.T) {
	code := `
l = "dd";

var dd = (1 + "aaal" + 32) * 2 / 321 + "dsadas";
ff = 100 + dd + dd;
lp = ff == "ewewq";

var kpd = dd + "d" + lp + (90 + (80 >> 434)) * 3;

function dasddddIIeee(a, b) {
    l = l + a + b;
    this["ppp"] = function() {
        l = 100 + "dsd" + 'ss';
    }
    ;
    this["lp"] = 100 + this.l;

}
;
var lppp = new dasddddIIeee(1,2);
lppp['ppp']();
try {
    dasd = '{\'a\':"100",\"b\":{\'a\':100}}'
    dd = JSON["parse"](dasd);
    cyout("====", dd['b']['a']);
    dd = JSON["stringify"](dd);
    cyout("====", dd);
    dd['l'] = this
    dd = JSON["stringify"](dde);
} catch(a) {
    cyout("==== 重新");
}

cyout("==== 结束 ====", dd);

 
`
	promise.CyJSInit()

	dt := lexer.New(code)
	kk := (*dt).Input()

	kpp := parse.New(kk)
	fff := kpp.ParseAst()
	dtt := &ast.Program{Body: fff}
	dtt.StatementNode()
	dd := PrushData(dtt)
	fmt.Println(dd)
	ss := LoadStr(dd)
	dyy := ss.(*ast.Program)
	evaluator.StartEval(dyy.Body)
	promise.Done()
}
