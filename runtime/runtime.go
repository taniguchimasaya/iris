package runtime

import (
	e "github.com/ta2gch/iris/runtime/environment"
	i "github.com/ta2gch/iris/runtime/ilos/instance"
)

func Init() {
	e.TopLevel.Macro.Define(i.NewSymbol("LAMBDA"), i.NewFunction(lambda))
	e.TopLevel.Macro.Define(i.NewSymbol("QUOTE"), i.NewFunction(quote))
	e.TopLevel.Function.Define(i.NewSymbol("THROW"), i.NewFunction(throw))
	e.TopLevel.Macro.Define(i.NewSymbol("CATCH"), i.NewFunction(catch))
}
