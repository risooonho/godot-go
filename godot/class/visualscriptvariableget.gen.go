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

//func NewVisualScriptVariableGetFromPointer(ptr gdnative.Pointer) VisualScriptVariableGet {
func NewVisualScriptVariableGetFromPointer(ptr gdnative.Pointer) VisualScriptVariableGet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptVariableGet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptVariableGet struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptVariableGet) BaseClass() string {
	return "VisualScriptVariableGet"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptVariableGet) GetVariable() gdnative.String {
	//log.Println("Calling VisualScriptVariableGet.GetVariable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptVariableGet", "get_variable")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/
func (o *VisualScriptVariableGet) SetVariable(name gdnative.String) {
	//log.Println("Calling VisualScriptVariableGet.SetVariable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptVariableGet", "set_variable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
