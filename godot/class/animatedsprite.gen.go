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

//func NewAnimatedSpriteFromPointer(ptr gdnative.Pointer) AnimatedSprite {
func NewAnimatedSpriteFromPointer(ptr gdnative.Pointer) AnimatedSprite {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimatedSprite{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Animations are created using a [SpriteFrames] resource, which can be configured in the editor via the SpriteFrames panel.
*/
type AnimatedSprite struct {
	Node2D
	owner gdnative.Object
}

func (o *AnimatedSprite) BaseClass() string {
	return "AnimatedSprite"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AnimatedSprite) X_IsPlaying() gdnative.Bool {
	//log.Println("Calling AnimatedSprite.X_IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "_is_playing")

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
	Args: [], Returns: void
*/
func (o *AnimatedSprite) X_ResChanged() {
	//log.Println("Calling AnimatedSprite.X_ResChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "_res_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false playing bool}], Returns: void
*/
func (o *AnimatedSprite) X_SetPlaying(playing gdnative.Bool) {
	//log.Println("Calling AnimatedSprite.X_SetPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(playing)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "_set_playing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimatedSprite) GetAnimation() gdnative.String {
	//log.Println("Calling AnimatedSprite.GetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "get_animation")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *AnimatedSprite) GetFrame() gdnative.Int {
	//log.Println("Calling AnimatedSprite.GetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "get_frame")

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
func (o *AnimatedSprite) GetOffset() gdnative.Vector2 {
	//log.Println("Calling AnimatedSprite.GetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "get_offset")

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
	Args: [], Returns: SpriteFrames
*/
func (o *AnimatedSprite) GetSpriteFrames() SpriteFrames {
	//log.Println("Calling AnimatedSprite.GetSpriteFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "get_sprite_frames")

	// Call the parent method.
	// SpriteFrames
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewSpriteFramesFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AnimatedSprite) IsCentered() gdnative.Bool {
	//log.Println("Calling AnimatedSprite.IsCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "is_centered")

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
	Args: [], Returns: bool
*/
func (o *AnimatedSprite) IsFlippedH() gdnative.Bool {
	//log.Println("Calling AnimatedSprite.IsFlippedH()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "is_flipped_h")

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
	Args: [], Returns: bool
*/
func (o *AnimatedSprite) IsFlippedV() gdnative.Bool {
	//log.Println("Calling AnimatedSprite.IsFlippedV()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "is_flipped_v")

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
        Return true if an animation if currently being played.
	Args: [], Returns: bool
*/
func (o *AnimatedSprite) IsPlaying() gdnative.Bool {
	//log.Println("Calling AnimatedSprite.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "is_playing")

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
        Play the animation set in parameter. If no parameter is provided, the current animation is played.
	Args: [{ true anim String}], Returns: void
*/
func (o *AnimatedSprite) Play(anim gdnative.String) {
	//log.Println("Calling AnimatedSprite.Play()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "play")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false animation String}], Returns: void
*/
func (o *AnimatedSprite) SetAnimation(animation gdnative.String) {
	//log.Println("Calling AnimatedSprite.SetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(animation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "set_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false centered bool}], Returns: void
*/
func (o *AnimatedSprite) SetCentered(centered gdnative.Bool) {
	//log.Println("Calling AnimatedSprite.SetCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(centered)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "set_centered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flip_h bool}], Returns: void
*/
func (o *AnimatedSprite) SetFlipH(flipH gdnative.Bool) {
	//log.Println("Calling AnimatedSprite.SetFlipH()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(flipH)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "set_flip_h")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flip_v bool}], Returns: void
*/
func (o *AnimatedSprite) SetFlipV(flipV gdnative.Bool) {
	//log.Println("Calling AnimatedSprite.SetFlipV()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(flipV)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "set_flip_v")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false frame int}], Returns: void
*/
func (o *AnimatedSprite) SetFrame(frame gdnative.Int) {
	//log.Println("Calling AnimatedSprite.SetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "set_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *AnimatedSprite) SetOffset(offset gdnative.Vector2) {
	//log.Println("Calling AnimatedSprite.SetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "set_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false sprite_frames SpriteFrames}], Returns: void
*/
func (o *AnimatedSprite) SetSpriteFrames(spriteFrames SpriteFrames) {
	//log.Println("Calling AnimatedSprite.SetSpriteFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(spriteFrames.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "set_sprite_frames")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stop the current animation (does not reset the frame counter).
	Args: [], Returns: void
*/
func (o *AnimatedSprite) Stop() {
	//log.Println("Calling AnimatedSprite.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
