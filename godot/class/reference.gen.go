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

//func NewReferenceFromPointer(ptr gdnative.Pointer) Reference {
func NewReferenceFromPointer(ptr gdnative.Pointer) Reference {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Reference{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base class for anything that keeps a reference count. Resource and many other helper objects inherit this. References keep an internal reference counter so they are only released when no longer in use.
*/
type Reference struct {
	Object
	owner gdnative.Object
}

func (o *Reference) BaseClass() string {
	return "Reference"
}

/*

	Args: [], Returns: bool
*/
func (o *Reference) InitRef() gdnative.Bool {
	//log.Println("Calling Reference.InitRef()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Reference", "init_ref")

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
        Increase the internal reference counter. Use this only if you really know what you are doing.
	Args: [], Returns: bool
*/
func (o *Reference) Reference() gdnative.Bool {
	//log.Println("Calling Reference.Reference()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Reference", "reference")

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
        Decrease the internal reference counter. Use this only if you really know what you are doing.
	Args: [], Returns: bool
*/
func (o *Reference) Unreference() gdnative.Bool {
	//log.Println("Calling Reference.Unreference()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Reference", "unreference")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}
