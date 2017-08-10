package instance

import (
	"github.com/ta2gch/iris/runtime/ilos"
	"github.com/ta2gch/iris/runtime/ilos/class"
)

//
// Symbol
//

type Symbol string

func (Symbol) Class() ilos.Class {
	return class.Symbol
}

func (i Symbol) GetSlotValue(key ilos.Instance, _ ilos.Class) (ilos.Instance, bool) {
	return nil, false
}

func (i Symbol) SetSlotValue(key ilos.Instance, value ilos.Instance, _ ilos.Class) bool {
	return false
}

func (i Symbol) String() string {
	return string(i)
}
