package banding

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math"
	"math/rand"
	"myvmp/etree"
	"myvmp/fs"
	"myvmp/http"
	"myvmp/object"
	"myvmp/parsejson"
	"myvmp/re"
	"myvmp/token"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func ParseInt(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	kk := (*dt[0]).Type()

	switch kk {
	case token.TYSTR:
		vlue := (*dt[0]).(*object.StringObject)
		num, err := strconv.Atoi(vlue.Value)
		if err != nil {
			return &object.NumericObject{Value: 0}

		}
		return &object.NumericObject{Value: float64(num)}
	case token.TYNUM:
		vlue := (*dt[0]).(*object.NumericObject)
		return &object.NumericObject{Value: float64(int(vlue.Value))}

	}
	return &object.NumericObject{Value: float64(0)}
}

func ParseFloat(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	kk := (*dt[0]).Type()
	switch kk {
	case token.TYSTR:
		vlue := (*dt[0]).(*object.StringObject)
		num, err := strconv.ParseFloat(vlue.Value, 64)
		if err != nil {
			return &object.NumericObject{Value: 0}

		}
		return &object.NumericObject{Value: num}
	case token.TYNUM:
		vlue := (*dt[0]).(*object.NumericObject)
		return &object.NumericObject{Value: vlue.Value}

	}
	return &object.NumericObject{Value: float64(0)}
}

func CharToStr(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args

	dtw := int((*dt[0]).(*object.NumericObject).Value)
	//fmt.Println(dtw)
	s := strconv.Itoa(dtw)
	//fmt.Println(s, utf8.RuneCountInString(string(dtw)))
	return &object.StringObject{Value: s}
}

func AppendArray(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	(*dt[0]).(*object.Environment).Value = append((*dt[0]).(*object.Environment).Value, dt[1])
	return *dt[1]
}

//
//func GetToken(dt []*object.Object) object.Object {
//	// 获取 window 对象
//
//	// 获取 window 对象
//	window := js.Global()
//	cdzfc := (*dt[0]).(*object.StringObject).Value
//	// 读取 window.mytoken 的值
//	window.Call("cycallback", cdzfc)
//
//	dte := &object.StringObject{Value: cdzfc}
//	return dte
//}

func Delete(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	ko := (*dt[0]).(*object.Environment)

	typej := ko.Type()
	var out object.Object
	switch typej {
	case token.Object:
		key := (*dt[1]).(*object.StringObject)
		out, _ = ko.Store.Get(key.Value)
		ko.Store.Delete(key.Value)
		//delete(ko.Store, key.Value)
	case token.ArrayE:
		key := int((*dt[1]).(*object.NumericObject).Value)
		out = *ko.Value[key]
		jhhi := ko.Value[0:key]
		jhhi2 := ko.Value[key+1 : len(ko.Value)]
		ko.Value = append(jhhi, jhhi2...)
	}

	return out
}

func CyPrint(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	for idx, v := range dt {
		fmt.Printf((*v).ToString())
		if idx != len(dt)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
	return &object.NULLObject{}
}

func Input(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	for idx, v := range dt {
		fmt.Printf((*v).ToString())
		if idx != len(dt)-1 {
			fmt.Printf(" ")
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	// 读取输入
	scanner.Scan()
	input := scanner.Text()

	// 检查是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return nil
	}

	// 输出用户输入的字符串
	return &object.StringObject{Value: input}
}

func GetChar(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	var din int
	if len(dt) == 1 {
		din = 0
	} else {
		kp := int((*dt[1]).(*object.NumericObject).Value)
		din = kp
	}
	strInfo := (*dt[0]).(*object.StringObject).Value
	runes := []rune(strInfo)
	return &object.NumericObject{
		Value: float64(runes[din]),
	}
}

func GetLength(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args
	kk := (*dt[0]).Type()
	switch kk {
	case token.TYSTR:
		vlue := (*dt[0]).(*object.StringObject)
		//fmt.Println(vlue.Value, len(vlue.Value), vlue.Value[0], vlue.Value[1])
		return &object.NumericObject{Value: float64(utf8.RuneCountInString(vlue.Value))}
	case token.ArrayE:
		vlue := (*dt[0]).(*object.Environment).Value
		return &object.NumericObject{Value: float64(len(vlue))}
	case token.Object:
		vlue := (*dt[0]).(*object.Environment).Store
		return &object.NumericObject{Value: float64(len(vlue.M))}

	}
	return &object.NumericObject{Value: float64(0)}
}

func Math_Ramdom(myfun *object.FunctionDeclarationObject) object.Object {
	//dt := myfun.Args
	randomValue := rand.Float64()

	return &object.NumericObject{Value: randomValue}
}
func Math_Pow(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args[0]
	dte := (*dt).(*object.NumericObject).Value
	dt2 := (*myfun.Args[1]).(*object.NumericObject).Value
	dd := math.Pow(dte, dt2)

	return &object.NumericObject{Value: dd}
}
func Math_Sqrt(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args[0]
	dte := (*dt).(*object.NumericObject).Value
	dd := math.Sqrt(dte)

	return &object.NumericObject{Value: dd}
}
func Objecte_setPrototypeOf(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args[0]
	dtold := myfun.Args[1]
	newOne := (*dt).(*object.Environment)
	oldOne := (*dtold).(*object.Environment)
	newOne.Outer = oldOne

	return &object.NULLObject{}
}
func JSON_stringify(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args[0]

	newOne := (*dt).(*object.Environment)

	return parsejson.JsonToStr(newOne)
}

func CONSOLE_log(myfun *object.FunctionDeclarationObject) object.Object {

	for idx, v := range myfun.Args {
		fmt.Printf((*v).ToString())
		if idx != len(myfun.Args)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
	return &object.NULLObject{}
}

func String_fromCharCode(myfun *object.FunctionDeclarationObject) object.Object {

	dtw := int((*myfun.Args[0]).(*object.NumericObject).Value)
	s := strconv.Itoa(dtw)
	return &object.StringObject{Value: s}
}

func String_strip(myfun *object.FunctionDeclarationObject) object.Object {

	dtw := (*myfun.Args[0]).(*object.StringObject).Value
	s := strings.TrimSpace(dtw)
	return &object.StringObject{Value: s}
}

func String_replace(myfun *object.FunctionDeclarationObject) object.Object {

	dtw := (*myfun.Args[0]).(*object.StringObject).Value
	dtw2 := (*myfun.Args[1]).(*object.StringObject).Value
	dtw3 := (*myfun.Args[2]).(*object.StringObject).Value
	s := strings.Replace(dtw, dtw2, dtw3, -1)
	return &object.StringObject{Value: s}
}

func String_split(myfun *object.FunctionDeclarationObject) object.Object {

	dtw := (*myfun.Args[0]).(*object.StringObject).Value
	dtw2 := (*myfun.Args[1]).(*object.StringObject).Value
	allslist := strings.Split(dtw, dtw2)
	ddd := object.NewArray()
	for _, vd := range allslist {
		d := &object.StringObject{Value: vd}
		var j object.Object = d
		ddd.Value = append(ddd.Value, &j)
	}

	return &ddd
}
func String_newbyte(myfun *object.FunctionDeclarationObject) object.Object {

	dtw := (*myfun.Args[0]).(*object.Environment).Value

	dnnnnnn := make([]byte, len(dtw))
	for idx, dt := range dtw {
		f := (*dt).Type()
		if f == token.TYNUM {
			b := (*dt).(*object.NumericObject).Value
			dnnnnnn[idx] = byte(b)
		} else {
			dnnnnnn[idx] = 0
		}

	}

	return &object.ByteObject{Value: dnnnnnn}
}

func String_decode(myfun *object.FunctionDeclarationObject) object.Object {

	dtw := (*myfun.Args[0]).(*object.ByteObject).Value
	var bian string = "utf-8"
	if len(myfun.Args) >= 2 {
		bian = (*myfun.Args[1]).(*object.StringObject).Value
	}
	if bian == "gbk" {
		decoder := simplifiedchinese.GBK.NewDecoder()

		// 使用GBK解码器将字节数组转换为字符串
		reader := transform.NewReader(bytes.NewReader(dtw), decoder)
		decodedBytes, _ := ioutil.ReadAll(reader)
		return &object.StringObject{Value: string(decodedBytes)}
	} else if bian == "utf-8" {
		return &object.StringObject{Value: string(dtw)}
	} else {
		return &object.StringObject{Value: string(dtw)}
	}

}

func String_encode(myfun *object.FunctionDeclarationObject) object.Object {

	dtw := (*myfun.Args[0]).(*object.StringObject).Value
	var bian string = "utf-8"
	if len(myfun.Args) >= 2 {
		bian = (*myfun.Args[1]).(*object.StringObject).Value
	}
	if bian == "gbk" {
		encoder := simplifiedchinese.GBK.NewEncoder()

		// 使用GBK解码器将字节数组转换为字符串
		var buf bytes.Buffer
		writer := transform.NewWriter(&buf, encoder)
		writer.Write([]byte(dtw))
		writer.Close()
		byteArray := buf.Bytes()
		return &object.ByteObject{Value: byteArray}
	} else if bian == "utf-8" {
		byteArray := []byte(dtw)
		return &object.ByteObject{Value: byteArray}
	} else {
		byteArray := []byte(dtw)
		return &object.ByteObject{Value: byteArray}
	}

}

func Date_now(myfun *object.FunctionDeclarationObject) object.Object {
	currentTime := time.Now()
	timestampMillis := currentTime.UnixNano() / int64(time.Millisecond)

	return &object.NumericObject{Value: float64(timestampMillis)}
}

func Date_sleep(myfun *object.FunctionDeclarationObject) object.Object {
	dd := int((*myfun.Args[0]).(*object.NumericObject).Value)
	time.Sleep(time.Duration(dd) * time.Millisecond)
	return &object.NULLObject{}
}

func Cyhttp_toJSON(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Env.Store
	jsdt, ok := dt.Get(token.Jsontext)
	if ok {
		return jsdt
	}

	text, _ := dt.Get(token.Text)
	dstest := text.(*object.StringObject)
	JsonDt := parsejson.ParseStrToJson(dstest.Value)
	dt.Set(token.Jsontext, JsonDt)
	return JsonDt
}
func Cyhttp_ReHeaders(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Env.Store
	jsdt, ok := dt.Get(token.Headerstext)
	if ok {
		return jsdt
	}
	text, _ := dt.Get(token.Headerstext2)
	dstest := text.(*object.StringObject)
	JsonDt := parsejson.ParseStrToJson(dstest.Value)
	dt.Set(token.Headerstext, JsonDt)
	return JsonDt
}

func Cyhttp_get(myfun *object.FunctionDeclarationObject) object.Object {
	url := (*myfun.Args[0]).(*object.StringObject).Value
	var huuu3 *object.Object
	if len(myfun.Args) < 2 {
		gg := object.NewObject()
		var h object.Object = &gg
		huuu3 = &h
	} else {
		huuu3 = myfun.Args[1]
	}
	timeOut := 40
	var huuu *object.Environment

	config := (*huuu3).(*object.Environment)
	huuue, ok := config.Store.Get(token.Headers)
	timeout, ok2 := config.Store.Get(token.Timeout)
	ParamsObj, ok3 := config.Store.Get(token.Params)
	Allow_redirectsTrue, ok4 := config.Store.Get(token.Allow_redirects)
	ProxiesStr, ok5 := config.Store.Get(token.Proxies)
	proxyStr := ""
	if ok5 {
		proxyStr = ProxiesStr.(*object.StringObject).Value
	}
	isAllow := true

	if ok4 {
		isAllow = Allow_redirectsTrue.(*object.BoolObject).Value
	}
	var Params map[string]string
	if ok3 {
		Params = map[string]string{}
		gyyy := ParamsObj.(*object.Environment)
		Keys := []string{}
		for key, _ := range gyyy.Store.M {
			Keys = append(Keys, key)
		}

		for _, key2 := range Keys {
			deee, _ := gyyy.Store.Get(key2)
			Params[key2] = deee.ToString()

		}
	} else {
		Params = nil
	}

	if ok2 {
		dhhh := timeout.(*object.NumericObject).Value
		timeOut = int(dhhh)
	}

	if ok == false {
		dss := object.NewObject()
		huuu = &dss
	} else {
		huuu = (huuue).(*object.Environment)
	}
	headersOb := huuu.Store.M
	Keys := []string{}
	for key, _ := range headersOb {
		Keys = append(Keys, key)
	}
	headers := map[string]string{}
	for _, key2 := range Keys {
		deee, _ := huuu.Store.Get(key2)
		headers[key2] = deee.ToString()

	}
	req := http.GetHttp(url, headers, int64(timeOut), Params, isAllow, proxyStr)
	reqdt := object.NewObject()
	reqdt.Store.Set(token.Status, &object.StringObject{Value: req.Status})
	reqdt.Store.Set(token.Iserror, &object.BoolObject{Value: req.IsError})
	reqdt.Store.Set(token.Content, &object.ByteObject{Value: req.Content})
	reqdt.Store.Set(token.Text, &object.StringObject{Value: req.Text})
	reqdt.Store.Set(token.Headerstext2, &object.StringObject{Value: req.ReHeaders})
	cyhttp_toJSON := Cyhttp_toJSON
	icyhttp_toJSON := new_Func(&cyhttp_toJSON)
	i2cyhttp_toJSON := icyhttp_toJSON.(*object.FunctionDeclarationObject)
	i2cyhttp_toJSON.Env = &reqdt
	reqdt.Store.Set(token.Json, i2cyhttp_toJSON)

	cyhttp_ReHeaders := Cyhttp_ReHeaders
	icyhttp_ReHeaders := new_Func(&cyhttp_ReHeaders)
	i2cyhttp_ReHeaders := icyhttp_ReHeaders.(*object.FunctionDeclarationObject)
	i2cyhttp_ReHeaders.Env = &reqdt
	reqdt.Store.Set(token.Cyhttp_ReHeaders, i2cyhttp_ReHeaders)
	return &reqdt
}

func Cyhttp_post(myfun *object.FunctionDeclarationObject) object.Object {
	url := (*myfun.Args[0]).(*object.StringObject).Value
	var huuu3 *object.Object
	if len(myfun.Args) < 2 {
		gg := object.NewObject()
		var h object.Object = &gg
		huuu3 = &h
	} else {
		huuu3 = myfun.Args[1]
	}
	timeOut := 40
	var huuu *object.Environment
	headers := map[string]string{}

	config := (*huuu3).(*object.Environment)
	huuue, ok := config.Store.Get(token.Headers)
	timeout, ok2 := config.Store.Get(token.Timeout)
	ParamsObj, ok3 := config.Store.Get(token.Params)
	Allow_redirectsTrue, ok4 := config.Store.Get(token.Allow_redirects)
	ProxiesStr, ok5 := config.Store.Get(token.Proxies)
	JsonObj, ok6 := config.Store.Get(token.Json)

	podata := ""
	if ok6 {
		podatjson := JsonObj.(*object.Environment)
		headers["Accept"] = "application/json"
		podata = parsejson.JsonToStr(podatjson).(*object.StringObject).Value
	} else {
		DataTest, ok7 := config.Store.Get(token.Data)
		if ok7 {
			podata = DataTest.ToString()
		}
	}

	proxyStr := ""
	if ok5 {
		proxyStr = ProxiesStr.(*object.StringObject).Value
	}
	isAllow := true

	if ok4 {
		isAllow = Allow_redirectsTrue.(*object.BoolObject).Value
	}
	var Params map[string]string
	if ok3 {
		Params = map[string]string{}
		gyyy := ParamsObj.(*object.Environment)
		Keys := []string{}
		for key, _ := range gyyy.Store.M {
			Keys = append(Keys, key)
		}

		for _, key2 := range Keys {
			deee, _ := gyyy.Store.Get(key2)
			Params[key2] = deee.ToString()

		}
	} else {
		Params = nil
	}

	if ok2 {
		dhhh := timeout.(*object.NumericObject).Value
		timeOut = int(dhhh)
	}

	if ok == false {
		dss := object.NewObject()
		huuu = &dss
	} else {
		huuu = (huuue).(*object.Environment)
	}
	headersOb := huuu.Store.M
	Keys := []string{}
	for key, _ := range headersOb {
		Keys = append(Keys, key)
	}
	for _, key2 := range Keys {
		deee, _ := huuu.Store.Get(key2)
		headers[key2] = deee.ToString()

	}
	req := http.PostHttp(url, headers, int64(timeOut), Params, isAllow, proxyStr, podata)
	reqdt := object.NewObject()
	reqdt.Store.Set(token.Status, &object.StringObject{Value: req.Status})
	reqdt.Store.Set(token.Iserror, &object.BoolObject{Value: req.IsError})
	reqdt.Store.Set(token.Text, &object.StringObject{Value: req.Text})
	reqdt.Store.Set(token.Content, &object.ByteObject{Value: req.Content})
	reqdt.Store.Set(token.Headerstext2, &object.StringObject{Value: req.ReHeaders})
	cyhttp_toJSON := Cyhttp_toJSON
	icyhttp_toJSON := new_Func(&cyhttp_toJSON)
	i2cyhttp_toJSON := icyhttp_toJSON.(*object.FunctionDeclarationObject)
	i2cyhttp_toJSON.Env = &reqdt
	reqdt.Store.Set(token.Json, i2cyhttp_toJSON)

	cyhttp_ReHeaders := Cyhttp_ReHeaders
	icyhttp_ReHeaders := new_Func(&cyhttp_ReHeaders)
	i2cyhttp_ReHeaders := icyhttp_ReHeaders.(*object.FunctionDeclarationObject)
	i2cyhttp_ReHeaders.Env = &reqdt
	reqdt.Store.Set(token.Cyhttp_ReHeaders, i2cyhttp_ReHeaders)
	return &reqdt
}

func JSON_parse(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args[0]
	jsonStr := (*dt).(*object.StringObject)
	ddd := parsejson.ParseStrToJson(jsonStr.Value)

	return ddd
}
func Objecte_keys(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args[0]
	dtInfo := (*dt).(*object.Environment)
	allk := object.NewArray()

	for key, _ := range dtInfo.Store.M {
		sko := &object.StringObject{Value: key}
		var kp object.Object = sko
		allk.Value = append(allk.Value, &kp)
	}
	return &allk
}
func NewEnv(eg *object.Environment) *object.Environment {
	s := object.NewSafeMap()
	env := &object.Environment{Store: s, Outer: eg, TypeInfo: token.THIS}
	return env
}
func new_Func(ddd *func(*object.FunctionDeclarationObject) object.Object) object.Object {
	d := &object.FunctionDeclarationObject{IsNative: 1, NativeBody: ddd}
	return d
}

func Init() (map[string]object.Object, map[string]*object.Environment) {
	rand.Seed(time.Now().UnixNano())

	getlength := GetLength
	input := Input
	cyPrint := CyPrint
	getChar := GetChar
	charToStr := CharToStr
	appendArray := AppendArray
	parseInt := ParseInt
	parseFloat := ParseFloat
	deletee := Delete
	//getToken := GetToken
	gy := make(map[string]object.Object)
	gy[token.GetLength] = new_Func(&getlength)
	gy[token.Print] = new_Func(&cyPrint)
	gy[token.GetChar] = new_Func(&getChar)
	gy[token.CharToStr] = new_Func(&charToStr)
	gy[token.AppendArray] = new_Func(&appendArray)
	gy[token.Delete] = new_Func(&deletee)
	gy[token.ParseInt] = new_Func(&parseInt)
	gy[token.ParseFloat] = new_Func(&parseFloat)
	gy[token.Input] = new_Func(&input)
	//gy[token.GetToken] = &getToken

	dxhs := make(map[string]*object.Environment)

	// Math ---------------------
	Math := NewEnv(nil)
	dxhs[token.Math] = Math

	// Math_Ramdom ---------------
	math_Ramdom := Math_Ramdom
	Math.Store.Set(token.Math_random, new_Func(&math_Ramdom))

	// Math_Pow ---------------
	math_Pow := Math_Pow
	Math.Store.Set(token.Math_Pow, new_Func(&math_Pow))
	// Math_Sqrt ---------------
	math_Sqrt := Math_Sqrt
	Math.Store.Set(token.Math_Sqrt, new_Func(&math_Sqrt))

	// Object --------------------
	Object := NewEnv(nil)
	dxhs[token.Objecte] = Object
	// Object_setPrototypeOf --------------------
	objecte_setPrototypeOf := Objecte_setPrototypeOf
	Object.Store.Set(token.Objecte_setPrototypeOf, new_Func(&objecte_setPrototypeOf))

	// Object_keys --------------------
	ojecte_keys := Objecte_keys
	Object.Store.Set(token.Objecte_keys, new_Func(&ojecte_keys))

	// JSON --------------------
	JSON := NewEnv(nil)
	dxhs[token.JSON] = JSON

	// JSON_stringify --------------------
	jSON_stringify := JSON_stringify
	JSON.Store.Set(token.JSON_stringify, new_Func(&jSON_stringify))

	// JSON_parse --------------------
	jSON_parse := JSON_parse
	JSON.Store.Set(token.JSON_parse, new_Func(&jSON_parse))

	// CONSOLE --------------------
	CONSOLE := NewEnv(nil)
	dxhs[token.CONSOLE] = CONSOLE

	// CONSOLE_log --------------------
	cONSOLE_log := CONSOLE_log
	CONSOLE.Store.Set(token.CONSOLE_log, new_Func(&cONSOLE_log))

	// String --------------------
	String := NewEnv(nil)
	dxhs[token.String] = String

	// String_fromCharCode --------------------
	string_fromCharCode := String_fromCharCode
	String.Store.Set(token.String_fromCharCode, new_Func(&string_fromCharCode))

	// String_strip --------------------
	string_strip := String_strip
	String.Store.Set(token.String_strip, new_Func(&string_strip))
	// String_replace --------------------
	string_replace := String_replace
	String.Store.Set(token.String_replace, new_Func(&string_replace))
	// String_split --------------------
	string_split := String_split
	String.Store.Set(token.String_split, new_Func(&string_split))
	// String_decode --------------------
	string_decode := String_decode
	String.Store.Set(token.String_decode, new_Func(&string_decode))
	// String_encode --------------------
	string_encode := String_encode
	String.Store.Set(token.String_encode, new_Func(&string_encode))
	// String_newbyte --------------------
	string_newbyte := String_newbyte
	String.Store.Set(token.String_newbyte, new_Func(&string_newbyte))

	// Date --------------------
	Date := NewEnv(nil)
	dxhs[token.Date] = Date

	// Date_now --------------------
	date_now := Date_now
	Date.Store.Set(token.Date_now, new_Func(&date_now))
	// Date_sleep --------------------
	date_sleep := Date_sleep
	Date.Store.Set(token.Date_sleep, new_Func(&date_sleep))

	// Fs --------------------
	Fs := NewEnv(nil)
	dxhs[token.Fs] = Fs

	// Fs_open --------------------
	fs_open := fs.Fs_open
	Fs.Store.Set(token.Fs_open, new_Func(&fs_open))
	// Fs_cmd --------------------
	Fs_cmd := fs.Fs_cmd
	Fs.Store.Set(token.Fs_cmd, new_Func(&Fs_cmd))

	// Cyhttp --------------------
	Cyhttp := NewEnv(nil)
	dxhs[token.Cyhttp] = Cyhttp

	// Cyhttp_get --------------------
	cyhttp_get := Cyhttp_get
	Cyhttp.Store.Set(token.Cyhttp_get, new_Func(&cyhttp_get))
	// Cyhttp_get --------------------
	cyhttp_post := Cyhttp_post
	Cyhttp.Store.Set(token.Cyhttp_post, new_Func(&cyhttp_post))

	// Etree --------------------
	Etree := NewEnv(nil)
	dxhs[token.Etree] = Etree
	// Cyhttp_get --------------------
	etree_HTML := etree.Etree_HTML
	Etree.Store.Set(token.Etree_HTML, new_Func(&etree_HTML))

	// Re --------------------
	Re := NewEnv(nil)
	dxhs[token.Re] = Re
	// Re_findall --------------------
	re_findall := re.Re_findall
	Re.Store.Set(token.Re_findall, new_Func(&re_findall))
	// Re_findall --------------------
	re_sub := re.Re_sub
	Re.Store.Set(token.Re_sub, new_Func(&re_sub))

	return gy, dxhs
}
