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

//func NewPinJoint2DFromPointer(ptr gdnative.Pointer) PinJoint2D {
func NewPinJoint2DFromPointer(ptr gdnative.Pointer) PinJoint2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PinJoint2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Pin Joint for 2D Rigid Bodies. It pins two bodies (rigid or static) together.
*/
type PinJoint2D struct {
	Joint2D
	owner gdnative.Object
}

func (o *PinJoint2D) BaseClass() string {
	return "PinJoint2D"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PinJoint2D) GetSoftness() gdnative.Float {
	//log.Println("Calling PinJoint2D.GetSoftness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PinJoint2D", "get_softness")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false softness float}], Returns: void
*/
func (o *PinJoint2D) SetSoftness(softness gdnative.Float) {
	//log.Println("Calling PinJoint2D.SetSoftness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(softness)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PinJoint2D", "set_softness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
