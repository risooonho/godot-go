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

//func NewGeometryInstanceFromPointer(ptr gdnative.Pointer) GeometryInstance {
func NewGeometryInstanceFromPointer(ptr gdnative.Pointer) GeometryInstance {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GeometryInstance{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base node for geometry based visual instances. Shares some common functionality like visibility and custom materials.
*/
type GeometryInstance struct {
	VisualInstance
	owner gdnative.Object
}

func (o *GeometryInstance) BaseClass() string {
	return "GeometryInstance"
}

/*
        Undocumented
	Args: [], Returns: enum.GeometryInstance::ShadowCastingSetting
*/

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *GeometryInstance) GetExtraCullMargin() gdnative.Float {
	//log.Println("Calling GeometryInstance.GetExtraCullMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_extra_cull_margin")

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
	Args: [{ false flag int}], Returns: bool
*/
func (o *GeometryInstance) GetFlag(flag gdnative.Int) gdnative.Bool {
	//log.Println("Calling GeometryInstance.GetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_flag")

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
	Args: [], Returns: float
*/
func (o *GeometryInstance) GetLodMaxDistance() gdnative.Float {
	//log.Println("Calling GeometryInstance.GetLodMaxDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_lod_max_distance")

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
	Args: [], Returns: float
*/
func (o *GeometryInstance) GetLodMaxHysteresis() gdnative.Float {
	//log.Println("Calling GeometryInstance.GetLodMaxHysteresis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_lod_max_hysteresis")

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
	Args: [], Returns: float
*/
func (o *GeometryInstance) GetLodMinDistance() gdnative.Float {
	//log.Println("Calling GeometryInstance.GetLodMinDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_lod_min_distance")

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
	Args: [], Returns: float
*/
func (o *GeometryInstance) GetLodMinHysteresis() gdnative.Float {
	//log.Println("Calling GeometryInstance.GetLodMinHysteresis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_lod_min_hysteresis")

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
	Args: [], Returns: Material
*/
func (o *GeometryInstance) GetMaterialOverride() Material {
	//log.Println("Calling GeometryInstance.GetMaterialOverride()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_material_override")

	// Call the parent method.
	// Material
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewMaterialFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false shadow_casting_setting int}], Returns: void
*/
func (o *GeometryInstance) SetCastShadowsSetting(shadowCastingSetting gdnative.Int) {
	//log.Println("Calling GeometryInstance.SetCastShadowsSetting()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(shadowCastingSetting)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_cast_shadows_setting")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false margin float}], Returns: void
*/
func (o *GeometryInstance) SetExtraCullMargin(margin gdnative.Float) {
	//log.Println("Calling GeometryInstance.SetExtraCullMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_extra_cull_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flag int} { false value bool}], Returns: void
*/
func (o *GeometryInstance) SetFlag(flag gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling GeometryInstance.SetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_flag")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *GeometryInstance) SetLodMaxDistance(mode gdnative.Float) {
	//log.Println("Calling GeometryInstance.SetLodMaxDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_lod_max_distance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *GeometryInstance) SetLodMaxHysteresis(mode gdnative.Float) {
	//log.Println("Calling GeometryInstance.SetLodMaxHysteresis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_lod_max_hysteresis")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *GeometryInstance) SetLodMinDistance(mode gdnative.Float) {
	//log.Println("Calling GeometryInstance.SetLodMinDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_lod_min_distance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *GeometryInstance) SetLodMinHysteresis(mode gdnative.Float) {
	//log.Println("Calling GeometryInstance.SetLodMinHysteresis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_lod_min_hysteresis")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false material Material}], Returns: void
*/
func (o *GeometryInstance) SetMaterialOverride(material Material) {
	//log.Println("Calling GeometryInstance.SetMaterialOverride()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_material_override")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
