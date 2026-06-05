package setter

import "github.com/tidwall/sjson"

func Setjson(data string) string {
	update, res := sjson.Set(data, "name", "diva")

	if res != nil {
		panic(res)
	}

	return update

}
