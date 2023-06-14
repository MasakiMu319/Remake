package main

import "strings"

type router struct {
	trees map[string]*node
}

type node struct {
	path     string
	children map[string]*node
	handler  handlefunc
}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}

// addRoute is an inner method, so we don't need check the value
// of method param when user use it.
func (r *router) addRoute(method string, route string, handler handlefunc) {
	if route == "" {
		panic("Web route can't be empty")
	}
	// 处理根节点不存在的情况
	root, ok := r.trees[method]
	if !ok {
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}

	if route[0] != '/' {
		panic("the begin of route must be '/'")
	}
	if route != "/" && route[len(route)-1] == '/' {
		panic("the end of route can't be '/'")
	}
	if route == "/" {
		// resolve duplicate route
		if root.handler != nil {
			panic("the route can't be duplicated")
		}
		root.handler = handler
		return
	}
	// 对于 strings.Split 方法会导致第一个匹配上的 sep 左右两个元素都被视为有效分割内容
	// 例如 /user/home => "", "user", "home"
	route = route[1:]
	// 以达到匹配根路径的情况
	segs := strings.Split(route, "/")
	for _, seg := range segs {
		if seg == "" {
			panic("web route can't contain continuous '/'")
		}
		child := root.childOrCreate(seg)
		root = child
	}
	if root.handler != nil {
		panic("the route can't be duplicated")
	}
	root.handler = handler
}

func (x *node) childOrCreate(seg string) *node {
	if x.children == nil {
		x.children = map[string]*node{}
	}
	child, ok := x.children[seg]
	if !ok {
		child = &node{
			path: seg,
		}
		x.children[seg] = child
	}
	return child
}
