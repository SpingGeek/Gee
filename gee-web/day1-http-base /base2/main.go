// @Author Gopher  // 代码作者
// @Date 2025/1/31 20:41:00  // 代码编写时间
// @Desc 实现了一个简单的 HTTP 服务器，可以处理两种特定的路径 / 和 /hello，对于其他路径则返回 404 错误

package main // 定义该文件为一个 Go 可执行程序（main 包）

import (
	"fmt"      // 导入 fmt 包，用于格式化输出
	"log"      // 导入 log 包，用于记录日志
	"net/http" // 导入 net/http 包，用于创建 HTTP 服务
)

// 定义一个 Engine 结构体（不包含任何字段），用于实现 http.Handler 接口
type Engine struct {
}

// 实现 http.Handler 接口的 ServeHTTP 方法，这样 Engine 就可以作为 HTTP 处理器使用
func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 根据请求的 URL 路径判断不同的处理逻辑
	switch req.URL.Path {
	case "/": // 如果路径是根路径（"/"）
		// 向客户端写入响应，返回请求的路径信息
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello": // 如果路径是 "/hello"
		// 遍历请求头中的所有键值对，返回给客户端
		for k, v := range req.Header {
			// 将每个请求头键值对写入响应
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default: // 如果路径是其他不支持的路径
		// 返回 404 错误信息，告知客户端未找到对应的路径
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// main 函数是程序的入口点
func main() {
	// 创建一个 Engine 类型的指针（Engine 结构体本身没有字段，所以这里只是一个简单的类型）
	engin := new(Engine)

	// 启动 HTTP 服务器，监听端口 9999，将请求交给 engin 处理
	// http.ListenAndServe 是阻塞式的，它会一直监听直到出现错误
	log.Fatal(http.ListenAndServe(":9999", engin)) // 使用 log.Fatal 来捕获错误，并打印到标准输出
}
