package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"study-server/app/libs/encry"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
	"strconv"
	"strings"
)

// 捕获所有的请求参数
type Request struct {
	c    *gin.Context
	Info struct {
		NowTime    int64  // 当前时间戳
		TodayStart int64  // 今日开始时间戳
		TodayEnd   int64  // 今日结束时间戳
		Ip         string // 请求IP
		UserAgent  string // 浏览器标识
	}
	Param map[string]interface{}
}

// 入口方法
func Handle(c *gin.Context) {
	global := Request{c: c}
	global.header()
	global.getAllParam()
	if global.Param == nil {
		global.c.Abort()
	} else {
		global.addParam()
		if global.sign() {
			global.c.Next()
		} else {
			global.c.Abort()
		}
	}
}

// 构建所有的请求参数
func (r *Request) getAllParam() {
	r.Param = r.getParam()
}

// 取所有请求参数
func (r *Request) getParam() map[string]interface{} {
	data := make(map[string]interface{})
	for k, v := range r.c.Request.URL.Query() { // 收集Url上面的参数
		data[k] = v[0]
	}
	if r.c.Request.Method != "GET" { // 收集非GET请求的 json 参数
		if strings.Contains(r.c.Request.Header.Get("Content-Type"), "json") { // this.c.ContentType() == "application/json"
			param, _ := io.ReadAll(r.c.Request.Body)
			r.c.Request.Body = io.NopCloser(bytes.NewBuffer(param)) // 数据重新写回 body
			e := json.Unmarshal(param, &data)
			if e != nil {
				if config.App.Server.Template {
					r.c.HTML(400, "400.tpl", gin.H{"msg": e.Error()})
				} else {
					r.c.JSON(400, gin.H{"msg": e.Error(), "data": nil})
				}
				return nil
			}
		} else {
			if strings.Contains(r.c.Request.URL.Path, "/api") && !strings.Contains(r.c.Request.URL.Path, "upload") { // api，且不是上传接口，只能传json
				if config.App.Server.Template {
					r.c.HTML(405, "405.tpl", gin.H{"msg": "Only Support Json Content Type"})
				} else {
					r.c.JSON(405, gin.H{"code": 405, "msg": "Only Support Json Content Type", "data": nil})
				}
				return nil
			} else { // 收集非GET请求的 form-data 参数
				_ = r.c.Request.ParseMultipartForm(128)
				for k, v := range r.c.Request.PostForm {
					data[k] = v[0]
				}
			}
		}
	}
	return data
}

// 设置头部信息
func (r *Request) header() {
	r.c.Header("Go-Server", "gin/"+gin.Version)
}

// 追加参数
func (r *Request) addParam() {
	var ip string
	if r.c.Request.Header.Get("ALI-CDN-REAL-IP") != "" {
		ip = r.c.Request.Header.Get("ALI-CDN-REAL-IP")
	} else {
		ip = r.c.ClientIP()
	}
	var user_agent string
	if len(r.c.Request.Header["User-Agent"]) != 0 {
		user_agent = r.c.Request.Header["User-Agent"][0]
	}
	t, ts, te := utils.GetTodayTime()

	r.Info.NowTime = t
	r.Info.TodayStart = ts
	r.Info.TodayEnd = te
	r.Info.Ip = ip
	r.Info.UserAgent = user_agent

	r.c.Set("_data", &r.Param)
	r.c.Set("_req", r)
}

// 签名效验
func (r *Request) sign() bool {
	data := r.Param
	clientSign := data["sign"]
	if clientSign != nil {
		delete(data, "sign")
		for k, v := range data {
			if k[0:1] == "_" { // 忽略 _ 开头的参数
				delete(data, k)
			}
			if k == "num" {
				T := fmt.Sprintf(`%T`, v) // 取变量类型
				if T == "float64" {       // json 类型
					data[k] = string(rune(v.(float64))) // 字母转ascii
				} else {
					iv, _ := strconv.Atoi(v.([]string)[0])
					data[k] = string(rune(iv)) // 字母转ascii
				}
			}
		}
		t := utils.ParamToString(data["t"])
		sign := encry.MD5(utils.HttpBuildQuery(data))
		it, _ := strconv.ParseInt(t, 10, 64)
		if utils.GetTime()-it > 60 {
			r.c.JSON(200, gin.H{"code": -1, "msg": "sign timeout"})
			return false
		}
		if sign != clientSign {
			r.c.JSON(200, gin.H{"code": -1, "msg": "sign error"})
			return false
		}
		data["sign"] = sign
	}
	return true
}

// GetInput 取请求参数 参数key为*取所有
func GetInput(c *gin.Context, key string) interface{} {
	data := GetAllData(c)
	if data == nil {
		return nil
	}
	if key == "*" {
		return data.Param
	} else {
		return data.Param[key]
	}
}

//GetAllData 取所有请求数据
func GetAllData(c *gin.Context) *Request {
	v, ok := c.Get("_data")
	if ok {
		return v.(*Request)
	}
	return nil
}
