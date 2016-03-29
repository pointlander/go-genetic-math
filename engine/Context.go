package engine

type Context struct {
	variables map[string]float64
}

func NewContext() *Context {
	return &Context{
		variables: map[string]float64{},
	}
}

func (this *Context) GetVariable(variable string) float64 {
	return this.variables[variable]
}

func (this *Context) SetVariable(variable string, value float64) {
    this.variables[variable] = value
}
