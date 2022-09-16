package encry

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	uuid "github.com/satori/go.uuid" // go get github.com/satori/go.uuid
	"net/url"
	"strconv"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
	"time"
)

var tokenKey = "441479573" // token签名签发密钥

// base64解码。需要解码的字符串
func Base64Decode(str string) ([]byte, string) {
	b, _ := base64.StdEncoding.DecodeString(str)
	return b, string(b)
}

// base64编码，需要编码的字节码
func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// URL编码
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// URL解码
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

// map/结构体 转 json
func JsonEncode(src interface{}) ([]byte, error) {
	return json.Marshal(src)
}

// json 转 map/结构体
func JsonDecode(str string, dst interface{}) error {
	return json.Unmarshal([]byte(str), &dst)
}

// 生成md5
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 创建token。用户id,token到期时间戳(-1为永不过期,-24为当天过后过期)，其他数据
func EncryptToken(id string, expire int64, data map[string]interface{}) string {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["id"] = id
	data["time"] = utils.GetTime()
	data["expire"] = expire
	if config.App.Other.TokenKey != "" {
		tokenKey = config.App.Other.TokenKey
	}
	b, _ := json.Marshal(data) // map转json
	str := string(b)
	data["sign"] = MD5(str + tokenKey)
	b, _ = json.Marshal(data) // map转json
	return Base64Encode(b)
}

// 解析token，返回token解析后的内容。用户token
func DecryptToken(token string) map[string]interface{} {
	baseToken, _ := Base64Decode(token) //base64解码
	param := make(map[string]interface{})
	err := json.Unmarshal(baseToken, &param) // json 转 map
	if err != nil {
		return nil
	}
	var sign = ""
	if param["sign"] != nil {
		sign = param["sign"].(string)
		delete(param, "sign")
	}
	data, _ := json.Marshal(param) // map转json
	if config.App.Other.TokenKey != "" {
		tokenKey = config.App.Other.TokenKey
	}
	if MD5(string(data)+tokenKey) == sign { // 效验sign
		if param["time"] == nil {
			param["time"] = float64(utils.GetTime())
		}
		if param["expire"] == nil {
			param["expire"] = float64(utils.GetTime())
		}
		timestamp := int64(param["time"].(float64))
		expire := int64(param["expire"].(float64))
		day, _ := strconv.Atoi(time.Unix(timestamp, 0).In(time.FixedZone("CST", 8*3600)).Format("02"))
		if expire != -24 && utils.GetNow().Day() != day { // token不是当天的 1 = 1
			return nil
		}
		if expire != -1 && expire > utils.GetTime() { // token 过期
			return nil
		}
		return param
	}
	return nil
}

// 生成UUID
func UUID() string {
	u := uuid.NewV4()
	return u.String()
}

// 解析UUID
func UUIDDecode(str string) (bool, error) {
	_, err := uuid.FromString(str)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 密码
func Password(pwd string) string {
	if config.App.Other.TokenKey != "" {
		tokenKey = config.App.Other.TokenKey
	}
	return MD5(tokenKey + pwd)
}
