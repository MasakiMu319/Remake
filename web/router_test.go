package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestRouter_AddRoute(t *testing.T) {
	testRoutes := []struct {
		method string
		path   string
	}{
		{
			method: http.MethodGet,
			path:   "/",
		},
		{
			method: http.MethodGet,
			path:   "/user",
		},
		{
			method: http.MethodGet,
			path:   "/user/home",
		},
		{
			method: http.MethodGet,
			path:   "/order/detail",
		},
		{
			method: http.MethodPost,
			path:   "/order/create",
		},
		{
			method: http.MethodPost,
			path:   "/login",
		},
	}

	mockHandler := func(c *Context) {}
	r := newRouter()
	for _, tr := range testRoutes {
		r.addRoute(tr.method, tr.path, mockHandler)
	}

	wantRouter := &router{
		trees: map[string]*node{
			http.MethodGet: &node{
				path: "/",
				children: map[string]*node{
					"user": &node{
						path: "user",
						children: map[string]*node{
							"home": &node{
								path:    "home",
								handler: mockHandler,
							},
						},
						handler: mockHandler,
					},
					"order": &node{
						path: "order",
						children: map[string]*node{
							"detail": &node{
								path:    "detail",
								handler: mockHandler,
							},
						},
					},
				},
				handler: mockHandler,
			},
			http.MethodPost: &node{
				path: "/",
				children: map[string]*node{
					"order": &node{
						path: "order",
						children: map[string]*node{
							"create": &node{
								path:    "create",
								handler: mockHandler,
							},
						},
					},
					"login": &node{
						path:    "login",
						handler: mockHandler,
					},
				},
			},
		},
	}

	msg, ok := r.equal(wantRouter)
	assert.True(t, ok, msg)

	r = newRouter()
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "", mockHandler)
	}, "the route can't be empty")
	r = newRouter()
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/a//b", mockHandler)
	}, "the route can't have continuous '/'")

	r = newRouter()
	r.addRoute(http.MethodGet, "/", mockHandler)
	assert.Panicsf(t, func() {
		r.addRoute(http.MethodGet, "/", mockHandler)
	}, "the route can't be duplicated")
}

func (x *router) equal(y *router) (string, bool) {
	for k, v := range x.trees {
		dest, ok := y.trees[k]
		if !ok {
			return fmt.Sprintf("target method tree doesn't exist"), false
		}
		msg, ok := v.equal(dest)
		if !ok {
			return msg, false
		}
	}
	return "", true
}

func (x *node) equal(y *node) (string, bool) {
	if x.path != y.path {
		return fmt.Sprintf("node path is not equal"), false
	}
	if len(x.children) != len(y.children) {
		return fmt.Sprintf("node's children number is not equal"), false
	}

	xHandler := reflect.ValueOf(x.handler)
	yHandler := reflect.ValueOf(y.handler)
	if xHandler != yHandler {
		return fmt.Sprintf("node handler is not equal"), false
	}

	for path, c := range x.children {
		dest, ok := y.children[path]
		if !ok {
			return fmt.Sprintf("targe children node: %s doesn't exist", path), false
		}
		msg, ok := c.equal(dest)
		if !ok {
			return msg, false
		}
	}
	return "", true
}
