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

//func NewWorldFromPointer(ptr gdnative.Pointer) World {
func NewWorldFromPointer(ptr gdnative.Pointer) World {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := World{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Class that has everything pertaining to a world. A physics space, a visual scenario and a sound space. Spatial nodes register their resources into the current world.
*/
type World struct {
	Resource
	owner gdnative.Object
}

func (o *World) BaseClass() string {
	return "World"
}

/*
        Undocumented
	Args: [], Returns: PhysicsDirectSpaceState
*/
func (o *World) GetDirectSpaceState() PhysicsDirectSpaceState {
	//log.Println("Calling World.GetDirectSpaceState()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World", "get_direct_space_state")

	// Call the parent method.
	// PhysicsDirectSpaceState
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewPhysicsDirectSpaceStateFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Environment
*/
func (o *World) GetEnvironment() Environment {
	//log.Println("Calling World.GetEnvironment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World", "get_environment")

	// Call the parent method.
	// Environment
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewEnvironmentFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Environment
*/
func (o *World) GetFallbackEnvironment() Environment {
	//log.Println("Calling World.GetFallbackEnvironment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World", "get_fallback_environment")

	// Call the parent method.
	// Environment
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewEnvironmentFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: RID
*/
func (o *World) GetScenario() gdnative.Rid {
	//log.Println("Calling World.GetScenario()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World", "get_scenario")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: RID
*/
func (o *World) GetSpace() gdnative.Rid {
	//log.Println("Calling World.GetSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World", "get_space")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false env Environment}], Returns: void
*/
func (o *World) SetEnvironment(env Environment) {
	//log.Println("Calling World.SetEnvironment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(env.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World", "set_environment")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false env Environment}], Returns: void
*/
func (o *World) SetFallbackEnvironment(env Environment) {
	//log.Println("Calling World.SetFallbackEnvironment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(env.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World", "set_fallback_environment")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
