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

//func NewCapsuleMeshFromPointer(ptr gdnative.Pointer) CapsuleMesh {
func NewCapsuleMeshFromPointer(ptr gdnative.Pointer) CapsuleMesh {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CapsuleMesh{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Class representing a capsule-shaped [PrimitiveMesh].
*/
type CapsuleMesh struct {
	PrimitiveMesh
	owner gdnative.Object
}

func (o *CapsuleMesh) BaseClass() string {
	return "CapsuleMesh"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *CapsuleMesh) GetMidHeight() gdnative.Float {
	//log.Println("Calling CapsuleMesh.GetMidHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleMesh", "get_mid_height")

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
	Args: [], Returns: int
*/
func (o *CapsuleMesh) GetRadialSegments() gdnative.Int {
	//log.Println("Calling CapsuleMesh.GetRadialSegments()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleMesh", "get_radial_segments")

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
	Args: [], Returns: float
*/
func (o *CapsuleMesh) GetRadius() gdnative.Float {
	//log.Println("Calling CapsuleMesh.GetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleMesh", "get_radius")

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
	Args: [], Returns: int
*/
func (o *CapsuleMesh) GetRings() gdnative.Int {
	//log.Println("Calling CapsuleMesh.GetRings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleMesh", "get_rings")

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
	Args: [{ false mid_height float}], Returns: void
*/
func (o *CapsuleMesh) SetMidHeight(midHeight gdnative.Float) {
	//log.Println("Calling CapsuleMesh.SetMidHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(midHeight)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleMesh", "set_mid_height")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false segments int}], Returns: void
*/
func (o *CapsuleMesh) SetRadialSegments(segments gdnative.Int) {
	//log.Println("Calling CapsuleMesh.SetRadialSegments()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(segments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleMesh", "set_radial_segments")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radius float}], Returns: void
*/
func (o *CapsuleMesh) SetRadius(radius gdnative.Float) {
	//log.Println("Calling CapsuleMesh.SetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleMesh", "set_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rings int}], Returns: void
*/
func (o *CapsuleMesh) SetRings(rings gdnative.Int) {
	//log.Println("Calling CapsuleMesh.SetRings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(rings)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleMesh", "set_rings")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
