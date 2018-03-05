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

//func NewSkyFromPointer(ptr gdnative.Pointer) Sky {
func NewSkyFromPointer(ptr gdnative.Pointer) Sky {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Sky{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The base class for [PanoramaSky] and [ProceduralSky].
*/
type Sky struct {
	Resource
	owner gdnative.Object
}

func (o *Sky) BaseClass() string {
	return "Sky"
}

/*
        Undocumented
	Args: [], Returns: enum.Sky::RadianceSize
*/

/*
        Undocumented
	Args: [{ false size int}], Returns: void
*/
func (o *Sky) SetRadianceSize(size gdnative.Int) {
	//log.Println("Calling Sky.SetRadianceSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sky", "set_radiance_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
