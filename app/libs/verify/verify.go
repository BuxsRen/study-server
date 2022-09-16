package verify

import (
	"encoding/json"
	"fmt"
	"reflect"
	"study-server/app/libs/utils"
)

// Bind 验证data参数，需要go版本支持泛型
/*

@Example:
	type Param struct {
		Id int `bind:"required,between=4:5" json:"id"`
	}

	data := map[string]interface{}{
		"id": 111,
	}

	param,_ = verify.Bind(data, Param{})
*/
func Bind[T any](data map[string]interface{}, stu T) (T, error) {
	b, e := json.Marshal(data)

	if e != nil {
		return stu, e
	}

	_ = json.Unmarshal(b, &stu)

	t := reflect.TypeOf(stu)

	for i := 0; i < t.NumField(); i++ {

		tag := t.Field(i).Tag.Get("bind")
		param := t.Field(i).Tag.Get("json")

		veritfy(data, param, tag)

		t1 := t.Field(i).Type.String()
		t2 := fmt.Sprintf("%T", data[param])

		if t1 != t2 {
			utils.ExitError(
				fmt.Sprintf(
					"参数:%s,提交的类型是:%s,期望的类型是:%s",
					param, t2, t1),
				-1)
		}

	}
	return stu, nil
}
