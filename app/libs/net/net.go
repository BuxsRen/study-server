package net

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type net struct {
	c           *http.Client
	req         *http.Request
	returnToMap bool
}

// 网络类
/**
 * @Example:
	c := net.New("http://127.0.0.1","POST","a=1&b=2")
	s ,_ := c.Do()
	fmt.Println(s)
*/
func New(url, method, data string) *net {
	var n = net{}
	n.c = &http.Client{}
	n.req, _ = http.NewRequest(method, url, strings.NewReader(data))
	n.req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return &n
}

// 设置请求地址
/**
 * @param scheme string 协议 http/https
 * @param host string 域名
 * @Example:
	c := net.New()
	c.SetRequestUrl("http","127.0.0.1")
*/
func (n *net) SetRequestUrl(scheme, host string) *net {
	n.req.URL.Scheme = scheme
	n.req.Host = host
	return n
}

// 设置请求头
/**
 * @param key string 键
 * @param val string 值
 * @Example:
	c := net.New()
	c.SetHeader("Content-Type","application/json")
*/
func (n *net) SetHeader(key, val string) *net {
	if n.req != nil {
		n.req.Header.Set(key, val)
	}
	return n
}

// 设置请求方式.GET/POST 等
func (n *net) SetMethod(method string) *net {
	if n.req != nil {
		n.req.Method = method
	}
	return n
}

// 设置返回数据是否json转map
func (n *net) SetReturnToMap(is bool) *net {
	if n.req != nil {
		n.returnToMap = is
	}
	return n
}

// 发送请求，并返回请求数据
func (n *net) Do() (interface{}, error) {
	if n.req == nil {
		return nil, errors.New("请先初始化:")
	}
	res, err := n.c.Do(n.req)
	if err != nil {
		return nil, err
	}
	body, e := io.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}
	defer res.Body.Close()
	if n.returnToMap {
		data := make(map[string]interface{})
		e := json.Unmarshal(body, &data)
		if e != nil {
			return nil, e
		}
		return data, nil
	}
	return string(body), nil
}
