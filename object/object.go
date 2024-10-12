package object

import (
	"encoding/hex"
	"myvmp/ast"
	"myvmp/token"
	"strconv"
	"strings"
	"sync"
)

// SafeMap 是一个线程安全的map结构
type SafeMap struct {
	mu sync.RWMutex
	M  map[string]Object
}

// NewSafeMap 创建一个新的SafeMap实例
func NewSafeMap() *SafeMap {
	return &SafeMap{
		M: make(map[string]Object),
	}
}

// Set 在SafeMap中设置键值对
func (sm *SafeMap) Set(key string, value Object) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.M[key] = value
}

// Get 从SafeMap中获取指定键的值
func (sm *SafeMap) Get(key string) (Object, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.M[key]
	return value, ok
}

// Delete 从SafeMap中删除指定键的值
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.M, key)
}

type Object interface {
	Type() string
	ToString() string
}

type NumericObject struct {
	Value float64
}

func (ob *NumericObject) Type() string {
	return token.TYNUM
}

func (ob *NumericObject) ToString() string {
	str := strconv.FormatFloat(ob.Value, 'f', -1, 64)
	return str
}

type StringObject struct {
	Value string
	Key   string
}

func (ob *StringObject) Type() string {

	return token.TYSTR
}

func (ob *StringObject) Slice(start int, end int) string {
	return ob.Value[start:end]
}

func (ob *StringObject) ToString() string {
	return ob.Value
}

type ByteObject struct {
	Value []byte
	Key   string
}

func (ob *ByteObject) Type() string {

	return token.BYTE
}

func (ob *ByteObject) ToString() string {
	hexString := hex.EncodeToString(ob.Value)
	return hexString
}

type NULLObject struct {
}

func (ob *NULLObject) Type() string {
	return token.BULLE
}
func (ob *NULLObject) ToString() string {
	return token.NNNN
}

type NanObject struct {
}

func (ob *NanObject) Type() string {
	return token.NANINFO
}
func (ob *NanObject) ToString() string {
	return token.NANINFO
}

type Environment struct {
	Value    []*Object
	Store    *SafeMap
	Outer    *Environment
	Key      string
	TypeInfo string
}

func (ob *Environment) Type() string {
	if ob.TypeInfo == "" {
		return token.TYNUM

	}
	return ob.TypeInfo
}
func (ob *Environment) ToString() string {
	if ob.TypeInfo == token.ArrayE || ob.TypeInfo == token.Object {
		return token.YOUZ + ob.TypeInfo + token.ZUOZ
	}
	return token.YOUZ + token.THIS + token.ZUOZ
}

func (ob *Environment) Slice(start int, end int) []*Object {
	return ob.Value[start:end]
}

type BoolObject struct {
	Value bool
}

func (ob *BoolObject) Type() string {
	return token.BOOL
}
func (ob *BoolObject) ToString() string {
	if ob.Value == true {
		return token.TRUE
	}
	return token.FALSE
}

type BreakObject struct {
}

func (ob *BreakObject) Type() string {
	return token.BREAK
}
func (ob *BreakObject) ToString() string {

	return token.BREAK
}

type ContinueObject struct {
}

func (ob *ContinueObject) Type() string {
	return token.CONTINUE
}
func (ob *ContinueObject) ToString() string {

	return token.CONTINUE
}

type FunctionDeclarationObject struct {
	Params     []*ast.Statement
	Args       []*Object
	Body       ast.Statement
	NativeBody *func(*FunctionDeclarationObject) Object
	Env        *Environment
	IsNative   int
	Callthis   int
	BindType   string
	BindOb     any
}

func (ob *FunctionDeclarationObject) Type() string {
	return token.FUNCTION
}
func (ob *FunctionDeclarationObject) ToString() string {

	return token.FUNCTION
}

type ReturnStatementObject struct {
	Value Object
}

func new_Func(ddd *func(*FunctionDeclarationObject) Object) Object {
	d := &FunctionDeclarationObject{IsNative: 1, NativeBody: ddd}
	return d
}
func (ob *ReturnStatementObject) Type() string {
	return token.RETURN
}
func (ob *ReturnStatementObject) ToString() string {

	return token.RETURN
}

func array_push(myfun *FunctionDeclarationObject) Object {
	Listg := myfun.Env
	args := myfun.Args
	for _, v := range args {
		Listg.Value = append(Listg.Value, v)
	}

	return &NumericObject{Value: float64(len(Listg.Value))}
}

func array_pop(myfun *FunctionDeclarationObject) Object {
	Listg := myfun.Env
	if len(Listg.Value) == 0 {
		return &NumericObject{}
	}
	out := Listg.Value[len(Listg.Value)-1]
	Listg.Value = Listg.Value[:len(Listg.Value)-1]

	return *out
}

func array_join(myfun *FunctionDeclarationObject) Object {
	Listg := myfun.Env
	ds := []string{}
	jst := (*myfun.Args[0]).(*StringObject).Value
	for _, v := range Listg.Value {
		ds = append(ds, (*v).ToString())
	}

	return &StringObject{
		Value: strings.Join(ds, jst),
	}
}

func NewArray() Environment {
	dtte := NewEnv(nil)
	dtte.TypeInfo = token.ArrayE
	dtte.Value = make([]*Object, 0)

	Push := array_push
	dtte.Store.Set(token.PUSH, new_Func(&Push))

	Pop := array_pop
	dtte.Store.Set(token.POP, new_Func(&Pop))

	Join := array_join
	dtte.Store.Set(token.JION, new_Func(&Join))
	return *dtte
}
func NewObject() Environment {
	dtte := NewEnv(nil)
	dtte.TypeInfo = token.Object
	return *dtte
}
func NewEnv(eg *Environment) *Environment {
	//s := make(map[string]object.Object)
	s := NewSafeMap()
	env := &Environment{Store: s, Outer: eg, TypeInfo: token.ENV}
	return env
}
