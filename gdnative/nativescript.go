// Package gdnative provides a wrapper around Godot's nativescript
// extension. It exists to provide a way to use Go as an alternative scripting
// language from GDScript.
package gdnative

/*
#include <nativescript/godot_nativescript.h>
#include "nativescript.gen.h"
#include "nativescript.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

type PropertyUsageFlags int

func (p PropertyUsageFlags) getBase() C.godot_property_usage_flags {
	return C.godot_property_usage_flags(p)
}

const (
	PropertyUsageStorage             PropertyUsageFlags = 1
	PropertyUsageEditor              PropertyUsageFlags = 2
	PropertyUsageNetwork             PropertyUsageFlags = 4
	PropertyUsageEditorHelper        PropertyUsageFlags = 8
	PropertyUsageCheckable           PropertyUsageFlags = 16  //used for editing global variables
	PropertyUsageChecked             PropertyUsageFlags = 32  //used for editing global variables
	PropertyUsageInternationalized   PropertyUsageFlags = 64  //hint for internationalized strings
	PropertyUsageGroup               PropertyUsageFlags = 128 //used for grouping props in the editor
	PropertyUsageCategory            PropertyUsageFlags = 256
	PropertyUsageStoreIfNonZero      PropertyUsageFlags = 512  //only store if nonzero
	PropertyUsageStoreIfNonOne       PropertyUsageFlags = 1024 //only store if false
	PropertyUsageNoInstanceState     PropertyUsageFlags = 2048
	PropertyUsageRestartIfChanged    PropertyUsageFlags = 4096
	PropertyUsageScriptVariable      PropertyUsageFlags = 8192
	PropertyUsageStoreIfNull         PropertyUsageFlags = 16384
	PropertyUsageAnimateAsTrigger    PropertyUsageFlags = 32768
	PropertyUsageUpdateAllIfModified PropertyUsageFlags = 65536

	PropertyUsageDefault     PropertyUsageFlags = PropertyUsageStorage | PropertyUsageEditor | PropertyUsageNetwork
	PropertyUsageDefaultIntl PropertyUsageFlags = PropertyUsageStorage | PropertyUsageEditor | PropertyUsageNetwork | PropertyUsageInternationalized
	PropertyUsageNoEditor    PropertyUsageFlags = PropertyUsageStorage | PropertyUsageNetwork
)

// CreateFunc will be called when we need to create a new instance of a class.
// Takes the instance object, user data.
type CreateFunc func(Object) string

// CreateFuncRegistry is a mapping of instance creation functions. This map is
// used whenever a CreateFunc is registered. It is also used to look up a
// Creation function when Godot asks Go to create a new class instance.
var CreateFuncRegistry = map[string]CreateFunc{}

// DestroyFunc will be called when the object is destroyed. Takes the instance
// object, method data, user data. The method data is generally the class name,
// and the user data is generally the class instance id to destroy.
type DestroyFunc func(Object, string, string)

// DestroyFuncRegistry is a mapping of instance destroy functions. This map is
// used whenever a DestroyFunc is registered. It is also used to look up a
// Destroy function when Godot asks Go to destroy a class instance.
var DestroyFuncRegistry = map[string]DestroyFunc{}

// FreeFunc will be called when we should free the instance from memory. Takes
// in method data. The method data is generally the class name to be freed.
type FreeFunc func(string)

// FreeFuncRegistry is a mapping of instance free functions. This map is used
// whenever a FreeFunc is registered. It is also used to look up a Free
// function when Godot asks Go to free a class instance.
var FreeFuncRegistry = map[string]FreeFunc{}

// MethodFunc will be called when a method attached to an instance is called.
type MethodFunc func(Object, string, string, int, []Variant) Variant

// MethodFuncRegistry is a mapping of instance method functions. This map is
// used whenever a MethofFunc is registered. It is also used to look up a Method
// function when Godot asks Go to call a class method.
var MethodFuncRegistry = map[string]MethodFunc{}

// InstanceCreateFunc is a structure that contains the instance creation function
// that will be called when Godot asks Go to create a new instance of a class.
type InstanceCreateFunc struct {
	base       C.godot_instance_create_func
	CreateFunc CreateFunc
	MethodData string
	FreeFunc   FreeFunc
}

func (i *InstanceCreateFunc) getBase() C.godot_instance_create_func {
	return i.base
}

// InstanceDestroyFunc is a structure that contains the instance destruction function
// that will be called when Godot asks Go to destroy a class instance.
type InstanceDestroyFunc struct {
	base        C.godot_instance_destroy_func
	DestroyFunc DestroyFunc
	MethodData  string
	FreeFunc    FreeFunc
}

func (i *InstanceDestroyFunc) getBase() C.godot_instance_destroy_func {
	return i.base
}

// InstanceMethod is a structure that contains an instance method function that
// will be called when Godot asks Go to call a method on a class.
type InstanceMethod struct {
	base       C.godot_instance_method
	Method     MethodFunc
	MethodData string
	FreeFunc   FreeFunc
}

func (i *InstanceMethod) getBase() C.godot_instance_method {
	return i.base
}

type MethodAttributes struct {
	base    C.godot_method_attributes
	RPCType MethodRpcMode
}

// NativeScript is a wrapper for the NativeScriptAPI.
var NativeScript = &nativeScript{}

// nativeScript is a structure that wraps the NativeScriptAPI methods.
type nativeScript struct {
	api *C.godot_gdnative_ext_nativescript_api_struct

	// Handle is a pointer to the gdnative handler. It must be passed to any
	// Godot nativescript functions. This will be populated when 'godot_nativescript_init'
	// is called by Godot upon script initialization.
	handle unsafe.Pointer
}

// RegisterClass will register the given class with Godot. This will make it
// available to be attached to a Node in Godot. The name of the class that you
// provide here will be the name that you specify when you attach a NativeScript
// script to a Node. The base parameter that you specify will be what the class
// should be inheriting from (e.g. Node2D, Node, etc.).
//
// The create and destroy function parameters are C structs that contain
// function pointers to C methods that will be called when the object is created
// or destroyed. Because of the pointer passing rules, Go code can not pass a
// function value directly to C, so a gateway function should be used. More
// information on using a gateway function can be found here:
// https://github.com/golang/go/wiki/cgo#function-variables
func (n *nativeScript) RegisterClass(name, base string, createFunc *InstanceCreateFunc, destroyFunc *InstanceDestroyFunc) {
	// Construct the C struct based on the Go struct wrappers
	createFunc.base.create_func = (C.create_func)(unsafe.Pointer(C.cgo_gateway_create_func))
	createFunc.base.method_data = unsafe.Pointer(C.CString(createFunc.MethodData))
	createFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))
	destroyFunc.base.destroy_func = (C.destroy_func)(unsafe.Pointer(C.cgo_gateway_destroy_func))
	destroyFunc.base.method_data = unsafe.Pointer(C.CString(destroyFunc.MethodData))
	destroyFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))

	// Register our Create and Destroy functions in a Go map, so the correct
	// function can be called when cgo_gateway_<type>_func is called.
	CreateFuncRegistry[createFunc.MethodData] = createFunc.CreateFunc
	FreeFuncRegistry[createFunc.MethodData] = createFunc.FreeFunc
	DestroyFuncRegistry[destroyFunc.MethodData] = destroyFunc.DestroyFunc
	FreeFuncRegistry[destroyFunc.MethodData] = destroyFunc.FreeFunc

	// Register the class with Godot.
	C.go_godot_nativescript_register_class(
		n.api,
		n.handle,
		C.CString(name),
		C.CString(base),
		createFunc.base,
		destroyFunc.base,
	)
}

// RegisterToolClass will register the given class with Godot as a tool. Refer to
// the 'RegisterClass' method for more information on how to use this.
func (n *nativeScript) RegisterToolClass(name, base string, createFunc *InstanceCreateFunc, destroyFunc *InstanceDestroyFunc) {
	// Construct the C struct based on the Go struct wrappers
	createFunc.base.create_func = (C.create_func)(unsafe.Pointer(C.cgo_gateway_create_func))
	createFunc.base.method_data = unsafe.Pointer(C.CString(createFunc.MethodData))
	createFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))
	destroyFunc.base.destroy_func = (C.destroy_func)(unsafe.Pointer(C.cgo_gateway_destroy_func))
	destroyFunc.base.method_data = unsafe.Pointer(C.CString(destroyFunc.MethodData))
	destroyFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))

	// Register our Create and Destroy functions in a Go map, so the correct
	// function can be called when cgo_gateway_<type>_func is called.
	CreateFuncRegistry[createFunc.MethodData] = createFunc.CreateFunc
	FreeFuncRegistry[createFunc.MethodData] = createFunc.FreeFunc
	DestroyFuncRegistry[destroyFunc.MethodData] = destroyFunc.DestroyFunc
	FreeFuncRegistry[destroyFunc.MethodData] = destroyFunc.FreeFunc

	// Register the class with Godot.
	C.go_godot_nativescript_register_tool_class(
		n.api,
		n.handle,
		C.CString(name),
		C.CString(base),
		createFunc.base,
		destroyFunc.base,
	)
}

// RegisterMethod will register the given method with Godot and associate it with
// the given class name. The name parameter is the name of the class you wish to
// attach this method to. The funcName parameter is the name of the function you
// want to register. The attributes and method are what will actually be called
// when Godot calls the method on the object.
func (n *nativeScript) RegisterMethod(name, funcName string, attributes *MethodAttributes, method *InstanceMethod) {
	// Construct the C struct based on the Go struct wrappers
	attributes.base.rpc_type = attributes.RPCType.getBase()
	method.base.method = (C.method)(unsafe.Pointer(C.cgo_gateway_method_func))
	method.base.method_data = unsafe.Pointer(C.CString(method.MethodData))
	method.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))

	// Register the Method function in a Go map, so the correct function can
	// be called when cgo_gateway_<type>_func is called.
	MethodFuncRegistry[method.MethodData] = method.Method
	FreeFuncRegistry[method.MethodData] = method.FreeFunc

	// Register the method with Godot.
	C.go_godot_nativescript_register_method(
		n.api,
		n.handle,
		C.CString(name),
		C.CString(funcName),
		attributes.base,
		method.base,
	)
}

// RegisterProperty will register the given property with Godot and associate it
// with the given class name. The name parameter is the name of the class you wish
// to attach this property to. The path is the name of the property itself. The
// attributes and setter/getter methods are what will be called when Godot gets
// or sets this property.
func (n *nativeScript) RegisterProperty(name, path string, attributes *C.godot_property_attributes, setFunc C.godot_property_set_func, getFunc C.godot_property_get_func) {
	C.go_godot_nativescript_register_property(
		n.api,
		n.handle,
		C.CString(name),
		C.CString(path),
		attributes,
		setFunc,
		getFunc,
	)
}

// RegisterSignal will register the given signal with Godot.
func (n *nativeScript) RegisterSignal(name string, signal *C.godot_signal) {
	C.go_godot_nativescript_register_signal(
		n.api,
		n.handle,
		C.CString(name),
		signal,
	)
}

// nativeScriptInit will be called when `godot_nativescript_init` is called by
// Godot. You can use `SetNativeScriptInit` to set the function that will be called
// when NativeScript initializes.
var nativeScriptInit func()

// SetNativeScriptInit will configure the given function to be called when
// `godot_nativescript_init` is called by Godot upon NativeScript initialization.
// This is used so you can define a function that will run to register all of the
// classes that you want exposed to Godot.
func SetNativeScriptInit(initFunc func()) {
	nativeScriptInit = initFunc
}

/** Script entry (Registering all the classes and stuff) **/
// godot_nativescript_init is the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes,
// etc. The `unsafe.Pointer` type is used to represent a void C pointer.
//export godot_nativescript_init
func godot_nativescript_init(hdl unsafe.Pointer) {
	log.Println("Initializing NativeScript")
	NativeScript.handle = hdl

	// Call the user-defined nativeScriptInit function
	if nativeScriptInit == nil {
		err := "NativeScript initialization function was not set! Use `gdnative.SetNativeScriptInit` to define the function that will run to register classes."
		log.Println(err)
		Log.Error(err)

		return
	}
	nativeScriptInit()
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in nativescript.c. It will be ultimately called by
// Godot, where it will pass us the Godot object and the MethodData defined in
// CreateFunc. We will need to return UserData, which can be used to track the
// actual instance that was created.
//export go_create_func
func go_create_func(godotObject *C.godot_object, methodData unsafe.Pointer) unsafe.Pointer {
	// Convert the method data into a Go string.
	methodDataString := unsafeToGoString(methodData)
	log.Println("Create function called for:", methodDataString)

	// Look up the creation function in our CreateFuncRegistry for the function
	// to call.
	constructor := CreateFuncRegistry[methodDataString]

	// Call the constructor and return the user data string. The user data
	// returned by the create func will be passed to the method function as
	// userData.
	userData := constructor(Object{base: godotObject})

	return unsafe.Pointer(C.CString(userData))
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in nativescript.c. It will use the userData passed to it,
// which is a CString of the instance id, which we can use to delete the instance
// from the instance registry. This will make the instance available to be garbage
// collected.
//export go_destroy_func
func go_destroy_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer) {
	// Convert the method data and user data into a Go string
	methodDataString := unsafeToGoString(methodData)
	userDataString := unsafeToGoString(userData)
	log.Println("Destroy function called for:", methodDataString)

	// Look up the destroy function in our DestroyFuncRegistry for the function
	// to call.
	destructor := DestroyFuncRegistry[methodDataString]

	// Call the destructor function. We pass the methodData and userData to
	// the destructor so it knows which class and instance to destroy.
	destructor(Object{base: godotObject}, methodDataString, userDataString)
}

//export go_free_func
func go_free_func(methodData unsafe.Pointer) {
	// Convert the method data and user data into a Go string
	methodDataString := unsafeToGoString(methodData)
	log.Println("Free function called for:", methodDataString)

	// Look up the free function in our FreeFuncRegistry for the function
	// to call.
	freer := FreeFuncRegistry[methodDataString]

	// Call the free function. We pass the methodData to the free
	// function so it knows which clas to free.
	freer(methodDataString)
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in nativescript.c.
//export go_method_func
func go_method_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant) C.godot_variant {
	// Convert the method data and user data into a Go string
	methodDataString := unsafeToGoString(methodData)
	userDataString := unsafeToGoString(userData)

	// Create a slice of Variants for the arguments
	variantArgs := []Variant{}

	// Get the size of each godot_variant object pointer.
	log.Println("  Getting size of argument pointer")
	size := unsafe.Sizeof(*args)

	// Panic if something's wrong.
	if int(numArgs) > 50 {
		panic("Too many arguments. Invalid method.")
	}

	// If we have arguments, append the first argument.
	if int(numArgs) > 0 {
		arg := *args
		// Loop through all our arguments.
		for i := 0; i < int(numArgs); i++ {
			// Convert the variant into a Go Variant
			variant := Variant{base: arg}

			// Append the variant to our list of variants
			variantArgs = append(variantArgs, variant)

			// Convert the pointer into a uintptr so we can perform artithmetic on it.
			arrayPtr := uintptr(unsafe.Pointer(args))

			// Add the size of the godot_variant pointer to our array pointer to get the position
			// of the next argument.
			arg = (*C.godot_variant)(unsafe.Pointer(arrayPtr + size))
		}
	}

	// Look up the method function in our MethodFuncRegistry for the function
	// to call.
	method := MethodFuncRegistry[methodDataString]

	// Call the method
	ret := method(Object{base: godotObject}, methodDataString, userDataString, int(numArgs), variantArgs)

	return *ret.base
}

////godot_variant go_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args)
////export go_method_func
//func go_method_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant) C.godot_variant {
//
//	// Get the object instance based on the instance string given in userData.
//	methodString := unsafeToGoString(methodData)
//	instanceString := unsafeToGoString(userData)
//	class := instanceRegistry[instanceString]
//	classValue := reflect.ValueOf(class)
//
//	log.Println("Method was called!")
//	log.Println("  godotObject:", godotObject)
//	log.Println("  numArgs:", int(numArgs))
//	log.Println("  args:", args)
//	log.Println("  instance:", class)
//	log.Println("  instanceString (userData):", instanceString)
//	log.Println("  methodString (methodData):", methodString)
//
//	// Create a slice of godot_variant arguments
//	goArgsSlice := []reflect.Value{}
//
//	// Get the size of each godot_variant object pointer.
//	log.Println("  Getting size of argument pointer")
//	size := unsafe.Sizeof(*args)
//
//	// Panic if something's wrong.
//	if int(numArgs) > 50 {
//		panic("Too many arguments. Invalid method.")
//	}
//
//	// If we have arguments, append the first argument.
//	log.Println("  Checking if method had arguments")
//	if int(numArgs) > 0 {
//		log.Println("    It does!")
//		arg := *args
//		// Loop through all our arguments.
//		for i := 0; i < int(numArgs); i++ {
//			// Check to see what type this variant is
//			variantType := C.godot_variant_get_type(arg)
//			log.Println("Argument variant type:", variantType)
//
//			// TODO: Handle all variant types.
//			switch variantType {
//			case C.godot_variant_type(VariantTypeBool):
//				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsBool(arg)))
//			case C.godot_variant_type(VariantTypeInt):
//				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsInt(arg)))
//			case C.godot_variant_type(VariantTypeReal):
//				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsReal(arg)))
//			case C.godot_variant_type(VariantTypeString):
//				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsString(arg)))
//			default:
//				log.Fatal("Unknown type of argument")
//			}
//
//			// Convert the pointer into a uintptr so we can perform artithmetic on it.
//			arrayPtr := uintptr(unsafe.Pointer(args))
//
//			// Add the size of the godot_variant pointer to our array pointer to get the position
//			// of the next argument.
//			arg = (*C.godot_variant)(unsafe.Pointer(arrayPtr + size))
//		}
//	}
//
//	// Use the method string to get the class name and method name.
//	log.Println("  Getting class name and method name...")
//	classMethodSlice := strings.Split(methodString, "::")
//	className := classMethodSlice[0]
//	methodName := classMethodSlice[1]
//	log.Println("    Class Name: ", className)
//	log.Println("    Method Name: ", methodName)
//
//	// Look up the registered class so we can find out how many arguments it takes
//	// and their types.
//	log.Println("  Look up the registered class and its method")
//	regClass := classRegistry[className]
//	if regClass == nil {
//		log.Fatal("  This class has not been registered! Class name: ", className, " Method name: ", methodName)
//	}
//	regMethod := regClass.methods[methodName]
//
//	log.Println("  Registered method arguments:", regMethod.arguments)
//	log.Println("  Arguments to pass:", goArgsSlice)
//
//	// Check to ensure the method has the same number of arguments we expect
//	if len(regMethod.arguments)-1 != int(numArgs) {
//		Log.Error("Invalid number of arguments. Expected ", numArgs, " arguments. (Got ", len(regMethod.arguments), ")")
//		panic("Invalid number of arguments.")
//	}
//
//	// Get the value of the class, so we can call methods on it.
//	method := classValue.MethodByName(methodName)
//	rawRet := method.Call(goArgsSlice)
//	log.Println(method)
//
//	var ret *C.godot_variant
//	var nonptrrtn C.godot_variant
//
//	if len(regMethod.returns) == 0 {
//		C.godot_variant_new_nil(&nonptrrtn)
//		return nonptrrtn
//	} else if len(regMethod.returns) > 1 {
//		panic("Too many return values from method! Methods called from within Godot should only return a single value.")
//	}
//
//	// Convert our returned value into a Godot Variant.
//	rawRetInterface := rawRet[0].Interface()
//	switch regMethod.returns[0].String() {
//
//	case "bool":
//		ret = boolAsVariant(rawRetInterface.(bool))
//
//	case "int64":
//		ret = intAsVariant(rawRetInterface.(int64))
//
//	case "int32":
//		ret = intAsVariant(int64(rawRetInterface.(int32)))
//
//	case "int":
//		ret = intAsVariant(int64(rawRetInterface.(int)))
//
//	case "uint64":
//		ret = uIntAsVariant(rawRetInterface.(uint64))
//
//	case "uint32":
//		ret = uIntAsVariant(uint64(rawRetInterface.(uint32)))
//
//	case "uint":
//		ret = uIntAsVariant(uint64(rawRetInterface.(uint)))
//
//	case "float64":
//		ret = realAsVariant(rawRetInterface.(float64))
//
//	case "string":
//		ret = stringAsVariant(rawRetInterface.(string))
//
//	default:
//		panic("The return was not valid. Should be Godot Variant or built-in Go type. Received: " + regMethod.returns[0].String())
//	}
//
//	return *ret
//}
