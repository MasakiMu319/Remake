package main

import "net/http"

type Context struct {
	resp http.ResponseWriter
	req  *http.Request
}
