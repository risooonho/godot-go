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

//func NewVisualScriptOperatorFromPointer(ptr gdnative.Pointer) VisualScriptOperator {
func NewVisualScriptOperatorFromPointer(ptr gdnative.Pointer) VisualScriptOperator {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptOperator{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptOperator struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptOperator) BaseClass() string {
	return "VisualScriptOperator"
}

/*
        Undocumented
	Args: [], Returns: enum.Variant::Operator
*/

/*
        Undocumented
	Args: [], Returns: enum.Variant::Type
*/

/*
        Undocumented
	Args: [{ false op int}], Returns: void
*/
func (o *VisualScriptOperator) SetOperator(op gdnative.Int) {
	//log.Println("Calling VisualScriptOperator.SetOperator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(op)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptOperator", "set_operator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type int}], Returns: void
*/
func (o *VisualScriptOperator) SetTyped(aType gdnative.Int) {
	//log.Println("Calling VisualScriptOperator.SetTyped()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptOperator", "set_typed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
