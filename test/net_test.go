package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestNetGet(t *testing.T) {
	url := "https://www.baidu.com"
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}

func TestPing(t *testing.T) {
	start := time.Now()
	url := "http://example.com"

	// 创建一个HTTP客户端
	client := &http.Client{}

	// 使用HEAD方法创建一个请求
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 发送请求并等待响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer io.Copy(io.Discard, resp.Body) // 确保读取Body直到结束
	defer resp.Body.Close()              // 关闭响应体

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("URL %s is reachable. Status: %s\n", url, resp.Status)
	} else {
		fmt.Printf("URL %s returned status: %s\n", url, resp.Status)
	}

	// 打印请求所花费的时间
	fmt.Printf("Request took %v\n", time.Since(start))
}

func TestUrlParse(t *testing.T) {
	// 定义一个 URL 字符串
	urlString := "http://www.example.com/path?query=123&param=value"

	// 解析 URL
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("无效的 URL:", err)
		return
	}

	// 检查并打印 URL 的各个部分
	fmt.Println("Scheme:", parsedURL.Scheme)     // 协议
	fmt.Println("Host:", parsedURL.Host)         // 主机名
	fmt.Println("Path:", parsedURL.Path)         // 路径
	fmt.Println("RawQuery:", parsedURL.RawQuery) // 查询字符串

	// 解析查询参数
	query := parsedURL.Query()
	fmt.Println("Query Params:", query)
}

/*
Scheme: http
Host: www.example.com
Path: /path
RawQuery: query=123&param=value
Query Params: map[param:[value] query:[123]]
*/
/*
Scheme:
Host:
Path: www.example.com/path
RawQuery: query=123&param=value
Query Params: map[param:[value] query:[123]]
*/

func TestNet(t *testing.T) {
	url := "localhost:8080"
	parsedURL, err := parse(url)
	// 检查并打印 URL 的各个部分
	fmt.Println("Scheme:", parsedURL.Scheme)     // 协议
	fmt.Println("Host:", parsedURL.Host)         // 主机名
	fmt.Println("Path:", parsedURL.Path)         // 路径
	fmt.Println("RawQuery:", parsedURL.RawQuery) // 查询字符串

	// 解析查询参数
	query := parsedURL.Query()
	fmt.Println("Query Params:", query)
	if err != nil {
		fmt.Println(err)
		return
	}
	if parsedURL.Scheme == "" {
		url = "http://" + url
		parsedURL, err = parse(url)
	}
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//fmt.Println(parsedURL)
}

func parse(str string) (*url.URL, error) {
	// 解析 URL
	parsedURL, err := url.Parse(str)
	if err != nil {
		return new(url.URL), err
	}
	return parsedURL, nil
}

func TestPing2(t *testing.T) {
	//url := "http://localhost:8081" // 假设你的服务运行在8080端口
	url := "localhost:8081"

	// 创建一个HTTP客户端
	client := &http.Client{}

	// 使用GET方法创建一个请求
	resp, err := client.PostForm(url, nil)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer io.Copy(io.Discard, resp.Body) // 确保读取Body直到结束
	defer resp.Body.Close()              // 关闭响应体

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Localhost is reachable. Status:", resp.Status)
	} else {
		fmt.Println("Localhost returned status:", resp.Status)
	}
}

func addSchemeToURL(rawURL string) string {
	if strings.HasPrefix(rawURL, "http://") || strings.HasPrefix(rawURL, "https://") {
		return rawURL
	}
	return "http://" + rawURL
}

func TestGET(t *testing.T) {
	//url := "localhost:8080?a=aa&b=bb"
	url := "baidu.com"
	url = addSchemeToURL(url)
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	var data map[string]string
	err = json.Unmarshal(all, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	// 使用json.MarshalIndent格式化输出
	indentedJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to indented JSON:", err)
		return
	}

	// 打印格式化的JSON
	fmt.Println(string(indentedJson))
}

type Greeting struct {
	Message string `json:"message"`
}

func TestTest(t *testing.T) {
	g := Greeting{Message: "Hello, world!"}
	jsonString, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonString))
}
