package interpreter

import "github.com/JosephNaberhaus/agnostic/code"

type Value struct {
	Value any
}

type List []Value

func (l *List) Get(i int64) Value {
	return (*l)[int(i)]
}

func (l *List) Set(i int64, value Value) {
	(*l)[i] = value
}

func (l *List) Pop() Value {
	value := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return value
}

func (l *List) Push(value Value) {
	*l = append(*l, value)
}

func (l *List) Length() Value {
	return Value{Value: int64(len(*l))}
}

type Map map[Value]Value

func (m *Map) Get(key Value) Value {
	return (*m)[key]
}

func (m *Map) Put(key Value, value Value) {
	(*m)[key] = value
}

type Set map[Value]struct{}

func (s *Set) Contains(value Value) bool {
	_, ok := (*s)[value]
	return ok
}

func (s *Set) Add(value Value) {
	(*s)[value] = struct{}{}
}

type Model struct {
	Properties map[string]Value
	Methods    map[string]*code.FunctionDef
}

func NewModel(definition *code.ModelDef) Value {
	value := new(Model)

	for _, field := range definition.Fields {
		value.Properties[field.Name] = Value{}
	}

	for _, method := range definition.Methods {
		value.Methods[method.Name] = method
	}

	return Value{Value: value}
}
