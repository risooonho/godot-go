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

//func NewSliderJointFromPointer(ptr gdnative.Pointer) SliderJoint {
func NewSliderJointFromPointer(ptr gdnative.Pointer) SliderJoint {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SliderJoint{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Slides across the x-axis of the [Pivot] object.
*/
type SliderJoint struct {
	Joint
	owner gdnative.Object
}

func (o *SliderJoint) BaseClass() string {
	return "SliderJoint"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *SliderJoint) X_GetLowerLimitAngular() gdnative.Float {
	//log.Println("Calling SliderJoint.X_GetLowerLimitAngular()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SliderJoint", "_get_lower_limit_angular")

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
	Args: [], Returns: float
*/
func (o *SliderJoint) X_GetUpperLimitAngular() gdnative.Float {
	//log.Println("Calling SliderJoint.X_GetUpperLimitAngular()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SliderJoint", "_get_upper_limit_angular")

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
	Args: [{ false lower_limit_angular float}], Returns: void
*/
func (o *SliderJoint) X_SetLowerLimitAngular(lowerLimitAngular gdnative.Float) {
	//log.Println("Calling SliderJoint.X_SetLowerLimitAngular()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(lowerLimitAngular)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SliderJoint", "_set_lower_limit_angular")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false upper_limit_angular float}], Returns: void
*/
func (o *SliderJoint) X_SetUpperLimitAngular(upperLimitAngular gdnative.Float) {
	//log.Println("Calling SliderJoint.X_SetUpperLimitAngular()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(upperLimitAngular)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SliderJoint", "_set_upper_limit_angular")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int}], Returns: float
*/
func (o *SliderJoint) GetParam(param gdnative.Int) gdnative.Float {
	//log.Println("Calling SliderJoint.GetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SliderJoint", "get_param")

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
	Args: [{ false param int} { false value float}], Returns: void
*/
func (o *SliderJoint) SetParam(param gdnative.Int, value gdnative.Float) {
	//log.Println("Calling SliderJoint.SetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromFloat(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SliderJoint", "set_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
