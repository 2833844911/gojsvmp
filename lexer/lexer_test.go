package lexer

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	code := `   l = "dd";

var dd = (1+"aaal"+32)*2/321+ "dsadas";
ff = 100+ dd+ dd;
lp = ff == "ewewq";
if (lp ){
	lp  = 100;
}else if((90>(80*434))|| (20>10)) {
lp  = 200;
}else{
if ((90>3)){
lp = 20;
};
};
var kpd = dd+"d"+lp + (90+(80>>434)) *3;
var ddde = {"a":100,b:dd};
for (var j=1; j<10000; j=(j+1)){
	
	var ds = [1,2,3];
function dasdII(a,b){
function dasdIIeee(a,b){
	l = a + b;
return l;
};
	
		l = a + b;

		
	};
	if ((j == 50)){
		ddd = ds[2] + 5.8990;

		continue;
	};
		this["ddde"]["a"]  = ddde["a"] + 5.8990;
	dasdII(ddde["a"], ds[1]);
	lp = lp+1;

};
function dasddddIIeee(a,b){
	l = l+a + b;
this["ppp"] = function(){
l= 100+"dsd"+'ss';
};
this["lp"] = 100+this.l;

};
lp;


ddde["a"];
var lppp = new dasddddIIeee(1,2);
lppp['ppp']();
lppp["l"];
1 ^ 90;
`
	dt := New(code)
	kk := (*dt).Input()
	fmt.Println(kk)
}
