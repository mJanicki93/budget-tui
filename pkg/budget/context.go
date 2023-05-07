package budget

type Context map[string]any

func NewContext() Context {
	return Context{}
}

func (ctx Context) AddToContext(key string, value any) {
	ctx[key] = value
}

func (ctx Context) GetFromContext(key string) any {
	return ctx[key]
}

func (ctx Context) AddMap(values map[string]any) {
	for name, value := range values {
		ctx[name] = value
	}
}
