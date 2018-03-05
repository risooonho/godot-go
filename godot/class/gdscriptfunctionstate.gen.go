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

//func NewGDScriptFunctionStateFromPointer(ptr gdnative.Pointer) GDScriptFunctionState {
func NewGDScriptFunctionStateFromPointer(ptr gdnative.Pointer) GDScriptFunctionState {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GDScriptFunctionState{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type GDScriptFunctionState struct {
	Reference
	owner gdnative.Object
}

func (o *GDScriptFunctionState) BaseClass() string {
	return "GDScriptFunctionState"
}

/*
        Undocumented
	Args: [], Returns: Variant
*/
func (o *GDScriptFunctionState) X_SignalCallback() gdnative.Variant {
	//log.Println("Calling GDScriptFunctionState.X_SignalCallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDScriptFunctionState", "_signal_callback")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{False true extended_check bool}], Returns: bool
*/
func (o *GDScriptFunctionState) IsValid(extendedCheck gdnative.Bool) gdnative.Bool {
	//log.Println("Calling GDScriptFunctionState.IsValid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(extendedCheck)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDScriptFunctionState", "is_valid")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{Null true arg Variant}], Returns: Variant
*/
func (o *GDScriptFunctionState) Resume(arg gdnative.Variant) gdnative.Variant {
	//log.Println("Calling GDScriptFunctionState.Resume()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVariant(arg)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDScriptFunctionState", "resume")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}
