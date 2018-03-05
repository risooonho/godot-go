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

//func NewTextureFromPointer(ptr gdnative.Pointer) Texture {
func NewTextureFromPointer(ptr gdnative.Pointer) Texture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Texture{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A texture works by registering an image in the video hardware, which then can be used in 3D models or 2D [Sprite] or GUI [Control].
*/
type Texture struct {
	Resource
	owner gdnative.Object
}

func (o *Texture) BaseClass() string {
	return "Texture"
}

/*

	Args: [{ false canvas_item RID} { false position Vector2} {1,1,1,1 true modulate Color} {False true transpose bool} {Null true normal_map Texture}], Returns: void
*/
func (o *Texture) Draw(canvasItem gdnative.Rid, position gdnative.Vector2, modulate gdnative.Color, transpose gdnative.Bool, normalMap Texture) {
	//log.Println("Calling Texture.Draw()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromRid(canvasItem)
	ptrArguments[1] = gdnative.NewPointerFromVector2(position)
	ptrArguments[2] = gdnative.NewPointerFromColor(modulate)
	ptrArguments[3] = gdnative.NewPointerFromBool(transpose)
	ptrArguments[4] = gdnative.NewPointerFromObject(normalMap.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "draw")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false canvas_item RID} { false rect Rect2} { false tile bool} {1,1,1,1 true modulate Color} {False true transpose bool} {Null true normal_map Texture}], Returns: void
*/
func (o *Texture) DrawRect(canvasItem gdnative.Rid, rect gdnative.Rect2, tile gdnative.Bool, modulate gdnative.Color, transpose gdnative.Bool, normalMap Texture) {
	//log.Println("Calling Texture.DrawRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 6, 6)
	ptrArguments[0] = gdnative.NewPointerFromRid(canvasItem)
	ptrArguments[1] = gdnative.NewPointerFromRect2(rect)
	ptrArguments[2] = gdnative.NewPointerFromBool(tile)
	ptrArguments[3] = gdnative.NewPointerFromColor(modulate)
	ptrArguments[4] = gdnative.NewPointerFromBool(transpose)
	ptrArguments[5] = gdnative.NewPointerFromObject(normalMap.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "draw_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false canvas_item RID} { false rect Rect2} { false src_rect Rect2} {1,1,1,1 true modulate Color} {False true transpose bool} {Null true normal_map Texture} {True true clip_uv bool}], Returns: void
*/
func (o *Texture) DrawRectRegion(canvasItem gdnative.Rid, rect gdnative.Rect2, srcRect gdnative.Rect2, modulate gdnative.Color, transpose gdnative.Bool, normalMap Texture, clipUv gdnative.Bool) {
	//log.Println("Calling Texture.DrawRectRegion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 7, 7)
	ptrArguments[0] = gdnative.NewPointerFromRid(canvasItem)
	ptrArguments[1] = gdnative.NewPointerFromRect2(rect)
	ptrArguments[2] = gdnative.NewPointerFromRect2(srcRect)
	ptrArguments[3] = gdnative.NewPointerFromColor(modulate)
	ptrArguments[4] = gdnative.NewPointerFromBool(transpose)
	ptrArguments[5] = gdnative.NewPointerFromObject(normalMap.GetBaseObject())
	ptrArguments[6] = gdnative.NewPointerFromBool(clipUv)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "draw_rect_region")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: Image
*/
func (o *Texture) GetData() Image {
	//log.Println("Calling Texture.GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "get_data")

	// Call the parent method.
	// Image
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewImageFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Texture) GetFlags() gdnative.Int {
	//log.Println("Calling Texture.GetFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "get_flags")

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
        Return the texture height.
	Args: [], Returns: int
*/
func (o *Texture) GetHeight() gdnative.Int {
	//log.Println("Calling Texture.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "get_height")

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
        Return the texture size.
	Args: [], Returns: Vector2
*/
func (o *Texture) GetSize() gdnative.Vector2 {
	//log.Println("Calling Texture.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "get_size")

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
        Return the texture width.
	Args: [], Returns: int
*/
func (o *Texture) GetWidth() gdnative.Int {
	//log.Println("Calling Texture.GetWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "get_width")

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

	Args: [], Returns: bool
*/
func (o *Texture) HasAlpha() gdnative.Bool {
	//log.Println("Calling Texture.HasAlpha()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "has_alpha")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false flags int}], Returns: void
*/
func (o *Texture) SetFlags(flags gdnative.Int) {
	//log.Println("Calling Texture.SetFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Texture", "set_flags")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
