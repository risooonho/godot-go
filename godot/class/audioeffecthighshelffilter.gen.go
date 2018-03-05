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

//func NewAudioEffectHighShelfFilterFromPointer(ptr gdnative.Pointer) AudioEffectHighShelfFilter {
func NewAudioEffectHighShelfFilterFromPointer(ptr gdnative.Pointer) AudioEffectHighShelfFilter {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectHighShelfFilter{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AudioEffectHighShelfFilter struct {
	AudioEffectFilter
	owner gdnative.Object
}

func (o *AudioEffectHighShelfFilter) BaseClass() string {
	return "AudioEffectHighShelfFilter"
}
