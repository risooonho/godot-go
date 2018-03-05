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

//func NewPhysics2DTestMotionResultFromPointer(ptr gdnative.Pointer) Physics2DTestMotionResult {
func NewPhysics2DTestMotionResultFromPointer(ptr gdnative.Pointer) Physics2DTestMotionResult {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Physics2DTestMotionResult{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type Physics2DTestMotionResult struct {
	Reference
	owner gdnative.Object
}

func (o *Physics2DTestMotionResult) BaseClass() string {
	return "Physics2DTestMotionResult"
}

/*
        Undocumented
	Args: [], Returns: Object
*/
func (o *Physics2DTestMotionResult) GetCollider() Object {
	//log.Println("Calling Physics2DTestMotionResult.GetCollider()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_collider")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewObjectFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Physics2DTestMotionResult) GetColliderId() gdnative.Int {
	//log.Println("Calling Physics2DTestMotionResult.GetColliderId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_collider_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: RID
*/
func (o *Physics2DTestMotionResult) GetColliderRid() gdnative.Rid {
	//log.Println("Calling Physics2DTestMotionResult.GetColliderRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_collider_rid")

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
	Args: [], Returns: int
*/
func (o *Physics2DTestMotionResult) GetColliderShape() gdnative.Int {
	//log.Println("Calling Physics2DTestMotionResult.GetColliderShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_collider_shape")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *Physics2DTestMotionResult) GetColliderVelocity() gdnative.Vector2 {
	//log.Println("Calling Physics2DTestMotionResult.GetColliderVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_collider_velocity")

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
func (o *Physics2DTestMotionResult) GetCollisionNormal() gdnative.Vector2 {
	//log.Println("Calling Physics2DTestMotionResult.GetCollisionNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_collision_normal")

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
func (o *Physics2DTestMotionResult) GetCollisionPoint() gdnative.Vector2 {
	//log.Println("Calling Physics2DTestMotionResult.GetCollisionPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_collision_point")

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
func (o *Physics2DTestMotionResult) GetMotion() gdnative.Vector2 {
	//log.Println("Calling Physics2DTestMotionResult.GetMotion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_motion")

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
func (o *Physics2DTestMotionResult) GetMotionRemainder() gdnative.Vector2 {
	//log.Println("Calling Physics2DTestMotionResult.GetMotionRemainder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DTestMotionResult", "get_motion_remainder")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}
