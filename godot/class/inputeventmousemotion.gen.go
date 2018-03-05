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

//func NewInputEventMouseMotionFromPointer(ptr gdnative.Pointer) InputEventMouseMotion {
func NewInputEventMouseMotionFromPointer(ptr gdnative.Pointer) InputEventMouseMotion {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventMouseMotion{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Contains mouse motion information. Supports relative, absolute positions and speed. See [method Node._input].
*/
type InputEventMouseMotion struct {
	InputEventMouse
	owner gdnative.Object
}

func (o *InputEventMouseMotion) BaseClass() string {
	return "InputEventMouseMotion"
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *InputEventMouseMotion) GetRelative() gdnative.Vector2 {
	//log.Println("Calling InputEventMouseMotion.GetRelative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventMouseMotion", "get_relative")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *InputEventMouseMotion) GetSpeed() gdnative.Vector2 {
	//log.Println("Calling InputEventMouseMotion.GetSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventMouseMotion", "get_speed")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false relative Vector2}], Returns: void
*/
func (o *InputEventMouseMotion) SetRelative(relative gdnative.Vector2) {
	//log.Println("Calling InputEventMouseMotion.SetRelative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(relative)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventMouseMotion", "set_relative")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false speed Vector2}], Returns: void
*/
func (o *InputEventMouseMotion) SetSpeed(speed gdnative.Vector2) {
	//log.Println("Calling InputEventMouseMotion.SetSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(speed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventMouseMotion", "set_speed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
