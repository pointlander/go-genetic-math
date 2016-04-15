package engine

type Context struct {
	variables map[string]float64
}

var EmptyContext *Context = nil

func NewContext() *Context {
	return &Context{
		variables: map[string]float64{},
	}
}

func (node *Context) GetVariable(variable string) float64 {
	return node.variables[variable]
}

func (node *Context) SetVariable(variable string, value float64) {
	node.variables[variable] = value
}
