package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// 将 interface{} 类型转成 string 类型
func ParamToString(val interface{}) string {
	var str string
	T := fmt.Sprintf(`%T`, val)
	switch T {
	case "float64":str = strconv.FormatFloat(val.(float64),'f',-1,64)
	case "[]string":str = val.([]string)[0]
	case "string":str = val.(string)
	case "int":str = strconv.Itoa(val.(int))
	case "int64":str = strconv.Itoa(int(val.(int64)))
	case "[]uint8":str = string(val.([]uint8))
	case "time.Time":str = val.(time.Time).Format("2006-01-02 15:04:05")
	}
	return str
}

// map 转 url 请求参数
func HttpBuildQuery(data map[string]interface{}) string {
	var uri url.URL
	q := uri.Query()
	for k,v := range data {
		q.Add(k, ParamToString(v))
	}
	return q.Encode()
}
