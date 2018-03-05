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

//func NewCanvasModulateFromPointer(ptr gdnative.Pointer) CanvasModulate {
func NewCanvasModulateFromPointer(ptr gdnative.Pointer) CanvasModulate {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CanvasModulate{}
	obj.SetBaseObject(owner)

	return obj
}

/*
[code]CanvasModulate[/code] tints the canvas elements using its assigned [code]color[/code].
*/
type CanvasModulate struct {
	Node2D
	owner gdnative.Object
}

func (o *CanvasModulate) BaseClass() string {
	return "CanvasModulate"
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *CanvasModulate) GetColor() gdnative.Color {
	//log.Println("Calling CanvasModulate.GetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CanvasModulate", "get_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *CanvasModulate) SetColor(color gdnative.Color) {
	//log.Println("Calling CanvasModulate.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CanvasModulate", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
