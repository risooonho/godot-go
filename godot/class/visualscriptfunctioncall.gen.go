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

//func NewVisualScriptFunctionCallFromPointer(ptr gdnative.Pointer) VisualScriptFunctionCall {
func NewVisualScriptFunctionCallFromPointer(ptr gdnative.Pointer) VisualScriptFunctionCall {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptFunctionCall{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptFunctionCall struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptFunctionCall) BaseClass() string {
	return "VisualScriptFunctionCall"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *VisualScriptFunctionCall) X_GetArgumentCache() gdnative.Dictionary {
	//log.Println("Calling VisualScriptFunctionCall.X_GetArgumentCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "_get_argument_cache")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false argument_cache Dictionary}], Returns: void
*/
func (o *VisualScriptFunctionCall) X_SetArgumentCache(argumentCache gdnative.Dictionary) {
	//log.Println("Calling VisualScriptFunctionCall.X_SetArgumentCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(argumentCache)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "_set_argument_cache")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *VisualScriptFunctionCall) GetBasePath() gdnative.NodePath {
	//log.Println("Calling VisualScriptFunctionCall.GetBasePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "get_base_path")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptFunctionCall) GetBaseScript() gdnative.String {
	//log.Println("Calling VisualScriptFunctionCall.GetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "get_base_script")

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
	Args: [], Returns: String
*/
func (o *VisualScriptFunctionCall) GetBaseType() gdnative.String {
	//log.Println("Calling VisualScriptFunctionCall.GetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "get_base_type")

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
	Args: [], Returns: enum.Variant::Type
*/

/*
        Undocumented
	Args: [], Returns: enum.VisualScriptFunctionCall::CallMode
*/

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptFunctionCall) GetFunction() gdnative.String {
	//log.Println("Calling VisualScriptFunctionCall.GetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "get_function")

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
	Args: [], Returns: enum.VisualScriptFunctionCall::RPCCallMode
*/

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptFunctionCall) GetSingleton() gdnative.String {
	//log.Println("Calling VisualScriptFunctionCall.GetSingleton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "get_singleton")

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
	Args: [], Returns: int
*/
func (o *VisualScriptFunctionCall) GetUseDefaultArgs() gdnative.Int {
	//log.Println("Calling VisualScriptFunctionCall.GetUseDefaultArgs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "get_use_default_args")

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
	Args: [], Returns: bool
*/
func (o *VisualScriptFunctionCall) GetValidate() gdnative.Bool {
	//log.Println("Calling VisualScriptFunctionCall.GetValidate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "get_validate")

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
	Args: [{ false base_path NodePath}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetBasePath(basePath gdnative.NodePath) {
	//log.Println("Calling VisualScriptFunctionCall.SetBasePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(basePath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_base_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false base_script String}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetBaseScript(baseScript gdnative.String) {
	//log.Println("Calling VisualScriptFunctionCall.SetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(baseScript)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_base_script")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false base_type String}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetBaseType(baseType gdnative.String) {
	//log.Println("Calling VisualScriptFunctionCall.SetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(baseType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_base_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false basic_type int}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetBasicType(basicType gdnative.Int) {
	//log.Println("Calling VisualScriptFunctionCall.SetBasicType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(basicType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_basic_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetCallMode(mode gdnative.Int) {
	//log.Println("Calling VisualScriptFunctionCall.SetCallMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_call_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false function String}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetFunction(function gdnative.String) {
	//log.Println("Calling VisualScriptFunctionCall.SetFunction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(function)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_function")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetRpcCallMode(mode gdnative.Int) {
	//log.Println("Calling VisualScriptFunctionCall.SetRpcCallMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_rpc_call_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false singleton String}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetSingleton(singleton gdnative.String) {
	//log.Println("Calling VisualScriptFunctionCall.SetSingleton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(singleton)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_singleton")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount int}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetUseDefaultArgs(amount gdnative.Int) {
	//log.Println("Calling VisualScriptFunctionCall.SetUseDefaultArgs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_use_default_args")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *VisualScriptFunctionCall) SetValidate(enable gdnative.Bool) {
	//log.Println("Calling VisualScriptFunctionCall.SetValidate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionCall", "set_validate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
