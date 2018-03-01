package jsonparseresult

import (
	"log"
	"reflect"

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

/*
Returned by [method JSON.parse], [code]JSONParseResult[/code] contains decoded JSON or error information if JSON source not successfully parsed. You can check if JSON source was successfully parsed with [code]if json_result.error == OK[/code].
*/
type JSONParseResult struct {
	Reference
}

func (o *JSONParseResult) BaseClass() string {
	return "JSONParseResult"
}

/*
   Undocumented
*/
func (o *JSONParseResult) GetError() gdnative.Int {
	log.Println("Calling JSONParseResult.GetError()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_error", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *JSONParseResult) GetErrorLine() gdnative.Int {
	log.Println("Calling JSONParseResult.GetErrorLine()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_error_line", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *JSONParseResult) GetErrorString() gdnative.String {
	log.Println("Calling JSONParseResult.GetErrorString()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_error_string", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *JSONParseResult) GetResult() *Variant {
	log.Println("Calling JSONParseResult.GetResult()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_result", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *JSONParseResult) SetError(error gdnative.Int) {
	log.Println("Calling JSONParseResult.SetError()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(error)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_error", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *JSONParseResult) SetErrorLine(errorLine gdnative.Int) {
	log.Println("Calling JSONParseResult.SetErrorLine()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(errorLine)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_error_line", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *JSONParseResult) SetErrorString(errorString gdnative.String) {
	log.Println("Calling JSONParseResult.SetErrorString()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(errorString)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_error_string", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *JSONParseResult) SetResult(result *Variant) {
	log.Println("Calling JSONParseResult.SetResult()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(result)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_result", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   JSONParseResultImplementer is an interface for JSONParseResult objects.
*/
type JSONParseResultImplementer interface {
	Class
}