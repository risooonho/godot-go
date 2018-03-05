package class

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewIP_UnixFromPointer(ptr gdnative.Pointer) IP_Unix {
func NewIP_UnixFromPointer(ptr gdnative.Pointer) IP_Unix {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := IP_Unix{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Unix-specific implementation of IP support functions. See [IP].
*/
type IP_Unix struct {
	ip
	owner gdnative.Object
}

func (o *IP_Unix) BaseClass() string {
	return "IP_Unix"
}
