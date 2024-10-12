package promise

import (
	"myvmp/ast"
	"myvmp/object"
	"myvmp/token"
	"sync"
)

var AllInfo sync.WaitGroup
var paseData *func(dtt *ast.Statement, env *object.Environment) object.Object

func Promise_thenT(myfun *object.FunctionDeclarationObject) object.Object {
	//fmt.Println(myfun.Args[0], "thenjjjjjjjjjjjjjjjjj")
	dbb := myfun.Args
	fuenv := myfun.Env
	for idx, vkey := range myfun.Params {
		zzkey := (*vkey).(*ast.Identifier).Name
		//if (zzkey == "rerea"){
		//	fmt.Println()
		//}
		if idx >= len(dbb) {
			fuenv.Store.Set(zzkey, &object.NULLObject{})

			continue
		}
		fuenv.Store.Set(zzkey, *dbb[idx])

	}
	AllInfo.Add(1)
	go func() {
		defer AllInfo.Done()
		(*paseData)(&myfun.Body, myfun.Env)
	}()

	return &object.NULLObject{}
}

func Promise_then(myfun *object.FunctionDeclarationObject) object.Object {
	promise_thenT := Promise_thenT
	fundd := (*myfun.Args[0]).(*object.FunctionDeclarationObject)
	newEnv := NewEnv(myfun.Env, fundd.Params)
	myfun.Env.Store.Set(token.Cbb_a, &object.FunctionDeclarationObject{Params: fundd.Params, IsNative: 1, NativeBody: &promise_thenT, Env: newEnv, Body: fundd.Body})
	if len(myfun.Args) >= 2 {
		fundd2 := (*myfun.Args[1]).(*object.FunctionDeclarationObject)
		myfun.Env.Store.Set(token.Cbb_b, &object.FunctionDeclarationObject{Params: fundd2.Params, IsNative: 1, NativeBody: &promise_thenT, Env: newEnv, Body: fundd2.Body})
	}
	//dtold := myfun.Args[1]

	return newEnv
}

func NewEnv(eg *object.Environment, ddp []*ast.Statement) *object.Environment {
	s := object.NewSafeMap()
	promise_then := Promise_then
	env := &object.Environment{Store: s, Outer: eg, TypeInfo: token.ENV}

	s.Set(token.Promise_then, &object.FunctionDeclarationObject{Params: ddp, IsNative: 1, NativeBody: &promise_then, Env: env})
	s.Set(token.Cbb_a, &object.NULLObject{})
	s.Set(token.Cbb_b, &object.NULLObject{})
	return env
}

func Init(dofun *func(dtt *ast.Statement, env *object.Environment) object.Object, funcdt *object.Object, env *object.Environment) object.Object {
	fundd := (*funcdt).(*object.FunctionDeclarationObject)
	newEnv := NewEnv(env, fundd.Params)
	paseData = dofun
	fundd.Env = newEnv
	AllInfo.Add(1)

	go func() {
		defer AllInfo.Done()

		(*dofun)(&fundd.Body, newEnv)
	}()
	return newEnv
}
func CyJSInit() {
	AllInfo = sync.WaitGroup{}

}
func Done() {
	AllInfo.Wait()
}
