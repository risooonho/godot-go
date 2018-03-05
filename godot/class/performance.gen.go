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

//func NewperformanceFromPointer(ptr gdnative.Pointer) performance {
func NewPerformanceFromPointer(ptr gdnative.Pointer) performance {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := performance{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonPerformance() *performance {
	return &performance{}
}

/*
   This class provides access to a number of different monitors related to performance, such as memory usage, draw calls, and FPS. These are the same as the values displayed in the [i]Monitor[/i] tab in the editor's [i]Debugger[/i] panel. By using the [method get_monitor] method of this class, you can access this data from your code. Note that a few of these monitors are only available in debug mode and will always return 0 when used in a release build. Many of these monitors are not updated in real-time, so there may be a short delay between changes.
*/
var Performance = newSingletonPerformance()

/*
This class provides access to a number of different monitors related to performance, such as memory usage, draw calls, and FPS. These are the same as the values displayed in the [i]Monitor[/i] tab in the editor's [i]Debugger[/i] panel. By using the [method get_monitor] method of this class, you can access this data from your code. Note that a few of these monitors are only available in debug mode and will always return 0 when used in a release build. Many of these monitors are not updated in real-time, so there may be a short delay between changes.
*/
type performance struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *performance) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("Performance")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *performance) BaseClass() string {
	return "Performance"
}

/*
        Returns the value of one of the available monitors. You should provide one of this class's constants as the argument, like this: [codeblock] print(Performance.get_monitor(Performance.TIME_FPS)) # Prints the FPS to the console [/codeblock]
	Args: [{ false monitor int}], Returns: float
*/
func (o *performance) GetMonitor(monitor gdnative.Int) gdnative.Float {
	o.ensureSingleton()
	//log.Println("Calling Performance.GetMonitor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(monitor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Performance", "get_monitor")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}
