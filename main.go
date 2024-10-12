package main

import (
	"bufio"
	"fmt"
	"github.com/peterh/liner"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"myvmp/ast"
	"myvmp/evaluator"
	"myvmp/lexer"
	"myvmp/object"
	"myvmp/parse"
	"myvmp/promise"
	"os"
	"strings"
)

func getdt(code string, env *object.Environment) {

	promise.CyJSInit()
	dt := lexer.New(code)

	kk := (*dt).Input()

	fff := parse.NewParse(kk)
	dtt := &ast.Program{Body: fff}
	dtt.StatementNode()
	evaluator.StartEval(dtt.Body, env)
	promise.Done()

}

func cmd() {

	line := liner.NewLiner()
	defer line.Close()
	fmt.Println("Welcome to the interactive CyJsShell. Type 'exit' or 'exit()' to quit.")

	line.SetCtrlCAborts(true)

	history := []string{}

	env := object.NewEnv(nil)

	for {

		code, err := line.Prompt(">>>")

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		history = append(history, code)
		line.AppendHistory(code)
		// 去掉输入字符串的换行符
		code = code + ";"
		if code == "exit;" {
			break
		} else if code == "exit();" {
			break
		}

		// 捕获 getdt(code, env) 的异常
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			getdt(code, env)
		}()

	}
	f, err := os.Create(".liner_history")
	if err != nil {
		fmt.Println("Error creating history file:", err)
		return
	}
	defer f.Close()

	line.WriteHistory(f)
}

func doFile(code string) {
	evaluator.Eval(code)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("CyJs version 1.13\n")
		cmd()
		return
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	flineee, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading first line: %v\n", err)
		return
	}
	fline := strings.TrimSpace(flineee)
	firstLine := strings.Replace(fline, " ", "", -1)

	var content string
	if firstLine == "//--utf-8--" {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		content = string(data)
	} else if firstLine == "//--gbk--" {
		gbkReader := transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder())
		data, err := ioutil.ReadAll(gbkReader)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		content = string(data)
	} else {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		content = fline + "\n" + string(data)
	}

	content = strings.TrimSpace(content) + ";"
	doFile(content)
}
