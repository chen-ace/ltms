package router

import art "github.com/plar/go-adaptive-radix-tree"

type Router struct {
	tree art.Tree
}

func NewRouter() *Router {
	return &Router{
		tree: art.New(),
	}
}

func (r Router) Register(path string, handler func()) {
	r.tree.Insert(art.Key(path), handler)
}

func (r Router) Call(path string) {
	if handler, found := r.tree.Search(art.Key(path)); found {
		handler.(func())()
	}
}
