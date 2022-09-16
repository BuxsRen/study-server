package verify

import (
	"fmt"
	"strconv"
	"strings"
	"study-server/app/libs/utils"
)

// 参数对应中文解释
var attributes = map[string]string{
	"username": "用户名",
	"password": "密码",
	"mobile":   "手机号",
	"title":    "标题",
}

var reg Regexp

var rule = map[string]interface{}{
	"required": map[string]interface{}{ // "required"
		"info": "参数[%s]不能为空",
		"fun":  reg.Required,
	},
	"numeric": map[string]interface{}{ // "numeric"
		"info": "参数[%s]类型必须为数字",
		"fun":  reg.Numeric,
	},
	"string": map[string]interface{}{
		"info": "参数[%s]类型必须为字符串",
		"fun":  reg.String,
	},
	"between": map[string]interface{}{ // "between:5,10"
		"info": "参数[%s]的长度必须大于%d且小于%d",
		"fun":  reg.Between,
	},
	"size": map[string]interface{}{ // "size:5,10"
		"info": "参数[%s]必须大于%d且小于%d",
		"fun":  reg.Size,
	},
	"alpha": map[string]interface{}{
		"info": "参数[%s]必须是字母构成",
		"fun":  reg.Alpha,
	},
	"alpha_num": map[string]interface{}{
		"info": "参数[%s]必须是字母和数字构成",
		"fun":  reg.AlphaNum,
	},
	"alpha_dash": map[string]interface{}{
		"info": "参数[%s]必须是字母，数字或特殊字符其中一种构成",
		"fun":  reg.AlphaDash,
	},
	"alpha_dash_all": map[string]interface{}{
		"info": "参数[%s]必须全部由字母、数字、特殊字符构成",
		"fun":  reg.AlphaDashAll,
	},
	"password": map[string]interface{}{
		"info": "参数[%s]必须以字母开头，长度在6~18之间，只能包含字母、数字和下划线",
		"fun":  reg.Password,
	},
	"date": map[string]interface{}{
		"info": "参数[%s]必须是日期格式",
		"fun":  reg.Date,
	},
	"time": map[string]interface{}{
		"info": "参数[%s]必须是时间格式",
		"fun":  reg.Time,
	},
	"date_time": map[string]interface{}{
		"info": "参数[%s]必须是日期时间格式",
		"fun":  reg.DateTime,
	},
	"url": map[string]interface{}{
		"info": "参数[%s]必须是链接",
		"fun":  reg.Url,
	},
	"email": map[string]interface{}{
		"info": "参数[%s]必须是正确的邮箱地址",
		"fun":  reg.Email,
	},
	"mobile": map[string]interface{}{
		"info": "参数[%s]必须是正确的手机号",
		"fun":  reg.Mobile,
	},
	"id_number": map[string]interface{}{
		"info": "参数[%s]必须是正确的身份证号",
		"fun":  reg.IdNumber,
	},
}

func veritfy(src map[string]interface{}, param, item string) {
	for _, key := range strings.Split(item, ",") { // 需要验证的规则列表
		sub := strings.Split(key, "=") // 规则是否有子项
		var min, max int
		if len(sub) > 1 {
			key = sub[0]
			sub = strings.Split(sub[1], ":") // 子项的子项
			min, _ = strconv.Atoi(sub[0])
			max, _ = strconv.Atoi(sub[1])
		}
		if rule[key] == nil {
			continue
		}
		str := utils.ParamToString(src[param]) // 需要验证的参数
		veritfyToFun(key, param, str, min, max)
	}
}

func veritfyToFun(key, param, value string, num1, num2 int) {
	var str = param
	if attributes[param] != "" { // 请求参数中文对照
		str = attributes[param]
	}
	r := rule[key].(map[string]interface{})
	info := r["info"].(string)
	if key == "between" || key == "size" {
		fun := r["fun"].(func(string, int, int) bool)
		if !fun(value, num1, num2) { // 验证不通过，输出失败原因
			utils.ExitError(fmt.Sprintf(info, str, num1, num2), -1)
		}
	} else if key == "string" {
		fun := r["fun"].(func(interface{}) bool)
		if !fun(value) {
			utils.ExitError(fmt.Sprintf(info, str), -1)
		}
	} else {
		fun := r["fun"].(func(str string) bool)
		if !fun(value) {
			utils.ExitError(fmt.Sprintf(info, str), -1)
		}
	}
}
