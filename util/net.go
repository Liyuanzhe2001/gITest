package util

import (
	"encoding/json"
	"errors"
	"gITest/data"
	"io"
	"net/http"
	"strings"
)

var Net = &net{}

type net struct{}

// 检查url是否包含scheme
func addSchemeToURL(rawURL string) string {
	if strings.HasPrefix(rawURL, "http://") || strings.HasPrefix(rawURL, "https://") {
		return rawURL
	}
	return "http://" + rawURL
}

// 解析请求头
func parseHeaders(req *http.Request, headers []string) ([]byte, error) {
	for _, s := range headers {
		if !strings.Contains(s, ":") {
			return nil, errors.New("Headers format error")
		}
		split := strings.Split(s, ":")
		req.Header.Set(split[0], split[1])
	}
	// 将请求头转换为json格式
	headersJson, err := json.MarshalIndent(req.Header, "", "  ")
	if err != nil {
		return nil, err
	}
	return headersJson, nil
}

// GET 请求
func (n *net) GET(url string, headers []string) (data.GetData, error) {
	getData := data.GetData{}

	// 检查url
	url = addSchemeToURL(url)

	// 创建客户端
	client := &http.Client{}
	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return getData, err
	}
	getData.Url = url

	// 解析请求头
	headersJson, err := parseHeaders(req, headers)
	if err != nil {
		return getData, err
	}
	getData.Headers = string(headersJson)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return getData, err
	}
	defer resp.Body.Close()

	// 获取内容
	all, err := io.ReadAll(resp.Body)
	// 获取结果失败
	if err != nil {
		return getData, err
	}

	// 用于保存json结果
	var data map[string]string
	err = json.Unmarshal(all, &data)
	// 如果err不为空 说明返回的结果不能被解析成json
	if err != nil {
		getData.Response = string(all)
		return getData, nil
	}
	// 使用json.MarshalIndent格式化输出
	responseJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return getData, err
	}
	getData.Response = string(responseJson)

	return getData, nil
}

func (n *net) POST(url string, headers []string, body string) {

}
