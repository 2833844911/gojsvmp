package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"log"
	"strings"
)

func main() {
	// 示例HTML内容
	htmlContent := `
    <html>
        <body>
            <div>
                <h1>Hello, 世界!</h1>
                <p class="greeting">Welcome to Go!</p>
            </div>
        </body>
    </html>
    `

	// 解析HTML
	doc, err := htmlquery.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// 使用XPath获取h1标签的内容
	h1Node := htmlquery.FindOne(doc, "//h1")
	if h1Node != nil {
		fmt.Println("H1:", htmlquery.InnerText(h1Node))
	}

	// 使用XPath获取p标签的内容
	pNode := htmlquery.Find(doc, "//*/text()")
	if pNode != nil {
		fmt.Println("P:", htmlquery.InnerText(pNode[5]))
	}
}
