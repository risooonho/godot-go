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

//func NewHSliderFromPointer(ptr gdnative.Pointer) HSlider {
func NewHSliderFromPointer(ptr gdnative.Pointer) HSlider {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HSlider{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Horizontal slider. See [Slider]. This one goes from left (min) to right (max).
*/
type HSlider struct {
	Slider
	owner gdnative.Object
}

func (o *HSlider) BaseClass() string {
	return "HSlider"
}
