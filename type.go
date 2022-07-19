package qtype

import (
	"html/template"
	"net/http"
)

type H map[string]interface{}

// HandlerFunc 整个框架的核心类型
type HandlerFunc func(c DefaultContext)

// DefaultEngine 整个项目的核心模块
type DefaultEngine interface {
	GetView() DefaultView
	Get(key string) any
}

// DefaultRouter 默认路由模块，可以在框架启动进行替换
type DefaultRouter interface {
	Add(method string, pattern string, handler HandlerFunc)
	Handle(c DefaultContext)
}

// DefaultView 默认渲染模块，可以在框架启动时进行替换
type DefaultView interface {
	Render(name string, data map[string]interface{}) (string, error)

	SetFuncMap(funcMap template.FuncMap)
	LoadHtmlGlob(pattern string)
}

// DefaultContext 整个请求的完整上下文 负责挂载请求上的所有逻辑
type DefaultContext interface {
	// Next 中间件中负责进行下一个流程的方法
	Next()

	// Method 获取请求的方法
	Method() string
	// Path 获取请求的pathname
	Path() string

	// Writer 获取响应体
	Writer() http.ResponseWriter
	// Req 获取请求体
	Req() *http.Request

	SetHandler(handlers HandlerFunc)
	SetHandlers(handlers []HandlerFunc)
	// SetEngin 挂载整个engin
	SetEngin(engin DefaultEngine)
	// SetParams 设置URL上的动态params参数
	SetParams(params map[string]string)
	// SetHeader 设置响应头
	SetHeader(key string, value string) DefaultContext

	Body(key string) string
	Query(key string) string
	Param(key string) string
	Header(key string) string

	// Status 设置响应code
	Status(code int) DefaultContext
	// End 设置响应结果
	End(text string)
	// Json 设置json响应结果
	Json(obj interface{})
	// Fail 设置失败响应
	Fail(err string)
	// 设置渲染响应
	Render(name string, data map[string]interface{})
	// Redirect重定向
	Redirect(code int, location string)
}
