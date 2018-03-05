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

//func NewThemeFromPointer(ptr gdnative.Pointer) Theme {
func NewThemeFromPointer(ptr gdnative.Pointer) Theme {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Theme{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Theme for skinning controls. Controls can be skinned individually, but for complex applications it's more efficient to just create a global theme that defines everything. This theme can be applied to any [Control], and it and its children will automatically use it. Theme resources can be alternatively loaded by writing them in a .theme file, see docs for more info.
*/
type Theme struct {
	Resource
	owner gdnative.Object
}

func (o *Theme) BaseClass() string {
	return "Theme"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Theme) X_EmitThemeChanged() {
	//log.Println("Calling Theme.X_EmitThemeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "_emit_theme_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears theme [Color] at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: void
*/
func (o *Theme) ClearColor(name gdnative.String, aType gdnative.String) {
	//log.Println("Calling Theme.ClearColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "clear_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears theme constant at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: void
*/
func (o *Theme) ClearConstant(name gdnative.String, aType gdnative.String) {
	//log.Println("Calling Theme.ClearConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "clear_constant")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears [Font] at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: void
*/
func (o *Theme) ClearFont(name gdnative.String, aType gdnative.String) {
	//log.Println("Calling Theme.ClearFont()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "clear_font")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears icon at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: void
*/
func (o *Theme) ClearIcon(name gdnative.String, aType gdnative.String) {
	//log.Println("Calling Theme.ClearIcon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "clear_icon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears [StyleBox] at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: void
*/
func (o *Theme) ClearStylebox(name gdnative.String, aType gdnative.String) {
	//log.Println("Calling Theme.ClearStylebox()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "clear_stylebox")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets theme values to a copy of the default theme values.
	Args: [], Returns: void
*/
func (o *Theme) CopyDefaultTheme() {
	//log.Println("Calling Theme.CopyDefaultTheme()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "copy_default_theme")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the [Color] at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: Color
*/
func (o *Theme) GetColor(name gdnative.String, aType gdnative.String) gdnative.Color {
	//log.Println("Calling Theme.GetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns all of the [Color]s as a [PoolStringArray] filled with each [Color]'s name, for use in [method get_color], if Theme has [code]type[/code].
	Args: [{ false type String}], Returns: PoolStringArray
*/
func (o *Theme) GetColorList(aType gdnative.String) gdnative.PoolStringArray {
	//log.Println("Calling Theme.GetColorList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_color_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the constant at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: int
*/
func (o *Theme) GetConstant(name gdnative.String, aType gdnative.String) gdnative.Int {
	//log.Println("Calling Theme.GetConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_constant")

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
        Returns all of the constants as a [PoolStringArray] filled with each constant's name, for use in [method get_constant], if Theme has [code]type[/code].
	Args: [{ false type String}], Returns: PoolStringArray
*/
func (o *Theme) GetConstantList(aType gdnative.String) gdnative.PoolStringArray {
	//log.Println("Calling Theme.GetConstantList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_constant_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Font
*/
func (o *Theme) GetDefaultFont() Font {
	//log.Println("Calling Theme.GetDefaultFont()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_default_font")

	// Call the parent method.
	// Font
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewFontFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the [Font] at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: Font
*/
func (o *Theme) GetFont(name gdnative.String, aType gdnative.String) Font {
	//log.Println("Calling Theme.GetFont()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_font")

	// Call the parent method.
	// Font
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewFontFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns all of the [Font]s as a [PoolStringArray] filled with each [Font]'s name, for use in [method get_font], if Theme has [code]type[/code].
	Args: [{ false type String}], Returns: PoolStringArray
*/
func (o *Theme) GetFontList(aType gdnative.String) gdnative.PoolStringArray {
	//log.Println("Calling Theme.GetFontList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_font_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the icon [Texture] at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: Texture
*/
func (o *Theme) GetIcon(name gdnative.String, aType gdnative.String) Texture {
	//log.Println("Calling Theme.GetIcon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_icon")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewTextureFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns all of the icons as a [PoolStringArray] filled with each [Texture]'s name, for use in [method get_icon], if Theme has [code]type[/code].
	Args: [{ false type String}], Returns: PoolStringArray
*/
func (o *Theme) GetIconList(aType gdnative.String) gdnative.PoolStringArray {
	//log.Println("Calling Theme.GetIconList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_icon_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the icon [StyleBox] at [code]name[/code] if Theme has [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: StyleBox
*/
func (o *Theme) GetStylebox(name gdnative.String, aType gdnative.String) StyleBox {
	//log.Println("Calling Theme.GetStylebox()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_stylebox")

	// Call the parent method.
	// StyleBox
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewStyleBoxFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns all of the [StyleBox]s as a [PoolStringArray] filled with each [StyleBox]'s name, for use in [method get_stylebox], if Theme has [code]type[/code].
	Args: [{ false type String}], Returns: PoolStringArray
*/
func (o *Theme) GetStyleboxList(aType gdnative.String) gdnative.PoolStringArray {
	//log.Println("Calling Theme.GetStyleboxList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_stylebox_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns all of the [StyleBox] types as a [PoolStringArray] filled with each [StyleBox]'s type, for use in [method get_stylebox] and/or [method get_stylebox_list], if Theme has [code]type[/code].
	Args: [], Returns: PoolStringArray
*/
func (o *Theme) GetStyleboxTypes() gdnative.PoolStringArray {
	//log.Println("Calling Theme.GetStyleboxTypes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_stylebox_types")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns all of the types in [code]type[/code] as a [PoolStringArray] for use in any of the get_* functions, if Theme has [code]type[/code].
	Args: [{ false type String}], Returns: PoolStringArray
*/
func (o *Theme) GetTypeList(aType gdnative.String) gdnative.PoolStringArray {
	//log.Println("Calling Theme.GetTypeList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "get_type_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns [code]true[/code] if [Color] with [code]name[/code] is in [code]type[/code]. Returns [code]false[/code] if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: bool
*/
func (o *Theme) HasColor(name gdnative.String, aType gdnative.String) gdnative.Bool {
	//log.Println("Calling Theme.HasColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "has_color")

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
        Returns [code]true[/code] if constant with [code]name[/code] is in [code]type[/code]. Returns [code]false[/code] if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: bool
*/
func (o *Theme) HasConstant(name gdnative.String, aType gdnative.String) gdnative.Bool {
	//log.Println("Calling Theme.HasConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "has_constant")

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
        Returns [code]true[/code] if [Font] with [code]name[/code] is in [code]type[/code]. Returns [code]false[/code] if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: bool
*/
func (o *Theme) HasFont(name gdnative.String, aType gdnative.String) gdnative.Bool {
	//log.Println("Calling Theme.HasFont()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "has_font")

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
        Returns [code]true[/code] if icon [Texture] with [code]name[/code] is in [code]type[/code]. Returns [code]false[/code] if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: bool
*/
func (o *Theme) HasIcon(name gdnative.String, aType gdnative.String) gdnative.Bool {
	//log.Println("Calling Theme.HasIcon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "has_icon")

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
        Returns [code]true[/code] if [StyleBox] with [code]name[/code] is in [code]type[/code]. Returns [code]false[/code] if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String}], Returns: bool
*/
func (o *Theme) HasStylebox(name gdnative.String, aType gdnative.String) gdnative.Bool {
	//log.Println("Calling Theme.HasStylebox()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "has_stylebox")

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
        Sets Theme's [Color] to [code]color[/code] at [code]name[/code] in [code]type[/code]. Does nothing if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String} { false color Color}], Returns: void
*/
func (o *Theme) SetColor(name gdnative.String, aType gdnative.String, color gdnative.Color) {
	//log.Println("Calling Theme.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)
	ptrArguments[2] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets Theme's constant to [code]constant[/code] at [code]name[/code] in [code]type[/code]. Does nothing if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String} { false constant int}], Returns: void
*/
func (o *Theme) SetConstant(name gdnative.String, aType gdnative.String, constant gdnative.Int) {
	//log.Println("Calling Theme.SetConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)
	ptrArguments[2] = gdnative.NewPointerFromInt(constant)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "set_constant")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false font Font}], Returns: void
*/
func (o *Theme) SetDefaultFont(font Font) {
	//log.Println("Calling Theme.SetDefaultFont()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(font.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "set_default_font")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets Theme's [Font] to [code]font[/code] at [code]name[/code] in [code]type[/code]. Does nothing if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String} { false font Font}], Returns: void
*/
func (o *Theme) SetFont(name gdnative.String, aType gdnative.String, font Font) {
	//log.Println("Calling Theme.SetFont()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)
	ptrArguments[2] = gdnative.NewPointerFromObject(font.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "set_font")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets Theme's icon [Texture] to [code]texture[/code] at [code]name[/code] in [code]type[/code]. Does nothing if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String} { false texture Texture}], Returns: void
*/
func (o *Theme) SetIcon(name gdnative.String, aType gdnative.String, texture Texture) {
	//log.Println("Calling Theme.SetIcon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)
	ptrArguments[2] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "set_icon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets Theme's [StyleBox] to [code]stylebox[/code] at [code]name[/code] in [code]type[/code]. Does nothing if Theme does not have [code]type[/code].
	Args: [{ false name String} { false type String} { false texture StyleBox}], Returns: void
*/
func (o *Theme) SetStylebox(name gdnative.String, aType gdnative.String, texture StyleBox) {
	//log.Println("Calling Theme.SetStylebox()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(aType)
	ptrArguments[2] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Theme", "set_stylebox")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
