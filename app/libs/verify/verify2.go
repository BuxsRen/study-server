package verify

// 验证data参数
/**
 * @Example:

	data := map[string]interface{}{
		"name": "xxx",
	}

	VerifyData(data,map[string]string{
		"name":"required,between=4:5",
	})
*/
func VerifyData(data map[string]interface{},v map[string]string) {
	for param,item := range v { // 遍历请求参数列表
		veritfy(data,param,item)
	}
}