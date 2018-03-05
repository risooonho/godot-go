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

//func NewAudioStreamPlaybackFromPointer(ptr gdnative.Pointer) AudioStreamPlayback {
func NewAudioStreamPlaybackFromPointer(ptr gdnative.Pointer) AudioStreamPlayback {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamPlayback{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Can play, loop, pause a scroll through Audio. See [AudioStream] and [AudioStreamOGGVorbis] for usage.
*/
type AudioStreamPlayback struct {
	Reference
	owner gdnative.Object
}

func (o *AudioStreamPlayback) BaseClass() string {
	return "AudioStreamPlayback"
}
