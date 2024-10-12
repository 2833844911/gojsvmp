package banding_test

import (
	"fmt"
	"myvmp/evaluator"
	"testing"
)

func TestNative(t *testing.T) {
	code := `
data = [1]
for (var i in data){
cyout(i)
}
`
	ddd := evaluator.Eval(code)
	fmt.Println(ddd)

}
