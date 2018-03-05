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

//func NewCameraFromPointer(ptr gdnative.Pointer) Camera {
func NewCameraFromPointer(ptr gdnative.Pointer) Camera {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Camera{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Camera is a special node that displays what is visible from its current location. Cameras register themselves in the nearest [Viewport] node (when ascending the tree). Only one camera can be active per viewport. If no viewport is available ascending the tree, the Camera will register in the global viewport. In other words, a Camera just provides [i]3D[/i] display capabilities to a [Viewport], and, without one, a scene registered in that [Viewport] (or higher viewports) can't be displayed.
*/
type Camera struct {
	Spatial
	owner gdnative.Object
}

func (o *Camera) BaseClass() string {
	return "Camera"
}

/*
        If this is the current Camera, remove it from being current. If it is inside the node tree, request to make the next Camera current, if any.
	Args: [], Returns: void
*/
func (o *Camera) ClearCurrent() {
	//log.Println("Calling Camera.ClearCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "clear_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Get the camera transform. Subclassed cameras (such as CharacterCamera) may provide different transforms than the [Node] transform.
	Args: [], Returns: Transform
*/
func (o *Camera) GetCameraTransform() gdnative.Transform {
	//log.Println("Calling Camera.GetCameraTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_camera_transform")

	// Call the parent method.
	// Transform
	retPtr := gdnative.NewEmptyTransform()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransformFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Camera) GetCullMask() gdnative.Int {
	//log.Println("Calling Camera.GetCullMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_cull_mask")

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
	Args: [], Returns: enum.Camera::DopplerTracking
*/

/*
        Undocumented
	Args: [], Returns: Environment
*/
func (o *Camera) GetEnvironment() Environment {
	//log.Println("Calling Camera.GetEnvironment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_environment")

	// Call the parent method.
	// Environment
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewEnvironmentFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Camera) GetFov() gdnative.Float {
	//log.Println("Calling Camera.GetFov()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_fov")

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
func (o *Camera) GetHOffset() gdnative.Float {
	//log.Println("Calling Camera.GetHOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_h_offset")

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
	Args: [], Returns: enum.Camera::KeepAspect
*/

/*
        Undocumented
	Args: [], Returns: enum.Camera::Projection
*/

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Camera) GetSize() gdnative.Float {
	//log.Println("Calling Camera.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_size")

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
func (o *Camera) GetVOffset() gdnative.Float {
	//log.Println("Calling Camera.GetVOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_v_offset")

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
func (o *Camera) GetZfar() gdnative.Float {
	//log.Println("Calling Camera.GetZfar()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_zfar")

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
func (o *Camera) GetZnear() gdnative.Float {
	//log.Println("Calling Camera.GetZnear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "get_znear")

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
	Args: [], Returns: bool
*/
func (o *Camera) IsCurrent() gdnative.Bool {
	//log.Println("Calling Camera.IsCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "is_current")

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
        Returns [code]true[/code] if the given position is behind the Camera.
	Args: [{ false world_point Vector3}], Returns: bool
*/
func (o *Camera) IsPositionBehind(worldPoint gdnative.Vector3) gdnative.Bool {
	//log.Println("Calling Camera.IsPositionBehind()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(worldPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "is_position_behind")

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
        Make this camera the current Camera for the [Viewport] (see class description). If the Camera Node is outside the scene tree, it will attempt to become current once it's added.
	Args: [], Returns: void
*/
func (o *Camera) MakeCurrent() {
	//log.Println("Calling Camera.MakeCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "make_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns a normal vector from the screen point location directed along the camera. Orthogonal cameras are normalized. Perspective cameras account for perspective, screen width/height, etc.
	Args: [{ false screen_point Vector2}], Returns: Vector3
*/
func (o *Camera) ProjectLocalRayNormal(screenPoint gdnative.Vector2) gdnative.Vector3 {
	//log.Println("Calling Camera.ProjectLocalRayNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(screenPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "project_local_ray_normal")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns how a 2D coordinate in the Viewport rectangle maps to a 3D point in worldspace.
	Args: [{ false screen_point Vector2}], Returns: Vector3
*/
func (o *Camera) ProjectPosition(screenPoint gdnative.Vector2) gdnative.Vector3 {
	//log.Println("Calling Camera.ProjectPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(screenPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "project_position")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns a normal vector in worldspace, that is the result of projecting a point on the [Viewport] rectangle by the camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
	Args: [{ false screen_point Vector2}], Returns: Vector3
*/
func (o *Camera) ProjectRayNormal(screenPoint gdnative.Vector2) gdnative.Vector3 {
	//log.Println("Calling Camera.ProjectRayNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(screenPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "project_ray_normal")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns a 3D position in worldspace, that is the result of projecting a point on the [Viewport] rectangle by the camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
	Args: [{ false screen_point Vector2}], Returns: Vector3
*/
func (o *Camera) ProjectRayOrigin(screenPoint gdnative.Vector2) gdnative.Vector3 {
	//log.Println("Calling Camera.ProjectRayOrigin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(screenPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "project_ray_origin")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *Camera) SetCullMask(mask gdnative.Int) {
	//log.Println("Calling Camera.SetCullMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_cull_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 bool}], Returns: void
*/
func (o *Camera) SetCurrent(arg0 gdnative.Bool) {
	//log.Println("Calling Camera.SetCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Camera) SetDopplerTracking(mode gdnative.Int) {
	//log.Println("Calling Camera.SetDopplerTracking()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_doppler_tracking")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false env Environment}], Returns: void
*/
func (o *Camera) SetEnvironment(env Environment) {
	//log.Println("Calling Camera.SetEnvironment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(env.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_environment")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *Camera) SetFov(arg0 gdnative.Float) {
	//log.Println("Calling Camera.SetFov()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_fov")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ofs float}], Returns: void
*/
func (o *Camera) SetHOffset(ofs gdnative.Float) {
	//log.Println("Calling Camera.SetHOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_h_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Camera) SetKeepAspectMode(mode gdnative.Int) {
	//log.Println("Calling Camera.SetKeepAspectMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_keep_aspect_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the camera projection to orthogonal mode, by specifying a width and the [i]near[/i] and [i]far[/i] clip planes in worldspace units. (As a hint, 2D games often use this projection, with values specified in pixels)
	Args: [{ false size float} { false z_near float} { false z_far float}], Returns: void
*/
func (o *Camera) SetOrthogonal(size gdnative.Float, zNear gdnative.Float, zFar gdnative.Float) {
	//log.Println("Calling Camera.SetOrthogonal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromFloat(size)
	ptrArguments[1] = gdnative.NewPointerFromFloat(zNear)
	ptrArguments[2] = gdnative.NewPointerFromFloat(zFar)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_orthogonal")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the camera projection to perspective mode, by specifying a [i]FOV[/i] Y angle in degrees (FOV means Field of View), and the [i]near[/i] and [i]far[/i] clip planes in worldspace units.
	Args: [{ false fov float} { false z_near float} { false z_far float}], Returns: void
*/
func (o *Camera) SetPerspective(fov gdnative.Float, zNear gdnative.Float, zFar gdnative.Float) {
	//log.Println("Calling Camera.SetPerspective()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromFloat(fov)
	ptrArguments[1] = gdnative.NewPointerFromFloat(zNear)
	ptrArguments[2] = gdnative.NewPointerFromFloat(zFar)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_perspective")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *Camera) SetProjection(arg0 gdnative.Int) {
	//log.Println("Calling Camera.SetProjection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_projection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *Camera) SetSize(arg0 gdnative.Float) {
	//log.Println("Calling Camera.SetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ofs float}], Returns: void
*/
func (o *Camera) SetVOffset(ofs gdnative.Float) {
	//log.Println("Calling Camera.SetVOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(ofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_v_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *Camera) SetZfar(arg0 gdnative.Float) {
	//log.Println("Calling Camera.SetZfar()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_zfar")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *Camera) SetZnear(arg0 gdnative.Float) {
	//log.Println("Calling Camera.SetZnear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "set_znear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns how a 3D point in worldspace maps to a 2D coordinate in the [Viewport] rectangle.
	Args: [{ false world_point Vector3}], Returns: Vector2
*/
func (o *Camera) UnprojectPosition(worldPoint gdnative.Vector3) gdnative.Vector2 {
	//log.Println("Calling Camera.UnprojectPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(worldPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Camera", "unproject_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}
