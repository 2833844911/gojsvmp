package parse

import (
	"fmt"
	"myvmp/lexer"
	"testing"
)

func TestNextTokewn(t *testing.T) {
	code := `
for (var i in data){
cyout(1)
}
`

	dt := lexer.New(code)
	kk := (*dt).Input()
	dtt := NewParse(kk)
	fmt.Println(dtt)

}
