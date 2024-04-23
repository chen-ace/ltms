package router

import art "github.com/plar/go-adaptive-radix-tree"

type Router struct {
	tree art.Tree
}

type Request struct {
	data map[string]interface{}
}

type Response = Request

func (r Router) Call(path string, request Request) (Response, bool) {
	if handler, found := r.tree.Search(art.Key(path)); found {
		response := handler.(func(Request) Response)(request)
		return response, true
	}
	return Response{}, false
}

func (r Router) Register(path string, handler func(Request) Response) {
	r.tree.Insert(art.Key(path), handler)
}

func (req Request) Get(key string) interface{} {
	return req.data[key]
}

func (req Request) Set(key string, value interface{}) {
	req.data[key] = value
}
