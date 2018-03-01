package kinematiccollision

import (
	"log"
	"reflect"

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

/*
Contains collision data for KinematicBody collisions. When a [KinematicBody] is moved using [method KinematicBody.move_and_collide], it stops if it detects a collision with another body. If a collision is detected, a KinematicCollision object is returned. This object contains information about the collision, including the colliding object, the remaining motion, and the collision position. This information can be used to calculate a collision response.
*/
type KinematicCollision struct {
	Reference
}

func (o *KinematicCollision) BaseClass() string {
	return "KinematicCollision"
}

/*
   Undocumented
*/
func (o *KinematicCollision) GetCollider() *Object {
	log.Println("Calling KinematicCollision.GetCollider()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetColliderId() gdnative.Int {
	log.Println("Calling KinematicCollision.GetColliderId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider_id", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetColliderMetadata() *Variant {
	log.Println("Calling KinematicCollision.GetColliderMetadata()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider_metadata", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetColliderShape() *Object {
	log.Println("Calling KinematicCollision.GetColliderShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider_shape", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetColliderShapeIndex() gdnative.Int {
	log.Println("Calling KinematicCollision.GetColliderShapeIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider_shape_index", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetColliderVelocity() *Vector3 {
	log.Println("Calling KinematicCollision.GetColliderVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider_velocity", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetLocalShape() *Object {
	log.Println("Calling KinematicCollision.GetLocalShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_local_shape", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetNormal() *Vector3 {
	log.Println("Calling KinematicCollision.GetNormal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_normal", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetPosition() *Vector3 {
	log.Println("Calling KinematicCollision.GetPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_position", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetRemainder() *Vector3 {
	log.Println("Calling KinematicCollision.GetRemainder()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_remainder", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *KinematicCollision) GetTravel() *Vector3 {
	log.Println("Calling KinematicCollision.GetTravel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_travel", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   KinematicCollisionImplementer is an interface for KinematicCollision objects.
*/
type KinematicCollisionImplementer interface {
	Class
}