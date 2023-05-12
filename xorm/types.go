package xorm

import (
	"reflect"

	"zhuochengs/goxorm/core"
)

var (
	ptrPkType = reflect.TypeOf(&core.PK{})
	pkType    = reflect.TypeOf(core.PK{})
)
