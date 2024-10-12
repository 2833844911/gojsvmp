package require

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"myvmp/object"
	"os"
	"strings"
)

type RequireInfo struct {
	DtInfo map[string]*object.Environment
	IsDo   string
}

func ReadFile(Path string) string {

	file, err := os.Open(Path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return ""
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	flineee, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading first line: %v\n", err)
		return ""
	}
	fline := strings.TrimSpace(flineee)
	firstLine := strings.Replace(fline, " ", "", -1)

	var content string
	if firstLine == "//--utf-8--" {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return ""
		}
		content = string(data)
	} else if firstLine == "//--gbk--" {
		gbkReader := transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder())
		data, err := ioutil.ReadAll(gbkReader)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return ""
		}
		content = string(data)
	} else {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return ""
		}
		content = fline + "\n" + string(data)
	}

	content = strings.TrimSpace(content) + ";"
	return content
}
