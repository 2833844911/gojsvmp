package parsejson

import (
	"fmt"
	"myvmp/object"
	"testing"
)

func TestParseJson(t *testing.T) {
	// JSON字符串
	jsonString := `{
        "name": "Alice",
        "age": 30,
        "address": {
            "city": "Wonderland",
            "zip": "12345"
        },
        "hobbies": ["reading", "chess"],
        "extra": {
            "nickname": "Ally",
            "isActive": true
        }
    }`

	dada := ParseStrToJson(jsonString)
	dadadd := JsonToStr(dada.(*object.Environment))
	fmt.Println(dadadd)
}
