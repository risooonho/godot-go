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

//func NewAudioEffectPannerFromPointer(ptr gdnative.Pointer) AudioEffectPanner {
func NewAudioEffectPannerFromPointer(ptr gdnative.Pointer) AudioEffectPanner {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectPanner{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Determines how much of an audio signal is sent to the left and right buses.
*/
type AudioEffectPanner struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectPanner) BaseClass() string {
	return "AudioEffectPanner"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectPanner) GetPan() gdnative.Float {
	//log.Println("Calling AudioEffectPanner.GetPan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPanner", "get_pan")

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
	Args: [{ false cpanume float}], Returns: void
*/
func (o *AudioEffectPanner) SetPan(cpanume gdnative.Float) {
	//log.Println("Calling AudioEffectPanner.SetPan()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(cpanume)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPanner", "set_pan")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
