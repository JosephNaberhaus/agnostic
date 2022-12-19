package interpreter

import "github.com/JosephNaberhaus/agnostic/code"

type callable struct {
	model    *Model
	function *code.FunctionDef
}

type callableMapper struct {
	runtime runtime
}

func (c callableMapper) MapFunction(original *code.Function) callable {
	model, function := c.runtime.functionByName(original.Name)
	return callable{
		function: function,
		model:    model,
	}
}

func (c callableMapper) MapFunctionProperty(original *code.FunctionProperty) callable {
	ofModel := code.MapValueNoError[Value](original.Of, &valueMapper{runtime: c.runtime}).Value.(*Model)
	return callable{
		model:    ofModel,
		function: ofModel.Methods[original.Name],
	}
}
