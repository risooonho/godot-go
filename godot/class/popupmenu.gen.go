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

//func NewPopupMenuFromPointer(ptr gdnative.Pointer) PopupMenu {
func NewPopupMenuFromPointer(ptr gdnative.Pointer) PopupMenu {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PopupMenu{}
	obj.SetBaseObject(owner)

	return obj
}

/*
PopupMenu is the typical Control that displays a list of options. They are popular in toolbars or context menus.
*/
type PopupMenu struct {
	Popup
	owner gdnative.Object
}

func (o *PopupMenu) BaseClass() string {
	return "PopupMenu"
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *PopupMenu) X_GetItems() gdnative.Array {
	//log.Println("Calling PopupMenu.X_GetItems()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "_get_items")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *PopupMenu) X_GuiInput(arg0 InputEvent) {
	//log.Println("Calling PopupMenu.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Array}], Returns: void
*/
func (o *PopupMenu) X_SetItems(arg0 gdnative.Array) {
	//log.Println("Calling PopupMenu.X_SetItems()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "_set_items")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *PopupMenu) X_SubmenuTimeout() {
	//log.Println("Calling PopupMenu.X_SubmenuTimeout()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "_submenu_timeout")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add a new checkable item with text "label". An id can optionally be provided, as well as an accelerator. If no id is provided, one will be created from the index. Note that checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
	Args: [{ false label String} {-1 true id int} {0 true accel int}], Returns: void
*/
func (o *PopupMenu) AddCheckItem(label gdnative.String, id gdnative.Int, accel gdnative.Int) {
	//log.Println("Calling PopupMenu.AddCheckItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(label)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)
	ptrArguments[2] = gdnative.NewPointerFromInt(accel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_check_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false shortcut ShortCut} {-1 true id int} {False true global bool}], Returns: void
*/
func (o *PopupMenu) AddCheckShortcut(shortcut ShortCut, id gdnative.Int, global gdnative.Bool) {
	//log.Println("Calling PopupMenu.AddCheckShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(shortcut.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(id)
	ptrArguments[2] = gdnative.NewPointerFromBool(global)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_check_shortcut")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add a new checkable item with text "label" and icon "texture". An id can optionally be provided, as well as an accelerator. If no id is provided, one will be created from the index. Note that checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
	Args: [{ false texture Texture} { false label String} {-1 true id int} {0 true accel int}], Returns: void
*/
func (o *PopupMenu) AddIconCheckItem(texture Texture, label gdnative.String, id gdnative.Int, accel gdnative.Int) {
	//log.Println("Calling PopupMenu.AddIconCheckItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromString(label)
	ptrArguments[2] = gdnative.NewPointerFromInt(id)
	ptrArguments[3] = gdnative.NewPointerFromInt(accel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_icon_check_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false texture Texture} { false shortcut ShortCut} {-1 true id int} {False true global bool}], Returns: void
*/
func (o *PopupMenu) AddIconCheckShortcut(texture Texture, shortcut ShortCut, id gdnative.Int, global gdnative.Bool) {
	//log.Println("Calling PopupMenu.AddIconCheckShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromObject(shortcut.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromInt(id)
	ptrArguments[3] = gdnative.NewPointerFromBool(global)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_icon_check_shortcut")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add a new item with text "label" and icon "texture". An id can optionally be provided, as well as an accelerator keybinding. If no id is provided, one will be created from the index.
	Args: [{ false texture Texture} { false label String} {-1 true id int} {0 true accel int}], Returns: void
*/
func (o *PopupMenu) AddIconItem(texture Texture, label gdnative.String, id gdnative.Int, accel gdnative.Int) {
	//log.Println("Calling PopupMenu.AddIconItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromString(label)
	ptrArguments[2] = gdnative.NewPointerFromInt(id)
	ptrArguments[3] = gdnative.NewPointerFromInt(accel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_icon_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false texture Texture} { false shortcut ShortCut} {-1 true id int} {False true global bool}], Returns: void
*/
func (o *PopupMenu) AddIconShortcut(texture Texture, shortcut ShortCut, id gdnative.Int, global gdnative.Bool) {
	//log.Println("Calling PopupMenu.AddIconShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromObject(shortcut.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromInt(id)
	ptrArguments[3] = gdnative.NewPointerFromBool(global)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_icon_shortcut")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add a new item with text "label". An id can optionally be provided, as well as an accelerator keybinding. If no id is provided, one will be created from the index.
	Args: [{ false label String} {-1 true id int} {0 true accel int}], Returns: void
*/
func (o *PopupMenu) AddItem(label gdnative.String, id gdnative.Int, accel gdnative.Int) {
	//log.Println("Calling PopupMenu.AddItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(label)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)
	ptrArguments[2] = gdnative.NewPointerFromInt(accel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add a separator between items. Separators also occupy an index.
	Args: [], Returns: void
*/
func (o *PopupMenu) AddSeparator() {
	//log.Println("Calling PopupMenu.AddSeparator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_separator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false shortcut ShortCut} {-1 true id int} {False true global bool}], Returns: void
*/
func (o *PopupMenu) AddShortcut(shortcut ShortCut, id gdnative.Int, global gdnative.Bool) {
	//log.Println("Calling PopupMenu.AddShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(shortcut.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(id)
	ptrArguments[2] = gdnative.NewPointerFromBool(global)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_shortcut")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds an item with a submenu. The submenu is the name of a child PopupMenu node that would be shown when the item is clicked. An id can optionally be provided, but if is isn't provided, one will be created from the index.
	Args: [{ false label String} { false submenu String} {-1 true id int}], Returns: void
*/
func (o *PopupMenu) AddSubmenuItem(label gdnative.String, submenu gdnative.String, id gdnative.Int) {
	//log.Println("Calling PopupMenu.AddSubmenuItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(label)
	ptrArguments[1] = gdnative.NewPointerFromString(submenu)
	ptrArguments[2] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "add_submenu_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clear the popup menu, in effect removing all items.
	Args: [], Returns: void
*/
func (o *PopupMenu) Clear() {
	//log.Println("Calling PopupMenu.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Return the accelerator of the item at index "idx". Accelerators are special combinations of keys that activate the item, no matter which control is focused.
	Args: [{ false idx int}], Returns: int
*/
func (o *PopupMenu) GetItemAccelerator(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling PopupMenu.GetItemAccelerator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_accelerator")

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
        Return the amount of items.
	Args: [], Returns: int
*/
func (o *PopupMenu) GetItemCount() gdnative.Int {
	//log.Println("Calling PopupMenu.GetItemCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_count")

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
        Return the icon of the item at index "idx".
	Args: [{ false idx int}], Returns: Texture
*/
func (o *PopupMenu) GetItemIcon(idx gdnative.Int) Texture {
	//log.Println("Calling PopupMenu.GetItemIcon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_icon")

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
        Return the id of the item at index "idx".
	Args: [{ false idx int}], Returns: int
*/
func (o *PopupMenu) GetItemId(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling PopupMenu.GetItemId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_id")

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
        Find and return the index of the item containing a given id.
	Args: [{ false id int}], Returns: int
*/
func (o *PopupMenu) GetItemIndex(id gdnative.Int) gdnative.Int {
	//log.Println("Calling PopupMenu.GetItemIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_index")

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
        Return the metadata of an item, which might be of any type. You can set it with [method set_item_metadata], which provides a simple way of assigning context data to items.
	Args: [{ false idx int}], Returns: Variant
*/
func (o *PopupMenu) GetItemMetadata(idx gdnative.Int) gdnative.Variant {
	//log.Println("Calling PopupMenu.GetItemMetadata()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_metadata")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*

	Args: [{ false idx int}], Returns: ShortCut
*/
func (o *PopupMenu) GetItemShortcut(idx gdnative.Int) ShortCut {
	//log.Println("Calling PopupMenu.GetItemShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_shortcut")

	// Call the parent method.
	// ShortCut
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewShortCutFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the submenu name of the item at index "idx".
	Args: [{ false idx int}], Returns: String
*/
func (o *PopupMenu) GetItemSubmenu(idx gdnative.Int) gdnative.String {
	//log.Println("Calling PopupMenu.GetItemSubmenu()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_submenu")

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
        Return the text of the item at index "idx".
	Args: [{ false idx int}], Returns: String
*/
func (o *PopupMenu) GetItemText(idx gdnative.Int) gdnative.String {
	//log.Println("Calling PopupMenu.GetItemText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_text")

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

	Args: [{ false idx int}], Returns: String
*/
func (o *PopupMenu) GetItemTooltip(idx gdnative.Int) gdnative.String {
	//log.Println("Calling PopupMenu.GetItemTooltip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "get_item_tooltip")

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
	Args: [], Returns: bool
*/
func (o *PopupMenu) IsHideOnCheckableItemSelection() gdnative.Bool {
	//log.Println("Calling PopupMenu.IsHideOnCheckableItemSelection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "is_hide_on_checkable_item_selection")

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
func (o *PopupMenu) IsHideOnItemSelection() gdnative.Bool {
	//log.Println("Calling PopupMenu.IsHideOnItemSelection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "is_hide_on_item_selection")

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
func (o *PopupMenu) IsHideOnStateItemSelection() gdnative.Bool {
	//log.Println("Calling PopupMenu.IsHideOnStateItemSelection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "is_hide_on_state_item_selection")

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
        Return whether the item at index "idx" has a checkbox. Note that checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
	Args: [{ false idx int}], Returns: bool
*/
func (o *PopupMenu) IsItemCheckable(idx gdnative.Int) gdnative.Bool {
	//log.Println("Calling PopupMenu.IsItemCheckable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "is_item_checkable")

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
        Return the checkstate status of the item at index "idx".
	Args: [{ false idx int}], Returns: bool
*/
func (o *PopupMenu) IsItemChecked(idx gdnative.Int) gdnative.Bool {
	//log.Println("Calling PopupMenu.IsItemChecked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "is_item_checked")

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
        Return whether the item at index "idx" is disabled. When it is disabled it can't be selected, or its action invoked.
	Args: [{ false idx int}], Returns: bool
*/
func (o *PopupMenu) IsItemDisabled(idx gdnative.Int) gdnative.Bool {
	//log.Println("Calling PopupMenu.IsItemDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "is_item_disabled")

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
        Return whether the item is a separator. If it is, it would be displayed as a line.
	Args: [{ false idx int}], Returns: bool
*/
func (o *PopupMenu) IsItemSeparator(idx gdnative.Int) gdnative.Bool {
	//log.Println("Calling PopupMenu.IsItemSeparator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "is_item_separator")

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
        Removes the item at index "idx" from the menu. Note that the indexes of items after the removed item are going to be shifted by one.
	Args: [{ false idx int}], Returns: void
*/
func (o *PopupMenu) RemoveItem(idx gdnative.Int) {
	//log.Println("Calling PopupMenu.RemoveItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "remove_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *PopupMenu) SetHideOnCheckableItemSelection(enable gdnative.Bool) {
	//log.Println("Calling PopupMenu.SetHideOnCheckableItemSelection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_hide_on_checkable_item_selection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *PopupMenu) SetHideOnItemSelection(enable gdnative.Bool) {
	//log.Println("Calling PopupMenu.SetHideOnItemSelection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_hide_on_item_selection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *PopupMenu) SetHideOnStateItemSelection(enable gdnative.Bool) {
	//log.Println("Calling PopupMenu.SetHideOnStateItemSelection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_hide_on_state_item_selection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the accelerator of the item at index "idx". Accelerators are special combinations of keys that activate the item, no matter which control is focused.
	Args: [{ false idx int} { false accel int}], Returns: void
*/
func (o *PopupMenu) SetItemAccelerator(idx gdnative.Int, accel gdnative.Int) {
	//log.Println("Calling PopupMenu.SetItemAccelerator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(accel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_accelerator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set whether the item at index "idx" has a checkbox. Note that checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
	Args: [{ false idx int} { false enable bool}], Returns: void
*/
func (o *PopupMenu) SetItemAsCheckable(idx gdnative.Int, enable gdnative.Bool) {
	//log.Println("Calling PopupMenu.SetItemAsCheckable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_as_checkable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Mark the item at index "idx" as a separator, which means that it would be displayed as a mere line.
	Args: [{ false idx int} { false enable bool}], Returns: void
*/
func (o *PopupMenu) SetItemAsSeparator(idx gdnative.Int, enable gdnative.Bool) {
	//log.Println("Calling PopupMenu.SetItemAsSeparator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_as_separator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the checkstate status of the item at index "idx".
	Args: [{ false idx int} { false checked bool}], Returns: void
*/
func (o *PopupMenu) SetItemChecked(idx gdnative.Int, checked gdnative.Bool) {
	//log.Println("Calling PopupMenu.SetItemChecked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromBool(checked)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_checked")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets whether the item at index "idx" is disabled or not. When it is disabled it can't be selected, or its action invoked.
	Args: [{ false idx int} { false disabled bool}], Returns: void
*/
func (o *PopupMenu) SetItemDisabled(idx gdnative.Int, disabled gdnative.Bool) {
	//log.Println("Calling PopupMenu.SetItemDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false idx int} { false icon Texture}], Returns: void
*/
func (o *PopupMenu) SetItemIcon(idx gdnative.Int, icon Texture) {
	//log.Println("Calling PopupMenu.SetItemIcon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromObject(icon.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_icon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the id of the item at index "idx".
	Args: [{ false idx int} { false id int}], Returns: void
*/
func (o *PopupMenu) SetItemId(idx gdnative.Int, id gdnative.Int) {
	//log.Println("Calling PopupMenu.SetItemId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_id")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the metadata of an item, which might be of any type. You can later get it with [method get_item_metadata], which provides a simple way of assigning context data to items.
	Args: [{ false idx int} { false metadata Variant}], Returns: void
*/
func (o *PopupMenu) SetItemMetadata(idx gdnative.Int, metadata gdnative.Variant) {
	//log.Println("Calling PopupMenu.SetItemMetadata()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVariant(metadata)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_metadata")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false idx int} { false state int}], Returns: void
*/
func (o *PopupMenu) SetItemMultistate(idx gdnative.Int, state gdnative.Int) {
	//log.Println("Calling PopupMenu.SetItemMultistate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(state)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_multistate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false idx int} { false shortcut ShortCut} {False true global bool}], Returns: void
*/
func (o *PopupMenu) SetItemShortcut(idx gdnative.Int, shortcut ShortCut, global gdnative.Bool) {
	//log.Println("Calling PopupMenu.SetItemShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromObject(shortcut.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromBool(global)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_shortcut")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the submenu of the item at index "idx". The submenu is the name of a child PopupMenu node that would be shown when the item is clicked.
	Args: [{ false idx int} { false submenu String}], Returns: void
*/
func (o *PopupMenu) SetItemSubmenu(idx gdnative.Int, submenu gdnative.String) {
	//log.Println("Calling PopupMenu.SetItemSubmenu()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromString(submenu)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_submenu")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the text of the item at index "idx".
	Args: [{ false idx int} { false text String}], Returns: void
*/
func (o *PopupMenu) SetItemText(idx gdnative.Int, text gdnative.String) {
	//log.Println("Calling PopupMenu.SetItemText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_text")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false idx int} { false tooltip String}], Returns: void
*/
func (o *PopupMenu) SetItemTooltip(idx gdnative.Int, tooltip gdnative.String) {
	//log.Println("Calling PopupMenu.SetItemTooltip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromString(tooltip)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "set_item_tooltip")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false idx int}], Returns: void
*/
func (o *PopupMenu) ToggleItemChecked(idx gdnative.Int) {
	//log.Println("Calling PopupMenu.ToggleItemChecked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "toggle_item_checked")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false idx int}], Returns: void
*/
func (o *PopupMenu) ToggleItemMultistate(idx gdnative.Int) {
	//log.Println("Calling PopupMenu.ToggleItemMultistate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PopupMenu", "toggle_item_multistate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
