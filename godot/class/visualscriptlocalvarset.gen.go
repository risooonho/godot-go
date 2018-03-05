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

//func NewVisualScriptLocalVarSetFromPointer(ptr gdnative.Pointer) VisualScriptLocalVarSet {
func NewVisualScriptLocalVarSetFromPointer(ptr gdnative.Pointer) VisualScriptLocalVarSet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptLocalVarSet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptLocalVarSet struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptLocalVarSet) BaseClass() string {
	return "VisualScriptLocalVarSet"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptLocalVarSet) GetVarName() gdnative.String {
	//log.Println("Calling VisualScriptLocalVarSet.GetVarName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptLocalVarSet", "get_var_name")

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
	Args: [], Returns: enum.Variant::Type
*/

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/
func (o *VisualScriptLocalVarSet) SetVarName(name gdnative.String) {
	//log.Println("Calling VisualScriptLocalVarSet.SetVarName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptLocalVarSet", "set_var_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type int}], Returns: void
*/
func (o *VisualScriptLocalVarSet) SetVarType(aType gdnative.Int) {
	//log.Println("Calling VisualScriptLocalVarSet.SetVarType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptLocalVarSet", "set_var_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
