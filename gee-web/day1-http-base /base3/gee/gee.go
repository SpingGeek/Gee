// @Author Gopher
// @Date 2025/1/31 20:56:00
// @Desc
package gee // 定义包名为 gee，表示这个代码实现了一个简单的 Web 框架

import (
	"fmt"      // 导入 fmt 包，用于格式化输出
	"net/http" // 导入 net/http 包，用于创建 HTTP 服务和处理 HTTP 请求
)

// HandlerFunc 定义了请求处理函数的类型，它接收 http.ResponseWriter 和 *http.Request
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现了 http.Handler 接口的 ServeHTTP 方法
// 它负责管理路由，并且可以根据请求的 URL 路径调用相应的处理函数
type Engine struct {
	router map[string]HandlerFunc // 路由表，存储不同请求方式和路径对应的处理函数
}

// New 是 Engine 类型的构造函数，返回一个新创建的 Engine 实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)} // 初始化一个空的路由表
}

// addRoute 用于将请求方法（GET、POST 等）和 URL 路径映射到具体的处理函数
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	// 将请求方法和路径组合成一个唯一的键
	key := method + "-" + pattern
	// 将请求方法和路径对应的处理函数添加到路由表中
	engine.router[key] = handler
}

// GET 用于将 GET 请求与处理函数绑定
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	// 调用 addRoute 将 GET 请求与对应的处理函数绑定
	engine.addRoute("GET", pattern, handler)
}

// POST 用于将 POST 请求与处理函数绑定
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	// 调用 addRoute 将 POST 请求与对应的处理函数绑定
	engine.addRoute("POST", pattern, handler)
}

// Run 启动一个 HTTP 服务器，监听给定的地址并开始接受请求
func (engine *Engine) Run(addr string) (err error) {
	// 启动 HTTP 服务，监听 addr 地址，并使用 Engine 实现的 ServeHTTP 方法来处理请求
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 实现了 http.Handler 接口的方法
// 它根据请求的 HTTP 方法和路径从路由表中查找并执行相应的处理函数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 根据请求方法和路径构造一个唯一的路由键
	key := req.Method + "-" + req.URL.Path
	// 查找路由表中是否有对应的处理函数
	if handler, ok := engine.router[key]; ok {
		// 如果找到了处理函数，就调用它来处理请求
		handler(w, req)
	} else {
		// 如果没有找到对应的处理函数，则返回 404 错误
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
