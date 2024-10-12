package fs

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"myvmp/object"
	"myvmp/token"
	"os"
	"os/exec"
)

type FileMy struct {
	File     *os.File
	Encoding string
	Ms       string
}

func Fs_read(myfun *object.FunctionDeclarationObject) object.Object {
	getbding := myfun.BindOb.(*FileMy)
	reader := bufio.NewReader(getbding.File)
	var content string
	if getbding.Encoding == "utf-8" {
		data, _ := ioutil.ReadAll(reader)
		content = string(data)
	} else if getbding.Encoding == "gbk" {
		gbkReader := transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder())
		data, _ := ioutil.ReadAll(gbkReader)
		content = string(data)
	} else {
		// 默认utf-8
		data, _ := ioutil.ReadAll(reader)
		content = string(data)
	}
	return &object.StringObject{Value: content}
}

func Fs_cmd(myfun *object.FunctionDeclarationObject) object.Object {
	ds := (*myfun.Args[0]).(*object.StringObject).Value
	cmd := exec.Command("sh", "-c", ds)
	// 获取命令输出，包括标准错误输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return &object.BoolObject{Value: false}
	}
	return &object.StringObject{Value: string(output)}
}

func Fs_readCont(myfun *object.FunctionDeclarationObject) object.Object {
	getbding := myfun.BindOb.(*FileMy)
	reader := bufio.NewReader(getbding.File)
	data, _ := ioutil.ReadAll(reader)
	return &object.ByteObject{Value: data}
}

func Fs_close(myfun *object.FunctionDeclarationObject) object.Object {
	ds := myfun.BindOb.(*FileMy).File.Close()
	if ds != nil {
		return &object.BoolObject{Value: false}
	}
	return &object.BoolObject{Value: true}
}

func Fs_write(myfun *object.FunctionDeclarationObject) object.Object {
	ds := myfun.BindOb.(*FileMy).File
	getbding := myfun.BindOb.(*FileMy)
	dtt := *myfun.Args[0]
	if dtt.Type() == token.BYTE {
		wz := dtt.(*object.ByteObject).Value
		ds.Write(wz)
	} else {
		wz := dtt.(*object.StringObject).Value
		var writer *os.File
		var err error
		if getbding.Encoding == "utf-8" {
			writer = ds
			_, err = io.WriteString(writer, wz)
		} else if getbding.Encoding == "gbk" {
			writer2 := transform.NewWriter(ds, simplifiedchinese.GBK.NewEncoder())
			_, err = io.WriteString(writer2, wz)
			writer2.Close()
		} else {
			writer = ds
			_, err = io.WriteString(writer, wz)
		}

		if err != nil {
			return &object.BoolObject{Value: false}
		}
	}

	return &object.BoolObject{Value: true}
}

func new_Func(ddd *func(*object.FunctionDeclarationObject) object.Object, typeS string, Obj any) object.Object {
	d := &object.FunctionDeclarationObject{IsNative: 1, NativeBody: ddd, BindType: typeS, BindOb: Obj}
	return d
}
func Fs_open(myfun *object.FunctionDeclarationObject) object.Object {
	dt := myfun.Args[0]

	fsDt := object.NewEnv(nil)
	fsDt.Store.Set(token.Iserror, &object.BoolObject{Value: true})
	filePAth := (*dt).(*object.StringObject).Value
	ms := "r"
	encoding := "utf-8"
	if len(myfun.Args) >= 2 {
		configOb := (*myfun.Args[1]).(*object.Environment)
		enc, ok := configOb.Store.Get(token.Fs_encoding)
		if ok {
			encoding = enc.ToString()
		}
		mse, ok := configOb.Store.Get(token.Fs_ms)
		if ok {
			ms = mse.ToString()
		}

	}
	var file *os.File
	var err error
	if ms == "w" {
		file, err = os.OpenFile(filePAth, os.O_WRONLY|os.O_CREATE, 0666)
	} else if ms == "a" {
		file, err = os.OpenFile(filePAth, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	} else {
		file, err = os.Open(filePAth)
	}

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return fsDt
	}
	fsDt.Store.Set(token.Iserror, &object.BoolObject{Value: false})
	fs_read := Fs_read
	fs_readCont := Fs_readCont
	fs_close := Fs_close
	fs_write := Fs_write
	dd := &FileMy{File: file, Encoding: encoding, Ms: ms}
	fsDt.Store.Set(token.Fs_read, new_Func(&fs_read, token.File, dd))
	fsDt.Store.Set(token.Fs_readCont, new_Func(&fs_readCont, token.File, dd))
	fsDt.Store.Set(token.Fs_close, new_Func(&fs_close, token.File, dd))
	fsDt.Store.Set(token.Fs_write, new_Func(&fs_write, token.File, dd))
	return fsDt
}
