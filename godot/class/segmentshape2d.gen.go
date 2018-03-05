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

//func NewSegmentShape2DFromPointer(ptr gdnative.Pointer) SegmentShape2D {
func NewSegmentShape2DFromPointer(ptr gdnative.Pointer) SegmentShape2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SegmentShape2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Segment shape for 2D collisions. Consists of two points, [code]a[/code] and [code]b[/code].
*/
type SegmentShape2D struct {
	Shape2D
	owner gdnative.Object
}

func (o *SegmentShape2D) BaseClass() string {
	return "SegmentShape2D"
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *SegmentShape2D) GetA() gdnative.Vector2 {
	//log.Println("Calling SegmentShape2D.GetA()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SegmentShape2D", "get_a")

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
func (o *SegmentShape2D) GetB() gdnative.Vector2 {
	//log.Println("Calling SegmentShape2D.GetB()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SegmentShape2D", "get_b")

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
	Args: [{ false a Vector2}], Returns: void
*/
func (o *SegmentShape2D) SetA(a gdnative.Vector2) {
	//log.Println("Calling SegmentShape2D.SetA()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(a)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SegmentShape2D", "set_a")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false b Vector2}], Returns: void
*/
func (o *SegmentShape2D) SetB(b gdnative.Vector2) {
	//log.Println("Calling SegmentShape2D.SetB()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(b)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SegmentShape2D", "set_b")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
