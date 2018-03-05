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

//func NewAudioEffectLowPassFilterFromPointer(ptr gdnative.Pointer) AudioEffectLowPassFilter {
func NewAudioEffectLowPassFilterFromPointer(ptr gdnative.Pointer) AudioEffectLowPassFilter {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectLowPassFilter{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Cuts frequencies higher than the [member cutoff_hz] and allows lower frequencies to pass.
*/
type AudioEffectLowPassFilter struct {
	AudioEffectFilter
	owner gdnative.Object
}

func (o *AudioEffectLowPassFilter) BaseClass() string {
	return "AudioEffectLowPassFilter"
}
