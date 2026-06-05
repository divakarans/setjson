package legacy

import (
	"encoding/json"
)

func Legacy(data string) interface{} {

	var parsedata interface{}

	err := json.Unmarshal([]byte(data), &parsedata)

	if err != nil {
		panic(err)
	}

	return parsedata

}
