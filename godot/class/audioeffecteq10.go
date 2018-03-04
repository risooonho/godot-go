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

//func NewAudioEffectEQ10FromPointer(ptr gdnative.Pointer) AudioEffectEQ10 {
func NewAudioEffectEQ10FromPointer(ptr gdnative.Pointer) AudioEffectEQ10 {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectEQ10{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Frequency bands : Band 1 : 31 Hz Band 2 : 62 Hz Band 3 : 125 Hz Band 4 : 250 Hz Band 5 : 500 Hz Band 6 : 1000 Hz Band 7 : 2000 Hz Band 8 : 4000 Hz Band 9 : 8000 Hz Band 10 : 16000 Hz See also [AudioEffectEQ], [AudioEffectEQ6], [AudioEffectEQ21].
*/
type AudioEffectEQ10 struct {
	AudioEffectEQ
	owner gdnative.Object
}

func (o *AudioEffectEQ10) BaseClass() string {
	return "AudioEffectEQ10"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioEffectEQ10) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *AudioEffectEQ10) GetBaseObject() gdnative.Object {
	return o.owner
}