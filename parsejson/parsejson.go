package parsejson

import (
	"encoding/json"
	"fmt"
	"myvmp/object"
	"myvmp/token"
)

func printJSON(data interface{}, envobject *object.Environment) *object.Environment {
	switch v := data.(type) {
	case map[string]interface{}:
		dataeenv := newObject()
		for key, value := range v {
			switch value.(type) {
			case string:
				dataeenv.Store.Set(key, &object.StringObject{
					Value: value.(string),
				})
			case int:
				dataeenv.Store.Set(key, &object.NumericObject{
					Value: float64(value.(int)),
				})
			case float64:
				dataeenv.Store.Set(key, &object.NumericObject{
					Value: value.(float64),
				})
			case bool:
				dataeenv.Store.Set(key, &object.BoolObject{
					Value: value.(bool),
				})
			default:
				dataeenv.Store.Set(key, printJSON(value, envobject))

			}

		}
		return dataeenv
	case []interface{}:
		dataeenv := newArray()
		for _, value := range v {
			switch value.(type) {
			case string:
				var dasda object.Object = &object.StringObject{
					Value: value.(string),
				}
				dataeenv.Value = append(dataeenv.Value, &dasda)
			case int:
				var dasda object.Object = &object.NumericObject{
					Value: float64(value.(int)),
				}
				dataeenv.Value = append(dataeenv.Value, &dasda)
			case float64:
				var dasda object.Object = &object.NumericObject{
					Value: value.(float64),
				}
				dataeenv.Value = append(dataeenv.Value, &dasda)
			case bool:
				var dasda object.Object = &object.BoolObject{
					Value: value.(bool),
				}
				dataeenv.Value = append(dataeenv.Value, &dasda)
			default:
				var dasda object.Object = printJSON(value, envobject)
				dataeenv.Value = append(dataeenv.Value, &dasda)
			}

		}
		return dataeenv
	default:
		fmt.Printf("%v\n", v)
	}
	return nil
}

func newObject() *object.Environment {
	dataeenv := object.NewObject()
	return &dataeenv

}
func newArray() *object.Environment {
	dataeenv := object.NewArray()
	return &dataeenv

}

func ParseStrToJson(jsonStr string) object.Object {

	var data map[string]interface{}
	var ListDt []interface{}
	var dsad object.Object
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		if errw := json.Unmarshal([]byte(jsonStr), &ListDt); errw != nil {
			fmt.Println("Error parsing JSON:", errw)
			return nil
		} else {
			dsad = printJSON(ListDt, nil)
		}

	} else {
		dsad = printJSON(data, nil)
	}

	return dsad
}

func stringifyTostr2(objec3 []*object.Object) []interface{} {
	dt := []interface{}{}
	for _, valuee := range objec3 {
		value := (*valuee)
		switch value.Type() {
		case token.TYNUM:
			dsd := value.(*object.NumericObject).Value
			dt = append(dt, dsd)
		case token.TYSTR:
			dsd := value.(*object.StringObject).Value
			dt = append(dt, dsd)
		case token.Object:
			dsd := value.(*object.Environment).Store.M
			dt = append(dt, stringifyTostr(dsd))
		case token.BOOL:
			dsd := value.(*object.BoolObject).Value
			dt = append(dt, dsd)
		case token.ArrayE:
			dsd := value.(*object.Environment).Value

			dt = append(dt, stringifyTostr2(dsd))
		default:
			dt = append(dt, token.YOUZ+value.Type()+token.ZUOZ)
		}
	}
	return dt
}

func stringifyTostr(objec3 map[string]object.Object) map[string]interface{} {
	dt := map[string]interface{}{}
	for key, value := range objec3 {
		switch value.Type() {
		case token.TYNUM:
			dsd := value.(*object.NumericObject).Value
			dt[key] = dsd
		case token.TYSTR:
			dsd := value.(*object.StringObject).Value
			dt[key] = dsd
		case token.Object:
			dsd := value.(*object.Environment).Store.M
			dt[key] = stringifyTostr(dsd)
		case token.BOOL:
			dsd := value.(*object.BoolObject).Value
			dt[key] = dsd
		case token.ArrayE:
			dsd := value.(*object.Environment).Value
			dt[key] = stringifyTostr2(dsd)
		default:
			continue
		}
	}
	return dt
}
func JsonToStr(objec3 *object.Environment) object.Object {
	switch objec3.Type() {
	case token.ArrayE:
		dsd := objec3.Value
		dgg := stringifyTostr2(dsd)
		jsonString, err := json.Marshal(dgg)
		if err != nil {
			//fmt.Println("Error marshaling to JSON:", err)
			return &object.StringObject{Value: "[]"}
		}
		return &object.StringObject{Value: string(jsonString)}
	case token.Object:
		dsad := objec3.Store.M
		dgg := stringifyTostr(dsad)
		jsonString, err := json.Marshal(dgg)
		if err != nil {
			//fmt.Println("Error marshaling to JSON:", err)
			return &object.StringObject{Value: "{}"}
		}
		return &object.StringObject{Value: string(jsonString)}

	}
	return &object.StringObject{Value: token.YOUZ + token.THIS + token.ZUOZ}

}
