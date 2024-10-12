package etree

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"myvmp/object"
	"myvmp/token"
	"strings"
)

func new_Func(ddd *func(*object.FunctionDeclarationObject) object.Object, typeS string, Obj any) object.Object {
	d := &object.FunctionDeclarationObject{IsNative: 1, NativeBody: ddd, BindType: typeS, BindOb: Obj}
	return d
}

func Etree_gethtml(myfun *object.FunctionDeclarationObject) object.Object {
	doc := myfun.BindOb.(*html.Node)
	ggg := htmlquery.OutputHTML(doc, true)

	return &object.StringObject{Value: ggg}
}

func Etree_xpath(myfun *object.FunctionDeclarationObject) object.Object {
	Etree := myfun.BindOb.(*html.Node)
	dtt := strings.TrimSpace((*myfun.Args[0]).(*object.StringObject).Value)
	gettext := 0
	if len(dtt) >= 7 && dtt[len(dtt)-7:] == "/text()" {
		gettext = 1
	} else {
		ddkk := strings.Split(dtt, "/")
		dc := ddkk[len(ddkk)-1][0:1]
		if dc == "@" {
			gettext = 1
		}
	}
	NodeList := htmlquery.Find(Etree, dtt)
	dss := object.NewArray()
	for _, ss := range NodeList {
		if gettext == 0 {
			parseHtml := object.NewEnv(nil)
			parseHtml.TypeInfo = token.THIS
			etreeXpath := Etree_xpath
			etreeXpath2 := new_Func(&etreeXpath, token.Etree, ss)
			parseHtml.Store.Set(token.Etree_xpath, etreeXpath2)

			etree_gethtml := Etree_gethtml
			etree_gethtml2 := new_Func(&etree_gethtml, token.Etree, ss)
			parseHtml.Store.Set(token.Etree_gethtml, etree_gethtml2)

			var hh object.Object = parseHtml
			dss.Value = append(dss.Value, &hh)

		} else {
			ddd := htmlquery.InnerText(ss)
			h := &object.StringObject{Value: ddd}
			var hh object.Object = h
			dss.Value = append(dss.Value, &hh)
		}

	}

	return &dss
}

func Etree_HTML(myfun *object.FunctionDeclarationObject) object.Object {
	html := (*myfun.Args[0]).ToString()
	doc, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return nil
	}
	parseHtml := object.NewEnv(nil)
	etreeXpath := Etree_xpath
	etreeXpath2 := new_Func(&etreeXpath, token.Etree, doc)
	parseHtml.Store.Set(token.Etree_xpath, etreeXpath2)
	etree_gethtml := Etree_gethtml
	etree_gethtml2 := new_Func(&etree_gethtml, token.Etree, doc)
	parseHtml.Store.Set(token.Etree_gethtml, etree_gethtml2)

	return parseHtml
}
