package main

import (
	"log"
	"net"
	"net/http"
)

type handlefunc func(c *Context)

type Server interface {
	http.Handler
	Start(addr string) error
}

var _ Server = &HttpServer{}

type HttpServer struct {
	*router
}

func newHTTPServer() *HttpServer {
	return &HttpServer{
		router: newRouter(),
	}
}

// ServeHTTP 处理请求的入口
func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		resp: w,
		req:  r,
	}
	// 查找路由，然后执行命中的路由逻辑
	s.serve(ctx)
}

func (s *HttpServer) serve(c *Context) {
	panic("implement me")
}

func (s *HttpServer) Start(addr string) error {
	//return http.ListenAndServe(addr, s)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// 中间执行一些回调操作，或者业务需要的前置条件
	return http.Serve(l, s)
}

// Get 是一个所谓 restful 风格的 API，但是这里抽象出来只是为了说明：接口可以保证尽可能的小，实现可以是复杂的
// 在这里就是 Server 接口作为核心 API 始终保持干净和必要，但 HttpServer 实现却是可以由用户自己决定扩展
func (s *HttpServer) Get(route string, handler handlefunc) {
	s.addRoute(http.MethodGet, route, handler)
}

func main() {
	s := &HttpServer{}
	err := s.Start(":8080")
	if err != nil {
		log.Fatalln("Launch failed")
	}
}
