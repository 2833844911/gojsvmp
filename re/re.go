package re

import (
	"fmt"
	"myvmp/object"
	"regexp"
)

func new_Func(ddd *func(*object.FunctionDeclarationObject) object.Object, typeS string, Obj any) object.Object {
	d := &object.FunctionDeclarationObject{IsNative: 1, NativeBody: ddd, BindType: typeS, BindOb: Obj}
	return d
}

func Re_findall(myfun *object.FunctionDeclarationObject) object.Object {
	myfunStr := (*myfun.Args[0]).(*object.StringObject).Value
	myStr := (*myfun.Args[1]).(*object.StringObject).Value
	re, err := regexp.Compile(myfunStr)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	// 查找所有匹配项
	matches := re.FindAllStringSubmatch(myStr, -1)
	dtList := object.NewArray()
	for _, match := range matches {
		if len(match) > 1 { // match[0] 是整个匹配，match[1] 是第一个括号表达式的匹配
			kk := &object.StringObject{Value: match[1]}
			var nnn object.Object = kk
			dtList.Value = append(dtList.Value, &nnn)
		} else {
			kk := &object.StringObject{Value: match[0]}
			var nnn object.Object = kk
			dtList.Value = append(dtList.Value, &nnn)
		}
	}

	return &dtList
}

func Re_sub(myfun *object.FunctionDeclarationObject) object.Object {
	myfunStr := (*myfun.Args[0]).(*object.StringObject).Value
	myStr := (*myfun.Args[1]).(*object.StringObject).Value
	myStr2 := (*myfun.Args[2]).(*object.StringObject).Value
	re, err := regexp.Compile(myfunStr)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	// 查找所有匹配项
	matches := re.ReplaceAllString(myStr, myStr2)

	return &object.StringObject{Value: matches}
}
