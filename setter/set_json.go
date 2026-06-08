package setter

import "github.com/tidwall/sjson"

func Setjson(jsonData string, fieldName string, value any) string {

	update, err := sjson.Set(jsonData, fieldName, value)

	if err != nil {
		panic(err)
	}

	return update
}
