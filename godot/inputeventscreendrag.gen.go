package godot

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

//func NewInputEventScreenDragFromPointer(ptr gdnative.Pointer) InputEventScreenDrag {
func newInputEventScreenDragFromPointer(ptr gdnative.Pointer) InputEventScreenDrag {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventScreenDrag{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Contains screen drag information. See [method Node._input].
*/
type InputEventScreenDrag struct {
	InputEvent
	owner gdnative.Object
}

func (o *InputEventScreenDrag) BaseClass() string {
	return "InputEventScreenDrag"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *InputEventScreenDrag) GetIndex() gdnative.Int {
	//log.Println("Calling InputEventScreenDrag.GetIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventScreenDrag", "get_index")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *InputEventScreenDrag) GetPosition() gdnative.Vector2 {
	//log.Println("Calling InputEventScreenDrag.GetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventScreenDrag", "get_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *InputEventScreenDrag) GetRelative() gdnative.Vector2 {
	//log.Println("Calling InputEventScreenDrag.GetRelative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventScreenDrag", "get_relative")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *InputEventScreenDrag) GetSpeed() gdnative.Vector2 {
	//log.Println("Calling InputEventScreenDrag.GetSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventScreenDrag", "get_speed")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false index int}], Returns: void
*/
func (o *InputEventScreenDrag) SetIndex(index gdnative.Int) {
	//log.Println("Calling InputEventScreenDrag.SetIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventScreenDrag", "set_index")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false position Vector2}], Returns: void
*/
func (o *InputEventScreenDrag) SetPosition(position gdnative.Vector2) {
	//log.Println("Calling InputEventScreenDrag.SetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventScreenDrag", "set_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false relative Vector2}], Returns: void
*/
func (o *InputEventScreenDrag) SetRelative(relative gdnative.Vector2) {
	//log.Println("Calling InputEventScreenDrag.SetRelative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(relative)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventScreenDrag", "set_relative")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false speed Vector2}], Returns: void
*/
func (o *InputEventScreenDrag) SetSpeed(speed gdnative.Vector2) {
	//log.Println("Calling InputEventScreenDrag.SetSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(speed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventScreenDrag", "set_speed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InputEventScreenDragImplementer is an interface that implements the methods
// of the InputEventScreenDrag class.
type InputEventScreenDragImplementer interface {
	InputEventImplementer
	GetIndex() gdnative.Int
	GetPosition() gdnative.Vector2
	GetRelative() gdnative.Vector2
	GetSpeed() gdnative.Vector2
	SetIndex(index gdnative.Int)
	SetPosition(position gdnative.Vector2)
	SetRelative(relative gdnative.Vector2)
	SetSpeed(speed gdnative.Vector2)
}
