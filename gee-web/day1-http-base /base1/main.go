// @Author Gopher
// @Date 2025/1/31 20:23:00
// @Desc 基于 net/http 封装实现的HTTP服务，提供基础路由处理功能

package main // 声明当前文件属于main包，表示这是一个可执行程序

// 导入依赖包
import (
	"fmt"      // 格式化输入输出
	"log"      // 日志记录
	"net/http" // HTTP协议处理
)

// main函数：程序入口，管理路由注册并启动HTTP服务
func main() {
	// 使用默认的多路复用器（DefaultServeMux）注册根路径"/"的处理函数
	http.HandleFunc("/", indexHandler)
	// 注册路径"/hello"的处理函数
	http.HandleFunc("/hello", helloHandler)

	// 启动HTTP服务器，监听本地9999端口
	// ListenAndServe会阻塞执行，直到服务出现错误（如端口冲突）
	// 使用log.Fatal捕获错误并记录，同时终止程序运行
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// indexHandler 处理根路径"/"的请求
// 参数说明：
//
//	w http.ResponseWriter: 用于构建HTTP响应
//	req *http.Request: 包含客户端请求的所有信息
func indexHandler(w http.ResponseWriter, req *http.Request) {
	// 将请求的URL路径写入响应
	// Fprintf的第一个参数为实现了io.Writer接口的对象，这里传入w即直接写入HTTP响应体
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// helloHandler 处理路径"/hello"的请求
// 功能：遍历并返回请求头中的所有字段
func helloHandler(w http.ResponseWriter, req *http.Request) {
	// 遍历请求头(req.Header)中的每个键值对
	// Header类型为map[string][]string，键为头字段名，值为对应的字符串数组
	for k, v := range req.Header {
		// 将每个头字段及其值格式化写入响应体
		// 示例输出：Header["User-Agent"] = ["curl/7.64.1"]
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
