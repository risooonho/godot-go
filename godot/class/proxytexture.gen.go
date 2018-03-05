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

//func NewProxyTextureFromPointer(ptr gdnative.Pointer) ProxyTexture {
func NewProxyTextureFromPointer(ptr gdnative.Pointer) ProxyTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ProxyTexture{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ProxyTexture struct {
	Texture
	owner gdnative.Object
}

func (o *ProxyTexture) BaseClass() string {
	return "ProxyTexture"
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *ProxyTexture) GetBase() Texture {
	//log.Println("Calling ProxyTexture.GetBase()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProxyTexture", "get_base")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewTextureFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false base Texture}], Returns: void
*/
func (o *ProxyTexture) SetBase(base Texture) {
	//log.Println("Calling ProxyTexture.SetBase()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(base.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProxyTexture", "set_base")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
